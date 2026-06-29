# Tarefas da Sprint 2: Mock Webservice (sprint2_tasks.md)

Este documento descreve detalhadamente as tarefas técnicas individuais necessárias para a execução da Sprint 2 pelo Programador.

---

## 1. Configuração do Swagger / OpenAPI
- [x] **Tarefa 1.1**: Adicionar dependências do Swagger ao projeto Go.
  - **Instruções**: Instalar as bibliotecas `swaggo/swag`, `swaggo/http-swagger` (para servir o UI no Chi router) e rodar o utilitário de geração de código do Swagger.
  - **Critério de Aceitação**: Dependências adicionadas ao `go.mod`.
- [x] **Tarefa 1.2**: Escrever anotações básicas de metadados no `main.go`.
  - **Instruções**: Adicionar comentários de documentação no ponto de entrada definindo título da API, versão, descrição, host base e esquemas de segurança JWT.
  - **Critério de Aceitação**: Arquivo `main.go` anotado.

---

## 2. Roteador, Middlewares e CORS
- [x] **Tarefa 2.1**: Implementar Middleware de CORS e Logging.
  - **Instruções**: No roteador Chi, adicionar o middleware nativo de CORS permitindo conexões da URL do frontend (`http://localhost:3000`), suportando métodos `GET, POST, PUT, DELETE, OPTIONS` e cabeçalhos `Authorization, Content-Type`.
  - **Critério de Aceitação**: Chamadas preflight (OPTIONS) retornando cabeçalhos CORS corretos.
- [x] **Tarefa 2.2**: Criar estrutura de Rotas REST.
  - **Instruções**: Mapear todas as rotas da API em sub-rotas estruturadas (ex: `/api/auth/*`, `/api/admin/*`, `/api/patients/*`).
  - **Critério de Aceitação**: Roteador compilando sem erros.

---

## 3. Endpoints de Autenticação e Usuários Admin (Mock)
- [x] **Tarefa 3.1**: Implementar Handlers Mock de Login e Senha.
  - **Instruções**:
    - `POST /api/auth/login`: Validar entrada e retornar JSON com token JWT e dados do usuário com perfil Admin ou Psicólogo.
    - `POST /api/auth/change-password`: Retornar mensagem de sucesso.
    - Documentar ambos no Swagger com os schemas de Request e Response correspondentes.
  - **Critério de Aceitação**: Endpoints respondendo adequadamente e listados no Swagger.
- [x] **Tarefa 3.2**: Implementar Handlers Mock de Administração Master de Usuários.
  - **Instruções**:
    - `GET /api/admin/users`: Retornar lista de psicólogos cadastrados.
    - `POST /api/admin/users`: Retornar o psicólogo recém-criado com ID (UUID) gerado.
    - `DELETE /api/admin/users/{id}`: Retornar HTTP 204.
    - Documentar todos no Swagger sob a tag `Administração`.
  - **Critério de Aceitação**: Teste manual das rotas via Swagger funcionando.

---

## 4. Endpoints de Pacientes e Anamnese (Mock)
- [x] **Tarefa 4.1**: Implementar Handlers Mock de Gestão de Pacientes.
  - **Instruções**:
    - `GET /api/patients`: Retornar array JSON mockado com 3 pacientes (ex: Ana Souza, Carlos Lima, Roberta Silva), simulando busca textual pelo nome se fornecido parâmetro de query `q`.
    - `GET /api/patients/{id}`: Retornar os detalhes completos de um paciente.
    - `POST /api/patients`: Simular criação retornando os dados passados com UUID gerado.
    - `PUT /api/patients/{id}`: Simular atualização.
    - `DELETE /api/patients/{id}`: Retornar HTTP 204.
    - `GET /api/patients/{id}/pdf`: Retornar um arquivo PDF mockado (cabeçalho de PDF estático ou string simulada).
  - **Critério de Aceitação**: Todos os métodos respondendo e documentados no Swagger.
- [x] **Tarefa 4.2**: Implementar Handlers Mock da Primeira Análise (Anamnese).
  - **Instruções**:
    - `GET /api/patients/{id}/analysis`: Retornar a ficha de primeira análise preenchida para o paciente.
    - `PUT /api/patients/{id}/analysis`: Receber os dados clínicos e retornar sucesso.
  - **Critério de Aceitação**: Requisições de anamnese mockadas funcionando.

---

## 5. Endpoints de Sessões e Integrações (Mock)
- [x] **Tarefa 5.1**: Implementar Handlers Mock de Sessões / Prontuário.
  - **Instruções**:
    - `GET /api/patients/{id}/sessions`: Retornar lista cronológica de sessões associadas ao paciente.
    - `POST /api/patients/{id}/sessions`: Simular agendamento de nova sessão.
    - `PUT /api/patients/{id}/sessions/{sid}`: Simular atualização das notas e status da sessão.
    - `DELETE /api/patients/{id}/sessions/{sid}`: Retornar HTTP 204.
  - **Critério de Aceitação**: Rotas de evolução respondendo e documentadas.
- [x] **Tarefa 5.2**: Implementar Handlers Mock de Integração de Agendas.
  - **Instruções**:
    - `GET /api/integrations/google/connect` e `GET /api/integrations/outlook/connect`: Simular o redirecionamento OAuth2.
    - `GET /api/integrations/google/callback` e `GET /api/integrations/outlook/callback`: Retornar tela ou JSON de sucesso na autenticação.
  - **Critério de Aceitação**: Fluxo simulado listado no Swagger.
