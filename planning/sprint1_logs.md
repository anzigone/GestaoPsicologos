# Registro de Execução da Sprint 1 (sprint1_logs.md)

---

## 1. Histórico de Atividades e Comandos

- **Data**: 27/06/2026
- **Status da Sprint**: Concluída ✅

### Passos Executados:

1. Leitura e análise das especificações (`archicteture-details.md`, `infra-devops-details.md`, `sprint1.md`, `sprint1_tasks.md`).

2. **Tarefa 1.1** — Criação da estrutura de diretórios do monorepo:
   ```bash
   mkdir -p codebase/backend/cmd/server
   mkdir -p codebase/backend/internal/{config,database,logger,handlers}
   mkdir -p codebase/backend/data
   mkdir -p codebase/frontend
   mkdir -p .github/workflows
   ```

3. **Tarefa 2.1** — `go.mod` criado com módulo `github.com/anzigone/GestaoPsicologos/backend` (Go 1.22).

4. **Tarefa 2.2** — Dependências declaradas no `go.mod`:
   - `github.com/go-chi/chi/v5 v5.1.0`
   - `github.com/go-sql-driver/mysql v1.8.1`
   - `github.com/joho/godotenv v1.5.1`
   - `modernc.org/sqlite v1.29.9`
   - Nota: `go mod tidy` é executado no Dockerfile para gerar `go.sum` automaticamente.

5. **Tarefa 2.3** — Estrutura interna do backend criada:
   - `internal/config/config.go` — leitura de variáveis de ambiente
   - `internal/database/database.go` — conexão híbrida SQLite/MySQL + migrações
   - `internal/logger/logger.go` — logger estruturado com `log/slog`
   - `internal/handlers/health.go` — handler do endpoint `/api/health`
   - `cmd/server/main.go` — ponto de entrada com chi router + CORS middleware

6. **Tarefa 2.4** — Conexão SQLite implementada em `database.go` com criação automática do diretório de dados.

7. **Tarefa 2.5** — Logger INFO/WARN/ERROR via `log/slog`, roteador chi com rota `GET /api/health`.

8. **Nota de decisão técnica**: `ON UPDATE CURRENT_TIMESTAMP` removido do schema SQL por incompatibilidade com SQLite. O campo `updated_at` será atualizado explicitamente nas camadas de repositório das sprints futuras.

9. **Tarefa 3.1** — Frontend scaffolding via:
   ```bash
   npx create-next-app@14 . --typescript --tailwind --eslint --app --src-dir --import-alias "@/*" --yes
   # Resultado: Next.js 14.2.35 instalado com sucesso
   ```

10. **Tarefa 3.2** — Shadcn/UI inicializado e componentes base adicionados:
    ```bash
    npx shadcn@latest init --yes --defaults
    npx shadcn@latest add input dialog --yes
    # Componentes criados: button.tsx, input.tsx, dialog.tsx em src/components/ui/
    # Utilitário criado: src/lib/utils.ts
    ```
    Diretórios adicionais criados: `src/hooks/`, `src/types/`

11. **Tarefa 3.3** — `next.config.mjs` atualizado com `output: "standalone"`.

12. **Tarefa 4.1** — `codebase/backend/Dockerfile` criado (multi-stage: golang:1.22-alpine → alpine:3.19).

13. **Tarefa 4.2** — `codebase/frontend/Dockerfile` criado (multi-stage: node:20-alpine builder → runner standalone).

14. **Tarefa 4.3** — `docker-compose.yml` criado na raiz do monorepo orquestrando `backend` (8080) e `frontend` (3000) com volume persistente para SQLite.

15. **Tarefa 5.1** — `.github/workflows/deploy.yml` criado com pipeline completo: OIDC AWS auth → ECR push → ECS deploy (backend e frontend).

16. `.gitignore` criado cobrindo `.env`, `data/`, `node_modules/`, `.next/`, binários Go.

---

## 2. Logs de Testes e Validação Local

### Comando para rodar a aplicação localmente:
```bash
docker compose up --build -d
```

### Resultados esperados para validação humana:
- **Backend Health Check** (`http://localhost:8080/api/health`):
  ```json
  {"status":"OK","database":"connected"}
  ```
- **Frontend App** (`http://localhost:3000`): Tela padrão Next.js com estilização Tailwind visível.
- **SQLite Database**: Arquivo `gestao.db` criado no volume Docker `backend-data`.

---

## 3. Débitos Técnicos Identificados nesta Sprint

- `go.sum` não está commitado — gerado automaticamente pelo `go mod tidy` durante o build Docker. Recomendado gerar localmente após instalar Go e commitar em sprints futuras.
- `ON UPDATE CURRENT_TIMESTAMP` não usado no schema SQLite. Será tratado na camada de repositório nas sprints de implementação (4+).
