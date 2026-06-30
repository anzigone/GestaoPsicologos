# Sprint 7: Dashboard Financeiro e Atendimentos (sprint7.md)

## 1. Escopo e Objetivos
O objetivo desta sprint é ativar os relatórios gerenciais e contábeis de faturamento e sessões. Para isso, criaremos a tabela de sessões no banco de dados e desenvolveremos as queries de agregação no backend para alimentar o Dashboard com dados reais do psicólogo logado.

**Objetivos principais**:
- Executar a migração da tabela `sessions` no SQLite.
- Implementar endpoints de agregação analítica e contábil no backend.
- Conectar os cartões KPI, gráficos e tabelas de lançamentos na tela `/dashboard` do frontend.
- Viabilizar filtros por status de pagamento e pesquisa por paciente na tabela de transações.

---

## 2. Entregáveis da Sprint
Ao final desta sprint, os seguintes entregáveis deverão estar prontos:
- Tabela `sessions` configurada no SQLite.
- Endpoints de agregação analítica implementados no backend:
  - `GET /api/dashboard/stats`: Retorna os valores contábeis acumulados do mês (Faturamento, Total de Atendimentos, Pacientes Ativos, Pendências).
  - `GET /api/dashboard/charts`: Retorna dados formatados para gráfico de faturamento de 6 meses e volume por paciente.
  - `GET /api/dashboard/transactions`: Lista as transações (sessões) com dados estruturados de data, valor, paciente e status.
- Tela `/dashboard` no frontend integrada e operando em tempo real com os dados consolidados do backend.

---

## 3. Checkpoints de Validação Humana
Para validar a conclusão desta Sprint, o humano realizará os seguintes testes:

1. **Subir os serviços**:
   - Executar na raiz do projeto:
     ```bash
     docker compose up --build -d
     ```
2. **Carga de Dados de Teste**:
   - Cadastrar dois pacientes com tarifas de consulta diferentes (ex: R$ 100,00 e R$ 150,00).
   - Cadastrar sessões para eles no banco de dados:
     - 2 sessões como "pago" (R$ 100,00 + R$ 150,00 = R$ 250,00 de faturamento).
     - 1 sessão como "pendente" (R$ 150,00 a receber).
3. **Validação dos KPIs**:
   - Abrir a página `/dashboard`.
   - Esperado:
     - Faturamento Total: "R$ 250,00".
     - Valor a Receber: "R$ 150,00".
     - Pacientes Ativos: "2".
     - Atendimentos: "3".
4. **Validação de Gráficos e Filtros**:
   - Verificar os dados plotados nos gráficos correspondentes.
   - Filtrar a tabela de transações por "Pendente".
   - Esperado: Exibição apenas da sessão de R$ 150,00 que não foi paga.
5. **Encerramento da Sprint**:
   - Realizar commit com a mensagem `feat: sprint 7 dashboard concluido`, criar a tag Git `v0.7.0` e fazer push para a branch `main`.
    - fazer o push https://github.com/anzigone/GestaoPsicologos.git
