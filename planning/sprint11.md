# Sprint 11: Deploy Completo na AWS (sprint11.md)

## 1. Escopo e Objetivos

O objetivo desta sprint é realizar o deploy completo da aplicação na nuvem AWS, estabelecendo toda a infraestrutura de produção, segurança e observabilidade. O sistema deverá estar totalmente operacional em ambiente de produção ao final desta sprint.

**Objetivos principais**:
- Criar e configurar a infraestrutura AWS (VPC, subnets, security groups, NAT Gateway).
- Configurar o banco de dados RDS MySQL em produção (Single-AZ, db.t3.micro Free Tier).
- Migrar dados do SQLite de desenvolvimento para MySQL RDS.
- Configurar AWS ECR (Elastic Container Registry) para armazenar as imagens Docker.
- Implementar as task definitions do ECS Fargate para Frontend e Backend.
- Configurar o Application Load Balancer (ALB) para roteamento de tráfego.
- Implementar o AWS CloudFront como CDN e proxy reverso.
- Configurar SSL/TLS com AWS Certificate Manager (ACM).
- Integrar AWS Secrets Manager para armazenamento seguro de credenciais.
- Configurar CloudWatch Logs para observabilidade centralizada.
- Validar o pipeline CI/CD do GitHub Actions para deploy automático.

---

## 2. Entregáveis da Sprint

Ao final desta sprint, os seguintes entregáveis deverão estar prontos:

- **Infraestrutura AWS**:
  - VPC com 2 subnets públicas e 2 privadas.
  - Security Groups para ALB, ECS Fargate e RDS.
  - NAT Gateway configurado em modo otimizado (FinOps).
  - Internet Gateway e tabelas de rota.

- **Banco de Dados**:
  - RDS MySQL 8.0 Free Tier provisionado em subnets privadas.
  - Schema e dados migrados do SQLite para MySQL com sucesso.
  - Backups automáticos ativados.

- **Container Registry**:
  - ECR repositories criados para `gestao-frontend` e `gestao-backend`.
  - Imagens Docker multi-stage compiladas e pushadas com sucesso.

- **Orquestração de Containers**:
  - ECS Cluster criado (Fargate).
  - Task Definitions para Frontend (0.25 vCPU, 0.5GB RAM) e Backend (0.25 vCPU, 0.5GB RAM).
  - Services ECS com auto-scaling mínimo 1 e máximo 3 replicas.

- **Load Balancing e CDN**:
  - Application Load Balancer (ALB) configurado com target groups.
  - CloudFront distribuição criada com comportamentos para `/api/*` (ALB) e raiz (Frontend).
  - Certificado SSL/TLS ativo no CloudFront.

- **Segurança e Secrets**:
  - AWS Secrets Manager com a chave `gestao-psicologos/prod-secrets` contendo:
    - `DB_PASSWORD` (senha do RDS MySQL).
    - `JWT_SECRET` (chave de assinatura JWT).
    - `GOOGLE_CLIENT_ID` e `GOOGLE_CLIENT_SECRET`.
    - `MICROSOFT_CLIENT_ID` e `MICROSOFT_CLIENT_SECRET`.

- **Observabilidade**:
  - CloudWatch Log Groups criados: `/ecs/gestao-frontend` e `/ecs/gestao-backend`.
  - Alarmes configurados para CPU/Memória > 85% por 5+ minutos.
  - Dashboard CloudWatch visualizando métricas principais.

- **CI/CD**:
  - GitHub Actions workflow (`.github/workflows/deploy.yml`) testado e funcional.
  - Pipeline capaz de compilar, fazer push das imagens ao ECR e atualizar task definitions no ECS.

---

## 3. Tarefas Detalhadas

### 3.1. Preparação e Planejamento
- [ ] Revisar a especificação de infra-devops-details.md.
- [ ] Preparar checklist de permissões IAM necessárias.
- [ ] Documentar os nomes de recursos AWS a serem criados (naming convention).
- [ ] Preparar scripts de migração de dados SQLite → MySQL.

### 3.2. Criação de Infraestrutura de Rede
- [ ] Criar VPC com CIDR block `10.0.0.0/16`.
- [ ] Criar 2 subnets públicas (`10.0.1.0/24` e `10.0.2.0/24`).
- [ ] Criar 2 subnets privadas (`10.0.10.0/24` e `10.0.20.0/24`).
- [ ] Criar e associar Internet Gateway.
- [ ] Criar NAT Gateway em modo otimizado (alternativa: usar VPC Endpoints para economia).
- [ ] Configurar tabelas de rota (públicas e privadas).

### 3.3. Configuração de Security Groups
- [ ] Criar Security Group para ALB (entrada: HTTP 80, HTTPS 443 de 0.0.0.0/0).
- [ ] Criar Security Group para ECS Fargate (entrada: porta 8080/3000 do ALB).
- [ ] Criar Security Group para RDS MySQL (entrada: porta 3306 do ECS, saída: qualquer).

### 3.4. Banco de Dados RDS
- [ ] Provisionar instância RDS MySQL 8.0 (db.t3.micro ou db.t4g.micro).
- [ ] Configurar localização em subnets privadas.
- [ ] Associar Security Group de RDS.
- [ ] Ativar backups automáticos (retenção 7 dias).
- [ ] Criar banco de dados `gestao_psicologos` na instância.
- [ ] Executar script de migração DDL (tabelas users, patients, sessions, etc.).
- [ ] Validar conectividade e dados no RDS.

### 3.5. ECR e Container Registry
- [ ] Criar ECR repository para `gestao-frontend`.
- [ ] Criar ECR repository para `gestao-backend`.
- [ ] Configurar política de retenção de imagens (ex: últimas 5 imagens).
- [ ] Testar push manual de imagens para validar permissões.

### 3.6. ECS Fargate Setup
- [ ] Criar ECS Cluster com nome `gestao-cluster`.
- [ ] Criar Task Definition para Backend:
  - Imagem: ECR Backend.
  - CPU: 0.25 vCPU, Memória: 0.5GB.
  - Variáveis de ambiente: `DB_DRIVER=mysql`, `FRONTEND_URL`, `JWT_SECRET` (via Secrets Manager).
  - Logs: CloudWatch (`/ecs/gestao-backend`).
- [ ] Criar Task Definition para Frontend:
  - Imagem: ECR Frontend.
  - CPU: 0.25 vCPU, Memória: 0.5GB.
  - Variáveis de ambiente: `NEXT_PUBLIC_API_URL`.
  - Logs: CloudWatch (`/ecs/gestao-frontend`).
- [ ] Criar Service Backend com Auto Scaling (min: 1, max: 3, target CPU: 70%).
- [ ] Criar Service Frontend com Auto Scaling (min: 1, max: 3, target CPU: 70%).

### 3.7. Application Load Balancer (ALB)
- [ ] Criar ALB em subnets públicas.
- [ ] Associar Security Group de ALB.
- [ ] Criar target group para Backend (porta 8080, health check `/api/health`).
- [ ] Criar target group para Frontend (porta 3000, health check `/`).
- [ ] Associar listener HTTP 80 com redirecionamento para HTTPS 443.
- [ ] Criar listener HTTPS 443 com certificado ACM.

### 3.8. CloudFront e ACM
- [ ] Solicitar certificado SSL/TLS no AWS Certificate Manager para domínio da aplicação (ex: `app.gestao-psicologos.com`).
- [ ] Validar propriedade do domínio no ACM.
- [ ] Criar distribuição CloudFront:
  - Origin 1: ALB (para `/api/*`).
  - Origin 2: Frontend ALB (para raiz e demais caminhos).
  - Comportamento padrão: Frontend com cache.
  - Comportamento `/api/*`: Backend sem cache, CORS ativo.
  - Certificado: ACM (HTTPS obrigatório).
- [ ] Configurar headers de segurança (HSTS, X-Frame-Options, etc.).

### 3.9. Secrets Manager e IAM
- [ ] Criar secret no AWS Secrets Manager com nome `gestao-psicologos/prod-secrets`.
- [ ] Armazenar credenciais:
  - `DB_PASSWORD`: Senha do RDS.
  - `JWT_SECRET`: Chave aleatória forte (mín. 32 bytes).
  - `GOOGLE_CLIENT_ID` e `GOOGLE_CLIENT_SECRET`.
  - `MICROSOFT_CLIENT_ID` e `MICROSOFT_CLIENT_SECRET`.
- [ ] Criar IAM Role para Task Execution (permissão para ler Secrets Manager e ECR).
- [ ] Criar IAM Role para Task (permissão para SES para envio de e-mails, se necessário).

### 3.10. CloudWatch e Observabilidade
- [ ] Criar Log Groups: `/ecs/gestao-frontend` e `/ecs/gestao-backend`.
- [ ] Configurar CloudWatch Alarms:
  - Backend CPU > 85% por 5 min → SNS email.
  - Backend Memória > 85% por 5 min → SNS email.
  - RDS CPU > 80% por 5 min → SNS email.
  - ALB Target Unhealthy → SNS email.
- [ ] Criar CloudWatch Dashboard com widgets:
  - Gráfico de requisições HTTP por minuto.
  - Gráfico de latência P50/P95/P99 do ALB.
  - Contagem de erros 4xx e 5xx.
  - CPU e Memória dos containers.

### 3.11. CI/CD Integration
- [ ] Testar manualmente o push de uma branch para `main`.
- [ ] Validar execução do GitHub Actions workflow (`.github/workflows/deploy.yml`).
- [ ] Confirmar compilação, push ao ECR e atualização do ECS.
- [ ] Verificar implantação automática no cluster.

### 3.12. Validação e Testes de Produção
- [ ] Acessar a URL do CloudFront (ex: `https://app.gestao-psicologos.com`).
- [ ] Testar login com credenciais do admin master.
- [ ] Executar fluxo completo: criar psicólogo, paciente, sessão.
- [ ] Verificar logs no CloudWatch.
- [ ] Testar failover de container (pausar um container, verificar se o Auto Scaling escala).
- [ ] Validar backup automático do RDS.

---

## 4. Checkpoints de Validação Humana

Para validar a conclusão desta Sprint, o humano realizará os seguintes testes:

1. **Infraestrutura Provisionada**:
   - [ ] Acessar AWS Console e confirmar: VPC, subnets, security groups, RDS, ECS Cluster.
   - [ ] Executar query no RDS para confirmar dados migrados com sucesso.

2. **Conectividade e Load Balancing**:
   - [ ] Acessar a URL do CloudFront (ex: `https://app.gestao-psicologos.com`).
   - [ ] Esperado: Página de login carrega sem erros.

3. **Fluxo de Login e Autenticação**:
   - [ ] Login com credenciais do admin (`admin/admin`).
   - [ ] Esperado: Redirecionamento para welcome, autenticação com sucesso.

4. **Operações de Negócio**:
   - [ ] Criar um novo psicólogo via `/admin/users`.
   - [ ] Criar um paciente e uma sessão com o novo psicólogo.
   - [ ] Acessar `/dashboard` e confirmar KPIs e gráficos.
   - [ ] Esperado: Toda operação retorna dados do banco RDS em produção.

5. **Observabilidade**:
   - [ ] Acessar CloudWatch Dashboard.
   - [ ] Verificar logs de requisições em `/ecs/gestao-backend`.
   - [ ] Confirmar métricas de CPU e memória.

6. **Segurança**:
   - [ ] Verificar que a URL usa HTTPS (certificado válido).
   - [ ] Inspeccionar headers de segurança (HSTS, X-Frame-Options).
   - [ ] Confirmar que credenciais não aparecem em logs ou variáveis de ambiente visíveis.

7. **CI/CD Pipeline**:
   - [ ] Realizar um commit teste em uma branch, fazer PR para `main`, mergear.
   - [ ] Aguardar execução automática do GitHub Actions.
   - [ ] Confirmar que nova versão foi deployada automaticamente no ECS.

8. **Encerramento da Sprint**:
   - [ ] Realizar commit com a mensagem `feat: sprint 11 deploy aws completo concluido`.
   - [ ] Criar a tag Git `v0.11.0`.
   - [ ] Fazer push para a branch `main` do GitHub.

---

## 5. Notas Importantes

- **Custo**: Monitorar o AWS Cost Explorer durante a sprint para garantir otimizações e ativar AWS Budgets.
- **Free Tier**: Garantir que recursos estejam dentro do Free Tier (RDS, ECS Fargate, CloudFront tráfego limitado).
- **Backups**: Testar restauração de backup do RDS antes de finalizr a sprint.
- **Domínio**: Se usar domínio customizado, configurar DNS com Route 53 apontando para CloudFront.
- **SSL Renewal**: ACM renova automaticamente certificados, sem ação manual necessária.
