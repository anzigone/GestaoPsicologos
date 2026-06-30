# Tarefas da Sprint 9: Polimento e Ajustes (sprint9_tasks.md)

Este documento descreve detalhadamente as tarefas técnicas individuais necessárias para a execução da Sprint 9 pelo Programador.

---

## 1. Validação de Formulários (Frontend & Backend)
- [x] **Tarefa 1.1**: Implementar validação do lado do cliente no Frontend.
  - **Instruções**: Utilizar `zod` e `react-hook-form` nas telas de Login, Cadastro de Pacientes, Configurações do Psicólogo e Novo Usuário. Validar e-mail, formato de CRP, obrigatoriedade de campos e valores numéricos mínimos.
  - **Critério de Aceitação**: Formulários impedindo envio se houver dados inconsistentes e exibindo alertas claros em PT-BR.
  - **Status**: ✅ Concluído — Login (email + senha min 6), Admin Settings (nome, CRP regex, base_fee min 0), Admin Users (email, CRP, senha min 6), Novo Paciente (nome min 2, telefone obrigatório, valor min 0). Todos com mensagens de erro em PT-BR.
- [x] **Tarefa 1.2**: Validar payloads no Backend.
  - **Instruções**: No Go, criar structs de request com tags de validação (ex: `validate:"required,email"`) e rejeitar payloads inválidos com HTTP 400.
  - **Critério de Aceitação**: Tentativa de burlar validações enviando JSON incompleto para a API retornando erro estruturado.
  - **Status**: ✅ Concluído — `CreatePatient` rejeita nome vazio (400), fee negativa (400); `CreateSession` rejeita session_date vazio (400), status inválido (400), session_date em formato inválido (400). Todos retornam `{"error":"..."}`.

---

## 2. Refinamento de Interface (UX/UI)
- [x] **Tarefa 2.1**: Adicionar Feedbacks de Carregamento (Loading State).
  - **Instruções**: Adicionar propriedades de `disabled` e spinners nos botões de formulário quando a flag `isLoading` ou `isSubmitting` estiver ativa. Inserir componentes Skeleton do Shadcn enquanto os dados iniciais de perfil e pacientes são buscados na API.
  - **Critério de Aceitação**: Transição suave exibindo estados de carregamento em todas as requisições lentas.
  - **Status**: ✅ Concluído — Login ("Entrando..."), Novo Paciente ("Salvando..."), Cadastro de Paciente ("Salvando..."), Anamnese ("Salvando..."), Sessões ("Salvando...", "Criando Sessão"), Settings ("Salvando...", "Alterando..."), Admin Users ("Criando..."). Skeleton loaders na sidebar de pacientes, tabela de usuários, KPIs do dashboard e formulário de perfil.
- [x] **Tarefa 2.2**: Formatação de Moedas e Datas.
  - **Instruções**: Usar `Intl.NumberFormat('pt-BR', { style: 'currency', currency: 'BRL' })` para exibição de valores e `Intl.DateTimeFormat` para exibição de datas nas tabelas de transações e na timeline do prontuário.
  - **Critério de Aceitação**: Valores monetários exibidos como `R$ X,XX` e datas como `DD/MM/AAAA`.
  - **Status**: ✅ Concluído — Dashboard: `brl()` com `toLocaleString('pt-BR', { style: 'currency' })`; Pacientes: `toLocaleString('pt-BR', { style: 'currency' })` nos fees, `toLocaleDateString('pt-BR', { day: '2-digit', month: '2-digit', year: 'numeric' })` nas sessões; Admin Users: `toLocaleDateString('pt-BR')`.

---

## 3. Ajustes de Fuso Horário e Timezones
- [x] **Tarefa 3.1**: Normalizar Armazenamento de Datas no Backend.
  - **Instruções**: Salvar todas as datas no banco de dados em formato UTC ISO 8601 (`YYYY-MM-DDTHH:MM:SSZ`).
  - **Critério de Aceitação**: SQLite/MySQL persistindo datas exclusivamente em UTC.
  - **Status**: ✅ Concluído — `CreateSession` agora parseia `session_date` com `time.Parse(time.RFC3339)` e fallback `ParseInLocation(..., time.UTC)`, convertendo explicitamente para `.UTC().Format(time.RFC3339)` antes do INSERT. `created_at`/`updated_at` sempre em `time.Now().UTC()`.
- [x] **Tarefa 3.2**: Ajustar Exibição de Data Local no Frontend.
  - **Instruções**: Ao renderizar as datas recebidas da API, convertê-las para o fuso horário local do navegador do usuário.
  - **Critério de Aceitação**: Datas de agendamentos exibidas no fuso horário correto do usuário final.
  - **Status**: ✅ Concluído — `formatDate(iso)` usa `new Date(iso).toLocaleDateString('pt-BR', { day: '2-digit', month: '2-digit', year: 'numeric' })` que converte automaticamente de UTC para o fuso local do navegador.
