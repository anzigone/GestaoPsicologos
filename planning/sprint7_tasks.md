# Tarefas da Sprint 7: Dashboard (sprint7_tasks.md)

Este documento descreve detalhadamente as tarefas técnicas individuais necessárias para a execução da Sprint 7 pelo Programador.

---

## 1. Banco de Dados e Migração (Backend)
- [ ] **Tarefa 1.1**: Criar rotina de migração para a tabela `sessions`.
  - **Instruções**: No arquivo de inicialização do banco (`internal/database/database.go`), criar a tabela `sessions` executando as instruções DDL descritas em `business-rules-details.md`.
  - **Critério de Aceitação**: Tabela criada no banco de dados SQLite ao iniciar o backend.

---

## 2. Endpoints e Queries de Agregação (Backend)
- [ ] **Tarefa 2.1**: Implementar queries SQL de Agregações de Dashboard.
  - **Instruções**:
    - Query para calcular faturamento total de sessões do mês com status 'pago', contagem de sessões no mês, soma de valores com status 'pendente' e contagem de pacientes ativos vinculados ao psicólogo logado.
    - Query de faturamento dos últimos 6 meses agrupado por ano-mês.
    - Query contanto sessões por ID de paciente.
  - **Critério de Aceitação**: Queries de agregação compilando com sucesso e testadas.
- [ ] **Tarefa 2.2**: Implementar Handlers de Estatísticas e Gráficos no Go.
  - **Instruções**:
    - `GET /api/dashboard/stats`: Retornar os KPIs agregados em JSON.
    - `GET /api/dashboard/charts`: Retornar dados de faturamento mensal e de atendimentos por paciente.
    - Proteger ambas as rotas com o middleware de JWT.
  - **Critério de Aceitação**: Endpoints respondendo a dados calculados em tempo real do banco SQLite.
- [ ] **Tarefa 2.3**: Implementar Handler de Transações Financeiras Recentes.
  - **Instruções**: Mapear `GET /api/dashboard/transactions` retornando lista detalhada de sessões contendo data, valor da consulta do paciente associado, nome do paciente e status de pagamento.
  - **Critério de Aceitação**: Rota retornando a lista de transações em ordem cronológica decrescente.

---

## 3. Integração das Telas (Frontend)
- [ ] **Tarefa 3.1**: Integrar Cards de KPI do Dashboard.
  - **Instruções**: Fazer a chamada `GET /api/dashboard/stats` usando SWR para alimentar os cards de "Faturamento Total", "Atendimentos", "Pacientes Ativos" e "A Receber".
  - **Critério de Aceitação**: Cards do Dashboard exibindo valores corretos do webservice.
- [ ] **Tarefa 3.2**: Integrar Gráficos no Dashboard.
  - **Instruções**: Integrar a chamada `GET /api/dashboard/charts` com os gráficos do Recharts (ou biblioteca similar adotada) para renderizar a curva de crescimento financeiro e barras de sessões por paciente.
  - **Critério de Aceitação**: Gráficos dinâmicos renderizados com dados reais.
- [ ] **Tarefa 3.3**: Integrar Tabela de Transações Financeiras.
  - **Instruções**: Integrar a chamada `GET /api/dashboard/transactions` na tabela da tela de dashboard. Adicionar lógica de filtro do lado do cliente para filtrar a lista local por status (Pago/Pendente) e buscar por nome do paciente.
  - **Critério de Aceitação**: Tabela exibindo os dados das transações e filtros funcionando.
