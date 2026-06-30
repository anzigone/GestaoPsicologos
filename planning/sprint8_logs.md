# Registro de Execução da Sprint 8 (sprint8_logs.md)

Este documento registra cronologicamente todas as ações realizadas pelo Programador durante a Sprint 8, servindo de trilha de auditoria e documentação técnica da implantação.

---

## 1. Histórico de Atividades e Comandos

- **Data**: 30/06/2026
- **Status da Sprint**: ⏸️ Adiada — será retomada ao final do projeto

### Decisão de Adiamento:

A integração OAuth2 com Google Calendar/Meet e Microsoft Outlook/Graph API requer:
1. **Registro de aplicativo no Google Cloud Console** com credenciais OAuth2 (`client_id`, `client_secret`) e URI de redirecionamento homologado.
2. **Registro de aplicativo no Azure AD (Microsoft Entra)** com as mesmas exigências para o escopo `Calendars.ReadWrite`.

Ambas as plataformas exigem que o domínio de callback esteja publicamente acessível e verificado, o que não é viável no ambiente de desenvolvimento atual (localhost/Docker).

**Decisão**: Sprint 8 adiada para reavaliação ao final do projeto (após Sprint 13). Prosseguir diretamente para a Sprint 9.

---

## 2. Logs de Testes e Validação Local

### Resultados obtidos:
- **Criação da tabela oauth_tokens no SQLite**: ⏸️ Não executada (sprint adiada)
- **Redirecionamento e Callback OAuth2 (Google/Outlook)**: ⏸️ Não executada (sprint adiada)
- **Geração e gravação do link de Google Meet na sessão**: ⏸️ Não executada (sprint adiada)
- **Sincronização de criação/exclusão de eventos no Outlook**: ⏸️ Não executada (sprint adiada)

---

## 3. Débitos Técnicos Identificados nesta Sprint

- Toda a Sprint 8 constitui um débito técnico planejado a ser retomado após a entrega das funcionalidades core (Sprints 9–13).
- Os campos `sessions.meet_link` e `sessions.outlook_event_id` já existem na tabela `sessions` (criados na Sprint 7) e estarão prontos para uso quando a integração for implementada.
