# Sprint 4: Autenticação e Gestão de Usuários (sprint4.md)

## 1. Escopo e Objetivos
O objetivo desta sprint é implementar a segurança e a dinâmica real do sistema através da autenticação JWT, encriptação de senhas no banco de dados, proteção de rotas privadas e o CRUD completo de gerenciamento de psicólogos administrado pelo usuário Master.

**Objetivos principais**:
- Inicializar a persistência de usuários com banco de dados SQLite local.
- Criar a semente (seed) automática do usuário master `admin/admin` no primeiro boot do backend.
- Desenvolver a geração e validação de tokens JWT no backend e middleware de autenticação.
- Proteger as rotas no frontend usando Next.js Middleware, redirecionando usuários não autenticados.
- Habilitar o cadastro e remoção real de psicólogos pela conta master.

---

## 2. Entregáveis da Sprint
Ao final desta sprint, os seguintes entregáveis deverão estar prontos:
- Tabela `users` configurada no SQLite de desenvolvimento.
- Mecanismo de autenticação real operando no backend na rota `POST /api/auth/login` com criptografia de senha.
- Middleware Go para validação de JWT em rotas restritas de API.
- Middleware Next.js protegendo rotas privadas (`/dashboard`, `/patients`, `/admin/*`) no frontend.
- Tela de Gestão de Usuários no frontend (`/admin/users`) integrada à API e realizando inclusões/exclusões de psicólogos no banco de dados.

---

## 3. Checkpoints de Validação Humana
Para validar a conclusão desta Sprint, o humano realizará os seguintes testes:

1. **Subir os serviços**:
   - Executar na raiz do projeto:
     ```bash
     docker compose up --build -d
     ```
2. **Testar Proteção de Rotas**:
   - Tentar acessar diretamente `http://localhost:3000/patients` no navegador.
   - Esperado: Bloqueio de navegação e redirecionamento automático para a tela `/login`.
3. **Login do Administrador Master**:
   - Digitar usuário `admin` e senha `admin` na tela de login.
   - Esperado: Autenticação com sucesso e redirecionamento para o portal `/welcome`.
4. **Cadastrar Psicólogo**:
   - Acessar a página de gestão de usuários (`http://localhost:3000/admin/users`).
   - Clicar em "+ Adicionar Psicólogo", preencher Nome, E-mail, CRP e uma Senha Inicial, e salvar.
   - Esperado: O novo psicólogo é exibido na tabela.
5. **Autenticação do Psicólogo**:
   - Realizar logout do admin master e tentar efetuar login com as credenciais do novo psicólogo cadastrado.
   - Esperado: Login bem-sucedido e acesso liberado às rotas privadas.
6. **Encerramento da Sprint**:
   - Realizar commit com a mensagem `feat: sprint 4 autenticacao e gestao de usuarios concluido`, criar a tag Git `v0.4.0` e fazer push para a branch `main`.
