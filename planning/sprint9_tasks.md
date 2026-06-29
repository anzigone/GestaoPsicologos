# Tarefas da Sprint 9: Polimento e Ajustes (sprint9_tasks.md)

Este documento descreve detalhadamente as tarefas técnicas individuais necessárias para a execução da Sprint 9 pelo Programador.

---

## 1. Validação de Formulários (Frontend & Backend)
- [ ] **Tarefa 1.1**: Implementar validação do lado do cliente no Frontend.
  - **Instruções**: Utilizar `zod` e `react-hook-form` nas telas de Login, Cadastro de Pacientes, Configurações do Psicólogo e Novo Usuário. Validar e-mail, formato de CRP, obrigatoriedade de campos e valores numéricos mínimos.
  - **Critério de Aceitação**: Formulários impedindo envio se houver dados inconsistentes e exibindo alertas claros em PT-BR.
- [ ] **Tarefa 1.2**: Validar payloads no Backend.
  - **Instruções**: No Go, criar structs de request com tags de validação (ex: `validate:"required,email"`) e rejeitar payloads inválidos com HTTP 400.
  - **Critério de Aceitação**: Tentativa de burlar validações enviando JSON incompleto para a API retornando erro estruturado.

---

## 2. Refinamento de Interface (UX/UI)
- [ ] **Tarefa 2.1**: Adicionar Feedbacks de Carregamento (Loading State).
  - **Instruções**: Adicionar propriedades de `disabled` e spinners nos botões de formulário quando a flag `isLoading` ou `isSubmitting` estiver ativa. Inserir componentes Skeleton do Shadcn enquanto os dados iniciais de perfil e pacientes são buscados na API.
  - **Critério de Aceitação**: Transição suave exibindo estados de carregamento em todas as requisições lentas.
- [ ] **Tarefa 2.2**: Formatação de Moedas e Datas.
  - **Instruções**: Usar `Intl.NumberFormat('pt-BR', { style: 'currency', currency: 'BRL' })` para exibição de valores e `Intl.DateTimeFormat` para exibição de datas nas tabelas de transações e na timeline do prontuário.
  - **Critério de Aceitação**: Valores monetários exibidos como `R$ X,XX` e datas como `DD/MM/AAAA`.

---

## 3. Ajustes de Fuso Horário e Timezones
- [ ] **Tarefa 3.1**: Normalizar Armazenamento de Datas no Backend.
  - **Instruções**: Salvar todas as datas no banco de dados em formato UTC ISO 8601 (`YYYY-MM-DDTHH:MM:SSZ`).
  - **Critério de Aceitação**: SQLite/MySQL persistindo datas exclusivamente em UTC.
- [ ] **Tarefa 3.2**: Ajustar Exibição de Data Local no Frontend.
  - **Instruções**: Ao renderizar as datas recebidas da API, convertê-las para o fuso horário local do navegador do usuário.
  - **Critério de Aceitação**: Datas de agendamentos exibidas no fuso horário correto do usuário final.
