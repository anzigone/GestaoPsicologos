# Tarefas da Sprint 6: Ficha Consultante (sprint6_tasks.md)

Este documento descreve detalhadamente as tarefas técnicas individuais necessárias para a execução da Sprint 6 pelo Programador.

---

## 1. Banco de Dados e Migração (Backend)
- [ ] **Tarefa 1.1**: Criar rotinas de migração para `patients` e `first_analysis`.
  - **Instruções**: No arquivo de inicialização do banco (`internal/database/database.go`), criar as tabelas `patients` e `first_analysis` executando as instruções DDL descritas em `business-rules-details.md`.
  - **Critério de Aceitação**: Tabelas criadas no banco de dados SQLite ao iniciar o backend.

---

## 2. Implementação dos Endpoints REST (Backend)
- [ ] **Tarefa 2.1**: Desenvolver repositórios e rotas de Cadastro de Pacientes.
  - **Instruções**:
    - Implementar queries para listar pacientes filtrando por psicólogo (ID do JWT) e nome (query param `q`), buscar por ID, criar e atualizar.
    - Mapear nos endpoints: `GET /api/patients`, `GET /api/patients/{id}`, `POST /api/patients`, `PUT /api/patients/{id}`.
  - **Critério de Aceitação**: Endpoints respondendo a dados do banco SQLite com isolamento estrito.
- [ ] **Tarefa 2.2**: Desenvolver repositórios e rotas da Primeira Análise (Anamnese).
  - **Instruções**:
    - Criar query SQL com lógica de `Upsert` (tentar inserir, se der conflito na chave `patient_id`, atualizar todos os campos clínicos) para salvar os dados da Primeira Análise do paciente.
    - Mapear nos endpoints: `GET /api/patients/{id}/analysis` e `PUT /api/patients/{id}/analysis`.
  - **Critério de Aceitação**: Upsert da anamnese salvando e atualizando dados sem criar registros duplicados.

---

## 3. Integração de Telas e Busca (Frontend)
- [ ] **Tarefa 3.1**: Integrar Lista Lateral e Busca de Pacientes.
  - **Instruções**: No frontend, implementar um hook SWR que busca de `GET /api/patients?q={termo}` toda vez que o input de busca for digitado (com debounce de 300ms). Ao clicar em um paciente, atualizar o ID do paciente ativo no estado da página.
  - **Critério de Aceitação**: Lista exibindo pacientes reais do backend e filtrando corretamente.
- [ ] **Tarefa 3.2**: Integrar Criação de Novo Paciente.
  - **Instruções**: Fazer com que o clique em "+ Novo Paciente" limpe os estados do formulário ativo e exiba um formulário em branco. Ao salvar, fazer um POST para `/api/patients` e, em caso de sucesso, forçar a atualização da lista lateral e selecionar o novo paciente.
  - **Critério de Aceitação**: Fluxo de inserção de paciente operando perfeitamente.
- [ ] **Tarefa 3.3**: Integrar Formulário de Edição Cadastral e Cálculo de Variação (Aba 1).
  - **Instruções**: Carregar os dados pessoais na Aba 1 ao selecionar um paciente. No input do valor da consulta, disparar um cálculo dinâmico: comparar o valor digitado com a tarifa base do psicólogo (guardada no estado global do usuário logado) e exibir o desvio em porcentagem (`+X% acima` ou `-X% abaixo`). Ao salvar, disparar PUT para `/api/patients/{id}`.
  - **Critério de Aceitação**: Variação percentual calculada em tempo real na tela e persistência de dados pessoais funcional.
- [ ] **Tarefa 3.4**: Integrar Formulário de Primeira Análise (Aba 2).
  - **Instruções**: Carregar a anamnese clínica ao mudar para a Aba 2. Integrar todos os campos dos Accordions clínicas com o estado e disparar um PUT para `/api/patients/{id}/analysis` ao clicar em "Salvar Primeira Análise".
  - **Critério de Aceitação**: Salvamento e recuperação dos dados clínicos de anamnese funcionais através dos Accordions.
