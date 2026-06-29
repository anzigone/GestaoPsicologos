# Tarefas da Sprint 4: Autenticação e Gestão (sprint4_tasks.md)

Este documento descreve detalhadamente as tarefas técnicas individuais necessárias para a execução da Sprint 4 pelo Programador.

---

## 1. Persistência de Dados e Banco de Dados (Backend)
- [x] **Tarefa 1.1**: Criar rotina de migração da tabela `users`.
  - **Instruções**: No arquivo de inicialização do banco (`internal/database/database.go`), criar a tabela `users` executando a DDL descrita em `business-rules-details.md`.
  - **Critério de Aceitação**: Tabela criada no SQLite local ao iniciar a API.
- [x] **Tarefa 1.2**: Implementar semente do Admin Master.
  - **Instruções**: Ao iniciar o banco de dados, verificar se a tabela `users` está vazia. Se sim, inserir o usuário admin master com e-mail `admin@admin.com.br`, role `admin` e o hash correspondente à senha `admin`.
  - **Critério de Aceitação**: Registro inserido no banco com sucesso.

---

## 2. Autenticação e Hash de Senhas (Backend)
- [x] **Tarefa 2.1**: Implementar utilitário de encriptação e validação de senhas.
  - **Instruções**: Criar funções utilitárias em Go para encriptar senhas (usando Bcrypt ou SHA256 com salt) e para comparar senhas fornecidas com o hash persistido.
  - **Critério de Aceitação**: Teste de hash de senhas executando com sucesso.
- [x] **Tarefa 2.2**: Implementar geração de tokens JWT.
  - **Instruções**: Utilizar a biblioteca `golang-jwt/jwt` para gerar tokens assinados digitalmente contendo o ID do usuário e seu perfil (role), com validade de 24 horas.
  - **Critério de Aceitação**: Endpoint de login retornando o token assinado.
- [x] **Tarefa 2.3**: Criar Middleware HTTP de validação de JWT.
  - **Instruções**: Criar middleware para interceptar as requisições protegidas, ler o cabeçalho `Authorization: Bearer <token>`, validar a assinatura e extrair as claims (ID e papel) inserindo-as no contexto da requisição (`r.Context()`).
  - **Critério de Aceitação**: Chamadas sem token ou com token inválido em rotas protegidas retornando status `401 Unauthorized`.

---

## 3. Endpoints de Usuários e Controle Admin (Backend)
- [x] **Tarefa 3.1**: Implementar endpoint de Login.
  - **Instruções**: No handler de autenticação, implementar a validação de credenciais de login em `POST /api/auth/login`.
  - **Critério de Aceitação**: Sucesso no login retorna o token e dados do usuário; falha retorna erro `401`.
- [x] **Tarefa 3.2**: Implementar CRUD de Psicólogos no Backend.
  - **Instruções**: Implementar os endpoints `GET /api/admin/users`, `POST /api/admin/users` (encriptando a senha informada) e `DELETE /api/admin/users/{id}`. Proteger estas rotas exigindo que o usuário autenticado possua o perfil `admin` (Master).
  - **Critério de Aceitação**: Apenas administradores conseguem listar, criar ou remover psicólogos.

---

## 4. Proteção de Rotas e Integração de Telas (Frontend)
- [x] **Tarefa 4.1**: Configurar Login no Frontend.
  - **Instruções**: Integrar o formulário de login do Next.js para realizar o POST em `/api/auth/login`. Ao receber o token com sucesso, salvar o token em cookies seguros (para leitura pelo Middleware) ou em local storage/estado da app e redirecionar para a página `/welcome`.
  - **Critério de Aceitação**: Login realizado com sucesso navegando para o dashboard.
- [x] **Tarefa 4.2**: Implementar Middleware de Rotas Privadas no Next.js.
  - **Instruções**: Criar um arquivo `middleware.ts` na raiz da pasta `src` do frontend. Verificar se o token JWT está presente nos cookies ao acessar as rotas protegidas (`/dashboard`, `/patients`, `/admin`). Se não estiver presente ou estiver expirado, redirecionar o usuário para `/login`.
  - **Critério de Aceitação**: Redirecionamento funcional ao acessar rotas restritas sem credenciais.
- [x] **Tarefa 4.3**: Integrar Tela de Usuários (Admin Master) no Frontend.
  - **Instruções**: Modificar a página `/admin/users` para listar os psicólogos fazendo requisição real para o backend. Integrar o formulário do modal para cadastrar novos psicólogos e a ação de exclusão na tabela.
  - **Critério de Aceitação**: CRUD de psicólogos totalmente integrado na tela e persistindo no SQLite.
