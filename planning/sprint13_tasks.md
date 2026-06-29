# Sprint 13: Testes Finais - Tarefas Detalhadas (sprint13_tasks.md)

## Overview
Sprint final de testes de regressão, ajustes menores baseados em feedback de Sprint 12, documentação de usuário e preparação para operação contínua. **Não é sprint de desenvolvimento**, apenas validação e polimento final.

---

## Tarefas (QA/Suporte/Programador)

### Área: Testes de Regressão

- [ ] **T13.1** Executar suite completa de testes E2E
  - Rodar: `npx cypress run`
  - Esperado: 100% de testes passando
  - Documentar resultado e screenshots

- [ ] **T13.2** Executar teste de carga reduzido
  - 50 requisições simultâneas (vs 100-500 em S12)
  - Validar que sistema responde normalmente
  - Latência dentro dos limites? ✓

- [ ] **T13.3** Testar fluxo crítico completo
  - [ ] Login → Welcome
  - [ ] Criar psicólogo
  - [ ] Criar paciente
  - [ ] Registrar sessão
  - [ ] Gerar PDF
  - [ ] Visualizar Dashboard
  - Esperado: Sem erros ou degradações

- [ ] **T13.4** Testar rollback de Sprint 12
  - Se houve deploy de ajustes em S12
  - Validar que tudo funciona ainda
  - Sem regressões? ✓

### Área: Análise de Logs e Saúde do Sistema

- [ ] **T13.5** Revisar CloudWatch Logs (últimas 24h)
  - Backend:
    - Erros 5xx: Esperado < 5
    - Timeouts: Esperado 0
    - Avisos (WARN): Documentar quais
  - Frontend:
    - Erros JavaScript: Esperado 0
    - Warnings: Documentar quais

- [ ] **T13.6** Analisar métricas do dashboard CloudWatch
  - CPU Backend: Esperado < 30% na média
  - Memória Backend: Esperado < 50% na média
  - Latência P99: Esperado < 5 seg
  - Taxa de erro: Esperado < 1%
  - Documentar em relatório

- [ ] **T13.7** Validar saúde de serviços
  - RDS: Status "available"? ✓
  - ECS: Tasks "running"? ✓
  - ALB: Target health "healthy"? ✓
  - CloudFront: Distribution "Deployed"? ✓

### Área: Implementar Ajustes de Sprint 12

- [ ] **T13.8** Priorizar issues de Sprint 12
  - Críticos (segurança, quebra funcional): Implementar AGORA
  - Maiores (usabilidade, perf): Implementar se tempo permite
  - Menores (cosmética): Documentar para sprint futura

- [ ] **T13.9** Implementar correções críticas
  - Exemplo: Se houver XSS vulnerability → fixa
  - Exemplo: Se logout não funcionar → fixa
  - Testar isoladamente
  - Deploy via GitHub Actions

- [ ] **T13.10** Implementar correções maiores (se houver tempo)
  - Exemplo: Adicionar confirmação antes de deletar paciente
  - Exemplo: Melhorar mensagem de erro de validação
  - Testar isoladamente
  - Deploy via GitHub Actions

- [ ] **T13.11** Validar que não há regressões
  - Rodar E2E tests novamente
  - Esperado: Todas passando ainda
  - Se falhar: rollback ou corrigir

### Área: Melhorias de UX Menores

- [ ] **T13.12** Revisar feedback de usabilidade
  - Stakeholders encontraram algo não intuitivo?
  - Interface está confusa em alguma tela?
  - Mensagens estão claras?

- [ ] **T13.13** Implementar melhorias de UX
  - Exemplo: Adicionar tooltip em campo de CRP
  - Exemplo: Melhorar ordem de campos em formulário
  - Exemplo: Adicionar atalho Ctrl+S para salvar
  - Testar com usuários finais se possível

### Área: Documentação de Usuário Final

- [ ] **T13.14** Criar guia "Como Logar"
  - Screenshots passo-a-passo
  - Diferenciar admin vs psicólogo
  - Incluir opções de "Esqueci a Senha" (se implementado)
  - Formato: PDF ou página web

- [ ] **T13.15** Criar guia "Cadastro do Psicólogo"
  - Como preencher dados profissionais
  - Como definir tarifa padrão
  - Como conectar Google Meet
  - Como conectar Outlook
  - Screenshots com cada passo

- [ ] **T13.16** Criar guia "Cadastro de Pacientes"
  - Como criar novo paciente
  - Campos obrigatórios vs opcionais
  - Como preencher anamnese
  - Como exportar ficha em PDF
  - Screenshots

- [ ] **T13.17** Criar guia "Registrar Sessões"
  - Como criar nova sessão
  - Como anotar evolução
  - Como gerar link Google Meet
  - Como marcar como pago/pendente
  - Screenshots

- [ ] **T13.18** Criar guia "Consultar Dashboard"
  - Como ler KPIs
  - Como entender gráficos
  - Como usar filtros
  - Legenda de cores (Teal, Verde, Vermelho)
  - Screenshots

- [ ] **T13.19** Criar FAQ (Perguntas Frequentes)
  - P: Como recuperar senha?
    R: Página de login tem link "Esqueci a senha" (ou enviar para admin)
  - P: Como integrar Google Meet?
    R: Ir para Administração → Conectar Conta Google
  - P: Como exportar dados de um paciente?
    R: Ficha → Botão "Exportar em PDF"
  - P: O que acontece quando deletar um paciente?
    R: Paciente e todas as sessões são removidos permanentemente
  - P: Como suportar múltiplos consultantes ao mesmo tempo?
    R: Usar abas do navegador ou acessar via dispositivos diferentes

- [ ] **T13.20** Criar guia de troubleshooting
  - Problema: "Não consigo fazer login"
    Solução: Verificar caps lock, limpar cache, usar navegador diferente
  - Problema: "Dashboard carrega lentamente"
    Solução: Verificar internet, fechar abas desnecessárias
  - Problema: "Google Meet não está funcionando"
    Solução: Verificar se Google está conectado, reconectar conta
  - Problema: "Botão 'Salvar' não responde"
    Solução: Verificar internet, recarregar página, contatar suporte
  - Problema: "Não consigo deletar um paciente"
    Solução: Verificar permissões, contatar admin

### Área: Documentação de Suporte Técnico

- [ ] **T13.21** Criar template de ticket de suporte
  - Padrão: Nome | E-mail | Problema | Prioridade
  - Sistema: Usar GitHub Issues, Zendesk, ou similar

- [ ] **T13.22** Documentar SLA de resposta
  - Crítico (app down): 1 hora
  - Grave (feature quebrada): 4 horas
  - Normal (bug menor): 1 dia
  - Baixa (sugestão): 5 dias

- [ ] **T13.23** Criar template de comunicação de incidente
  - Status inicial
  - Status de investigação
  - Status de resolução
  - Comunicar a stakeholders

### Área: Validação de Conformidade Final

- [ ] **T13.24** Executar scan OWASP final
  - Nenhuma vulnerabilidade crítica/alta? ✓
  - Todas as recomendações implementadas? ✓

- [ ] **T13.25** Validar LGPD um mais vez
  - Criptografia em trânsito (HTTPS)? ✓
  - Criptografia em repouso (senhas bcrypt)? ✓
  - Isolamento de dados multi-tenancy? ✓
  - Direito ao esquecimento (delete cascade)? ✓
  - Nenhum dado sensível em logs? ✓

- [ ] **T13.26** Validar SSL/TLS
  - Certificado válido e não expirado? ✓
  - Renovação automática configurada? ✓
  - TLS 1.2+? ✓

### Área: Plano de Manutenção Contínua

- [ ] **T13.27** Documentar cronograma de manutenção
  - **Diária**:
    - Monitoramento: Checar dashboards CloudWatch
    - Logs: Revisar erros em CloudWatch
  - **Semanal**:
    - Review: Verificar trends de performance
    - Alertas: Analisa alarmes disparados
  - **Mensal**:
    - Testes: Executar suite E2E
    - Updates: Aplicar patches de segurança do SO
    - Backup: Validar integridade de backup RDS
  - **Trimestral**:
    - Review: Análise de custo e FinOps
    - Planejamento: Identificar melhorias

- [ ] **T13.28** Documentar processo de atualização
  - Go: Como atualizar versão de Go
  - Node: Como atualizar versão de Node
  - Dependências: Como fazer `go get -u ./...` ou `npm update`
  - Testar localmente
  - Deploy via CI/CD

- [ ] **T13.29** Documentar processo de hotfix
  - Bug crítico encontrado em produção
  - Criar branch `hotfix/*`
  - Corrigir e fazer PR
  - Mergear em `main`
  - Deploy automático via GitHub Actions
  - Tempo total: < 1 hora

- [ ] **T13.30** Documentar processo de scaling
  - Se CPU > 80% por > 10 min:
    - Aumentar max replicas ECS de 3 para 5
    - Monitorar por 30 min
    - Se melhorar, manter; senão, investigar
  - Se RDS CPU > 80%:
    - Upgrade para db.t4g.small
    - Testar performance
    - Monitorar por 1 hora

### Área: Handover para Times de Operação

- [ ] **T13.31** Preparar apresentação técnica
  - Arquitetura geral (diagrama)
  - Stack (Next.js, Go, MySQL, Fargate)
  - Fluxo de deploy (GitHub → Actions → ECR → ECS)
  - Monitoramento (CloudWatch)
  - Incident response (runbooks)

- [ ] **T13.32** Treinar times de operação
  - Explicar como acessar AWS Console
  - Mostrar CloudWatch Dashboard
  - Demonstrar como fazer rollback
  - Responder dúvidas

- [ ] **T13.33** Preparar documentação técnica final
  - README com links para:
    - Repositório GitHub
    - Documentação API (Swagger)
    - CloudWatch Dashboard
    - Runbooks
    - Contatos de suporte
  - Incluir diagrama de arquitetura
  - Incluir ERD (Entity Relationship Diagram) do banco

- [ ] **T13.34** Transferir posse de credenciais
  - AWS: Criar IAM users para times de operação
  - GitHub: Adicionar como maintainers
  - Secrets: Documentar como acessar Secrets Manager
  - Domínio: Transferir propriedade se necessário

### Área: Relatório Final Executivo

- [ ] **T13.35** Preparar relatório de status final
  - Projeto: Gestão Psicólogos
  - Data: [data atual]
  - Status: ✅ 100% Completo e em Produção
  - Sprints: 13 sprints concluídas
  - Métricas:
    - Testes E2E: 100% passando
    - Uptime: 99.9% (se houver dados)
    - Latência P99: < 5 seg
    - Taxa de erro: < 1%

- [ ] **T13.36** Documentar lessons learned
  - O que funcionou bem?
    - Abordagem spec-driven development
    - Sprints bem definidas
    - Testes automatizados
  - O que poderia melhorar?
    - Mais testes de integração desde o início
    - Melhor documentação de APIs
    - Feedback do usuário mais cedo
  - Recomendações para projetos futuros

- [ ] **T13.37** Documentar roadmap futuro
  - Features não implementadas em S13:
    - Autenticação via SMS 2FA
    - Relatório fiscal avançado
    - Mobile app nativa
  - Melhorias de performance:
    - Cache em Redis
    - CDN para imagens
  - Novas integrações:
    - Slack notifications
    - Zendesk integration

### Área: Encerramento Formal

- [ ] **T13.38** Criar release notes consolidadas
  - Versão: v0.13.0
  - Data: [data]
  - Resumo: Projeto completo em produção com validação 100%
  - Features em v0.13.0: (apenas ajustes)
  - Bugs corrigidos: (listar)

- [ ] **T13.39** Gerar tags Git finais
  - Commit: `feat: sprint 13 testes finais concluido`
  - Tag: `v0.13.0`
  - Push para `main`
  - Criar release no GitHub com notas

- [ ] **T13.40** Comunicar conclusão ao time
  - Email/Slack: Projeto Gestão Psicólogos está 100% completo
  - Incluir link para URL de produção
  - Incluir contatos de suporte
  - Agradecer ao time e stakeholders

---

## Checkpoints Internos

- [ ] Testes E2E: 100% passando
- [ ] Nenhuma regressão detectada
- [ ] Ajustes de S12 implementados e validados
- [ ] Documentação de usuário completa
- [ ] Documentação de suporte completa
- [ ] Plano de manutenção documentado
- [ ] Relatório executivo concluído
- [ ] Handover realizado com times de operação

---

## Notas para QA/Suporte

- **Prioridade**: Estabilidade > novos ajustes
- **Foco**: Documentação e transferência de conhecimento
- **Timeline**: Não apressar, tudo precisa estar correto
- **Comunicação**: Manter stakeholders atualizados

---

## Bloqueadores / Dependências

- Dependência: Sprint 12 deve estar 100% concluído
- Pré-requisito: Ambiente AWS estável
