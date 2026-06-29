# Tarefas da Sprint 3: Mock Frontend (sprint3_tasks.md)

Este documento descreve detalhadamente as tarefas técnicas individuais necessárias para a execução da Sprint 3 pelo Programador.

---

## 1. Rotas e Estrutura de Navegação (Layouts)
- [x] **Tarefa 1.1**: Configurar Rotas do App Router do Next.js.
  - **Instruções**: Estruturar as rotas `/login`, `/admin/settings`, `/patients` e `/dashboard` no diretório `./codebase/frontend/src/app`.
  - **Critério de Aceitação**: Rotas acessíveis via URL direta sem erros 404.
- [x] **Tarefa 1.2**: Implementar o Layout Global e Navbar/Sidebar.
  - **Instruções**: Criar um componente de navegação lateral (Sidebar) comum para a área logada contendo links rápidos com ícones Lucide (Dashboard, Pacientes, Minha Conta, Logout).
  - **Critério de Aceitação**: Navegação funcional entre as rotas através da Sidebar.

---

## 2. Implementação das Interfaces de Usuário (alta fidelidade)
- [x] **Tarefa 2.1**: Implementar a Tela de Login (`/login`).
  - **Instruções**: Centralizar o formulário em um card sobre um gradiente suave. Renderizar a imagem do logotipo da psicologia (`Simbolo-da-psicologia.jpg`) centralizada no topo. Adicionar campos de e-mail, senha e botão "Entrar".
  - **Critério de Aceitação**: Fidelidade ao wireframe `previa1.png`.
- [x] **Tarefa 2.2**: Implementar a Tela de Boas-Vindas (`/`).
  - **Instruções**: Criar a mensagem de saudação e os três cards interativos de atalho ("Administração do Psicólogo", "Ficha Consultante", "Dashboard").
  - **Critério de Aceitação**: Fidelidade ao wireframe `previa2.png`.
- [x] **Tarefa 2.3**: Implementar a Tela de Administração do Psicólogo (`/admin/settings`).
  - **Instruções**: Layout de duas colunas. Coluna esquerda: dados do psicólogo e tarifas padrão. Coluna direita: cartões de integração com Google Meet e Microsoft Outlook (mostrando status mockado "Conectado" ou "Conectar Conta").
  - **Critério de Aceitação**: Layout de duas colunas responsivo, fiel a `previa3.png`.
- [x] **Tarefa 2.4**: Implementar a Tela de Ficha do Consultante (`/patients`).
  - **Instruções**:
    - **Esquerda**: Lista lateral de pacientes com campo de busca de texto e o botão de ação **"+ Novo Paciente"**.
    - **Direita**: Tela detalhada com Abas do Shadcn UI:
      - *Aba 1 (Cadastro)*: Inputs pessoais e campo de valor de consulta (calcular variação em tempo real em relação ao valor padrão do psicólogo usando state do React).
      - *Aba 2 (Análise)*: Componente Accordion contendo as seções clínicas.
      - *Aba 3 (Evoluções)*: Timeline com blocos de notas de sessões e botão "Iniciar Nova Sessão".
  - **Critério de Aceitação**: Busca local de pacientes funcionando via estado do React e abas funcionais. Fiel a `previa4.png`.
- [x] **Tarefa 2.5**: Implementar a Tela de Dashboard (`/dashboard`).
  - **Instruções**: Adicionar cartões de KPI (Faturamento, Pacientes Ativos, etc.), um gráfico financeiro ilustrativo (usando `recharts`) e a tabela de lançamentos com badges de status de pagamento.
  - **Critério de Aceitação**: Dashboard visualmente rico e fiel a `previa5.png`.
- [x] **Tarefa 2.6**: Implementar a Tela de Usuários (Admin Master) (`/admin/users`).
  - **Instruções**: Tabela listando psicólogos mockados e botão para abrir modal de inclusão de novos psicólogos.
  - **Critério de Aceitação**: Modal abrindo e fechando de forma fluida.

---

## 3. Integração com Mock Backend
- [x] **Tarefa 3.1**: Configurar chamadas para a API mockada.
  - **Instruções**: Utilizar a biblioteca `fetch` nativo no Next.js apontando as chamadas de dados para `http://localhost:8080` (URL do backend do docker compose) para carregar os dados de pacientes e simular o login.
  - **Critério de Aceitação**: O frontend renderizando as listas que vêm da API Go da Sprint 2.
