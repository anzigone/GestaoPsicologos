# Registro de Execução da Sprint 5 (sprint5_logs.md)

Este documento registra cronologicamente todas as ações realizadas pelo Programador durante a Sprint 5, servindo de trilha de auditoria e documentação técnica da implantação.

---

## 1. Histórico de Atividades e Comandos

- **Data**: 29/06/2026
- **Status da Sprint**: ✅ Concluída e Validada

### Passos Executados:

1. Criado `codebase/backend/internal/handlers/psychologist.go` com os handlers `GetPsychologist` e `UpdatePsychologist` usando queries SQL diretas no SQLite. A query de seleção usa `COALESCE` para campos opcionais nulos. O isolamento por usuário é garantido pelo `userID` extraído do JWT via `mw.UserIDFromContext(r)`.

2. Adicionado endpoint `POST /api/auth/change-password` em `internal/handlers/auth.go` com tipo `ChangePasswordRequest`, validação do hash da senha atual via `auth.CheckPassword` e gravação do novo hash via `auth.HashPassword`.

3. Registradas as rotas no grupo protegido por JWT em `cmd/server/main.go`:
   ```
   r.Post("/api/auth/change-password", handlers.ChangePassword(db))
   r.Get("/api/psychologist", handlers.GetPsychologist(db))
   r.Put("/api/psychologist", handlers.UpdatePsychologist(db))
   ```

4. Reescrita a página `src/app/(app)/admin/settings/page.tsx` para integração real com a API:
   - `useEffect` carrega `GET /api/psychologist` ao montar a página
   - `handleSaveProfile` dispara `PUT /api/psychologist`
   - `handleChangePassword` dispara `POST /api/auth/change-password` com validação de confirmação e mínimo de 6 caracteres
   - Hook `useToast` com auto-dismiss em 3s para feedback em PT-BR

5. Build e restart via `docker compose up --build -d`.

---

## 2. Logs de Testes e Validação Local

### Comando executado para rodar a aplicação localmente:
```bash
docker compose up --build -d
```

### Resultados obtidos (validados via curl em 29/06/2026):

- **Leitura do perfil do psicólogo logado**: ✅ `GET /api/psychologist` retornou JSON com todos os campos do usuário autenticado (isolado por JWT sub)
- **Atualização cadastral persistida no SQLite**: ✅ `PUT /api/psychologist` com specialty="Terapia Cognitivo Comportamental", base_fee=180.00, crp="06/123456", phone="(11) 99999-1234" → dados confirmados em GET subsequente com `updated_at` renovado
- **Rejeição de senha incorreta**: ✅ `POST /api/auth/change-password` com senha errada → `{"error": "Senha atual incorreta"}` (HTTP 401)
- **Troca de senha validando a antiga**: ✅ `POST /api/auth/change-password` com senha correta → `{"message": "Senha alterada com sucesso"}`
- **Login utilizando as novas credenciais**: ✅ Login com nova senha retornou token JWT válido com dados atualizados; senha antiga rejeitada com `{"error": "Credenciais inválidas"}`

### Probes adicionais:
- `GET /api/psychologist` com token de usuário diferente (admin) → retorna dados do admin, não vaza dados de outro psicólogo ✅
- UI (`/admin/settings`): código implementado com carregamento, salvamento e troca de senha — extensão Chrome indisponível durante validação automatizada, testar manualmente se necessário.

---

## 3. Débitos Técnicos Identificados nesta Sprint
- **DT-S5-01**: UI não validada por automação de browser (extensão Chrome desconectada no momento da validação). Testar manualmente navegando em `/admin/settings` após login como psicólogo.
- **DT-S5-02**: `package_sessions` e `package_fee` são gravados pelo `PUT /api/psychologist` mas a UI da Sprint 5 não expõe esses campos (previstos para Sprint futura). Sem impacto funcional.
