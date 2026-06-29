# Tarefas da Sprint 8: Integração de Calendário (sprint8_tasks.md)

Este documento descreve detalhadamente as tarefas técnicas individuais necessárias para a execução da Sprint 8 pelo Programador.

---

## 1. Banco de Dados e Migração (Backend)
- [ ] **Tarefa 1.1**: Criar rotina de migração para a tabela `oauth_tokens`.
  - **Instruções**: No arquivo de inicialização do banco (`internal/database/database.go`), criar a tabela `oauth_tokens` executando as instruções DDL descritas em `business-rules-details.md`.
  - **Critério de Aceitação**: Tabela criada no banco de dados SQLite ao iniciar o backend.

---

## 2. Implementação do Fluxo OAuth2 (Backend)
- [ ] **Tarefa 2.1**: Implementar fluxo OAuth2 do Google.
  - **Instruções**: Utilizar o pacote oficial `golang.org/x/oauth2` e a API do Google Calendar. Criar rotas `/api/integrations/google/connect` (redireciona para Google com escopo `calendar.events`) e `/api/integrations/google/callback` (captura code, obtém tokens, salva ou atualiza na tabela `oauth_tokens`).
  - **Critério de Aceitação**: Fluxo OAuth funcional salvando os tokens associados ao psicólogo logado.
- [ ] **Tarefa 2.2**: Implementar fluxo OAuth2 da Microsoft.
  - **Instruções**: Utilizar o pacote `golang.org/x/oauth2` configurado com endpoints do Azure AD / Microsoft Graph. Criar rotas `/api/integrations/outlook/connect` (escopo `Calendars.ReadWrite`) e `/api/integrations/outlook/callback` (captura code e salva tokens).
  - **Critério de Aceitação**: Tokens do Azure AD persistidos com sucesso.
- [ ] **Tarefa 2.3**: Rotina de Renovação Automática de Tokens.
  - **Instruções**: Desenvolver função utilitária no repositório que verifica se o `access_token` do banco está próximo do vencimento (com base no campo `expiry`). Se sim, dispara requisição OAuth2 usando o `refresh_token` correspondente para obter e gravar um novo token válido.
  - **Critério de Aceitação**: Renovação de tokens funcional de forma transparente.

---

## 3. Sincronização e Ações de Agenda (Backend)
- [ ] **Tarefa 3.1**: Integração de Google Meet.
  - **Instruções**: Implementar função em Go que se conecta ao Google Calendar API usando o token ativo e cria um evento enviando `conferenceDataVersion=1` com a chave `hangoutsMeet`. Retornar a URL da reunião gerada.
  - **Critério de Aceitação**: Chamada gerando links do Google Meet com sucesso.
- [ ] **Tarefa 3.2**: Integração de Outlook Calendar (Graph API).
  - **Instruções**:
    - Implementar `GET /api/integrations/outlook/calendar` que lista os compromissos da agenda do Outlook.
    - Implementar alocador automático que dispara um POST para a API do Microsoft Graph `/me/events` ao criar uma sessão no sistema, salvando o `id` retornado na coluna `sessions.outlook_event_id`.
    - Implementar rotina que envia `DELETE /me/events/{event_id}` para o Graph se uma sessão for removida do banco.
  - **Critério de Aceitação**: Criação e remoção de compromissos no Outlook funcionando sincronizadas com as sessões locais.

---

## 4. Integração no Frontend
- [ ] **Tarefa 4.1**: Conectar Botões OAuth e Sincronização.
  - **Instruções**: Na página `/admin/settings` (Coluna Direita), fazer requisição à API para carregar o status de conexão das contas. Integrar os botões "Conectar Google" e "Conectar Outlook" para abrir o fluxo de autorização.
  - **Critério de Aceitação**: Interface exibindo status em tempo real e redirecionando corretamente.
- [ ] **Tarefa 4.2**: Renderizar Sincronização da Agenda.
  - **Instruções**: Na coluna de integração, adicionar um componente de lista ou mini-calendário que busca de `GET /api/integrations/outlook/calendar` para exibir os horários ocupados.
  - **Critério de Aceitação**: Visualização dos compromissos integrada na UI.
- [ ] **Tarefa 4.3**: Integrar "Gerar Link Google Meet" na Timeline de Sessões.
  - **Instruções**: Na Timeline do Prontuário, ao criar ou editar uma sessão, habilitar o botão "Gerar Link Google Meet". Ao ser clicado, envia requisição ao backend e atualiza a sessão exibindo o link resultante na tela.
  - **Critério de Aceitação**: Link de Meet associado e exibido após o clique na UI do paciente.
