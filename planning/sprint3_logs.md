# Registro de Execução da Sprint 3 (sprint3_logs.md)

Este documento registra cronologicamente todas as ações realizadas pelo Programador durante a Sprint 3, servindo de trilha de auditoria e documentação técnica da implantação.

---

## 1. Histórico de Atividades e Comandos

- **Data**: 29/06/2026
- **Status da Sprint**: Concluída ✅

### Passos Executados:

1. **Dependências adicionadas ao `package.json`**:
   - `recharts ^2.12.7` — gráfico de área no Dashboard.
   - `Dockerfile` alterado de `npm ci` → `npm install` para instalar nova dependência no build Docker.

2. **`src/app/globals.css`** — CSS variables atualizadas para tema teal:
   - `--primary: 175 84% 32%` (teal-600 `#0D9488`)
   - `--primary-foreground: 0 0% 100%`
   - Modo escuro configurado para componentes do Dashboard.

3. **`src/types/index.ts`** — Tipos TypeScript criados espelhando os modelos da API Go:
   - `User`, `Patient`, `FirstAnalysis`, `Session`, `LoginResponse`.

4. **`src/lib/api.ts`** — Cliente HTTP centralizado:
   - `api.get<T>()`, `api.post<T>()`, `api.put<T>()`, `api.delete()`
   - Base URL lida de `NEXT_PUBLIC_API_URL` com fallback `http://localhost:8080`.

5. **`src/components/layout/Sidebar.tsx`** — Sidebar `"use client"`:
   - Itens: Início, Agenda, Pacientes, Finanças, Configurações + Sair.
   - `usePathname()` para highlight do item ativo com borda teal esquerda.
   - Logo PsicoGestor com símbolo Ψ no topo.

6. **`src/app/(app)/layout.tsx`** — Route group `(app)`:
   - Layout com `Sidebar` à esquerda + `<main>` à direita.
   - Aplica-se a: `/dashboard`, `/patients`, `/admin/settings`, `/admin/users`.

7. **`src/app/page.tsx`** — Tela de Boas-Vindas (`/`):
   - Header escuro com nome/CRP da psicóloga e botão Sair.
   - 3 cards de acesso rápido: Minha Conta & Integrações, Pacientes & Prontuários, Finanças & Relatórios.
   - Fundo teal-50, fiel a `previa2.png`.

8. **`src/app/login/page.tsx`** — Tela de Login (`/login`):
   - Formulário centralizado em card branco sobre gradiente teal.
   - Símbolo da psicologia (`/simbolo-psicologia.jpg`) no topo.
   - Campos Email + Senha com ícones Lucide.
   - Chama `POST /api/auth/login` e redireciona para `/` no sucesso.
   - Fiel a `previa1.png`.

9. **`src/app/(app)/admin/settings/page.tsx`** — Configurações (`/admin/settings`):
   - Layout 2 colunas: dados do psicólogo + integrações.
   - Formulário editável: Nome, CRP, Especialidade, E-mail, Alterar Senha.
   - Cards Google Meet (status Conectado mockado) e Microsoft Outlook.
   - Fiel a `previa3.png`.

10. **`src/app/(app)/patients/page.tsx`** — Ficha do Consultante (`/patients`):
    - Sidebar esquerda com lista de pacientes + busca em tempo real via estado React.
    - Seletor de paciente integrado com `GET /api/patients`.
    - 3 abas implementadas manualmente com Tailwind: Cadastro, Primeira Análise, Evolução das Sessões.
    - Aba Análise: accordion com 14 seções clínicas, consome `GET /api/patients/{id}/analysis`.
    - Aba Evoluções: timeline de sessões com editor de notas, consome `GET /api/patients/{id}/sessions`.
    - Botão "Exportar Ficha (PDF)" aponta para `GET /api/patients/{id}/pdf`.
    - Fiel a `previa4.png`.

11. **`src/app/(app)/dashboard/page.tsx`** — Dashboard Financeiro (`/dashboard`):
    - Tema escuro (slate-900) fiel a `previa5.png`.
    - 4 cards KPI: Faturamento Total, Atendimentos, Pacientes Ativos, A Receber.
    - Gráfico AreaChart com recharts: Faturamento vs Meta (Jan–Jun) com dados mockados.
    - Tabela de transações recentes com badges Pago/Pendente.

12. **`src/app/(app)/admin/users/page.tsx`** — Usuários Admin (`/admin/users`):
    - Tabela carregando usuários de `GET /api/admin/users`.
    - Modal de criação com 5 campos, chama `POST /api/admin/users`.
    - Exclusão inline com `DELETE /api/admin/users/{id}`.

13. **`public/simbolo-psicologia.jpg`** — Copiado de `especification/files/Simbolo-da-psicologia.jpg`.

---

## 2. Logs de Testes e Validação Local

### Comando executado para rodar a aplicação localmente:
```bash
docker compose up --build -d
```

### Checklist de Validação (abrir no navegador após subir os containers):

**Teste 1 — Tela de Login**
```
http://localhost:3000/login
```
- **Esperado**: Card centralizado sobre fundo gradiente teal, logo da psicologia visível no topo, campos E-mail e Senha, botão "Entrar" teal.
- **Resultado**: ✅ OK
  - Card branco sobre gradiente teal: OK
  - Símbolo da Psicologia visível no topo: OK
  - Campo E-mail presente: OK
  - Campo Senha presente: OK
  - Botão "Entrar" teal presente: OK

**Teste 2 — Tela de Boas-Vindas**
```
http://localhost:3000/
```
- **Esperado**: Header escuro com nome da psicóloga, 3 cards de acesso rápido (Conta, Pacientes, Finanças), footer.
- **Resultado**: ✅ OK
  - Header com nome da psicóloga (Dra. Ana): OK
  - Card "Minha Conta & Integrações": OK
  - Card "Pacientes & Prontuários": OK
  - Card "Finanças & Relatórios": OK
  - Botão "Sair": OK

**Teste 3 — Configurações**
```
http://localhost:3000/admin/settings
```
- **Esperado**: Sidebar à esquerda (item "Configurações" ativo), layout 2 colunas com dados do psicólogo e cards de integração Google/Outlook.
- **Resultado**: ✅ OK
  - Sidebar com logo PsicoGestor: OK
  - Item "Configurações" presente na sidebar: OK
  - Item "Agenda" presente na sidebar: OK
  - Dados do psicólogo (campo CRP): OK
  - Dados do psicólogo (campo Especialidade): OK
  - Card Google Meet com status Conectado: OK
  - Card Microsoft Outlook Calendar: OK
  - Botão "Salvar Alterações": OK
  - **Observação**: A rota `/admin/settings` requer autenticação via cookie `auth_token`. O middleware de proteção de rotas foi implementado na Sprint 3 (antecipado em relação ao planejado para Sprint 4).

**Teste 4 — Ficha do Consultante**
```
http://localhost:3000/patients
```
- **Esperado**: Sidebar com lista de pacientes (Ana Souza, Carlos Lima, Roberta Silva carregados da API), abas funcionais, accordion na aba Análise, sessões na aba Evoluções.
- **Resultado**: ✅ OK
  - Sidebar com logo PsicoGestor: OK
  - Paciente "Ana Souza" carregado da API: OK
  - Paciente "Carlos Lima" carregado da API: OK
  - Paciente "Roberta Silva" carregado da API: OK
  - Aba "Cadastro" presente: OK
  - Aba "Primeira Análise" presente: OK
  - Aba "Evoluções das Sessões" presente: OK
  - Botão "Exportar Ficha (PDF)": OK
  - Campo busca (placeholder "Buscar paciente..."): OK — confirmado via interação no Teste 5

**Teste 5 — Busca de Pacientes**
- Digitar "Ana" no campo de busca da lista lateral.
- **Esperado**: Apenas "Ana Souza" listada em tempo real.
- **Resultado**: ✅ OK
  - Campo de busca encontrado: OK
  - "Ana Souza" visível após digitar "Ana": OK
  - "Carlos Lima" oculto após filtro: OK
  - "Roberta Silva" oculta após filtro: OK

**Teste 6 — Dashboard**
```
http://localhost:3000/dashboard
```
- **Esperado**: Tema escuro, 4 KPI cards, gráfico de área animado (Recharts), tabela de transações com badges Pago/Pendente.
- **Resultado**: ✅ OK
  - Sidebar com logo PsicoGestor: OK
  - KPI "Faturamento Total": OK
  - KPI "Atendimentos": OK
  - KPI "Pacientes Ativos": OK
  - KPI "A Receber": OK
  - Badge "Pago" na tabela de transações: OK
  - Badge "Pendente" na tabela de transações: OK
  - Dados do gráfico com meses (Jan, Fev, Mar...): OK

**Teste 7 — Usuários Admin**
```
http://localhost:3000/admin/users
```
- **Esperado**: Tabela com 2 psicólogos carregados da API, botão "+ Novo Psicólogo" abre modal fluido.
- **Resultado**: ❌ NOT OK
  - Sidebar com logo PsicoGestor: OK
  - Botão "Novo Psicólogo" presente: OK
  - Tabela carregada da API: OK
  - **Motivo do NOT OK**: A tabela exibe apenas **1 usuário** ("Administrador Master") em vez dos 2 psicólogos esperados. O handler `GET /api/admin/users` consulta o banco de dados real (não retorna dados mock), e o banco persistido do Docker volume contém apenas o usuário admin criado pelo seed. Nenhum psicólogo foi inserido no banco durante a Sprint 3. Para reproduzir o cenário esperado de "2 psicólogos", é necessário criar usuários via `POST /api/admin/users` ou popular o banco com dados de seed adicionais.

---

## 3. Débitos Técnicos Identificados nesta Sprint

- **Autenticação via cookie implementada antecipadamente**: O login salva o token JWT em cookie `auth_token` (max-age 24h) e o middleware `src/middleware.ts` já protege `/dashboard`, `/patients` e `/admin/*`, redirecionando para `/login` sem o cookie. Isso adianta parte do escopo da Sprint 4.
- **Tabela de usuários admin sem dados mock**: O handler `GET /api/admin/users` usa dados reais do banco. Em ambiente Docker com volume persistido, o banco contém apenas o admin do seed. Para demonstrar a tela com psicólogos pré-cadastrados, é necessário inserir dados manualmente ou adicionar seed de psicólogos de teste.
- **Busca de pacientes**: O filtro de busca opera localmente sobre os dados já carregados (não faz nova chamada à API com `?q=`). Consistente com o comportamento esperado nesta sprint.
- **Gráfico Dashboard com dados estáticos**: Os dados do gráfico (chartData) são mockados no componente. A integração real com `GET /api/dashboard/charts` ocorre na Sprint 7.
- **`package-lock.json` desatualizado**: Como o `recharts` foi adicionado ao `package.json` sem rodar `npm install` localmente, o lock file não foi atualizado. O Dockerfile usa `npm install` (sem `--ci`) para compensar.
