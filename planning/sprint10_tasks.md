# Tarefas da Sprint 10: Hardening (sprint10_tasks.md)

Este documento descreve detalhadamente as tarefas técnicas individuais necessárias para a execução da Sprint 10 pelo Programador.

---

## 1. Migração para Cookies HttpOnly (Backend & Frontend)
- [ ] **Tarefa 1.1**: Enviar Cookies HttpOnly no Login (Backend).
  - **Instruções**: No handler `POST /api/auth/login`, em vez de devolver o JWT no corpo do JSON, configurar um header `Set-Cookie` enviando o token na chave `token` com propriedades `HttpOnly=true`, `Path=/`, `SameSite=Lax` e `Secure=true` (em produção).
  - **Critério de Aceitação**: Cookie de sessão configurado na resposta do login.
- [ ] **Tarefa 1.2**: Ler JWT dos Cookies nos Middlewares (Backend).
  - **Instruções**: Atualizar o middleware de autenticação em Go para extrair o token do cookie da requisição (`r.Cookie("token")`) em vez do header `Authorization`.
  - **Critério de Aceitação**: Rota protegida da API validando a sessão a partir do cookie HTTP.
- [ ] **Tarefa 1.3**: Adaptar Next.js Middleware para validação de Cookies.
  - **Instruções**: No arquivo `middleware.ts` do frontend, ler o cookie `token` da requisição e realizar o redirecionamento. Configurar chamadas de API no frontend para enviar credenciais (`credentials: 'include'`) para garantir que o navegador repasse os cookies.
  - **Critério de Aceitação**: Roteamento privado do Next.js funcionando com base em cookies HttpOnly.

---

## 2. Leitura Segura de Variáveis (AWS Secrets Manager)
- [ ] **Tarefa 2.1**: Implementar Leitura de Parâmetros AWS.
  - **Instruções**: No backend Go, criar um módulo que, se detectar a variável `ENV=production`, utilize o SDK oficial da AWS (`aws-sdk-go-v2`) para ler o segredo do AWS Secrets Manager e inicializar as variáveis sensíveis do sistema (chaves do banco, chaves do OAuth, segredo JWT).
  - **Critério de Aceitação**: Leitura de segredos funcionando ao emular ambiente de produção.

---

## 3. Menu Mobile Responsivo (Frontend)
- [ ] **Tarefa 3.1**: Criar Gaveta do Menu Mobile.
  - **Instruções**: Adicionar o componente Sheet do Shadcn no layout da área logada. Configurar media queries do Tailwind (`md:hidden`) para ocultar a sidebar fixa em telas pequenas e exibir o botão hambúrguer. Ao clicar no botão, abrir o Sheet exibindo as opções da sidebar de navegação.
  - **Critério de Aceitação**: Navegação mobile funcionando perfeitamente em telas estreitas, sem quebras de layout.
