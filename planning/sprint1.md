# Sprint 1: Setup Inicial do Projeto (sprint1.md)

## 1. Escopo e Objetivos
O objetivo desta sprint é estabelecer toda a fundação e infraestrutura local do projeto, permitindo que os desenvolvedores tenham um ambiente de desenvolvimento unificado e funcional.

**Objetivos principais**:
- Definir a estrutura de diretórios monorepo.
- Inicializar os esqueletos do Frontend (Next.js) e Backend (Go).
- Dockerizar os componentes individualmente para desenvolvimento e produção.
- Criar a orquestração de desenvolvimento local usando Docker Compose.
- Desenhar o arquivo base de CI/CD do GitHub Actions.

---

## 2. Entregáveis da Sprint
Ao final desta sprint, os seguintes entregáveis deverão estar prontos no repositório:
- `./codebase/backend/`: Diretório contendo a aplicação Go inicializada, com rotas básicas (`chi` router) e suporte a drivers de banco de dados SQLite e MySQL.
- `./codebase/frontend/`: Diretório contendo o aplicativo Next.js 14+ com TypeScript, TailwindCSS e base do Shadcn/UI configurada.
- `./codebase/backend/Dockerfile` e `./codebase/frontend/Dockerfile`: Arquivos de dockerização multi-stage otimizados para produção (Alpine).
- `./docker-compose.yml`: Orquestração para build e execução local do ecossistema.
- `.github/workflows/deploy.yml`: O script do pipeline GitHub Actions inicializado.

---

## 3. Checkpoints de Validação Humana
Para validar a conclusão desta Sprint e autorizar a passagem para a próxima, o humano realizará os seguintes testes:

1. **Clonar e Executar**:
   - Clonar o repositório e na raiz do projeto executar:
     ```bash
     docker compose up --build -d
     ```
2. **Validação do Frontend**:
   - Abrir o navegador e acessar `http://localhost:3000`.
   - Esperado: Tela inicial do Next.js padrão ou tela de boas-vindas do sistema com estilização Tailwind visível.
3. **Validação do Backend**:
   - Acessar no navegador ou via curl: `http://localhost:8080/api/health`.
   - Esperado: Resposta JSON `{"status":"OK", "database":"connected"}`.
4. **Validação de Banco de Dados**:
   - Verificar a criação automática do arquivo SQLite em `./codebase/backend/data/gestao.db` (ou dentro do volume Docker correspondente).
5. **Encerramento da Sprint**:
   - Realizar commit com a mensagem `feat: sprint 1 setup inicial concluido`, criar a tag Git `v0.1.0` e fazer push para a branch `main` do GitHub.
