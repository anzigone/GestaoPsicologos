# Sprint 12: Validação de Produção - Tarefas Detalhadas (sprint12_tasks.md)

## Overview
Sprint de testes, validação, otimização de performance e conformidade em ambiente de produção AWS. Foco em qualidade e prontidão operacional, não em desenvolvimento de features.

---

## Tarefas (QA/DevOps/Programador)

### Área: Testes E2E Automatizados

- [ ] **T12.1** Preparar suite de testes E2E com Cypress
  - Instalar Cypress: `npm install -D cypress`
  - Criar pasta `cypress/e2e/`
  - Configurar `cypress.config.ts` com baseUrl da produção

- [ ] **T12.2** Implementar testes de autenticação
  - [ ] Teste: Login com admin master
  - [ ] Teste: Login com psicólogo
  - [ ] Teste: Logout
  - [ ] Teste: Acesso negado sem autenticação
  - [ ] Esperado: Todos passam

- [ ] **T12.3** Implementar testes de CRUD psicólogo
  - [ ] Teste: Criar psicólogo (admin)
  - [ ] Teste: Editar dados do psicólogo
  - [ ] Teste: Trocar senha
  - [ ] Teste: Deletar psicólogo (admin)
  - [ ] Esperado: Todos passam

- [ ] **T12.4** Implementar testes de CRUD paciente
  - [ ] Teste: Criar paciente
  - [ ] Teste: Buscar paciente por nome
  - [ ] Teste: Editar dados do paciente
  - [ ] Teste: Calcular variação percentual de tarifa
  - [ ] Teste: Deletar paciente
  - [ ] Esperado: Todos passam

- [ ] **T12.5** Implementar testes de sessões
  - [ ] Teste: Criar sessão
  - [ ] Teste: Editar notas de sessão
  - [ ] Teste: Registrar sessão como pago/pendente
  - [ ] Teste: Deletar sessão
  - [ ] Esperado: Todos passam

- [ ] **T12.6** Implementar testes de anamnese
  - [ ] Teste: Preencher primeira análise
  - [ ] Teste: Salvar anamnese
  - [ ] Teste: Carregar dados de anamnese
  - [ ] Teste: Editar seção de anamnese
  - [ ] Esperado: Todos passam

- [ ] **T12.7** Implementar testes de dashboard
  - [ ] Teste: Carregar dashboard
  - [ ] Teste: Validar KPIs (faturamento, atendimentos, pacientes)
  - [ ] Teste: Filtrar transações por status
  - [ ] Teste: Gráficos renderizam sem erros
  - [ ] Esperado: Todos passam

- [ ] **T12.8** Implementar testes de integração OAuth
  - [ ] Teste: Status de conexão Google (simulado)
  - [ ] Teste: Status de conexão Outlook (simulado)
  - [ ] Teste: Geração de link Google Meet (simulado)
  - [ ] Esperado: Todos passam

- [ ] **T12.9** Executar suite E2E completa
  - Rodar: `npx cypress run`
  - Gerar relatório HTML
  - Documentar resultados

### Área: Testes de Carga e Performance

- [ ] **T12.10** Preparar scripts de teste de carga com k6
  - Instalar k6 ou usar plataforma cloud
  - Criar script `loadtest.js` com cenários:
    - 100 requisições simultâneas de login
    - 50 psicólogos criando pacientes
    - Leitura massiva de dashboard
    - Busca e filtro de pacientes

- [ ] **T12.11** Executar teste de carga
  - Executar: `k6 run loadtest.js`
  - Coletar métricas:
    - Latência P50, P95, P99
    - Taxa de erro
    - Throughput (req/sec)
    - Consumo de CPU/Memória

- [ ] **T12.12** Analisar resultados de carga
  - Latência P99 < 5 segundos? ✓
  - Taxa de erro < 1%? ✓
  - ECS auto-scaling funcionou? ✓
  - Documentar gargalos encontrados

- [ ] **T12.13** Teste de estresse
  - Aumentar carga para 500 requisições simultâneas
  - Verificar comportamento sob extremo estresse
  - Documentar ponto de quebra

### Área: Testes de Segurança

- [ ] **T12.14** Executar verificação OWASP Top 10
  - [ ] A01: Injection (SQL, NoSQL) - Validado? ✓
  - [ ] A02: Broken Authentication - Validado? ✓
  - [ ] A03: Sensitive Data Exposure - Validado? ✓
  - [ ] A04: XXE - Validado? ✓
  - [ ] A05: Broken Access Control - Validado? ✓
  - [ ] A06: Security Misconfiguration - Validado? ✓
  - [ ] A07: XSS - Validado? ✓
  - [ ] A08: Insecure Deserialization - Validado? ✓
  - [ ] A09: Using Components with Known Vulnerabilities - Validado? ✓
  - [ ] A10: Insufficient Logging & Monitoring - Validado? ✓

- [ ] **T12.15** Executar dependency scanning
  - Backend: `go list -json -m all | nancy sleuth`
  - Frontend: `npm audit`
  - Documentar vulnerabilidades encontradas

- [ ] **T12.16** Validar certificado SSL/TLS
  - Usar `ssl-test.com` ou `testssl.sh`
  - Verificar:
    - Protocolo TLS 1.2+ apenas
    - Cipher suites fortes
    - HSTS header presente
    - CAA records corretos

- [ ] **T12.17** Verificar headers de segurança
  - Validar presença de:
    - `Strict-Transport-Security`
    - `X-Content-Type-Options: nosniff`
    - `X-Frame-Options: DENY`
    - `Content-Security-Policy` (se aplicável)

### Área: Testes de Conformidade (LGPD)

- [ ] **T12.18** Validar criptografia em trânsito
  - Todos os endpoints usam HTTPS? ✓
  - Certificado válido e não expirado? ✓

- [ ] **T12.19** Validar criptografia em repouso
  - Senhas com bcrypt? ✓
  - Tokens seguros (HttpOnly, SameSite)? ✓
  - Dados sensíveis não em logs? ✓

- [ ] **T12.20** Testar isolamento de dados (Multi-tenancy)
  - Psicólogo A consegue ver dados de Psicólogo B? Não ✓
  - Paciente A consegue ver dados de Paciente B (outro psicólogo)? Não ✓

- [ ] **T12.21** Testar direito ao esquecimento (LGPD)
  - Deletar usuário psicólogo
  - Validar cascade delete em pacientes, sessões
  - Validar que dados foram realmente removidos do banco

### Área: Testes de Disaster Recovery

- [ ] **T12.22** Simular falha de container ECS
  - Pausar manualmente um container do ECS
  - Verificar se Auto Scaling cria novo container
  - Aplicação continua responsiva? ✓
  - Tempo para recuperação? (Documentar)

- [ ] **T12.23** Testar restore de RDS
  - Criar snapshot manual do RDS
  - Restaurar em instância de teste
  - Validar integridade dos dados
  - Documentar tempo de restore

- [ ] **T12.24** Testar rollback de deployment
  - Fazer deploy de versão anterior via GitHub Actions
  - Verificar se aplicação volta a funcionar
  - Dados não foram perdidos? ✓

### Área: Otimizações de Performance

- [ ] **T12.25** Analisar queries lentas
  - Ativar query logging no RDS
  - Identificar queries com tempo > 1 segundo
  - Executar EXPLAIN PLAN

- [ ] **T12.26** Criar índices otimizados
  - Índice: `patients(psychologist_id, created_at DESC)`
  - Índice: `sessions(patient_id, session_date DESC)`
  - Índice: `sessions(status)` para filtros de dashboard
  - Validar redução de tempo de query

- [ ] **T12.27** Otimizar cache no CloudFront
  - Assets estáticos: TTL = 86400 (1 dia)
  - HTML: TTL = 300 (5 minutos)
  - `/api/*`: Sem cache (ou TTL curto)
  - Validar cache hit rate > 80% em monitoring

- [ ] **T12.28** Validar compressão Gzip/Brotli
  - Curl com header Accept-Encoding
  - Verificar se Content-Encoding: gzip retornado
  - Redução de tamanho > 60%? ✓

- [ ] **T12.29** Validar lazy loading no frontend
  - Components carregam on-demand? ✓
  - Imagens com lazy loading? ✓
  - First Contentful Paint (FCP) < 3 segundos? ✓

### Área: Otimizações de Custo (FinOps)

- [ ] **T12.30** Analisar AWS Cost Explorer
  - EC2: $0 (usando Fargate)
  - RDS: ~$15-20/mês (db.t3.micro)
  - NAT Gateway: ~$32/mês (se não otimizado, considerar alternativa)
  - CloudFront: Variável (depende de tráfego)
  - Total esperado: ~$50-100/mês

- [ ] **T12.31** Implementar otimizações de custo
  - [ ] Reduzir retenção de logs CloudWatch: 7 dias
  - [ ] Usar VPC Endpoints para AWS services (economizar NAT)
  - [ ] Validar que ECS está usando Free Tier Fargate

- [ ] **T12.32** Configurar AWS Budgets
  - Limite de custo: $100/mês
  - Alerta em 80%: Email para admin
  - Alerta em 100%: Email crítico

### Área: Documentação Operacional

- [ ] **T12.33** Criar Runbook de Deployment
  - Procedimento passo-a-passo
  - Validações pré-deploy
  - Monitoramento pós-deploy
  - Rollback se necessário

- [ ] **T12.34** Criar Runbook de Rollback
  - Como identificar versão anterior
  - Como fazer rollback via GitHub Actions
  - Como validar sucesso de rollback

- [ ] **T12.35** Criar Runbook de Incident Response
  - Aplicação lenta → checar CPU/Memória → escalar manual
  - Erro de banco → verificar RDS status → restore backup
  - Erro 5xx em logs → revisar logs → identificar bug → deploy hotfix

- [ ] **T12.36** Criar Guia de Escalabilidade
  - Aumentar ECS: editar max replicas de 3 para 5
  - Aumentar RDS: upgrade para db.t4g.small
  - Aumentar CloudFront: automático (sem ação necessária)

### Área: Monitoramento e Alertas

- [ ] **T12.37** Validar alarmes CloudWatch
  - Simular alta CPU no container
  - Verificar se SNS envia e-mail
  - Documentar tempo de alerta

- [ ] **T12.38** Criar dashboard de monitoramento
  - Taxa de erro HTTP 5xx
  - Latência P99
  - Conexões RDS ativas
  - Espaço em disco
  - Refresh: 5 minutos

- [ ] **T12.39** Configurar alertas críticos
  - Error rate > 5% → alerta (crítico)
  - Latência P99 > 10 sec → alerta (warning)
  - RDS CPU > 85% → alerta (warning)

### Área: Testes de Integração com Terceiros

- [ ] **T12.40** Testar Google Meet
  - Criar sessão com Google Meet
  - Verificar link é gerado
  - Acessar link → abre sala Meet
  - Documentar resultado

- [ ] **T12.41** Testar Outlook Calendar
  - Criar sessão
  - Verificar sincronização Outlook
  - Deletar sessão → verif deletar do calendário
  - Documentar resultado

### Área: Relatórios e Handover

- [ ] **T12.42** Gerar relatório de testes E2E
  - Incluir: totais, passar, falhas, tempo
  - Adicionar screenshots de sucesso
  - Exportar como HTML/PDF

- [ ] **T12.43** Gerar relatório de carga
  - Métricas: latência, throughput, erros
  - Comparar com baselines
  - Gráficos de evolução

- [ ] **T12.44** Gerar relatório de segurança
  - OWASP checklist
  - Vulnerabilidades encontradas e resolvidas
  - Conformidade LGPD validada

- [ ] **T12.45** Criar matriz de validação final
  - Item | Status | Evidência
  - Testes E2E | PASS | relatório HTML
  - Testes de Carga | PASS | métricas
  - Segurança | PASS | OWASP checklist
  - etc.

---

## Checkpoints Internos

- [ ] Testes E2E: 100% passando
- [ ] Teste de carga: P99 < 5 seg, erro rate < 1%
- [ ] Segurança: Nenhuma vuln crítica/alta
- [ ] Disaster Recovery: Falha e recuperação validadas
- [ ] Performance: Otimizações implementadas
- [ ] Custo: Dentro do orçamento estimado
- [ ] Documentação: Todos os runbooks completos

---

## Notas para QA/DevOps

- **Não alterar código** salvo correções críticas de bugs de segurança
- **Todos os testes devem ser repetíveis** e documentados
- **Feedback loop**: Issues encontradas vão para Sprint 13

---

## Bloqueadores / Dependências

- Dependência: Sprint 11 deve estar 100% concluído
- Pré-requisito: Ambiente AWS em produção estável
