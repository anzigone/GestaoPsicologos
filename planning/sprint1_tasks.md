# Tarefas da Sprint 1: Setup Inicial (sprint1_tasks.md)

Este documento descreve detalhadamente as tarefas técnicas individuais necessárias para a execução da Sprint 1 pelo Programador. Cada tarefa deve ser marcada como concluída apenas após a validação de seus respectivos critérios de aceitação.

---

## 1. Setup do Monorepo e Pastas
- [ ] **Tarefa 1.1**: Criar a estrutura base de pastas no repositório local.
  - **Instruções**: Criar a pasta `./codebase` e, dentro dela, criar as pastas `./codebase/backend` e `./codebase/frontend`.
  - **Critério de Aceitação**: Diretórios criados.

---

## 2. Setup Inicial do Backend (Golang)
- [ ] **Tarefa 2.1**: Inicializar o módulo Go.
  - **Instruções**: Dentro de `./codebase/backend`, executar `go mod init github.com/anzigone/GestaoPsicologos/backend` (ou o nome do repositório correspondente).
  - **Critério de Aceitação**: Arquivo `go.mod` criado.
- [ ] **Tarefa 2.2**: Instalar dependências básicas do Go.
  - **Instruções**: Rodar `go get github.com/go-chi/chi/v5`, `go get modernc.org/sqlite` e `go get github.com/go-sql-driver/mysql` para drivers de banco, e `go get github.com/joho/godotenv` para gerenciar `.env` local.
  - **Critério de Aceitação**: Dependências listadas no `go.mod` e `go.sum`.
- [ ] **Tarefa 2.3**: Criar estrutura interna do Backend.
  - **Instruções**: Criar os diretórios para a arquitetura do backend:
    - `./codebase/backend/cmd/server/main.go` (Ponto de entrada)
    - `./codebase/backend/internal/config/` (Leitura de variáveis)
    - `./codebase/backend/internal/database/` (Acesso a dados)
    - `./codebase/backend/internal/logger/` (Mecanismo de Logs)
    - `./codebase/backend/internal/handlers/` (Rotas e Handlers HTTP)
  - **Critério de Aceitação**: Estrutura física de pastas criada.
- [ ] **Tarefa 2.4**: Implementar conexão básica ao banco SQLite (DEV).
  - **Instruções**: Escrever o código de inicialização do banco lendo a variável `DB_DRIVER=sqlite` e `DATABASE_PATH`. Criar a pasta `./codebase/backend/data/` para armazenar o banco em desenvolvimento local.
  - **Critério de Aceitação**: Conexão estabelecida e arquivo `.db` gerado com sucesso durante a inicialização do app.
- [ ] **Tarefa 2.5**: Configurar o Logger e Roteador HTTP Chi.
  - **Instruções**: Setup de logging básico com níveis `INFO`, `WARN` e `ERROR` saindo para `Stdout`. Criar um roteador Chi configurando uma rota `GET /api/health` que responde JSON com o status da API e status da conexão com o banco de dados.
  - **Critério de Aceitação**: Endpoint respondendo `200 OK` com JSON válido.

---

## 3. Setup Inicial do Frontend (Next.js)
- [ ] **Tarefa 3.1**: Inicializar o projeto Next.js.
  - **Instruções**: Dentro de `./codebase/frontend`, criar a estrutura Next.js 14+ usando TypeScript, ESLint e TailwindCSS.
  - **Critério de Aceitação**: Presença de `package.json`, `tailwind.config.js`, `tsconfig.json` e o diretório `src/app/` com a página padrão.
- [ ] **Tarefa 3.2**: Configurar Shadcn/UI.
  - **Instruções**: Configurar os componentes primitivos do Shadcn/UI (botões, inputs e dialogs) na pasta de componentes.
  - **Critério de Aceitação**: Arquivo de configuração de componentes Shadcn pronto e diretório `src/components/ui/` estruturado.
- [ ] **Tarefa 3.3**: Configurar Output Standalone do Next.js.
  - **Instruções**: Ajustar o `next.config.js` para incluir a diretiva `output: 'standalone'` para garantir que o build Docker final seja enxuto.
  - **Critério de Aceitação**: `next.config.js` atualizado.

---

## 4. Dockerização e Orquestração Local
- [ ] **Tarefa 4.1**: Criar o Dockerfile do Backend.
  - **Instruções**: Criar o arquivo `./codebase/backend/Dockerfile` utilizando um build multi-stage baseado em `golang:1.22-alpine` e runtime `alpine:3.19`. Garantir a compilação com `CGO_ENABLED=0`.
  - **Critério de Aceitação**: Build da imagem Docker do backend executado com sucesso localmente.
- [ ] **Tarefa 4.2**: Criar o Dockerfile do Frontend.
  - **Instruções**: Criar o arquivo `./codebase/frontend/Dockerfile` utilizando build multi-stage baseado em `node:20-alpine`, gerando o build com suporte a `standalone` do Next.js.
  - **Critério de Aceitação**: Build da imagem Docker do frontend executado com sucesso localmente.
- [ ] **Tarefa 4.3**: Criar o arquivo docker-compose.yml.
  - **Instruções**: Criar `./docker-compose.yml` na raiz do monorepo, definindo os serviços `backend` e `frontend`, configurando as portas `8080:8080` e `3000:3000` respectivamente, passando as variáveis de ambiente necessárias e montando volumes locais para persistência do SQLite do backend.
  - **Critério de Aceitação**: Executar `docker compose up --build -d` inicializa ambos os containers de forma estável.

---

## 5. Pipeline de CI/CD Inicial
- [ ] **Tarefa 5.1**: Criar o workflow do GitHub Actions.
  - **Instruções**: Criar a pasta `.github/workflows/` na raiz do projeto e adicionar o arquivo `deploy.yml` contendo o workflow básico para build, empacotamento e deploy (conforme infra-devops-details.md).
  - **Critério de Aceitação**: Arquivo YAML sintaticamente válido.
