# Tarefas da Sprint 6: Ficha Consultante (sprint6_tasks.md)

Este documento descreve detalhadamente as tarefas técnicas individuais necessárias para a execução da Sprint 6 pelo Programador.

---

## 1. Banco de Dados e Migração (Backend)
- [x] **Tarefa 1.1**: Criar rotinas de migração para `patients` e `first_analysis`.
  - **Instruções**: No arquivo de inicialização do banco (`internal/database/database.go`), criar as tabelas `patients` e `first_analysis` executando as instruções DDL descritas em `business-rules-details.md`.
  - **Critério de Aceitação**: Tabelas criadas no banco de dados SQLite ao iniciar o backend.
  - **Implementado em**: `codebase/backend/internal/database/database.go` — função `Migrate()`. Tabelas `patients` (com campos: id, psychologist_id, name, phone, birthdate, age, profession, company, city, state, marital_status, consultation_fee, created_at, updated_at) e `first_analysis` (com 14 campos clínicos + patient_id + updated_at) adicionadas ao slice `statements`. FK com `ON DELETE CASCADE` em ambas as tabelas.

---

## 2. Implementação dos Endpoints REST (Backend)
- [x] **Tarefa 2.1**: Desenvolver repositórios e rotas de Cadastro de Pacientes.
  - **Instruções**:
    - Implementar queries para listar pacientes filtrando por psicólogo (ID do JWT) e nome (query param `q`), buscar por ID, criar e atualizar.
    - Mapear nos endpoints: `GET /api/patients`, `GET /api/patients/{id}`, `POST /api/patients`, `PUT /api/patients/{id}`.
  - **Critério de Aceitação**: Endpoints respondendo a dados do banco SQLite com isolamento estrito.
  - **Implementado em**: `codebase/backend/internal/handlers/patients.go` — handlers `ListPatients`, `GetPatient`, `CreatePatient`, `UpdatePatient`, `DeletePatient`, `ExportPatientPDF`. Isolamento por `psychologist_id` obtido via `mw.UserIDFromContext(r)`. Helper `scanPatient()` e constante `querySelectPatient` evitam duplicação. Rotas registradas em `cmd/server/main.go` via `r.Route("/api/patients", ...)`.
- [x] **Tarefa 2.2**: Desenvolver repositórios e rotas da Primeira Análise (Anamnese).
  - **Instruções**:
    - Criar query SQL com lógica de `Upsert` (tentar inserir, se der conflito na chave `patient_id`, atualizar todos os campos clínicos) para salvar os dados da Primeira Análise do paciente.
    - Mapear nos endpoints: `GET /api/patients/{id}/analysis` e `PUT /api/patients/{id}/analysis`.
  - **Critério de Aceitação**: Upsert da anamnese salvando e atualizando dados sem criar registros duplicados.
  - **Implementado em**: `codebase/backend/internal/handlers/analysis.go` — handlers `GetAnalysis` e `UpdateAnalysis`. Segurança via helper `patientBelongsTo()` que verifica se o paciente pertence ao psicólogo autenticado. Upsert via `INSERT ... ON CONFLICT(patient_id) DO UPDATE SET ...` nativo do SQLite. Retorna `FirstAnalysis` vazio (não 404) quando anamnese ainda não foi preenchida.

---

## 3. Integração de Telas e Busca (Frontend)
- [x] **Tarefa 3.1**: Integrar Lista Lateral e Busca de Pacientes.
  - **Instruções**: No frontend, implementar um hook SWR que busca de `GET /api/patients?q={termo}` toda vez que o input de busca for digitado (com debounce de 300ms). Ao clicar em um paciente, atualizar o ID do paciente ativo no estado da página.
  - **Critério de Aceitação**: Lista exibindo pacientes reais do backend e filtrando corretamente.
  - **Implementado em**: `codebase/frontend/src/app/(app)/patients/page.tsx` — `useEffect` com `debounceRef` de 300ms buscando `GET /api/patients?q=...`. Sidebar renderiza a lista ordenada alfabeticamente com highlight do paciente selecionado via borda teal.
- [x] **Tarefa 3.2**: Integrar Criação de Novo Paciente.
  - **Instruções**: Fazer com que o clique em "+ Novo Paciente" limpe os estados do formulário ativo e exiba um formulário em branco. Ao salvar, fazer um POST para `/api/patients` e, em caso de sucesso, forçar a atualização da lista lateral e selecionar o novo paciente.
  - **Critério de Aceitação**: Fluxo de inserção de paciente operando perfeitamente.
  - **Implementado em**: `codebase/frontend/src/app/(app)/patients/page.tsx` — modal controlado por `showModal`, formulário `newForm` com nome, telefone e valor da consulta. `handleCreatePatient()` faz POST, insere o paciente na lista local, seleciona-o e exibe toast de sucesso. Modal inclui preview dinâmico da variação tarifária.
- [x] **Tarefa 3.3**: Integrar Formulário de Edição Cadastral e Cálculo de Variação (Aba 1).
  - **Instruções**: Carregar os dados pessoais na Aba 1 ao selecionar um paciente. No input do valor da consulta, disparar um cálculo dinâmico: comparar o valor digitado com a tarifa base do psicólogo (guardada no estado global do usuário logado) e exibir o desvio em porcentagem (`+X% acima` ou `-X% abaixo`). Ao salvar, disparar PUT para `/api/patients/{id}`.
  - **Critério de Aceitação**: Variação percentual calculada em tempo real na tela e persistência de dados pessoais funcional.
  - **Implementado em**: `codebase/frontend/src/app/(app)/patients/page.tsx` — `baseFee` carregado de `GET /api/psychologist`. Funções `feeVariation()` e `feeVariationColor()` calculam a variação percentual e colorem o texto (teal para acima, laranja para abaixo). `handleSaveForm()` dispara PUT e atualiza estado local do paciente e da lista.
- [x] **Tarefa 3.4**: Integrar Formulário de Primeira Análise (Aba 2).
  - **Instruções**: Carregar a anamnese clínica ao mudar para a Aba 2. Integrar todos os campos dos Accordions clínicas com o estado e disparar um PUT para `/api/patients/{id}/analysis` ao clicar em "Salvar Primeira Análise".
  - **Critério de Aceitação**: Salvamento e recuperação dos dados clínicos de anamnese funcionais através dos Accordions.
  - **Implementado em**: `codebase/frontend/src/app/(app)/patients/page.tsx` — constante `ANALYSIS_SECTIONS` mapeia as 14 seções clínicas. Accordions controlados por `openAccordion`. Anamnese carregada via `GET /api/patients/{id}/analysis` ao selecionar um paciente. `handleSaveAnalysis()` dispara PUT e atualiza estado local.
