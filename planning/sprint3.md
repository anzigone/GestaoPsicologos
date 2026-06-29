# Sprint 3: Mock Frontend e Protótipo de Telas (sprint3.md)

## 1. Escopo e Objetivos
O objetivo desta sprint é desenvolver a interface de usuário (UI/UX) de alta fidelidade para todas as telas do sistema utilizando React, Next.js, TailwindCSS e Shadcn/UI, de acordo com as diretrizes do [layout-details.md](file:///D:/OneDrive/Git/SpecDrivenPsicologo/especification/layout-details.md). As páginas utilizarão dados mockados ou se conectarão aos endpoints mockados criados na Sprint 2.

**Objetivos principais**:
- Criar a estrutura de roteamento e layouts padrão (navegação e barras laterais).
- Implementar as interfaces responsivas com suporte a temas e estilização premium.
- Assegurar que todas as interações básicas e feedbacks de botões estejam visuais.
- Traduzir e manter todo o texto da interface estritamente em Português do Brasil (PT-BR).

---

## 2. Entregáveis da Sprint
Ao final desta sprint, os seguintes entregáveis deverão estar prontos:
- Páginas estruturadas e responsivas implementadas sob `./codebase/frontend/src/app/`:
  - Rota `/login`: Tela de login com o logotipo da psicologia.
  - Rota `/welcome` ou `/`: Portal de Boas-Vindas com atalhos interativos.
  - Rota `/admin/settings`: Configurações profissionais do psicólogo e integrações de agendas.
  - Rota `/patients`: Painel Mestre-Detalhe contendo a listagem com busca de consultantes, formulário de dados cadastrais (com cálculo de variação percentual), anamnese clínica e timeline de evoluções.
  - Rota `/dashboard`: Painel com métricas contábeis/financeiras e gráficos de acompanhamento.
  - Rota `/admin/users` (Apenas visual): Tabela de gerenciamento de psicólogos para o usuário Master.
- Fluxo de roteamento e navegação fluida de transição entre as páginas.

---

## 3. Checkpoints de Validação Humana
Para validar a conclusão desta Sprint, o humano realizará os seguintes testes:

1. **Executar a Aplicação**:
   - Subir os containers na raiz do projeto:
     ```bash
     docker compose up --build -d
     ```
2. **Navegar pelo Protótipo**:
   - Acessar `http://localhost:3000` no navegador.
   - Testar o fluxo de login (clicar em "Entrar") e navegar pelas opções da tela de Boas-Vindas.
3. **Validação da Ficha do Paciente**:
   - Acessar `/patients`. Selecionar um paciente fictício da lista, alterar o valor da consulta e verificar se o cálculo do desvio percentual é atualizado dinamicamente na tela.
   - Navegar pelas abas de "Primeira Análise" e "Histórico de Sessões" verificando a estética dos accordions e a timeline.
4. **Verificação Comparativa**:
   - Comparar o resultado visual no navegador com os wireframes contidos na pasta de especificações (`especification/files/previa1.png` a `previa5.png`).
5. **Encerramento da Sprint**:
   - Realizar commit com a mensagem `feat: sprint 3 mock frontend concluido`, criar a tag Git `v0.3.0` e fazer push para a branch `main`.
