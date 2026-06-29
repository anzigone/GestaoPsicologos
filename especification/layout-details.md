# Projeto: Gestão Psicólogos - Detalhamento do Layout (layout-details.md)

Este documento descreve detalhadamente o design visual, a identidade, a paleta de cores, os componentes de interface (Shadcn/UI) e a estrutura detalhada de cada tela. O sistema será totalmente traduzido e apresentado em **Português do Brasil (PT-BR)**.

---

## 1. Guia de Estilo (Design System)

### Paleta de Cores (Aparência Premium e Terapêutica)
Para transmitir confiança, calma e profissionalismo, utilizaremos uma paleta de cores com tons neutros suaves e acentos de verde menta/azul ardósia:

- **Fundo da Aplicação (Background)**:
  - Claro: `#F8FAFC` (Slate 50) com gradientes suaves em áreas específicas.
  - Escuro (opcional/tema escuro): `#0F172A` (Slate 900).
- **Cor Primária (Ações principais e marca)**:
  - `#0D9488` (Teal 600) e `#0F766E` (Teal 700) para estados de hover.
- **Cor Secundária (Elementos de navegação secundários)**:
  - `#475569` (Slate 600).
- **Acentos e Estados**:
  - **Sucesso (Financeiro recebido/Sessão paga)**: `#10B981` (Emerald 500).
  - **Alerta (Pendente/Não pago)**: `#F59E0B` (Amber 500).
  - **Erro (Falhas/Cancelamento)**: `#EF4444` (Red 500).
- **Superfícies (Cards e Painéis)**:
  - `#FFFFFF` com 80% de opacidade e efeito `backdrop-blur-md` (Glassmorphism sutil para uma estética moderna).
  - Bordas finas em `#E2E8F0` (Slate 200).

### Tipografia
- **Família de Fontes**: `Outfit` (do Google Fonts) como fonte principal para títulos e `Inter` para textos de leitura, proporcionando uma aparência limpa, geométrica e extremamente premium.
- **Escala de Texto**:
  - Títulos Principais: `text-3xl` (30px), peso `font-bold` (700).
  - Subtítulos de Seções: `text-xl` (20px), peso `font-semibold` (600).
  - Textos de Apoio/Inputs: `text-sm` (14px), peso `font-normal` (400) ou `font-medium` (500).

### Ícones
- Utilização da biblioteca **Lucide React** (ex: `Calendar`, `User`, `Users`, `DollarSign`, `TrendingUp`, `LogOut`, `Settings`, `FileText`, `Video`).

---

## 2. Estrutura Detalhada das Telas

### 2.1. Tela de Login (previa1.png)
- **Objetivo**: Acesso seguro para Psicólogos e Administrador Master.
- **Layout**:
  - Centralizado, com um card flutuante sobre um fundo degradê suave (`from-teal-50 to-slate-100`).
  - No topo do card, exibição do logotipo da psicologia ([Simbolo-da-psicologia.jpg](file:///D:/OneDrive/Git/SpecDrivenPsicologo/especification/files/Simbolo-da-psicologia.jpg)) em alta resolução, redimensionado para 80x80px, com bordas arredondadas e centralizado.
  - Título: "Gestão Psicológica" (`text-2xl font-bold text-slate-800 text-center`).
  - Formulário contendo:
    - Campo de e-mail (Input Shadcn com ícone `Mail`).
    - Campo de senha (Input Shadcn do tipo password com botão de alternar visibilidade).
    - Botão de envio "Entrar" (`bg-teal-600 hover:bg-teal-700 text-white w-full rounded-lg py-2`).
  - Mensagens de erro de validação (ex: "E-mail ou senha incorretos") exibidas em vermelho (`text-red-500 text-xs mt-1`).

### 2.2. Tela de Boas-Vindas / Dashboard Inicial (previa2.png)
- **Objetivo**: Primeiro contato do usuário psicólogo após realizar o login.
- **Layout**:
  - Barra superior (Header) contendo o nome do Psicólogo conectado, CRP e botão de logout.
  - Seção central com mensagem acolhedora: "Olá, Dr(a). [Nome]! O que deseja gerenciar hoje?".
  - Três grandes cards interativos (com efeito de hover de escala e sombra):
    1. **Administração do Psicólogo**: Ícone `Settings`, título "Minha Conta & Integrações", descrição de atalho.
    2. **Ficha Consultante**: Ícone `Users`, título "Pacientes & Prontuários", descrição de atalho.
    3. **Dashboard Financeiro**: Ícone `TrendingUp`, título "Finanças & Relatórios", descrição de atalho.

### 2.3. Tela de Administração do Psicólogo (previa3.png)
- **Objetivo**: Cadastro de dados profissionais, configuração de valores de consultas e conexões de agenda.
- **Layout**:
  - Duas colunas principais:
    - **Coluna Esquerda (Dados do Profissional)**:
      - Formulário com os campos: Nome do Psicólogo, CRP, Especialidade, E-mail Principal, Telefone Celular.
      - Seção de valores: Valor da Consulta Padrão (R$), Quantidade de Sessões do Pacote, Valor do Pacote (R$).
      - Formulário de troca de senha integrado na parte inferior (Senha Atual, Nova Senha).
    - **Coluna Direita (Integrações e Agendas)**:
      - Card **Google Meet**: Botão "Conectar Conta Google" (ou status "Conectado como e-mail@gmail.com" com botão para desconectar). Informações sobre criação automática do link do Google Meet nas sessões.
      - Card **Microsoft Outlook**: Botão "Conectar Calendário Outlook" (ou status "Conectado"). Visualização em miniatura do calendário sincronizado do Outlook mostrando os horários ocupados.
  - Botão de ação "Salvar Alterações" fixado no rodapé da página.

### 2.4. Tela da Ficha do Consultante / Prontuário (previa4.png)
- **Objetivo**: Cadastro completo de pacientes, listagem e busca de consultantes, e acompanhamento das sessões.
- **Layout Geral (Estrutura Mestre-Detalhe)**:
  - **Coluna Lateral Esquerda (Lista de Pacientes)**:
    - Ocupa 25% da largura da tela.
    - No topo: Campo de busca de texto ("Buscar paciente...") com ícone de lupa e botão destacado **"+ Novo Paciente"** (ícone `UserPlus`, cor Teal).
    - Corpo: Lista vertical de cards compactos dos pacientes cadastrados, exibindo o Nome Completo e uma etiqueta do status financeiro geral (ex: "Em dia" ou "Pendente"). O paciente selecionado fica destacado visualmente (fundo Teal 50 ou borda Teal 600).
  - **Coluna Central/Direita (Painel de Detalhes do Paciente)**:
    - Ocupa 75% da largura da tela.
    - Se nenhum paciente estiver selecionado (estado inicial): exibe uma tela vazia amigável com a mensagem "Selecione um paciente na lista ou crie um novo para ver o prontuário".
    - Se "+ Novo Paciente" for clicado: Exibe o formulário de cadastro limpo na Aba 1 e o botão "Salvar Novo Paciente" no rodapé.
    - Se um paciente for selecionado:
      - **Cabeçalho**: Nome do paciente em destaque (`text-2xl font-bold`), telefone principal, e botão proeminente de ação "Exportar Ficha Completa (PDF)" (que gera o relatório em PDF com primeira análise e histórico).
      - **Menu de Abas (Tabs do Shadcn)**:
        - **Aba 1: Cadastro & Dados**:
          - Formulário com os campos de cadastro do paciente:
            - *Obrigatórios*: Nome Completo, Telefone Celular.
            - *Opcionais*: Data de Nascimento, Idade, Profissão, Empresa, Cidade, Estado, Estado Civil.
            - *Financeiro*: Campo para preencher o **Valor da Consulta a ser Cobrada** (R$). O sistema deve calcular e exibir automaticamente, em tempo real, a porcentagem de variação acima ou abaixo em relação ao valor padrão cadastrado na "Administração do Psicólogo" (ex: "+15% acima do valor padrão" ou "-10% abaixo do valor padrão").
          - Botão "Salvar Alterações Cadastrais" no rodapé da aba.
        - **Aba 2: Primeira Análise**:
          - Seções expansíveis (Accordion do Shadcn) para preenchimento clínico estruturado:
            - Queixa Principal / Diagnóstico e Sintoma.
            - Fatores (Desenvolvimento, Situacionais, Biológicos/Médicos).
            - Pontos fortes / Recursos / Vícios.
            - Padrões (Estímulos, Pensamentos, Comportamentos, Afetos, Fisiológico).
            - Objetivos e Plano de Tratamento.
          - Botão "Salvar Primeira Análise" no rodapé da aba.
        - **Aba 3: Evolução das Sessões (Histórico)**:
          - No topo: Botão destacado "Iniciar Nova Sessão" (cria uma sessão para o dia de hoje, abrindo um bloco de anotações vazio) e botão "Gerar Link Google Meet" (gera o link do Meet e o anexa à sessão).
          - Corpo: Timeline vertical com cards de cada dia de atendimento. Cada card contém:
            - Data da sessão (editável/editada automaticamente), Hora e Status do Atendimento (Pendente / Pago).
            - Caixa de texto expansível (rich text leve ou text area grande) para anotações do psicólogo.
            - Botão "Salvar Notas" individual ou auto-save indicando status de gravação.
            - Os blocos de anotações históricos são totalmente persistentes e podem ser editados a qualquer momento.

### 2.5. Tela de Dashboard Geral / Relatórios (previa5.png)
- **Objetivo**: Visualização financeira e contábil, além de contagem de atendimentos.
- **Layout**:
  - **Linha de Indicadores Chave (Cards de KPI)**:
    - Faturamento Total no Mês (R$).
    - Consultas Realizadas vs. Canceladas.
    - Pacientes Ativos.
    - Valor a Receber (Consultas pendentes de pagamento).
  - **Gráficos e Listas**:
    - Gráfico de linha mostrando o crescimento do faturamento nos últimos 6 meses.
    - Gráfico de barras com a distribuição de atendimentos por paciente.
    - Tabela de lançamentos financeiros recentes com filtros rápidos por paciente e por status (Pago/Pendente).
  - **Ação**: Botão para abrir o modal de relatórios contábeis detalhados para exportação.

### 2.6. Tela de Usuários (Administrador Master)
- **Objetivo**: Controle de acesso para criação de novos psicólogos.
- **Acesso**: Apenas usuário logado com papel `admin` (admin/admin).
- **Layout**:
  - Interface simples de tabela contendo colunas: Nome, CRP, E-mail, Data de Cadastro, Ações.
  - Botão superior "+ Adicionar Psicólogo", que abre um modal com formulário para preenchimento de: Nome, E-mail, CRP e Senha Inicial.
  - Coluna de Ações contendo botão "Excluir" (abre diálogo de confirmação "Deseja realmente excluir este psicólogo? Seus pacientes e dados também serão removidos permanentemente").

---

## 3. Imagens de Wireframes e Prévia de Layout

As imagens de prévia serão geradas no formato `previa{n}.png` na pasta de arquivos de especificação:
1. `especification/files/previa1.png` -> Wireframe da Tela de Login.
2. `especification/files/previa2.png` -> Wireframe da Tela de Boas-Vindas / Atalhos.
3. `especification/files/previa3.png` -> Wireframe da Tela de Administração do Psicólogo.
4. `especification/files/previa4.png` -> Wireframe da Tela de Ficha do Consultante.
5. `especification/files/previa5.png` -> Wireframe da Tela de Dashboard Financeiro.
