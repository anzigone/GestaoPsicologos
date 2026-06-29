# Projeto: Gestão Psicólogos - Detalhamento da Arquitetura (archicteture-details.md)

Este documento descreve detalhadamente o setup de desenvolvimento, as versões de software, a estrutura do codebase, as configurações do Docker, o design da API, a segurança e a integração com agendas externas.

---

## 1. Stack Tecnológica e Versões Mínimas

### Frontend
- **Framework**: Next.js 14.2+ (App Router para roteamento moderno, carregamento otimizado e Server Actions).
- **Linguagem**: TypeScript 5.0+ (tipagem estática para maior segurança no código).
- **Estilização**: TailwindCSS 3.4+ (utilitários de CSS para desenvolvimento rápido de interfaces responsivas).
- **Componentes**: Shadcn/UI (baseado em Radix UI e Lucide React para ícones, fornecendo componentes acessíveis e customizáveis).
- **State Management / Data Fetching**: React Context para estados globais leves (ex: tema, sessão) e SWR (Stale-While-Revalidate) ou React Query para fetch/cache de dados do Backend.

### Backend
- **Linguagem/Framework**: Golang 1.22+ (alto desempenho, concorrência nativa e baixo consumo de memória).
- **Roteador HTTP**: `chi` v5 (roteador leve, compatível com a assinatura padrão de HTTP do Go, ideal para APIs REST).
- **Banco de Dados (Híbrido)**:
  - **Desenvolvimento Local (DEV)**: SQLite (driver: `modernc.org/sqlite` - driver em Go puro que elimina a necessidade de CGO).
  - **Produção (PROD - AWS)**: MySQL 8.0+ (driver: `github.com/go-sql-driver/mysql` - driver em Go puro).
  - O Backend selecionará dinamicamente o driver de banco através da variável de ambiente `DB_DRIVER` (`sqlite` ou `mysql`).
- **Documentação de API**: Swagger / OpenAPI 3.0 (usando `swaggo/swag` para gerar os documentos a partir dos comentários do código).

---

## 2. Estrutura do Codebase

A estrutura de diretórios do projeto no repositório será organizada da seguinte forma sob o diretório `./codebase`:

```text
./codebase/
├── backend/
│   ├── cmd/
│   │   └── server/
│   │       └── main.go         # Ponto de entrada do webservice
│   ├── internal/
│   │   ├── auth/               # Autenticação (JWT, hash SHA256)
│   │   ├── config/             # Leitura de variáveis de ambiente
│   │   ├── database/           # Conexão híbrida SQLite/MySQL, migrações e repositórios
│   │   ├── integration/        # Clientes HTTP para Google Meet e MS Outlook
│   │   ├── logger/             # Gerenciamento de logs (WARN, ERROR, INFO)
│   │   └── handlers/           # Controladores HTTP (BFF e API endpoints)
│   ├── docs/                   # Arquivos gerados do Swagger
│   ├── go.mod
│   ├── go.sum
│   └── Dockerfile
│
├── frontend/
│   ├── src/
│   │   ├── app/                # Estrutura do Next.js App Router
│   │   ├── components/         # Componentes reutilizáveis do Shadcn/UI
│   │   ├── hooks/              # Custom hooks (ex: useAuth, useSWR)
│   │   ├── lib/                # Funções utilitárias (ex: api client)
│   │   └── types/              # Definições de tipos TypeScript
│   ├── public/                 # Imagens, fontes e ícones estáticos
│   ├── package.json
│   ├── tailwind.config.js
│   ├── next.config.js
│   └── Dockerfile
│
└── docker-compose.yml          # Setup de desenvolvimento local
```

---

## 3. Banco de Dados & Estratégia de Migração Híbrida

A aplicação deve ser capaz de operar de forma transparente com **SQLite** localmente e **MySQL** na nuvem AWS. Para isso:
1. Usaremos tipos SQL genéricos compatíveis com ambos os dialetos.
2. A inicialização do banco utilizará `database/sql` mapeando o driver correspondente de acordo com as variáveis de ambiente:
   - Se `DB_DRIVER=sqlite`, inicializa conexão com o arquivo local e usa o driver `modernc.org/sqlite`.
   - Se `DB_DRIVER=mysql`, inicializa conexão de rede com o RDS MySQL usando `github.com/go-sql-driver/mysql`.

### Modelagem de Dados Base (Compatível com SQLite e MySQL)

```sql
-- Tabela de Usuários (Psicólogos e Master Admin)
CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(36) PRIMARY KEY,        -- UUID v4
    email VARCHAR(255) UNIQUE NOT NULL,-- E-mail de login do psicólogo/admin
    password_hash VARCHAR(255) NOT NULL, -- Hash da senha do usuário
    role VARCHAR(20) NOT NULL,         -- 'admin' (Master) ou 'psicologo'
    name VARCHAR(255) NOT NULL,
    crp VARCHAR(50),                   -- CRP (nulo para admin master)
    specialty VARCHAR(100),            -- Especialidade (nulo para admin master)
    phone VARCHAR(20),                 -- Telefone do psicólogo
    base_fee DECIMAL(10,2) DEFAULT 0.00, -- Valor padrão da consulta
    package_sessions INT,              -- Qtd de sessões do pacote
    package_fee DECIMAL(10,2),         -- Valor do pacote
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Tabela de Integrações OAuth2 (Google e Microsoft)
CREATE TABLE IF NOT EXISTS oauth_tokens (
    user_id VARCHAR(36) NOT NULL,
    provider VARCHAR(20) NOT NULL,     -- 'google' ou 'outlook'
    access_token TEXT NOT NULL,
    refresh_token TEXT,
    expiry DATETIME NOT NULL,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, provider),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Tabela de Sessões Ativas (Para controle de revogação/expiração de JWT se necessário)
CREATE TABLE IF NOT EXISTS active_sessions (
    id VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36) NOT NULL,
    token TEXT NOT NULL,
    expires_at DATETIME NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
```
*(Nota: O trecho `ON UPDATE CURRENT_TIMESTAMP` no MySQL é válido e para o SQLite faremos um trigger simples ou atualização via código na struct de repositório caso o driver local não o suporte nativamente).*

---

## 4. Segurança

- **Credenciais Sensíveis**: Nenhuma chave, credencial ou segredo do OAuth do Google/Microsoft ou chave JWT pode estar hardcoded. Devem ser fornecidas via variáveis de ambiente (`.env`).
- **Autenticação de Usuários**: A senha fornecida pelo usuário passa por um processo de hash usando **Bcrypt** ou **SHA256 com salt individual por usuário** antes de ser comparada/armazenada no banco.
- **Autenticação de APIs (Token JWT)**:
  - O Backend emitirá um token JWT assinado digitalmente no login bem-sucedido.
  - O token conterá os claims básicos (`sub` = ID do usuário, `role` = papel, `exp` = tempo de expiração de 24 horas).
  - O Frontend enviará o token no cabeçalho `Authorization: Bearer <token>` em todas as requisições protegidas.
- **CORS (Cross-Origin Resource Sharing)**:
  - O backend irá expor um middleware CORS configurado para aceitar requisições apenas da URL de origem do Frontend definida no `.env` (ex: `FRONTEND_URL=http://localhost:3000` em desenvolvimento).

---

## 5. API Design & Comunicação

A comunicação será via HTTP REST com payloads em JSON.

### Endpoints de Autenticação e Gestão de Usuários

#### 1. `POST /api/auth/login`
- **Descrição**: Autentica um usuário (psicólogo ou admin master).
- **Payload Entrada**:
  ```json
  {
    "email": "usuario@exemplo.com",
    "password": "senha_em_texto_plano"
  }
  ```
- **Resposta Sucesso (200 OK)**:
  ```json
  {
    "token": "eyJhbGciOi...",
    "user": {
      "id": "uuid-v4-do-usuario",
      "email": "usuario@exemplo.com",
      "name": "Nome do Usuário",
      "role": "admin"
    }
  }
  ```

#### 2. `POST /api/auth/change-password` (Autenticado)
- **Descrição**: Permite ao psicólogo logado trocar sua senha.
- **Payload Entrada**:
  ```json
  {
    "current_password": "senha_antiga",
    "new_password": "nova_senha"
  }
  ```
- **Resposta Sucesso (200 OK)**:
  ```json
  { "message": "Senha alterada com sucesso." }
  ```

#### 3. `GET /api/admin/users` (Autenticado - Apenas Admin Master)
- **Descrição**: Lista todos os psicólogos cadastrados no sistema.
- **Resposta Sucesso (200 OK)**:
  ```json
  [
    {
      "id": "uuid-v4-1",
      "email": "psicologo1@exemplo.com",
      "name": "Dr. Silva",
      "crp": "12/34567",
      "role": "psicologo"
    }
  ]
  ```

#### 4. `POST /api/admin/users` (Autenticado - Apenas Admin Master)
- **Descrição**: Cadastra um novo psicólogo.
- **Payload Entrada**:
  ```json
  {
    "email": "psicologo2@exemplo.com",
    "password": "senha_inicial_temporaria",
    "name": "Dra. Maria",
    "crp": "06/98765"
  }
  ```
- **Resposta Sucesso (201 Created)**:
  ```json
  {
    "id": "uuid-v4-2",
    "email": "psicologo2@exemplo.com",
    "name": "Dra. Maria",
    "crp": "06/98765",
    "role": "psicologo"
  }
  ```

#### 5. `DELETE /api/admin/users/{id}` (Autenticado - Apenas Admin Master)
- **Descrição**: Remove um usuário psicólogo do sistema.
- **Resposta Sucesso (204 No Content)**.

---

## 6. Integrações de Terceiros (Google Meet & MS Outlook)

Para oferecer reuniões integradas e sincronização de agendas, o backend lidará com as APIs oficiais do **Google Calendar** (para reuniões no Meet) e **Microsoft Graph** (para agenda do Outlook).

### Mecanismo de OAuth2 Seguro
1. O psicólogo navega até a tela de configurações no Frontend e clica em "Conectar com Google" ou "Conectar com Outlook".
2. O Frontend redireciona para o endpoint do Backend correspondente (ex: `/api/integrations/google/connect`).
3. O Backend gera um link OAuth2 com parâmetros de escopo adequados (ex: `https://www.googleapis.com/auth/calendar.events` para Google e `Calendars.ReadWrite` para Microsoft) e redireciona o usuário.
4. Após o consentimento do usuário, a plataforma externa o redireciona de volta para `/api/integrations/{provider}/callback` enviando um código temporário.
5. O Backend intercepta o código, faz a requisição de troca do código pelo `access_token` e `refresh_token` de longa duração, associa ao ID do psicólogo autenticado e salva os tokens na tabela `oauth_tokens` no banco de dados.
6. A partir daí, o Backend renovará o `access_token` de forma assíncrona usando o `refresh_token` sempre que estiver próximo do vencimento (`expiry`).

### Fluxo do Google Meet
- **Endpoint Utilizado**: Google Calendar API - `events.insert` com a opção `conferenceDataVersion=1`.
- **Payload de Criação**:
  ```json
  {
    "summary": "Consulta Psicológica - Paciente X",
    "start": { "dateTime": "2026-06-30T10:00:00Z" },
    "end": { "dateTime": "2026-06-30T11:00:00Z" },
    "conferenceData": {
      "createRequest": {
        "requestId": "uuid-v4-aleatorio",
        "conferenceSolutionKey": { "type": "hangoutsMeet" }
      }
    }
  }
  ```
- **Retorno**: O link gerado em `conferenceData.entryPoints[0].uri` (URL do Google Meet) será extraído pelo Backend e salvo na sessão do paciente no SQLite/MySQL, além de enviado por e-mail para o paciente.

### Fluxo do MS Outlook Calendar
- **Endpoint Utilizado**: Microsoft Graph API `/me/events`.
- **Operações suportadas**:
  - `GET /me/calendar/events` para retornar a lista de compromissos e renderizar o calendário no Frontend.
  - `POST /me/events` para criar um novo compromisso na agenda do Outlook.
  - `DELETE /me/events/{event_id}` para remover um compromisso alocado.

---

## 7. Observabilidade e Logs

O backend utilizará um sistema de logs simples, estruturado e de fácil análise:
- **Níveis de Log**:
  - `INFO`: Registros operacionais de rotina (ex: "Servidor iniciado na porta 8080", "Conexão com SQLite realizada").
  - `WARN`: Eventos inesperados que não impedem a execução (ex: "Falha ao tentar renovar token OAuth2 expirado do usuário X, nova tentativa em 5m").
  - `ERROR`: Falhas graves de processamento que exigem intervenção ou tratamento de erro (ex: "Erro ao abrir arquivo sqlite.db", "Falha de conexão com a API do Graph").
- **Saída**:
  - Console (`Stdout`/`Stderr`) para captura direta em sistemas de containers (CloudWatch Logs via ECS).
  - Arquivo local texto simples persistente `/var/log/gestao-psicologos/app.log` (apenas para desenvolvimento/ambientes que requeiram I/O local persistente).

---

## 8. Runtime e Dockerização (Desenvolvimento Local)

### Dockerfile do Backend (`./codebase/backend/Dockerfile`)
```dockerfile
# Estágio de Compilação
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# CGO_ENABLED=0 para compilação estática (Go puro, drivers modernos sqlite e mysql não precisam de CGO)
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/server

# Estágio de Execução
FROM alpine:3.19
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /app
COPY --from=builder /app/server /app/server
# Cria a pasta de dados onde o SQLite persistirá (utilizado localmente)
RUN mkdir -p /app/data && chmod 777 /app/data
EXPOSE 8080
VOLUME ["/app/data"]
CMD ["/app/server"]
```

### Dockerfile do Frontend (`./codebase/frontend/Dockerfile`)
```dockerfile
# Estágio de Dependências e Build
FROM node:20-alpine AS builder
WORKDIR /app
COPY package*.json ./
RUN npm ci
COPY . .
RUN npm run build

# Estágio de Execução
FROM node:20-alpine AS runner
WORKDIR /app
ENV NODE_ENV production
COPY --from=builder /app/public ./public
COPY --from=builder /app/.next/standalone ./.next/standalone
COPY --from=builder /app/.next/static ./.next/static
EXPOSE 3000
CMD ["node", ".next/standalone/server.js"]
```

### Docker Compose Local (`./docker-compose.yml`)
```yaml
version: '3.8'

services:
  backend:
    build:
      context: ./codebase/backend
      dockerfile: Dockerfile
    ports:
      - "8080:8080"
    environment:
      - PORT=8080
      - DB_DRIVER=sqlite
      - DATABASE_PATH=/app/data/gestao.db
      - JWT_SECRET=mudar_para_um_segredo_forte_em_producao
      - FRONTEND_URL=http://localhost:3000
    volumes:
      - backend-data:/app/data

  frontend:
    build:
      context: ./codebase/frontend
      dockerfile: Dockerfile
    ports:
      - "3000:3000"
    environment:
      - PORT=3000
      - NEXT_PUBLIC_API_URL=http://localhost:8080
    depends_on:
      - backend

volumes:
  backend-data:
```
