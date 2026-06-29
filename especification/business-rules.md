# Projeto: Gestao Psicologos

## Resumo: Gestão para profissionais da área de psicologia.

## Mecâninca e funcionamento

    a) Fluxo Comum

        1. A única tela acessível sem login é a própria tela de login. Todas as demais são área-logada.

        2. Administração do Psicologo: aqui o psicólogo terá sua pagina de administração onde de ter alguns cadastros e configurações
            - Nome do Psicologo
            - Numero do CRP
            - Email Calendário e Principal
            - Email para Enviar o meeting
            - Telefone Celular
            - Valor da Consulta Padrão
            - Valor do Pacote de Números de consultas e Valor
            - Especialialidade do Psicologo
            - Gestão de troca de senha de acesso

            2.1 Email Calendário e Principal: Cadastro e configuração de email e serviços de autenticação do gmail ou microsoft outlook.
                - Configuração e autenticação persistente do Microsoft Outlook
                - Criar Compromisso para alocação da Agenda do Outlook
                - O calendário dos compromissos do outlook deve ser mostrado no Front
                - Conexã segura com o Google meeting e configuração para ser possivel criar o Link da meeting, onde deve ser copiado e colocado para ser enviado por email para o consultante.

        3. Dashboard: Feature de relatório de atendimentos, por mês, por paciente, relatório de atendimentos e Feature de relatório contábil de ganhos.

        4. Ficha consultante: Cadastro do paciente e gestão de atendimentos do paciente.
            	- Nome e telefone são obrigatórios
		        - Opcional:
		            -Data de Nascimento, empresa que trabalha, cidade e estado, estado civil, idade, profissão
		        -Valor da consulta a ser cobrada (necessário apontar a porcentagem que está acima ou abaixo do valor padrão cobrado feito no cadastro do Psicólogo)
		        - Primeira Analise:
                    Queixa principal:
                    Diagnostico / Sintoma:
                    Influência do desenvolvimento:
                    Questões situacionais:
                    Fatores Biológicos, genéticos e médicos
                    Pontos forte / recursos
			
                    Vícios
                    
                    Estímulos
                    Pensamentos
                    Comportamentos
                    Afetos
                    Fisiológico
                    Objetivos de tratamento:
                    Plano de tratamento
			
			    - Em cada ficha de paciente precisa ter Blocos de anotações que possibilitam o Psicólogo escrever em cada dia de sessão, cada bloco precisa estar separado com a Data do atendimento e deve ser persistente, podendo ser editado a qualquer momento
			
			    - Possibilidade de exportar para PDF a ficha completa do paciente


    b) Primeiro acesso
        
        1. Cada usuário inicia a sessão com o usuario e senha previamente cadastrado pelo usuário master.

        2. Cada usuário inicia no sistema com uma paginas de Boas Vindas onde tem a opção com Botões para acesso
            Administração do Psicologo
            Ficha consultante
            Dashboard

    c) Usuario master

        1. Um usuário Master já nasce junto com o sistema.
            - Usuário e senha padrão: admin/admin
        
        2. Só o usuário Master tem acesso a tela de adicionar e remover usuário.
                Primeiro acesso do Psicologo
                Um psicologo não pode ver os consultantes, pacientes e configurações de outro pscicologo

## DETALHAMENTO
    Crie o arquivo business-rules-details.md com todos os pontos descritos aqui em um nível de profundidade adequado para iniciar o desenvolvimento, incluindo querys, modelagem de tabelas e toda informação necessária para início do código. Qualquer dúvida, conflito ou risco deve ser apontado nesse documento.
