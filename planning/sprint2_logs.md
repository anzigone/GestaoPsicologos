# Registro de Execução da Sprint 2 (sprint2_logs.md)

Este documento registra cronologicamente todas as ações realizadas pelo Programador durante a Sprint 2, servindo de trilha de auditoria e documentação técnica da implantação.

---

## 1. Histórico de Atividades e Comandos

- **Data**: 29/06/2026
- **Status da Sprint**: Concluída ✅

### Passos Executados:

1. **Tarefa 1.1** — Dependências do Swagger adicionadas ao `go.mod`:
   - `github.com/swaggo/swag v1.16.6`
   - `github.com/swaggo/http-swagger/v2 v2.0.2`
   - `github.com/swaggo/files/v2 v2.0.0`
   - Dependências transitivas (go-openapi/spec, mailru/easyjson, etc.) também adicionadas ao `go.mod`.

2. **Tarefa 1.2** — `main.go` anotado com metadados Swagger:
   - `@title`, `@version`, `@description`, `@host`, `@BasePath` no topo do arquivo.
   - Esquema de segurança `@securityDefinitions.apikey BearerAuth` definido.
   - Rota `GET /swagger/*` adicionada ao roteador Chi usando `httpSwagger.Handler`.

3. **Nota de arquitetura**: O diretório `docs/` (gerado pelo `swag init`) **não é commitado**. A geração ocorre durante o build Docker via:
   ```bash
   RUN go install github.com/swaggo/swag/cmd/swag@v1.16.6
   RUN swag init -g cmd/server/main.go --output docs
   ```
   Este comportamento é intencional para evitar artefatos gerados no repositório.

4. **Tarefa 2.1** — Middleware CORS implementado em `main.go`:
   - `corsMiddleware(cfg.FrontendURL)` permite origin `http://localhost:3000`.
   - Métodos permitidos: `GET, POST, PUT, DELETE, OPTIONS`.
   - Headers permitidos: `Content-Type, Authorization`.
   - Responde `204 No Content` para preflight `OPTIONS`.
   - Middleware de logging do Chi (`middleware.Logger`) e `middleware.Recoverer` adicionados.

5. **Tarefa 2.2** — Estrutura de rotas REST mapeada em `main.go`:
   - `/api/health` — Health check
   - `/api/auth/login`, `/api/auth/change-password` — Autenticação
   - `/api/admin/users` (GET, POST, DELETE) — Administração
   - `/api/patients` (CRUD completo + `/pdf`, `/analysis`, `/sessions`) — Pacientes
   - `/api/integrations/*` (Google e Outlook) — Integrações

6. **Tarefa 3.1** — `internal/handlers/auth.go` criado:
   - `POST /api/auth/login` → retorna token JWT mockado + `models.User` de "Dra. Ana Beatriz Santos"
   - `POST /api/auth/change-password` → retorna `{"message": "Senha alterada com sucesso"}`
   - Ambos anotados com Swagger (`@Summary`, `@Tags`, `@Param`, `@Success`, `@Failure`, `@Router`).

7. **Tarefa 3.2** — `internal/handlers/admin.go` criado:
   - `GET /api/admin/users` → retorna 2 psicólogos mockados (Dra. Ana + Dr. Carlos)
   - `POST /api/admin/users` → retorna `models.User` com UUID gerado
   - `DELETE /api/admin/users/{id}` → retorna HTTP 204
   - Tag Swagger: `Administração`.

8. **Tarefa 4.1** — `internal/handlers/patients.go` criado:
   - 3 pacientes mockados: Ana Souza, Carlos Lima, Roberta Silva
   - `GET /api/patients` → lista com filtro textual por `?q=`
   - `GET /api/patients/{id}` → detalhes do paciente
   - `POST /api/patients` → criação com UUID fixo
   - `PUT /api/patients/{id}` → atualização mockada
   - `DELETE /api/patients/{id}` → HTTP 204
   - `GET /api/patients/{id}/pdf` → PDF mínimo válido em bytes (PDF-1.4 estático)

9. **Tarefa 4.2** — `internal/handlers/analysis.go` criado:
   - `GET /api/patients/{id}/analysis` → retorna `models.FirstAnalysis` completa (TAG, TCC, etc.)
   - `PUT /api/patients/{id}/analysis` → aceita body e retorna análise atualizada

10. **Tarefa 5.1** — `internal/handlers/sessions.go` criado:
    - 3 sessões mockadas para Ana Souza (2 pagas + 1 pendente futura)
    - `GET /api/patients/{id}/sessions` → lista mockada
    - `POST /api/patients/{id}/sessions` → cria sessão com MeetLink e OutlookEventID mockados
    - `PUT /api/patients/{id}/sessions/{sid}` → atualização mockada
    - `DELETE /api/patients/{id}/sessions/{sid}` → HTTP 204

11. **Tarefa 5.2** — `internal/handlers/integrations.go` criado:
    - `GET /api/integrations/google/connect` → retorna URL de autorização OAuth2 mockada do Google
    - `GET /api/integrations/google/callback` → retorna `{"provider":"google","status":"conectado",...}`
    - `GET /api/integrations/outlook/connect` → retorna URL de autorização OAuth2 mockada da Microsoft
    - `GET /api/integrations/outlook/callback` → retorna `{"provider":"outlook","status":"conectado",...}`
    - `DELETE /api/integrations/{provider}/disconnect` → retorna mensagem de sucesso

12. **`internal/models/models.go`** criado com os structs base:
    - `User`, `Patient`, `FirstAnalysis`, `Session`, `ErrorResponse`, `MessageResponse`
    - Todas com tags Swagger `@Description` e `example:`.

---

## 2. Logs de Testes e Validação Local

### Comando executado para rodar o backend localmente:
```bash
docker compose up backend --build -d
```

### Checklist de Validação (executar após subir o container):

**Teste 1 — Swagger UI**
```
GET http://localhost:8080/swagger/index.html
```
- **Esperado**: Página HTML do Swagger UI carregando com todas as rotas listadas nas tags: Autenticação, Administração, Pacientes, Primeira Análise, Sessões, Integrações.
- **Resultado (29/06/2026)**: ✅ OK — HTTP 200, Content-Type: `text/html; charset=utf-8`, página Swagger UI confirmada.

**Teste 2 — Login Mock**
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email": "dr.ana@email.com", "password": "senha123"}'
```
- **Esperado**: HTTP 200 com `{"token": "eyJ...", "user": {"id": "550e8400...", "name": "Dra. Ana Beatriz Santos", ...}}`
- **Resultado (29/06/2026)**: ✅ OK — HTTP 200, token JWT retornado (`eyJhbGciOiJIUzI1NiIs...`), user `Dra. Ana Beatriz Santos` com role `psicologo`, CRP `06/123456`, `base_fee: 200`.

**Teste 3 — Listagem de Pacientes**
```bash
curl http://localhost:8080/api/patients
```
- **Esperado**: HTTP 200 com array de 3 pacientes: Ana Souza, Carlos Lima, Roberta Silva.
- **Resultado (29/06/2026)**: ✅ OK — HTTP 200, array com 3 pacientes: `Ana Souza` (id: 660e8400...), `Carlos Lima` (id: 661e8400...), `Roberta Silva` (id: 662e8400...).

**Teste 4 — Busca de Pacientes por Nome**
```bash
curl "http://localhost:8080/api/patients?q=Ana"
```
- **Esperado**: HTTP 200 com array contendo somente "Ana Souza".
- **Resultado (29/06/2026)**: ✅ OK — HTTP 200, array com 1 resultado: `Ana Souza`. Filtro case-insensitive funcional.

**Teste 5 — PDF do Prontuário**
```bash
curl http://localhost:8080/api/patients/660e8400-e29b-41d4-a716-446655440001/pdf \
  -o prontuario_mock.pdf
```
- **Esperado**: Arquivo `prontuario_mock.pdf` gerado com Content-Type `application/pdf`.
- **Resultado (29/06/2026)**: ✅ OK — HTTP 200, Content-Type: `application/pdf`, 298 bytes, header PDF válido (`%PDF` confirmado nos magic bytes).

**Teste 6 — Listagem de Sessões**
```bash
curl http://localhost:8080/api/patients/660e8400-e29b-41d4-a716-446655440001/sessions
```
- **Esperado**: HTTP 200 com 3 sessões mockadas (2 pagas + 1 pendente).
- **Resultado (29/06/2026)**: ✅ OK — HTTP 200, 3 sessões: `20/06/2026 pago`, `06/06/2026 pago`, `10/07/2026 pendente` (futura). MeetLinks Google Meet presentes em todas.

**Teste 7 — Listagem de Usuários Admin**
```bash
curl http://localhost:8080/api/admin/users
```
- **Esperado**: HTTP 200 com 2 psicólogos: Dra. Ana Beatriz Santos + Dr. Carlos Eduardo Lima.
- **Resultado (29/06/2026)**: ✅ OK — HTTP 200, 2 usuários: `Dra. Ana Beatriz Santos` e `Dr. Carlos Eduardo Lima`.

**Teste 8 — Integração Google Connect**
```bash
curl http://localhost:8080/api/integrations/google/connect
```
- **Esperado**: HTTP 200 com `{"redirect_url": "https://accounts.google.com/..."}`.
- **Resultado (29/06/2026)**: ✅ OK — HTTP 200, `redirect_url` retornado apontando para `accounts.google.com/o/oauth2/auth` com parâmetros `client_id`, `redirect_uri`, `scope=calendar` e `access_type=offline`.

---

## 3. Débitos Técnicos Identificados nesta Sprint

- **`docs/` não commitado**: O pacote Swagger gerado por `swag init` é produzido durante o Docker build. Para desenvolvimento local fora do Docker, executar `go install github.com/swaggo/swag/cmd/swag@v1.16.6 && swag init -g cmd/server/main.go --output docs` na pasta `codebase/backend/`.
- **Sem autenticação JWT real**: Todos os endpoints estão sem middleware de validação de JWT. Isso é intencional até a Sprint 4 (implementação real de auth).
- **Dados mockados em memória**: Todas as respostas são dados estáticos hardcoded. A camada de repositório real será implementada nas Sprints 4-9.
