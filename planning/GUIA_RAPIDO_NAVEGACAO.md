# 🎯 GUIA RÁPIDO: DOCUMENTAÇÃO DO PROJETO COMPLETA

## 📚 Como Navegar a Documentação

### **ESPECIFICAÇÃO** (`/especification/`)

| Arquivo | Conteúdo | Usar Para |
|---------|----------|-----------|
| `layout-details.md` | Design system, paleta cores, estrutura de telas (5 previews) | Validação visual frontend |
| `business-rules-details.md` | Modelagem DDL, queries SQL, lógica de negócio | Implementação backend |
| `archicteture-details.md` | Stack (Next.js+Go), Docker, estrutura codebase | Setup infraestrutura dev |
| `infra-devops-details.md` | AWS (VPC, ECS, RDS, CloudFront), CI/CD GitHub Actions | Deploy produção |
| `files/previa*.png` | Wireframes de 5 telas principais | Validação UX/UI |

---

### **PLANEJAMENTO** (`/planning/`)

#### 🗂️ Visão Geral

| Arquivo | Propósito | Leitura |
|---------|-----------|---------|
| `planning.md` | Resumo de 13 fases do projeto (LEIA PRIMEIRO) | 📄 5 min |
| `RESUMO_AJUSTES_COMPLETOS.md` | Sumário executivo dos ajustes realizados | 📄 5 min |
| `MELHORIAS_SUGERIDAS.md` | 18 sugestões priorizadas com matriz | 📄 10 min |

#### 🚀 Sprints (13 Total)

**Padrão de cada sprint:**
- `sprintN.md` - Escopo, objetivos, entregáveis (2-3 páginas)
- `sprintN_tasks.md` - 40-45 tarefas detalhadas para programador
- `sprintN_logs.md` - Registro de execução (preenchido durante sprint)

**Sprints 1-10** (Já existentes):
```
✅ Sprint 1: Setup Inicial (infra dev local)
✅ Sprint 2: Mock Webservice (Swagger)
✅ Sprint 3: Mock Frontend (UI/UX)
✅ Sprint 4: Autenticação (JWT)
✅ Sprint 5: Admin Psicólogo (CRUD)
✅ Sprint 6: Ficha Consultante + PDF Export 📄 NEW
✅ Sprint 7: Dashboard Financeiro
✅ Sprint 8: Calendário + E-mail 📧 NEW
✅ Sprint 9: Polimento (validações, UX)
✅ Sprint 10: Hardening (segurança, cookies)
```

**Sprints 11-13** (Recém criadas):
```
✅ Sprint 11: Deploy AWS (VPC, ECS, RDS, CloudFront)
✅ Sprint 12: Validação Produção (testes, otimizações)
✅ Sprint 13: Testes Finais (regressão, docs)
```

---

## 🎓 FLUXO DE LEITURA RECOMENDADO

### 📖 Para o **PROGRAMADOR** (vai implementar):

1. Leia: `planning/planning.md` (entenda as 4 fases)
2. Leia: `planning/sprint1.md` (conheça a primeira sprint)
3. Leia: `planning/sprint1_tasks.md` (veja as 40+ tarefas)
4. Leia: `especification/archicteture-details.md` (setup inicial)
5. Leia: `planning/MELHORIAS_SUGERIDAS.md` (contexto de requirements)
6. **Comece Sprint 1!**

---

### 🎯 Para o **PLANEJADOR** (vai revisar):

1. Leia: `planning/RESUMO_AJUSTES_COMPLETOS.md` (visão geral)
2. Leia: `planning/planning.md` (valide estrutura de 13 sprints)
3. Leia: `planning/MELHORIAS_SUGERIDAS.md` (decide prioridade)
4. Verifique: `planning/sprint*.md` (valide checkpoints)
5. Comunique ao time!

---

### 👥 Para **STAKEHOLDERS**:

1. Leia: `planning/planning.md` (resumo executivo)
2. Leia: `planning/MELHORIAS_SUGERIDAS.md` (se há impacto em requirements)
3. Veja: `especification/files/previa*.png` (visualização de resultado final)
4. Pronto para discutir timeline e riscos!

---

### 🔒 Para o **REVISOR DE CÓDIGO** (code review final):

1. Leia: `planning/MELHORIAS_SUGERIDAS.md` (checklist de validação)
2. Leia: `especification/business-rules-details.md` (regras de negócio)
3. Leia: `especification/archicteture-details.md` (padrões técnicos)
4. Compare implementação com specs detalhados

---

## 📊 ORGANIZAÇÃO POR ÁREA

### **Frontend (Next.js + Shadcn/UI)**
- Especificação: `layout-details.md`
- Tasks: `sprint3_tasks.md` (UI), `sprint4_tasks.md` (Auth), etc.
- Validação: `previa*.png` (comparação visual)

### **Backend (Go + Chi Router)**
- Especificação: `business-rules-details.md`
- Tasks: `sprint2_tasks.md` (API), `sprint4_tasks.md` (Auth), etc.
- Documentação: Swagger (`localhost:8080/swagger`)

### **Banco de Dados (SQLite Dev / MySQL Prod)**
- Especificação: `business-rules-details.md` (DDL)
- Tasks: `sprint4_tasks.md` (tabela users), `sprint6_tasks.md` (patients), etc.
- Migração: `sprint11_tasks.md` (SQLite → MySQL)

### **DevOps / Infraestrutura (AWS)**
- Especificação: `infra-devops-details.md`
- Tasks: `sprint1_tasks.md` (Docker local), `sprint11_tasks.md` (AWS)
- CI/CD: `.github/workflows/deploy.yml`

### **QA / Validação**
- E2E: `sprint12_tasks.md` (Cypress)
- Carga: `sprint12_tasks.md` (k6)
- Segurança: `sprint12_tasks.md` (OWASP)
- Regressão: `sprint13_tasks.md`

---

## ⚡ QUICK LINKS

### 📝 Arquivos Críticos (Ler Primeiro)

```
planning/planning.md                    ← Visão geral (5 min)
planning/RESUMO_AJUSTES_COMPLETOS.md   ← O que foi feito (5 min)
planning/MELHORIAS_SUGERIDAS.md         ← Recomendações (10 min)
```

### 🚀 Sprint Inicial (Sprint 1)

```
planning/sprint1.md                     ← Objetivos e entregáveis
planning/sprint1_tasks.md               ← 16 tarefas detalhadas
especification/archicteture-details.md  ← Stack e setup
```

### 🎨 Design e UI (Sprint 3)

```
especification/layout-details.md        ← Guia de estilo + 5 telas
especification/files/previa*.png        ← Wireframes para comparação
```

### 🔒 Segurança e Auth (Sprint 4)

```
especification/business-rules-details.md ← DDL users, queries
planning/sprint4_tasks.md                ← JWT, login, proteção
```

### 📊 Dashboard (Sprint 7)

```
especification/business-rules-details.md ← Queries de agregação
planning/sprint7_tasks.md                ← Implementação KPIs e gráficos
```

### ☁️ AWS Deploy (Sprint 11)

```
especification/infra-devops-details.md  ← Arquitetura AWS completa
planning/sprint11.md                     ← Escopo detalhado
planning/sprint11_tasks.md               ← 43 tarefas passo-a-passo
```

---

## 🔄 CICLO DE VIDA DE UMA SPRINT

Para cada sprint `N`:

1. **Planejador** revisa: `sprintN.md` (entregáveis e checkpoints)
2. **Programador** revisa: `sprintN_tasks.md` (tarefas detalhadas)
3. **Programador** executa: Implementa conforme tarefas
4. **Programador** documenta: Preenche `sprintN_logs.md`
5. **QA** valida: Executa checkpoints de validação humana
6. **Planejador** aprova: Valida entregáveis
7. **Git**: Commit + Tag + Push (ex: `v0.1.0`)

---

## 🎯 CHECKLIST DE VALIDAÇÃO FINAL

- [ ] Li `planning/planning.md` (entendo as 4 fases)
- [ ] Li `planning/sprint1.md` (pronto para começar)
- [ ] Revisei `planning/MELHORIAS_SUGERIDAS.md` (entendo requirements extras)
- [ ] Entendo que há 13 sprints total (não 10)
- [ ] Tenho acesso a todos os arquivos em `/especification/`
- [ ] Tenho acesso a todos os wireframes em `/especification/files/`
- [ ] Entendo que PDF e E-mail são críticos (sugestões 1 e 2)
- [ ] Entendo que testes são transversais (não apenas S12-S13)
- [ ] Confirmei que `planning.md` foi corrigido (não há S8 duplicada)
- [ ] Pronto para iniciar desenvolvimento!

---

## 🆘 DÚVIDAS FREQUENTES

**P: Por onde começo?**  
R: Comece por `planning/planning.md` para entender o big picture.

**P: Onde estão os wireframes?**  
R: Em `especification/files/previa1.png` até `previa5.png`.

**P: Como somar detalhes de uma tarefa?**  
R: Abra `planning/sprintN_tasks.md` para a sprint correspondente.

**P: Onde está a especificação de segurança?**  
R: Em `especification/archicteture-details.md` (seção 4 - Segurança).

**P: Quantas sprints são ao todo?**  
R: 13 sprints (era 10 + adicionadas 3).

**P: Há sugestões de melhoria?**  
R: Sim, veja `planning/MELHORIAS_SUGERIDAS.md` com 18 recomendações.

**P: Quando é o deploy AWS?**  
R: Sprint 11 (após Sprint 10 pronto).

---

## 📞 Suporte

Se tiver dúvidas sobre:
- **Especificação**: Consulte arquivos em `/especification/`
- **Planejamento**: Consulte `planning.md`
- **Tarefas de Sprint**: Consulte `sprintN_tasks.md`
- **Sugestões de Melhoria**: Consulte `MELHORIAS_SUGERIDAS.md`
- **Resumo de Ajustes**: Consulte `RESUMO_AJUSTES_COMPLETOS.md`

---

**Última Atualização**: 27 de Junho, 2026  
**Status**: ✅ Documentação Completa e Pronta para Desenvolvimento
