# Sprint 11: Deploy AWS Completo - Tarefas Detalhadas (sprint11_tasks.md)

## Overview
Sprint de deploy completo da infraestrutura na nuvem AWS com foco em provisionar recursos, configurar segurança e validar CI/CD em produção.

---

## Tarefas (DevOps/Programador)

### Área: Preparação e Planejamento

- [ ] **T11.1** Revisar documentação de infra-devops
  - Ler completamente `especification/infra-devops-details.md`
  - Documentar naming convention dos recursos AWS
  - Preparar credential AWS (access key + secret)

- [ ] **T11.2** Preparar scripts de migração de dados
  - Criar script SQL para migrar schema de SQLite para MySQL
  - Incluir dados de teste (admin master, psicólogos, pacientes)
  - Testar script localmente com SQLite → MySQL

- [ ] **T11.3** Configurar credenciais e permissões IAM
  - Criar IAM user com permissões necessárias ou usar STS AssumeRole
  - Configurar AWS CLI com credenciais
  - Testar acesso com `aws sts get-caller-identity`

### Área: Infraestrutura de Rede (VPC)

- [ ] **T11.4** Criar VPC
  - Criar VPC com CIDR `10.0.0.0/16`
  - Nome: `gestao-vpc`
  - Habilitar DNS hostnames

- [ ] **T11.5** Criar subnets públicas
  - Subnet 1: `10.0.1.0/24` em AZ us-east-1a
  - Subnet 2: `10.0.2.0/24` em AZ us-east-1b
  - Nomes: `gestao-public-subnet-1a`, `gestao-public-subnet-1b`

- [ ] **T11.6** Criar subnets privadas
  - Subnet 1: `10.0.10.0/24` em AZ us-east-1a
  - Subnet 2: `10.0.20.0/24` em AZ us-east-1b
  - Nomes: `gestao-private-subnet-1a`, `gestao-private-subnet-1b`

- [ ] **T11.7** Criar Internet Gateway
  - Criar IGW com nome `gestao-igw`
  - Associar à VPC `gestao-vpc`

- [ ] **T11.8** Criar NAT Gateway (FinOps otimizado)
  - Alocar Elastic IP
  - Criar NAT Gateway em subnet pública (1a)
  - Nome: `gestao-nat`

- [ ] **T11.9** Configurar Route Tables
  - Route table pública: `0.0.0.0/0 -> IGW`
  - Route table privada: `0.0.0.0/0 -> NAT Gateway`
  - Associar subnets às suas respectivas route tables

### Área: Security Groups

- [ ] **T11.10** Criar Security Group para ALB
  - Nome: `gestao-alb-sg`
  - Entrada: HTTP 80 de `0.0.0.0/0`
  - Entrada: HTTPS 443 de `0.0.0.0/0`
  - Saída: Tudo

- [ ] **T11.11** Criar Security Group para ECS
  - Nome: `gestao-ecs-sg`
  - Entrada: Porta 8080 do ALB SG
  - Entrada: Porta 3000 do ALB SG
  - Saída: Tudo (para acessar internet/APIs)

- [ ] **T11.12** Criar Security Group para RDS
  - Nome: `gestao-rds-sg`
  - Entrada: Porta 3306 do ECS SG
  - Saída: Nenhuma (ou padrão VPC)

### Área: Banco de Dados (RDS MySQL)

- [ ] **T11.13** Provisionar instância RDS MySQL
  - Engine: MySQL 8.0
  - Classe: db.t3.micro (Free Tier) ou db.t4g.micro
  - Armazenamento: 20 GB (Free Tier)
  - Multi-AZ: Não (para FinOps)
  - Backup: 7 dias de retenção
  - Nome da instância: `gestao-db`

- [ ] **T11.14** Configurar RDS em subnets privadas
  - DB Subnet Group: subnets privadas 1a e 1b
  - Security Group: `gestao-rds-sg`
  - Público acessível: Não

- [ ] **T11.15** Criar banco de dados inicial
  - Conectar à instância RDS
  - Executar script de criação de schema e dados
  - Criar databases: `gestao_psicologos`
  - Validar dados importados com `SELECT COUNT(*) FROM users;`

- [ ] **T11.16** Configurar backups e snapshots
  - Ativar backup automático (7 dias)
  - Testar snapshot manual
  - Documentar processo de restore

### Área: Container Registry (ECR)

- [ ] **T11.17** Criar ECR repository para Frontend
  - Nome: `gestao-frontend`
  - Scan images on push: Habilitado
  - Tag mutability: Habilitado

- [ ] **T11.18** Criar ECR repository para Backend
  - Nome: `gestao-backend`
  - Scan images on push: Habilitado
  - Tag mutability: Habilitado

- [ ] **T11.19** Testar push de imagens ao ECR
  - Build backend: `docker build -t <ecr-url>/gestao-backend:latest ./codebase/backend`
  - Push: `docker push <ecr-url>/gestao-backend:latest`
  - Repetir para frontend
  - Validar no console AWS

### Área: ECS Fargate

- [ ] **T11.20** Criar ECS Cluster
  - Nome: `gestao-cluster`
  - Tipo: Fargate
  - VPC: `gestao-vpc`

- [ ] **T11.21** Criar Task Definition para Backend
  - Família: `gestao-backend`
  - CPU: 0.25 vCPU
  - Memória: 0.5 GB
  - Imagem: ECR backend latest
  - Port: 8080
  - Logs: CloudWatch `/ecs/gestao-backend`
  - Variáveis de ambiente:
    - `DB_DRIVER=mysql`
    - `DB_HOST=<rds-endpoint>`
    - `DB_USER=admin`
    - `DB_PASSWORD=<from Secrets Manager>`
    - `JWT_SECRET=<from Secrets Manager>`
    - `FRONTEND_URL=https://app.gestao-psicologos.com`

- [ ] **T11.22** Criar Task Definition para Frontend
  - Família: `gestao-frontend`
  - CPU: 0.25 vCPU
  - Memória: 0.5 GB
  - Imagem: ECR frontend latest
  - Port: 3000
  - Logs: CloudWatch `/ecs/gestao-frontend`
  - Variáveis:
    - `NEXT_PUBLIC_API_URL=https://app.gestao-psicologos.com/api`

- [ ] **T11.23** Criar Service Backend no ECS
  - Nome: `gestao-backend-service`
  - Task Definition: `gestao-backend`
  - Desejado: 1 tarefa
  - Min: 1, Max: 3 (Auto Scaling)
  - Subnets privadas
  - Security Group: `gestao-ecs-sg`

- [ ] **T11.24** Criar Service Frontend no ECS
  - Nome: `gestao-frontend-service`
  - Task Definition: `gestao-frontend`
  - Desejado: 1 tarefa
  - Min: 1, Max: 3 (Auto Scaling)
  - Subnets privadas
  - Security Group: `gestao-ecs-sg`

### Área: Load Balancer e Roteamento

- [ ] **T11.25** Criar Application Load Balancer (ALB)
  - Nome: `gestao-alb`
  - Tipo: Application
  - VPC: `gestao-vpc`
  - Subnets públicas: 1a e 1b
  - Security Group: `gestao-alb-sg`

- [ ] **T11.26** Criar Target Group Backend
  - Nome: `gestao-backend-tg`
  - Protocolo: HTTP
  - Porta: 8080
  - Health check: `/api/health`
  - Intervalo: 30s, threshold: 2 healthy

- [ ] **T11.27** Criar Target Group Frontend
  - Nome: `gestao-frontend-tg`
  - Protocolo: HTTP
  - Porta: 3000
  - Health check: `/`
  - Intervalo: 30s, threshold: 2 healthy

- [ ] **T11.28** Registrar ECS Services nos Target Groups
  - Backend service → Backend TG
  - Frontend service → Frontend TG
  - Validar registração no console

- [ ] **T11.29** Configurar Listeners no ALB
  - Listener HTTP 80: Redirecionar para HTTPS 443
  - Listener HTTPS 443:
    - `/api/*` → Backend TG
    - `/` → Frontend TG (default)

### Área: CloudFront e SSL/TLS

- [ ] **T11.30** Requisitar certificado SSL no ACM
  - Domínio: `app.gestao-psicologos.com` (ou usar `*.gestao-psicologos.com`)
  - Validação: DNS (recomendado)
  - Região: us-east-1

- [ ] **T11.31** Validar propriedade do domínio
  - Adicionar registros DNS CNAME conforme ACM instruir
  - Aguardar validação (pode levar minutos a horas)
  - Confirmar status "Issued" no ACM console

- [ ] **T11.32** Criar distribuição CloudFront
  - Nome: `gestao-distribution`
  - Origin 1: ALB DNS name (para `/api/*`)
  - Origin 2: ALB DNS name (para raiz)
  - Behavior padrão: Origin 2, Cache enabled
  - Behavior `/api/*`: Origin 1, Cache disabled, Forward all query strings
  - Certificado SSL: ACM certificate
  - Domínio alternativo: `app.gestao-psicologos.com`

- [ ] **T11.33** Configurar headers de segurança no CloudFront
  - Adicionar Security Headers via Lambda@Edge (opcional) ou origin headers
  - `Strict-Transport-Security: max-age=31536000`
  - `X-Content-Type-Options: nosniff`
  - `X-Frame-Options: DENY`

### Área: Secrets Manager

- [ ] **T11.34** Criar Secret no AWS Secrets Manager
  - Name: `gestao-psicologos/prod-secrets`
  - Tipo: Other
  - Conteúdo (JSON):
    ```json
    {
      "DB_PASSWORD": "senha_forte_aqui",
      "JWT_SECRET": "chave_jwt_forte_32_bytes_minimo",
      "GOOGLE_CLIENT_ID": "xxx.apps.googleusercontent.com",
      "GOOGLE_CLIENT_SECRET": "xxx",
      "MICROSOFT_CLIENT_ID": "xxx",
      "MICROSOFT_CLIENT_SECRET": "xxx"
    }
    ```

- [ ] **T11.35** Testar leitura de secrets
  - Atualizar código backend para ler de Secrets Manager
  - Testar localmente com credenciais mockadas
  - Testar em produção via ECS Task Execution Role

### Área: CloudWatch e Observabilidade

- [ ] **T11.36** Criar Log Groups no CloudWatch
  - `/ecs/gestao-backend`
  - `/ecs/gestao-frontend`
  - Retenção: 30 dias

- [ ] **T11.37** Configurar CloudWatch Alarms
  - CPU Backend > 85% por 5 min → SNS email
  - Memória Backend > 85% por 5 min → SNS email
  - ALB Target Unhealthy → SNS email
  - RDS CPU > 80% por 5 min → SNS email

- [ ] **T11.38** Criar CloudWatch Dashboard
  - Widgets:
    - Requisições HTTP por minuto (ALB)
    - Latência P50/P95/P99
    - Erros 4xx/5xx
    - CPU/Memória ECS
    - Conexões RDS ativas

### Área: CI/CD e GitHub Actions

- [ ] **T11.39** Configurar OIDC com AWS
  - Criar IAM role com OIDC provider do GitHub
  - Role name: `GithubActionsDeployRole`
  - Permissões: ECR push, ECS update service

- [ ] **T11.40** Validar GitHub Actions workflow
  - Fazer commit em branch `main`
  - Aguardar execução de `.github/workflows/deploy.yml`
  - Validar:
    - Build successful
    - Push to ECR successful
    - ECS service updated

- [ ] **T11.41** Testar deploy automático
  - Fazer pequena mudança no código
  - Commit em `main`
  - Validar que nova versão foi deployada em produção (via versão de imagem)

### Área: DNS e Acesso Final

- [ ] **T11.42** Configurar DNS no Route 53 (se usar)
  - Criar record A alias apontando para CloudFront distribution
  - Domínio: `app.gestao-psicologos.com`
  - Target: CloudFront distribution

- [ ] **T11.43** Testar acesso à aplicação
  - Acessar `https://app.gestao-psicologos.com` no navegador
  - Esperado: Página de login carrega
  - Verificar certificado SSL válido (cadeado verde)

---

## Checkpoints Internos

- [ ] VPC e subnets criadas com sucesso
- [ ] RDS MySQL acessível e dados importados
- [ ] ECR repositories criados e imagens pushed
- [ ] ECS Cluster, Services e Task Definitions criados
- [ ] ALB e Target Groups registrados
- [ ] CloudFront distribuição ativa com certificado SSL
- [ ] Secrets Manager com credenciais armazenadas
- [ ] CloudWatch Logs e Alarms configurados
- [ ] GitHub Actions CI/CD funcional
- [ ] Aplicação acessível via HTTPS no domínio

---

## Notas para DevOps/Programador

- **Custos**: Monitorar AWS Cost Explorer durante provisioning
- **Free Tier**: Validar que recursos estão dentro do Free Tier
- **Backups**: Testar restore de RDS antes de finalizar sprint
- **SSL Renewal**: ACM renova automaticamente, sem ação necessária

---

## Bloqueadores / Dependências

- Dependência: Sprint 10 deve estar 100% concluído
- Pré-requisito: Credenciais AWS e domínio registrado
