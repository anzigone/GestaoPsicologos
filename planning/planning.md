# Projeto: Gestao Psicologos

## Resumo: Resumo das fases de entrega do projeto com checkpoints de validação humana.

Em cada arquivo de sprint, você deve necessáriamente detalhas minisciosamente quais validações humanas precisam ser feitas para validar a entrega e então partir para a próxima sprint.

** Importante: ** Realizar um commit, tag e push to remote do projeto após a validação de cada sprint.

## Sprints (13 Sprints - Ciclo Completo)

### Fase 1: Especificação e Prototipagem

    1. Setup inicial: Arquivos de configuração, instalação de frameworks, criação de dockerfiles e setup de CI/CD.
        - Entregável: codebase inicial, repositório configurado e ambiente pronto para desenvolvimento.

    2. Mock Webservice: Asset do webservice (BFF) com toda a documentação Swagger entregue, com métodos mockado (sem lógica real). O endpoint deve ser deployável.
        - Entregável: build & run do webservice para análise e conferência humana dos endpoints.

    3. Mock Frontend: Asset do frontend com as telas protótipo de todas as áreas da app, ainda sem autenticação.
        - Entregável: telas visualmente fiéis ao resultado final para validação e ajustes humanos.

### Fase 2: Funcionalidades Core

    4. Autenticação e Gestão: Telas internas protegidas por login.
        - Entregáveis:
            - Todo o mecanismo de autenticação pronto e testável.
            - App possível de login com o usuário admin.
            - Tela de gestão dos usuários pronta e funcional.

    5. Administração do Psicólogo: Tela de visualização e administração para psicólogos.
        - Entregáveis:
            - Visualização da tela
            - Todo o cadastro do Psicologo.
            - Troca de senha segura.
             
    6. Ficha Consultante: Ficha de cadastro completa.
        - Entregáveis:
            - Navegação no cadastro, visualização de toda a tela.
            - Cadastro do Paciente consultante.
            - Preenchimento da Anamnese (Primeira Análise).
            - **NOVO**: Exportação de prontuário em PDF com primeira análise e histórico de sessões.

    7. Dashboard Financeiro: Toda a gestão dos pacientes e financeira.
        - Entregáveis:
            - Tela de visualização completa com KPIs.
            - Gráficos de faturamento e volume de atendimentos.
            - Tabela de transações com filtros.

### Fase 3: Integrações Externas e Polimento

    8. Integração com Calendário: Configuração e integração com Google Meeting e Microsoft Outlook.
        - Entregáveis:
            - Autenticação OAuth2 com Google e Microsoft.
            - Criação automática de eventos no Outlook ao registrar sessões.
            - Geração de link Google Meet nas sessões.
            - **NOVO**: Envio de e-mail com convite e link de reunião para psicólogo e paciente.

    9. Polimento e Ajustes Gerais: Refinamento de UX, validações e formatações.
        - Entregável: 
            - Validações robustas de formulários.
            - Feedback visual de carregamento.
            - Normalização de fusos horários.
            - Formatação de moedas e datas em PT-BR.

    10. Hardening e Prontidão para Produção: Segurança e otimização para deploy.
        - Entregáveis:
            - Autenticação migrada para Cookies HttpOnly e Secure.
            - Resolução de N+1 em queries de API.
            - Integração com AWS Secrets Manager.
            - Menu Mobile funcional (Drawer/Sheet).
            - Testes de performance implementados.

### Fase 4: Deploy e Validação

    11. Deploy Completo na AWS: Infraestrutura de produção total.
        - Entregáveis:
            - VPC, subnets, security groups configurados.
            - RDS MySQL provisionado e dados migrados.
            - ECR com imagens Docker compiladas.
            - ECS Fargate Services para Frontend e Backend.
            - Application Load Balancer (ALB) configurado.
            - CloudFront CDN com SSL/TLS.
            - AWS Secrets Manager para credenciais.
            - CloudWatch Logs e Alarms.
            - Pipeline CI/CD do GitHub Actions testado.

    12. Validação e Testes de Produção: Testes de carga, segurança e disaster recovery.
        - Entregáveis:
            - Testes E2E automatizados executados com sucesso.
            - Testes de carga e stress validados.
            - Validação OWASP Top 10 e conformidade LGPD.
            - Testes de disaster recovery (failover, restore de backup).
            - Otimizações de performance implementadas.
            - Documentação operacional completa (runbooks, troubleshooting).

    13. Testes Finais e Ajustes: Ajustes menores, testes de regressão, documentação de usuário.
        - Entregáveis:
            - Testes de regressão com 100% de sucesso.
            - Ajustes baseados em feedback de Sprint 12.
            - Documentação de usuário final (guias, FAQ).
            - Documentação de suporte técnico.
            - Plano de manutenção contínua.
            - Handover para times de operação.
            - Relatório executivo final do projeto.


    ## DETALHAMENTO
Planejador: Crie os arquivos de planejamento de sprint conforme definido em documento anterior. Faça isso de forma iterativa com o humano, declarando sua visão sobre cada sprint e aguardando a confirmação antes da criação do arquivo. Essa é a última especificação antes do desenvolvedor, então seu detalhamento precisa ser minucioso no nível de desenvolvimento para que não haja interpretação vaga das tasks e implementação que serão realizadas. Todo o contexto refinado aqui é importante.
Programador: Siga o desenvolvimento passo a passo com o humano. No final de cada sprint, você deve instrui-lo de como validar por conta própria os entregáveis de cada sprint.
