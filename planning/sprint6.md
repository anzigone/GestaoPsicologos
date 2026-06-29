# Sprint 6: Ficha Consultante e Anamnese (sprint6.md)

## 1. Escopo e Objetivos
O objetivo desta sprint é implementar o gerenciamento completo dos pacientes (consultantes) e a ficha clínica estruturada de Primeira Análise (anamnese). A persistência física das tabelas será ativada no SQLite, e a interface mestre-detalhe será completamente integrada à API.

**Objetivos principais**:
- Executar a migração das tabelas `patients` e `first_analysis`.
- Implementar as rotas de busca, inserção e alteração de pacientes com isolamento por psicólogo logado.
- Desenvolver a persistência rica do formulário de anamnese (Primeira Análise) composto por mais de 10 seções clínicas.
- Integrar a listagem lateral com busca textual e o formulário de cadastro com cálculo dinâmico de tarifa.

---

## 2. Entregáveis da Sprint
Ao final desta sprint, os seguintes entregáveis deverão estar prontos:
- Tabelas `patients` e `first_analysis` criadas no banco de dados SQLite.
- Endpoints de API implementados e integrados no backend:
  - `GET /api/patients?q={filtro}`: Retorna a lista de pacientes do psicólogo, filtrando opcionalmente por nome.
  - `POST /api/patients` e `PUT /api/patients/{id}`: Criação e atualização de dados cadastrais.
  - `GET /api/patients/{id}/analysis` e `PUT /api/patients/{id}/analysis`: Carregamento e Upsert (salvamento/atualização) da anamnese clínica.
- Interface `/patients` no frontend integrada e funcional, permitindo buscar, listar, cadastrar novos consultantes, alterar dados de cadastro (calculando a variação percentual sobre o valor base do psicólogo) e preencher a anamnese clínica.

---

## 3. Checkpoints de Validação Humana
Para validar a conclusão desta Sprint, o humano realizará os seguintes testes:

1. **Subir os serviços**:
   - Executar na raiz do projeto:
     ```bash
     docker compose up --build -d
     ```
2. **Criação de Paciente**:
   - Logar como psicólogo e navegar até `/patients`.
   - Clicar em "+ Novo Paciente". Preencher "Nome: Carlos Drummond", "Telefone: 21999998888", "Valor da Consulta: R$ 135,00" (considerando que o valor padrão da consulta do psicólogo é R$ 150,00). Clicar em "Salvar Novo Paciente".
   - Esperado: Carlos Drummond aparece na lista lateral imediatamente.
3. **Cálculo da Variação Percentual**:
   - Clicar sobre "Carlos Drummond" na lista lateral.
   - Observar o campo do valor da consulta na Aba de Cadastro.
   - Esperado: Indicação clara de variação de tarifa na tela: "-10% abaixo do valor padrão" (calculado a partir de R$ 135,00 vs R$ 150,00).
4. **Preenchimento da Primeira Análise**:
   - Clicar na aba "Primeira Análise".
   - Expandir a seção "Queixa Principal / Diagnóstico e Sintoma" e digitar dados clínicos de teste.
   - Preencher a seção "Plano de Tratamento" e clicar em "Salvar Primeira Análise".
   - Recarregar a página (F5), clicar no paciente e reabrir as seções da Primeira Análise.
   - Esperado: Os dados inseridos persistem e são carregados adequadamente.
5. **Encerramento da Sprint**:
   - Realizar commit com a mensagem `feat: sprint 6 ficha consultante concluido`, criar a tag Git `v0.6.0` e fazer push para a branch `main`.
