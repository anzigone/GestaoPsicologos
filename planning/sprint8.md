# Sprint 8: Integração de Calendário (sprint8.md)

## 1. Escopo e Objetivos
O objetivo desta sprint é estabelecer a integração profunda com os ecossistemas externos da Google (Calendar/Meet) e Microsoft (Outlook Calendar/Graph API) através de fluxos OAuth2 de produção seguros. O psicólogo poderá gerenciar e sincronizar seus atendimentos diretamente nestas plataformas.

**Objetivos principais**:
- Executar a migração da tabela `oauth_tokens` no SQLite.
- Desenvolver os endpoints de redirecionamento e callbacks OAuth2 para Google e Microsoft no backend.
- Criar a rotina assíncrona de atualização de tokens usando `refresh_token` no banco de dados.
- Implementar wrappers de cliente HTTP para as chamadas de API do Microsoft Graph e Google Calendar (inserção de Meet e eventos de agenda).
- Integrar a alocação e remoção automática de compromissos no Outlook a partir das sessões da ficha do paciente.

---

## 2. Entregáveis da Sprint
Ao final desta sprint, os seguintes entregáveis deverão estar prontos:
- Tabela `oauth_tokens` persistindo credenciais no SQLite.
- Endpoints de conexão e callback OAuth2 implementados no backend.
- Wrappers de integração no backend com:
  - Google Calendar API: Criando eventos com dados de conferência (URL do Google Meet) e preenchendo a coluna `sessions.meet_link`.
  - Microsoft Graph API: Criando compromissos, listando eventos de agenda e removendo compromissos ao excluir sessões.
- Interface `/admin/settings` integrada, contendo botões OAuth operacionais e mini-calendário sincronizado.
- Botão "Gerar Link Google Meet" integrado na timeline de sessões do paciente.

---

## 3. Checkpoints de Validação Humana
Para validar a conclusão desta Sprint, o humano realizará os seguintes testes:

1. **Subir os serviços**:
   - Executar na raiz do projeto:
     ```bash
     docker compose up --build -d
     ```
2. **Conexão OAuth2**:
   - Acessar `/admin/settings`.
   - Clicar em "Conectar Conta Google".
   - Esperado: Redirecionamento seguro para a tela de permissão da Google, retorno com sucesso e status atualizado para "Conectado como e-mail@gmail.com".
   - Executar o mesmo procedimento com "Conectar Calendário Outlook" para a conta Microsoft.
3. **Agendamento com Criação de Reunião e Bloco de Agenda**:
   - Acessar `/patients` e selecionar um paciente de teste.
   - Ir na aba "Evolução das Sessões" e clicar em "Iniciar Nova Sessão".
   - Clicar em "Gerar Link Google Meet".
   - Esperado:
     - A coluna ou card de atendimento é atualizada exibindo o link gerado (ex: `https://meet.google.com/abc-defg-hij`).
     - Acessar a agenda pessoal do Outlook do psicólogo externamente.
     - Esperado: Um novo compromisso "Consulta Psicológica - [Nome do Paciente]" está alocado no horário definido.
4. **Remoção de Compromisso**:
   - Excluir a sessão recém-criada na timeline.
   - Acessar a agenda externa do Outlook do psicólogo.
   - Esperado: O compromisso correspondente foi automaticamente deletado do calendário.
5. **Encerramento da Sprint**:
   - Realizar commit com a mensagem `feat: sprint 8 integracao de calendario concluido`, criar a tag Git `v0.8.0` e fazer push para a branch `main`.
