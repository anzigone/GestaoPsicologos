# Registro de Execução da Sprint 4 (sprint4_logs.md)

Este documento registra cronologicamente todas as ações realizadas pelo Programador durante a Sprint 4, servindo de trilha de auditoria e documentação técnica da implantação.

---

## 1. Histórico de Atividades e Comandos

- **Data**: 29/06/2026
- **Status da Sprint**: Concluída ✅

### Passos Executados:

1. Criada migração da tabela `users` em `internal/database/database.go` (função `Migrate`).
2. Implementado seed do Admin Master em `database.go` (função `Seed`) — insere `admin@admin.com.br` / senha `admin` se tabela estiver vazia.
3. Implementadas funções `HashPassword` e `CheckPassword` em `internal/auth/auth.go` usando SHA256+salt.
4. Implementada função `GenerateToken` com `golang-jwt/jwt` (24h de validade) em `auth.go`.
5. Criado middleware `JWTRequired` e `AdminOnly` em `internal/middleware/jwt.go`.
6. Implementado handler `Login` (POST /api/auth/login) em `internal/handlers/auth.go`.
7. Implementado handler `ChangePassword` (POST /api/auth/change-password) em `handlers/auth.go`.
8. Implementados handlers `ListUsers`, `CreateUser`, `DeleteUser` em `internal/handlers/admin.go`.
9. Configuradas rotas no `cmd/server/main.go` com grupos protegidos por JWT e AdminOnly.
10. Integrado formulário de login no frontend (`src/app/login/page.tsx`) — POST para `/api/auth/login`, salva cookie `auth_token`.
11. Criado `src/middleware.ts` no Next.js para proteger `/dashboard`, `/patients`, `/admin`.
12. Integrada tela `/admin/users` com CRUD real via `api.ts`.

### Correções aplicadas (29/06/2026):
- **Bug 1**: E-mail do admin alterado de `"admin"` para `"admin@admin.com.br"` em `database.go` — campo do formulário é `type="email"`, exigindo formato válido.
- **Bug 2**: Método `api.delete` em `lib/api.ts` corrigido para roteá-lo por `request()`, garantindo envio do header `Authorization: Bearer`.

---

## 2. Logs de Testes e Validação Local

### Comando para rodar a aplicação:
```bash
docker compose up --build
```

### Critérios de aceitação a validar:
- **Redirecionamento automático sem login**: Acessar `/dashboard` sem cookie → redireciona para `/login`
- **Login com admin**: `admin@admin.com.br` / `admin` → redireciona para dashboard
- **Inclusão de psicólogo via tela**: Modal em `/admin/users` → novo registro persiste no SQLite
- **Exclusão de psicólogo via tela**: Botão de lixeira → remove do banco e da tabela
- **Login com novo psicólogo**: Credenciais criadas pelo admin → acessa o sistema

---

## 3. Débitos Técnicos Identificados nesta Sprint

- **DT-S4-01 — Rota raiz `/` sem proteção de autenticação**
  - A página home (`/`) retorna HTTP 200 para usuários não autenticados. O middleware Next.js (`src/middleware.ts`) cobre apenas `/dashboard`, `/patients` e `/admin`. Quando a home passar a exibir dados reais do usuário (Sprints futuras), ela deverá ser adicionada ao matcher de rotas protegidas.
  - **Impacto atual**: Baixo — a página ainda exibe dados mockados. Nenhuma informação sensível é exposta.
  - **Ação recomendada**: Adicionar `/` (ou criar rota `/welcome`) ao `PROTECTED_PATHS` e ao `matcher` no momento em que a home consumir dados do JWT.

- **DT-S4-02 — Botão "Sair" sem implementação**
  - O botão `Sair` em `src/app/page.tsx` (linha 36) não possui handler `onClick`. Clicar nele não limpa o cookie `auth_token` nem redireciona para `/login`, impossibilitando o logout do sistema.
  - **Impacto atual**: Alto — qualquer usuário que queira sair precisa apagar o cookie manualmente ou aguardar expiração do JWT (24h).
  - **Ação recomendada**: Implementar função de logout que execute `document.cookie = 'auth_token=; max-age=0; path=/'` seguido de `router.push('/login')`. Prioridade alta para a Sprint 5.

- **DT-S4-03 — Dados do usuário hardcoded na home**
  - `src/app/page.tsx` exibe "Dra. Ana Beatriz Santos / CRP 06/123456" de forma estática, sem ler as informações do JWT ou de uma chamada à API.
  - **Impacto atual**: Baixo — comportamento esperado nesta sprint (dados reais entram em sprint posterior).
  - **Ação recomendada**: Decodificar o JWT do cookie ou chamar `GET /api/auth/me` (a implementar) para preencher nome e CRP dinamicamente.

- **DT-S4-04 — Diretório `codebase/backend;C` com nome inválido**
  - Existe um diretório com nome `backend;C` em `codebase/`, provavelmente criado por um comando mal digitado. Não afeta o build Docker, mas polui o repositório.
  - **Impacto atual**: Nenhum funcional.
  - **Ação recomendada**: Verificar conteúdo e remover se estiver vazio ou contiver apenas arquivos duplicados.
