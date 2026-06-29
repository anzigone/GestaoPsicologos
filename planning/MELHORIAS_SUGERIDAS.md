# 📋 Sugestões de Melhoria para as Sprints Existentes

## Resumo Executivo

Este documento consolida todas as sugestões de melhorias, ajustes e optimizações para as Sprints 1-10, identificadas durante a revisão completa da especificação e plano. Cada sugestão é prioritizada e acompanhada de justificativa.

---

## 🔴 CRÍTICAS (Implementar Obrigatoriamente)

### Sprint 6: Ficha Consultante - Adicionar Exportação em PDF

**Problema**: O business-rules-details.md especifica "Exportar Ficha Completa (PDF)" na Aba 1, mas não há referência clara em sprint6.md.

**Solução Proposta**:
- Adicionar tarefa T6.X: "Implementar geração de PDF da ficha do paciente"
- Endpoint Backend: `GET /api/patients/{id}/export-pdf`
- Incluir no PDF:
  - Dados cadastrais (nome, telefone, profissão, etc.)
  - Primeira Análise completa
  - Timeline de sessões (últimas 10) com datas, notas e status
- Usar biblioteca Go: `gofpdf` ou `fpdf`
- Frontend: Botão "Exportar Ficha (PDF)" no header da ficha do paciente
- Validação humana: Baixar PDF e verificar formatação

**Impacto**: Alto | **Esforço**: Médio | **Prioridade**: ⭐⭐⭐

---

### Sprint 8: Integração com Calendário - Adicionar Envio de E-mail

**Problema**: business-rules-details.md menciona "O webservice envia um e-mail... contendo o convite e o link", mas Sprint 8 não especifica isso.

**Solução Proposta**:
- Adicionar tarefa T8.X: "Implementar envio de e-mail de convite"
- Quando criar sessão com Google Meet ou Outlook:
  - Gerar e-mail HTML com:
    - Nome do psicólogo
    - Data e hora da sessão
    - Link do Google Meet (se disponível)
    - Informações de contato do psicólogo
  - Enviar para:
    - Psicólogo (confirmação)
    - Paciente (convite com link da reunião)
- Usar AWS SES (Simple Email Service) em produção
- Usar nodemailer ou similar em desenvolvimento (com mock)
- Frontend: Checkbox "Enviar e-mail de convite" (padrão checked)
- Validação humana: Criar sessão e validar e-mail recebido

**Impacto**: Alto | **Esforço**: Médio | **Prioridade**: ⭐⭐⭐

---

## 🟠 IMPORTANTES (Recomendado)

### Sprints 1-10: Testes Automatizados como Atividade Transversal

**Problema**: Nenhuma sprint especifica estratégia de testes unitários, de integração ou E2E durante o desenvolvimento. Testes aparecem apenas em Sprint 12 (fim do projeto).

**Solução Proposta**:
- Cada sprint deve incluir tarefas de testes:
  - Sprint 1: Setup de infrastructure de testes (Jest, Vitest para frontend; Go testing para backend)
  - Sprint 2: Testes unitários dos handlers mockados
  - Sprint 3: Testes de componentes Shadcn/UI
  - Sprint 4: Testes de autenticação (JWT, login/logout)
  - Sprint 5: Testes de perfil do psicólogo (CRUD)
  - Sprint 6: Testes de pacientes e anamnese
  - Sprint 7: Testes de dashboard e agregações
  - Sprint 8: Testes de integração OAuth2 (simulados)
  - Sprint 9: Testes E2E com Cypress (smoke tests)
  - Sprint 10: Testes de segurança (HttpOnly cookies, CSRF)
- Adicionar na sprint_tasks.md de cada uma:
  - Exemplo para Sprint 4: `T4.X Escrever testes unitários de autenticação`
- Validação humana em cada sprint: Rodar `npm test` ou `go test ./...` com cobertura > 80%

**Impacto**: Alto | **Esforço**: Alto (distribuído) | **Prioridade**: ⭐⭐⭐

---

### Sprints 2-10: Documentação de API em Swagger Mantida Atualizada

**Problema**: Sprint 2 menciona Swagger, mas não especifica responsabilidade de manutenção nas sprints subsequentes.

**Solução Proposta**:
- Cada sprint que adiciona endpoints deve incluir:
  - Tarefa: "Atualizar documentação Swagger"
  - Adicionar comentários Swagger nos handlers Go (`// @Router /api/... [GET]`)
  - Executar `swag init ./cmd/server` para regenerar docs
  - Validação: Acessar `http://localhost:8080/swagger/index.html` e verificar novos endpoints
- Exemplo para Sprint 4:
  - T4.Y: "Documentar endpoints de autenticação em Swagger"
  - T4.Z: "Documentar endpoints de gestão de usuários em Swagger"

**Impacto**: Médio | **Esforço**: Baixo | **Prioridade**: ⭐⭐

---

### Sprint 3: Mock Frontend - Adicionar Validação Visual com Wireframes

**Problema**: Sprint 3 menciona comparação com wireframes (verificação Comparativa), mas não formaliza critério de aceitação.

**Solução Proposta**:
- Adicionar checklist de validação visual:
  ```
  [ ] Login (previa1.png): Posicionamento, cores, tamanho do logotipo coincidem?
  [ ] Dashboard Inicial (previa2.png): 3 cards interativos com ícones corretos?
  [ ] Admin Settings (previa3.png): Layout 2 colunas, campos corretos?
  [ ] Ficha Paciente (previa4.png): Mestre-detalhe, abas visíveis?
  [ ] Dashboard Financeiro (previa5.png): KPIs, gráficos presentes?
  ```
- Frontend dev: Comparar lado-a-lado com wireframes
- Checkpoint humano: "As telas estão 95%+ fiéis aos wireframes? Ajustes cosmética são aceitáveis."

**Impacto**: Médio | **Esforço**: Baixo | **Prioridade**: ⭐⭐

---

### Sprint 9: Polimento - Adicionar Internacionalização (i18n) Base

**Problema**: Projeto está 100% em PT-BR, mas sem infraestrutura de i18n, dificultará futuras expansões para EN.

**Solução Proposta**:
- Adicionar tarefa T9.X: "Implementar infraestrutura base de i18n"
- Frontend: Instalar `next-i18next` ou `i18next`
- Criar arquivo `public/locales/pt-BR/common.json` com todas as strings
- Converter hardcoded strings para uso de `t('key')`
- Backend: Utilizar padrão simples de enums para mensagens
- Validação: Aplicação continua 100% em PT-BR, mas estrutura pronta para EN/ES

**Impacto**: Baixo (futuro) | **Esforço**: Médio | **Prioridade**: ⭐ (para sprint futura, não obrigatório em S9)

---

## 🟡 SUGESTÕES (Considerar)

### Sprint 10: Hardening - Adicionar Rate Limiting

**Problema**: Não há menção a proteção contra brute force ou DDoS aplicacional.

**Solução Proposta**:
- Adicionar tarefa T10.X: "Implementar rate limiting em endpoints sensíveis"
- Backend Go: Usar `github.com/go-chi/chi/middleware.ThrottleBacklog` ou similar
- Endpoints com rate limit:
  - `POST /api/auth/login`: Máx 5 tentativas por minuto por IP
  - `POST /api/admin/users`: Máx 10 por minuto por usuário admin
  - `GET /api/patients`: Máx 100 por minuto por usuário
- Frontend: Desabilitar botão por 3-5 segundos após múltiplas tentativas
- Validação: Fazer múltiplas requisições e verificar 429 (Too Many Requests)

**Impacto**: Médio | **Esforço**: Baixo | **Prioridade**: ⭐⭐

---

### Sprint 10: Hardening - Adicionar Monitoramento de Sessão

**Problema**: JWT expira em 24h, mas não há mecanismo para detectar sessões comprometidas ou logout em tempo real.

**Solução Proposta**:
- Adicionar tabela `active_sessions` (já no DDL do architecture-details.md)
- Quando login bem-sucedido: Registrar sessão com expiração
- Quando logout: Marcar sessão como inativa
- Middleware: Validar que JWT ainda está em `active_sessions`
- Bônus: Permitir logout em todos os dispositivos (invalidar todas as sessões do usuário)
- Validação: Fazer logout em um device, verificar que outro device é forçado a fazer login novamente

**Impacto**: Médio | **Esforço**: Médio | **Prioridade**: ⭐⭐

---

### Sprint 7: Dashboard - Adicionar Gráfico de Evolução de Pacientes

**Problema**: Dashboard mostra faturamento e volume, mas não mostra evolução clínica dos pacientes.

**Solução Proposta**:
- Adicionar widget "Pacientes por Status Clínico"
- Mostrar:
  - Pacientes em tratamento inicial
  - Pacientes em progressão (> 5 sessões)
  - Pacientes em alta clínica
- Gráfico: Pizza ou anel (donut)
- Validação: Clicar em cada seção para filtrar tabela de pacientes

**Impacto**: Baixo (nice-to-have) | **Esforço**: Médio | **Prioridade**: ⭐

---

### Sprint 5: Administração do Psicólogo - Adicionar Configuração de Horários de Atendimento

**Problema**: Não há modo de configurar horários de disponibilidade do psicólogo (por exemplo: seg-sex 09:00-18:00).

**Solução Proposta**:
- Adicionar campos na tela `/admin/settings`:
  - Dias de atendimento (checkboxes: seg, ter, qua, qui, sex, sab, dom)
  - Horário de início e fim (time inputs)
  - Intervalo entre sessões (ex: 50 minutos para 60 min de sessão)
- Backend: Armazenar em tabela `psychologist_availability`
- Validação: Usar para sugerir horários ao criar sessão
- Checkpoint: Campo preenchido com padrão seg-sex 09:00-18:00

**Impacto**: Médio | **Esforço**: Médio | **Prioridade**: ⭐

---

## 🟢 OTIMIZAÇÕES (Implementar depois)

### Sprint 2: Mock Webservice - Adicionar Versioning de API

**Sugestão**: Adicionar `/api/v1/` como prefixo de versioning desde o início facilita manutenção futura.

**Implementação**: Usar `chi.NewRouter().Route("/api/v1", func(r chi.Router) { ... })`

**Benefício**: Pronto para `/api/v2/` no futuro se houver breaking changes

---

### Sprint 3: Mock Frontend - Adicionar Storybook para Componentes

**Sugestão**: Documentar componentes Shadcn/UI customizados com Storybook.

**Benefício**: Facilita onboarding de novos devs e testes visuais

---

### Sprint 4: Autenticação - Adicionar Recuperação de Senha

**Sugestão**: Implementar fluxo de "Esqueci minha senha" com e-mail de reset.

**Benefício**: Evita bloqueio de usuário com senha perdida

---

### Sprint 7: Dashboard - Adicionar Exportação de Relatório

**Sugestão**: Botão "Exportar Relatório" que gera PDF com:
- Período selecionado (mês, trimestre, ano)
- KPIs consolidados
- Gráficos (faturamento, volume, pacientes)
- Tabela de transações

**Benefício**: Psicólogo pode compartilhar relatório com contador/gestor

---

## 📊 Matriz de Priorização

| Sugestão | Categoria | Prioridade | Impacto | Esforço | Sprint Ideal |
|----------|-----------|------------|--------|--------|------------|
| PDF Exportação | Crítica | ⭐⭐⭐ | Alto | Médio | 6 |
| E-mail Convites | Crítica | ⭐⭐⭐ | Alto | Médio | 8 |
| Testes Transversais | Importante | ⭐⭐⭐ | Alto | Alto | 1-10 |
| Swagger Maintenance | Importante | ⭐⭐ | Médio | Baixo | 2-10 |
| Wireframe Validation | Importante | ⭐⭐ | Médio | Baixo | 3 |
| Rate Limiting | Sugestão | ⭐⭐ | Médio | Baixo | 10 |
| Monitoramento Sessão | Sugestão | ⭐⭐ | Médio | Médio | 10 |
| Gráfico Evolução | Sugestão | ⭐ | Baixo | Médio | 7 |
| Horários Atendimento | Sugestão | ⭐ | Médio | Médio | 5 |
| i18n Base | Otimização | ⭐ | Baixo | Médio | 9 (futura) |
| API Versioning | Otimização | ⭐ | Baixo | Baixo | 2 |
| Storybook | Otimização | ⭐ | Baixo | Médio | 3 |
| Recuperação Senha | Otimização | ⭐ | Médio | Médio | 4 |
| Relatório Exportável | Otimização | ⭐ | Médio | Médio | 7 |

---

## ✅ Recomendação Final

### Implementar Obrigatoriamente (Antes de Produção):

1. **PDF Exportação (Sprint 6)**
2. **E-mail Convites (Sprint 8)**
3. **Testes Transversais (Sprints 1-10)**
4. **Rate Limiting (Sprint 10)**

### Implementar Recomendado (Se houver tempo):

5. **Monitoramento de Sessão (Sprint 10)**
6. **Recuperação de Senha (Sprint 4 revisitado ou 9)**
7. **Horários de Atendimento (Sprint 5 revisitado)**

### Implementar em Sprints Futuras (Pós-Produção):

8. Gráfico de Evolução de Pacientes
9. Exportação de Relatório
10. Internacionalização (i18n)
11. Storybook

---

## 📝 Como Usar Este Documento

1. **Planejador**: Revisar sugestões **Críticas** e **Importantes** e atualizar sprint_tasks.md correspondentes
2. **Programador**: Implementar conforme documentado em cada tarefa
3. **QA**: Validar conforme checkpoints especificados
4. **Gestor de Projeto**: Priorizar conforme matriz, ajustando timeline se necessário

---

## 🔗 Documentos Relacionados

- `especification/business-rules-details.md` (para detalhes de funcionalidades)
- `especification/layout-details.md` (para validação visual)
- `planning/sprint*.md` (para contexto de cada sprint)
- `planning/sprint*_tasks.md` (para tarefas detalhadas)
