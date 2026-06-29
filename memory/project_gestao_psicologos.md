---
name: project-gestao-psicologos
description: Contexto do projeto Gestão Psicólogos — estado atual, stack, e próximos passos
metadata:
  type: project
---

App web para gestão de pacientes, atendimentos e financeiro para psicólogos.

**Stack**: Next.js 14 (TypeScript, TailwindCSS, Shadcn/UI) + Go 1.22 (chi, SQLite/MySQL híbrido) + Docker + AWS ECS Fargate

**Estado (27/06/2026)**:
- Especificação: 100% completa (`especification/`)
- Planejamento: 13 sprints documentadas (`planning/`)
- Sprint 1 (Setup): ✅ Concluída — codebase criado, pronto para validação humana
- Sprints 2-13: Não iniciadas

**Checkpoint Sprint 1** (validação humana pendente):
```bash
docker compose up --build -d
# http://localhost:8080/api/health → {"status":"OK","database":"connected"}
# http://localhost:3000 → tela Next.js com Tailwind
```

**Why:** Após validação, iniciar Sprint 2 (Mock Webservice com Swagger).
**How to apply:** Aguardar confirmação do usuário antes de iniciar Sprint 2.
