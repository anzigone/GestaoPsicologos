# Projeto: Gestão Psicólogos - Infraestrutura & DevOps (infra-devops-details.md)

Este documento detalha a arquitetura de produção na nuvem AWS, a estratégia de entrega contínua (CI/CD) via GitHub Actions, a gestão de segredos e o monitoramento em produção.

---

## 1. Controle de Código e Versionamento

- **Repositório**: Monorepo hospedado no GitHub no endereço [https://github.com/anzigone/GestaoPsicologos](https://github.com/anzigone/GestaoPsicologos).
- **Branching Model**:
  - `main`: Branch de produção. Modificações entram apenas por Pull Request aprovado e validado.
  - `develop` (opcional) ou branches de feature `feature/*`: Onde o desenvolvimento diário ocorre.
- **Estratégia de Versão (Versionamento Semântico)**:
  - Formato: `{release}.{version}.{build}` (ex: `0.1.0` no setup inicial).
  - O número de `build` será automaticamente incrementado pelo pipeline de CI/CD a cada execução bem-sucedida de deploy no GitHub Actions.
  - Cada entrega de Sprint incrementa a `{version}` ou `{release}` dependendo do escopo de entrega humana validada.

---

## 2. Arquitetura de Nuvem na AWS

O design de infraestrutura prioriza custo-eficiência (FinOps) utilizando serviços serverless e Free Tier da AWS.

```text
                  [ Usuário Final ]
                          │ (HTTPS)
                          ▼
                 [ AWS Route 53 / ACM ]
                          │
                          ▼
                [ AWS CloudFront CDN ]
                 ├── /api/* ──► [ AWS Application Load Balancer ]
                 │                      │
                 └── (Padrão) ──►       ▼ (Target Groups)
                                 [ AWS ECS Fargate Cluster ]
                                  ├── Service: gestao-frontend
                                  └── Service: gestao-backend
                                               │
                                               ▼ (Subnets Privadas)
                                        [ AWS RDS MySQL ]
```

### Detalhes dos Componentes AWS:
1. **Rede (VPC)**:
   - Uma VPC com 2 Subnets Públicas (para ALB e instâncias externas temporárias) e 2 Subnets Privadas (para rodar os containers ECS Fargate e o RDS MySQL de forma segura).
   - Um **NAT Gateway** em modo simplificado ou apenas comunicação interna local para controle de custos, usando VPC Endpoints para comunicação direta com serviços AWS (ECR, Secrets Manager, CloudWatch).
2. **CDN (AWS CloudFront)**:
   - Ponto de entrada público único. O tráfego HTTP é redirecionado para HTTPS através de certificados emitidos no **AWS ACM (Certificate Manager)**.
   - **Comportamentos (Behaviors)**:
     - `/api/*`: Encaminha as requisições HTTP REST diretamente para o **Application Load Balancer (ALB)** do backend, mantendo os cabeçalhos de autenticação e desabilitando cache.
     - `Default (*)`: Encaminha para o serviço do **Frontend** (Next.js) em execução no ECS Fargate.
3. **Orquestração de Containers (AWS ECS Fargate)**:
   - **Cluster ECS** compartilhado.
   - **Service Frontend (`gestao-frontend`)**:
     - Tipo de Inicialização: Fargate (Serverless, sem gerenciamento de EC2).
     - Recursos: 0.25 vCPU e 0.5 GB RAM.
     - Dimensionamento: Mínimo 1 container, escalando sob demanda.
   - **Service Backend (`gestao-backend`)**:
     - Tipo de Inicialização: Fargate.
     - Recursos: 0.25 vCPU e 0.5 GB RAM.
     - Integração com o Target Group do ALB.
4. **Banco de Dados (AWS RDS MySQL)**:
   - Instância MySQL 8.0 Free Tier (`db.t3.micro` ou `db.t4g.micro`, Single-AZ).
   - Localizado estritamente nas subnets privadas e acessível apenas pelo Security Group do container `gestao-backend`.

---

## 3. Segurança e Gestão de Segredos

- **Variables locais (.env)**: Arquivos `.env` locais no frontend e backend nunca são comitados (bloqueados via `.gitignore`).
- **AWS Secrets Manager / Systems Manager Parameter Store**:
  - Em produção, todas as credenciais sensíveis (senha do banco RDS, segredo do JWT, Client Secret do Google/Outlook) são armazenadas no AWS Secrets Manager sob o nome `gestao-psicologos/prod-secrets`.
- **Injeção de Variáveis na Task Definition do ECS**:
  - As variáveis de ambiente do container em produção são povoadas dinamicamente pelo ECS buscando as chaves seguras do Secrets Manager durante o boot do container, evitando que fiquem salvas em texto puro na imagem Docker.
- **Permissões IAM (Role-Based)**:
  - **Task Execution Role**: Concede permissão para o agente do ECS baixar imagens do ECR e ler segredos do Secrets Manager.
  - **Task Role**: Concede permissão para a aplicação Go acessar recursos AWS em execução (ex: integração com Simple Email Service - SES para envio de emails).

---

## 4. Pipeline de CI/CD (GitHub Actions)

O deploy contínuo será automatizado a cada push na branch `main` usando GitHub Actions. O arquivo de configuração ficará em `.github/workflows/deploy.yml`.

### Script de Integração e Deploy Contínuo (`.github/workflows/deploy.yml`)

```yaml
name: Deploy Gestao Psicologos to AWS

on:
  push:
    branches:
      - main

permissions:
  contents: read
  id-token: write  # Necessário para autenticação OIDC com AWS

env:
  AWS_REGION: us-east-1
  ECR_FRONTEND_REPOSITORY: gestao-frontend
  ECR_BACKEND_REPOSITORY: gestao-backend
  ECS_CLUSTER: gestao-cluster
  ECS_FRONTEND_SERVICE: gestao-frontend-service
  ECS_BACKEND_SERVICE: gestao-backend-service
  FRONTEND_TASK_DEF: .aws/task-def-frontend.json
  BACKEND_TASK_DEF: .aws/task-def-backend.json

jobs:
  build-and-deploy:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout Code
        uses: actions/checkout@v4

      # 1. Autenticação na AWS usando OIDC (seguro, sem armazenar chaves permanentes no Github)
      - name: Configure AWS Credentials
        uses: aws-actions/configure-aws-credentials@v4
        with:
          role-to-assume: arn:aws:iam::123456789012:role/GithubActionsDeployRole
          aws-region: ${{ env:AWS_REGION }}

      - name: Login to Amazon ECR
        id: login-ecr
        uses: aws-actions/amazon-ecr-login@v2

      # 2. Build e Push do Backend (Golang)
      - name: Build, Tag, and Push Backend Image to ECR
        id: build-backend
        env:
          ECR_REGISTRY: ${{ steps.login-ecr.outputs.registry }}
          IMAGE_TAG: ${{ github.sha }}
        run: |
          docker build -t $ECR_REGISTRY/$ECR_BACKEND_REPOSITORY:$IMAGE_TAG -t $ECR_REGISTRY/$ECR_BACKEND_REPOSITORY:latest ./codebase/backend
          docker push $ECR_REGISTRY/$ECR_BACKEND_REPOSITORY:$IMAGE_TAG
          docker push $ECR_REGISTRY/$ECR_BACKEND_REPOSITORY:latest
          echo "backend_image=$ECR_REGISTRY/$ECR_BACKEND_REPOSITORY:$IMAGE_TAG" >> $GITHUB_OUTPUT

      # 3. Build e Push do Frontend (Next.js)
      - name: Build, Tag, and Push Frontend Image to ECR
        id: build-frontend
        env:
          ECR_REGISTRY: ${{ steps.login-ecr.outputs.registry }}
          IMAGE_TAG: ${{ github.sha }}
        run: |
          docker build -t $ECR_REGISTRY/$ECR_FRONTEND_REPOSITORY:$IMAGE_TAG -t $ECR_REGISTRY/$ECR_FRONTEND_REPOSITORY:latest ./codebase/frontend
          docker push $ECR_REGISTRY/$ECR_FRONTEND_REPOSITORY:$IMAGE_TAG
          docker push $ECR_REGISTRY/$ECR_FRONTEND_REPOSITORY:latest
          echo "frontend_image=$ECR_REGISTRY/$ECR_FRONTEND_REPOSITORY:$IMAGE_TAG" >> $GITHUB_OUTPUT

      # 4. Atualização das Task Definitions com as novas imagens
      - name: Fill in the new Backend Image ID in the Amazon ECS Task Definition
        id: task-def-backend
        uses: aws-actions/amazon-ecs-render-task-definition@v1
        with:
          task-definition: ${{ env.BACKEND_TASK_DEF }}
          container-name: backend-container
          image: ${{ steps.build-backend.outputs.backend_image }}

      - name: Fill in the new Frontend Image ID in the Amazon ECS Task Definition
        id: task-def-frontend
        uses: aws-actions/amazon-ecs-render-task-definition@v1
        with:
          task-definition: ${{ env.FRONTEND_TASK_DEF }}
          container-name: frontend-container
          image: ${{ steps.build-frontend.outputs.frontend_image }}

      # 5. Deploy no AWS ECS (Fargate)
      - name: Deploy Amazon ECS Backend Task Definition
        uses: aws-actions/amazon-ecs-deploy-task-definition@v2
        with:
          task-definition: ${{ steps.task-def-backend.outputs.task-definition }}
          service: ${{ env.ECS_BACKEND_SERVICE }}
          cluster: ${{ env.ECS_CLUSTER }}
          wait-for-service-stability: true

      - name: Deploy Amazon ECS Frontend Task Definition
        uses: aws-actions/amazon-ecs-deploy-task-definition@v2
        with:
          task-definition: ${{ steps.task-def-frontend.outputs.task-definition }}
          service: ${{ env.ECS_FRONTEND_SERVICE }}
          cluster: ${{ env.ECS_CLUSTER }}
          wait-for-service-stability: true
```

---

## 5. Observabilidade e Logs (AWS CloudWatch)

- **Configuração de Logs do ECS**:
  - Os containers ECS Fargate utilizarão o driver de logs nativo `awslogs` para direcionar a saída padrão (`Stdout` e `Stderr`) diretamente para grupos de logs do CloudWatch:
    - `/ecs/gestao-frontend`
    - `/ecs/gestao-backend`
- **Métricas e Alertas**:
  - Criação de alarmes simples no CloudWatch para consumo de CPU/Memória acima de 85% por mais de 5 minutos.
  - Alerta de erros críticos de conexões do banco de dados monitorando logs com filtro de palavras-chave (ex: `"ERROR"` ou `"FATAL"`).
