# Registro de Execução da Sprint 6 (sprint6_logs.md)

Este documento registra cronologicamente todas as ações realizadas pelo Programador durante a Sprint 6, servindo de trilha de auditoria e documentação técnica da implantação.

---

## 1. Histórico de Atividades e Comandos

- **Data**: 29/06/2026
- **Status da Sprint**: ✅ Concluída e validada pelo humano em 29/06/2026

### Passos Executados:

1. **Migração das tabelas** (`codebase/backend/internal/database/database.go`):
   - Adicionadas tabelas `patients` e `first_analysis` à função `Migrate()` no slice `statements`.
   - `patients`: campos id, psychologist_id, name, phone, birthdate, age, profession, company, city, state, marital_status, consultation_fee, created_at, updated_at. FK `ON DELETE CASCADE` referenciando `users(id)`.
   - `first_analysis`: campos patient_id (PK), 14 campos clínicos em TEXT (main_complaint, symptom_diagnosis, developmental_influence, situational_issues, biological_factors, strengths_resources, addictions, stimuli, thoughts, behaviors, affects, physiological, treatment_goals, treatment_plan), updated_at. FK `ON DELETE CASCADE` referenciando `patients(id)`.

2. **Handlers de pacientes** (`codebase/backend/internal/handlers/patients.go`):
   - Criado `CreatePatientRequest` com todos os campos cadastrais e anotações Swagger.
   - Helper `scanPatient()` e constante `querySelectPatient` para evitar duplicação de scan entre GET e operações de escrita.
   - `ListPatients`: query `WHERE psychologist_id=? AND name LIKE ?` com `ORDER BY name ASC`. Parâmetro `q` transformado em padrão LIKE com curingas.
   - `GetPatient`: isolamento por `psychologist_id` na query de busca por ID.
   - `CreatePatient`: gera UUID via `auth.NewUUID()`, insere todos os campos, retorna paciente criado via `scanPatient`.
   - `UpdatePatient`: UPDATE com cláusula `WHERE id=? AND psychologist_id=?`, verifica `RowsAffected()` para retornar 404 se não encontrado.
   - `DeletePatient`: remoção com isolamento por psychologist_id.
   - `ExportPatientPDF`: retorna PDF mínimo válido como placeholder (implementação real na Sprint 11).

3. **Handlers de anamnese** (`codebase/backend/internal/handlers/analysis.go`):
   - Helper `patientBelongsTo()` verificando ownership antes de qualquer operação de anamnese.
   - `GetAnalysis`: `COALESCE` em todos os campos TEXT, retorna `FirstAnalysis` com `PatientID` preenchido e campos vazios quando não há registro (sem 404).
   - `UpdateAnalysis`: Upsert via `INSERT ... ON CONFLICT(patient_id) DO UPDATE SET ...` — sintaxe nativa SQLite. Atualiza `updated_at` explicitamente.

4. **Rotas registradas** (`codebase/backend/cmd/server/main.go`):
   ```
   r.Route("/api/patients", func(r chi.Router) {
       r.Get("/", handlers.ListPatients(db))
       r.Post("/", handlers.CreatePatient(db))
       r.Route("/{id}", func(r chi.Router) {
           r.Get("/", handlers.GetPatient(db))
           r.Put("/", handlers.UpdatePatient(db))
           r.Delete("/", handlers.DeletePatient(db))
           r.Get("/pdf", handlers.ExportPatientPDF())
           r.Get("/analysis", handlers.GetAnalysis(db))
           r.Put("/analysis", handlers.UpdateAnalysis(db))
           r.Route("/sessions", ...)
       })
   })
   ```

5. **Frontend** (`codebase/frontend/src/app/(app)/patients/page.tsx`):
   - Layout mestre-detalhe com sidebar (w-64) + painel principal (flex-1).
   - `baseFee` carregado de `GET /api/psychologist` para cálculo de variação.
   - Busca com debounce de 300ms via `useRef<ReturnType<typeof setTimeout>>`.
   - Modal "+ Novo Paciente" com formulário de 3 campos e preview de variação tarifária em tempo real.
   - Aba "Cadastro": formulário completo com 10 campos, variação tarifária colorida (teal/laranja), PUT ao salvar.
   - Aba "Primeira Análise": 14 accordions com `ANALYSIS_SECTIONS`, Upsert via PUT ao salvar.
   - Aba "Evolução das Sessões": integração mock com dados de `GET /api/patients/{id}/sessions` (implementação real na Sprint 7).
   - Toast de sucesso/erro com ícone e auto-dismiss em 3s.

---

## 2. Logs de Testes e Validação Local

### Comando executado para rodar a aplicação localmente:
```bash
docker compose up --build -d
```

### Resultados obtidos:
- **Criação de tabelas patients e first_analysis**: ✅ DDL adicionado à migração automática — tabelas criadas na inicialização do backend
- **Listagem e busca textual de pacientes na Sidebar**: ✅ `GET /api/patients?q=` com debounce de 300ms operacional
- **Isolamento por psicólogo**: ✅ Todos os endpoints filtram por `psychologist_id` extraído do JWT
- **Variação tarifária calculando de forma dinâmica**: ✅ Funções `feeVariation()` e `feeVariationColor()` calculam em tempo real no frontend
- **Anamnese salvando em Upsert no SQLite**: ✅ `INSERT ... ON CONFLICT ... DO UPDATE` sem duplicação de registros
- **Modal de criação de paciente**: ✅ POST + atualização da lista local + seleção automática do novo paciente

### Validação humana concluída (29/06/2026):
- ✅ Criação de paciente "Carlos Drummond" via modal — aparece na lista lateral
- ✅ Variação tarifária "-10% abaixo do valor padrão" exibida corretamente (R$ 135,00 vs R$ 150,00)
- ✅ Dados clínicos da Primeira Análise persistem após F5 (Upsert SQLite confirmado)
- ✅ Sprint encerrada com commit `feat: sprint 6 ficha consultante concluido`, tag `v0.6.0` e push para main

---

## 3. Débitos Técnicos Identificados nesta Sprint
- **DT-S6-01**: Aba "Evolução das Sessões" carrega dados do endpoint mock `GET /api/patients/{id}/sessions`. Botão "Nova Sessão" e "Salvar Alterações" ainda não funcionais. Implementação real está prevista para a Sprint 7.
- **DT-S6-02**: `ExportPatientPDF` retorna PDF mínimo válido como placeholder. Geração real do prontuário está prevista para a Sprint 11.
- **DT-S6-03**: Validação de browser não realizada via automação (extensão Chrome indisponível durante implementação). Testar manualmente navegando em `/patients` após login.
