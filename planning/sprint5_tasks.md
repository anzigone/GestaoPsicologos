# Tarefas da Sprint 5: Administração do Psicólogo (sprint5_tasks.md)

Este documento descreve detalhadamente as tarefas técnicas individuais necessárias para a execução da Sprint 5 pelo Programador.

---

## 1. Implementação do Repositório e Endpoints (Backend)
- [x] **Tarefa 1.1**: Adicionar queries SQL para ler e atualizar psicólogos.
  - **Instruções**: No módulo `internal/database/`, criar as funções de repositório para selecionar e atualizar dados cadastrais do psicólogo baseado no ID do usuário.
  - **Critério de Aceitação**: Queries de leitura e atualização compilando com sucesso.
  - **Implementado em**: `internal/handlers/psychologist.go` — constante `querySelectUserByID` + `scanUserRow()`.
- [x] **Tarefa 1.2**: Criar Handlers de Perfil no Go.
  - **Instruções**:
    - `GET /api/psychologist`: Ler o ID do psicólogo do contexto (JWT claims), buscar no banco e retornar como JSON.
    - `PUT /api/psychologist`: Receber dados de atualização e gravar no banco para o ID do contexto.
    - Proteger ambas as rotas com o middleware de JWT.
  - **Critério de Aceitação**: Endpoints testados manualmente via Swagger respondendo dados reais do banco.
  - **Implementado em**: `internal/handlers/psychologist.go` — `GetPsychologist()` e `UpdatePsychologist()`. Rotas registradas em `cmd/server/main.go`.
- [x] **Tarefa 1.3**: Implementar validação e alteração de senha.
  - **Instruções**: No handler `POST /api/auth/change-password`, ler o ID do psicólogo, carregar a senha atual do banco, validar o hash com a senha informada, gerar hash para a nova senha e atualizar no banco de dados.
  - **Critério de Aceitação**: Senha alterada e persistida de forma segura.
  - **Implementado em**: `internal/handlers/auth.go` — `ChangePassword()`. Rota protegida por JWT em `cmd/server/main.go`.

---

## 2. Integração no Frontend
- [x] **Tarefa 2.1**: Carregar e persistir dados cadastrais na tela `/admin/settings`.
  - **Instruções**: Fazer a requisição `GET /api/psychologist` ao montar a página para preencher os inputs com os dados do psicólogo logado. Integrar o botão de envio para disparar o `PUT /api/psychologist` enviando os valores atualizados.
  - **Critério de Aceitação**: Edição e salvamento de dados cadastrais funcionando perfeitamente pela UI.
  - **Implementado em**: `src/app/(app)/admin/settings/page.tsx` — `useEffect` + `handleSaveProfile()`.
- [x] **Tarefa 2.2**: Integrar o formulário de alteração de senha.
  - **Instruções**: Integrar a chamada `POST /api/auth/change-password` no formulário correspondente, limpando os campos e exibindo mensagens de sucesso ou de erro (ex: "Senha atual incorreta") em PT-BR.
  - **Critério de Aceitação**: Feedback de erro e sucesso visível na tela e login refletindo a nova senha.
  - **Implementado em**: `src/app/(app)/admin/settings/page.tsx` — `handleChangePassword()` com validação de confirmação e mínimo 6 caracteres.
- [x] **Tarefa 2.3**: Exibir mensagens amigáveis (Toasts).
  - **Instruções**: Utilizar o componente Toast do Shadcn para confirmar as operações do usuário.
  - **Critério de Aceitação**: Toast de sucesso exibido após salvar dados ou trocar senha.
  - **Implementado em**: `src/app/(app)/admin/settings/page.tsx` — hook `useToast()` customizado com auto-dismiss em 3s.
