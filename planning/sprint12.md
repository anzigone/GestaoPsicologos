# Sprint 12: Validação, Testes e Ajustes de Produção (sprint12.md)

## 1. Escopo e Objetivos

O objetivo desta sprint é realizar a validação completa do sistema em ambiente de produção AWS, executar testes de carga e estresse, identificar e corrigir problemas de performance ou segurança, e garantir que a aplicação está pronta para receber usuários finais.

**Objetivos principais**:
- Executar testes de carga e estresse na aplicação em produção.
- Validar todas as funcionalidades fim-a-fim no ambiente AWS.
- Identificar e corrigir gargalos de performance.
- Validar a segurança e conformidade de dados (GDPR, LGPD se aplicável).
- Otimizar custos de infraestrutura (FinOps).
- Preparar documentação operacional e runbooks para suporte.
- Realizar testes de disaster recovery (simulação de falhas).

---

## 2. Entregáveis da Sprint

Ao final desta sprint, os seguintes entregáveis deverão estar prontos:

- **Testes de Validação**:
  - Relatório de testes E2E automatizados (Cypress ou Playwright).
  - Relatório de testes de carga (k6, Apache JMeter ou similar).
  - Evidência de execução de testes de segurança (OWASP Top 10).

- **Documentação Operacional**:
  - Runbook de deployment e rollback.
  - Runbook de incident response para falhas comuns.
  - Documentação de acesso e permissões IAM.
  - Guia de escalabilidade manual e automática.

- **Performance Optimization**:
  - Relatório de otimizações implementadas (cache, CDN, índices de banco de dados).
  - Métricas de latência P50/P95/P99 consolidadas.
  - Recomendações de monitoramento contínuo.

- **Segurança**:
  - Checklist de segurança validado (OWASP, LGPD, encryption).
  - Relatório de testes de penetração (se realizado).
  - Configuração de WAF (Web Application Firewall) se necessário.

- **Cost Optimization**:
  - Relatório de custo mensal estimado da infraestrutura.
  - Recomendações de otimização de custos implementadas.

- **Preparação para Produção**:
  - Monitoramento e alertas configurados e testados.
  - Backup e restore procedures validados.
  - Plano de atualização de dependências definido.

---

## 3. Tarefas Detalhadas

### 3.1. Testes E2E (End-to-End)
- [ ] Preparar suite de testes E2E com Cypress ou Playwright:
  - Login e autenticação.
  - CRUD completo de psicólogos.
  - CRUD completo de pacientes.
  - Fluxo de sessões e evolução.
  - Dashboard e relatórios.
  - Integração OAuth2 (simulada se necessário).
- [ ] Executar testes E2E contra produção AWS.
- [ ] Documentar resultados e Screenshots de sucesso.

### 3.2. Testes de Carga
- [ ] Preparar scripts de teste de carga:
  - 100 requisições simultâneas de login.
  - 50 psicólogos criando pacientes simultaneamente.
  - Leitura massiva de dashboard (lista de pacientes, gráficos).
  - Consultas combinadas (busca + filtros).
- [ ] Executar testes com k6 ou Apache JMeter.
- [ ] Analisar métricas:
  - Latência P50, P95, P99.
  - Taxa de erro (4xx, 5xx).
  - Throughput (requisições por segundo).
  - Consumo de CPU/Memória dos containers.
- [ ] Documentar comportamento sob carga e identificar gargalos.

### 3.3. Testes de Segurança
- [ ] Executar verificação de OWASP Top 10:
  - Injection (SQL Injection, NoSQL Injection).
  - Broken Authentication.
  - Sensitive Data Exposure.
  - XML External Entities (XXE).
  - Broken Access Control.
  - Security Misconfiguration.
  - XSS (Cross-Site Scripting).
  - Insecure Deserialization.
  - Using Components with Known Vulnerabilities.
  - Insufficient Logging & Monitoring.
- [ ] Executar dependency scanning (npm audit, go mod vulnerabilities).
- [ ] Validar certificados SSL/TLS e HSTS.
- [ ] Verificar headers de segurança (X-Content-Type-Options, CSP, etc.).
- [ ] Documentar resultados e ações corretivas.

### 3.4. Testes de Conformidade (GDPR/LGPD)
- [ ] Validar criptografia de dados em trânsito (HTTPS).
- [ ] Validar criptografia de dados em repouso (senhas com bcrypt, tokens seguros).
- [ ] Verificar isolamento de dados entre psicólogos (multi-tenancy).
- [ ] Testar direito ao esquecimento (GDPR/LGPD):
  - Deletar usuário psicólogo e validar cascata (pacientes, sessões também deletam).
  - Validar que dados são realmente removidos do banco.
- [ ] Documentar conformidade implementada.

### 3.5. Testes de Disaster Recovery
- [ ] Simular falha de container Backend:
  - [ ] Pausar container manualmente.
  - [ ] Verificar se Auto Scaling cria novo container.
  - [ ] Confirmar que aplicação continua responsiva.
- [ ] Simular falha de RDS (simulação com snapshot):
  - [ ] Verificar se backup automático foi realizado.
  - [ ] Testar restauração de backup em instância de teste.
  - [ ] Validar integridade dos dados após restore.
- [ ] Testar rollback de deployment:
  - [ ] Fazer deploy de versão anterior via GitHub Actions.
  - [ ] Verificar se aplicação volta a funcionar sem perda de dados.

### 3.6. Otimizações de Performance
- [ ] Analisar logs de latência e identificar endpoints lentos.
- [ ] Implementar índices de banco de dados otimizados:
  - [ ] Índice em `patients(psychologist_id)`.
  - [ ] Índice em `sessions(patient_id, status)`.
  - [ ] Índice em `sessions(session_date)` para dashboard.
- [ ] Validar cache no CloudFront:
  - [ ] Assets estáticos (JS, CSS, imagens) com TTL = 1 dia.
  - [ ] HTML com TTL curto (5 minutos) ou sem cache.
  - [ ] `/api/*` sem cache.
- [ ] Otimizar queries SQL em agregações do dashboard:
  - [ ] Adicionar EXPLAIN PLAN para identificar full scans.
  - [ ] Refatorar queries se necessário.
- [ ] Testar compressão Gzip/Brotli em responses.
- [ ] Validar lazy loading de componentes no frontend.

### 3.7. Otimizações de Custo (FinOps)
- [ ] Acessar AWS Cost Explorer e analisar despesas por serviço.
- [ ] Validar:
  - [ ] RDS está em Free Tier (db.t3.micro ou db.t4g.micro).
  - [ ] ECS está em Free Tier (0.25 vCPU x 2 containers = 50 horas/mês gratuitas).
  - [ ] CloudFront transferência de dados está dentro de limites.
- [ ] Implementar otimizações sugeridas:
  - [ ] Reduzir retenção de logs no CloudWatch (ex: 7 dias em vez de 30).
  - [ ] Usar Reserved Instances para RDS se necessário.
  - [ ] Ativar S3 Intelligent-Tiering para backups antigos.
- [ ] Configurar AWS Budgets com alertas de custo máximo.

### 3.8. Documentação Operacional

#### 3.8.1 Runbook de Deployment
- [ ] Documentar procedimento de deployment:
  - Fazer commit em `main`.
  - GitHub Actions executa pipeline.
  - Validações automáticas (build, testes).
  - Push para ECR.
  - Update de task definition no ECS.
  - Monitoramento de health checks.

#### 3.8.2 Runbook de Rollback
- [ ] Documentar procedimento de rollback:
  - Identificar versão anterior estável (via git tags).
  - Forçar deploy da versão anterior via GitHub Actions ou ECS console.
  - Validar funcionalidade pós-rollback.

#### 3.8.3 Runbook de Incident Response
- [ ] Documentar resposta para cenários:
  - Aplicação lenta ou não responsiva.
  - Erro de banco de dados (conexão, corrupção).
  - Vazamento de memória (CPU/RAM elevados).
  - Ataque DDoS ou anomalia de tráfego.
  - Certificado SSL expirando.

#### 3.8.4 Guia de Escalabilidade
- [ ] Documentar como aumentar capacidade:
  - Escalar RDS (mudar instância para db.t4g.small).
  - Escalar ECS (aumentar max replicas de 3 para 5).
  - Distribuir tráfego entre múltiplas regiões (se necessário no futuro).

### 3.9. Monitoramento e Alertas
- [ ] Validar que todos os alarmes CloudWatch estão disparando corretamente:
  - [ ] Simular alta CPU no container (teste de estresse).
  - [ ] Confirmar que SNS envia e-mail de alerta.
- [ ] Configurar dashboard de monitoramento contínuo com:
  - [ ] Taxa de erro HTTP 5xx (alertar se > 1%).
  - [ ] Latência P99 (alertar se > 5 segundos).
  - [ ] Conexões de banco ativas (alertar se > 80%).
  - [ ] Espaço em disco RDS (alertar se > 80%).

### 3.10. Testes de Integração com Terceiros
- [ ] Validar integração Google Meet:
  - [ ] Criar sessão com Google Meet.
  - [ ] Verificar que link é gerado corretamente.
  - [ ] Testar acesso ao link (abre sala do Meet).
- [ ] Validar integração Outlook:
  - [ ] Criar sessão e sincronizar com Outlook.
  - [ ] Verificar evento no calendário do Outlook.
  - [ ] Deletar sessão e validar remoção do calendário.

### 3.11. Preparação para Suporte e Maintenance
- [ ] Documentar procedimento de upgrade de dependências:
  - Testar localmente.
  - Criar branch de feature.
  - Executar testes automatizados.
  - Deploy via pipeline.
- [ ] Preparar documentação de troubleshooting comum.
- [ ] Criar plano de backup e recuperação de longo prazo.

### 3.12. UAT (User Acceptance Testing) - Simulação com Usuários Reais
- [ ] Fornecer credenciais de teste para stakeholders.
- [ ] Solicitar feedback sobre:
  - Usabilidade da interface.
  - Performance percebida.
  - Funcionalidade de negócio.
- [ ] Documentar issues e priorizar correções.

---

## 4. Checkpoints de Validação Humana

Para validar a conclusão desta Sprint, o humano realizará os seguintes testes:

1. **Testes E2E Automatizados**:
   - [ ] Executar suite completa de testes E2E.
   - [ ] Esperado: Todos os testes passam (100% taxa de sucesso).
   - [ ] Revisar relatório e screenshots.

2. **Teste de Carga**:
   - [ ] Revisar relatório de carga:
     - Latência P99 < 5 segundos.
     - Taxa de erro < 1%.
     - Aplicação responsiva sob carga.

3. **Teste de Segurança**:
   - [ ] Revisar relatório OWASP.
   - [ ] Confirmar que não há vulnerabilidades críticas ou altas.
   - [ ] Validar que dependências não têm vulnerabilidades conhecidas.

4. **Teste de Disaster Recovery**:
   - [ ] Confirmar que simulação de falha foi executada com sucesso.
   - [ ] Verificar que aplicação se recuperou automaticamente.
   - [ ] Validar que dados não foram perdidos.

5. **Performance e Custo**:
   - [ ] Revisar relatório de otimizações implementadas.
   - [ ] Confirmar custo mensal está dentro do orçamento.
   - [ ] Verificar que métricas de performance estão dentro dos limites aceitáveis.

6. **Documentação Operacional**:
   - [ ] Revisar runbooks e procedimentos.
   - [ ] Confirmar que estão claros e testáveis por terceiros.

7. **UAT e Feedback**:
   - [ ] Coletar feedback de stakeholders.
   - [ ] Documentar issues encontradas (pequenos ajustes apenas, nada crítico).

8. **Encerramento da Sprint**:
   - [ ] Realizar commit com a mensagem `feat: sprint 12 validacao producao concluido`.
   - [ ] Criar a tag Git `v0.12.0`.
   - [ ] Fazer push para a branch `main` do GitHub.

---

## 5. Notas Importantes

- **Não Alterar Código em Sprint 12**: Esta sprint é principalmente de testes e validação. Apenas correções críticas de bugs de segurança devem ser implementadas.
- **Relatórios**: Todos os testes devem gerar relatórios exportáveis (PDF ou HTML) com evidências.
- **Comunicação**: Manter comunicação constante com stakeholders durante testes.
- **Feedback Loop**: Issues encontradas em Sprint 12 podem resultar em Sprint 13 (testes finais) para ajustes menores.
