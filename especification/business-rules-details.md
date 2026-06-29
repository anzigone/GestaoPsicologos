# Projeto: Gestão Psicólogos - Detalhamento das Regras de Negócio (business-rules-details.md)

Este documento descreve minuciosamente as regras de negócio do sistema, a modelagem completa das tabelas do banco de dados (SQLite e MySQL), os fluxos lógicos fundamentais e as queries SQL necessárias para o desenvolvimento do webservice.

---

## 1. Isolamento de Dados (Multi-tenancy)

O sistema opera sob o modelo de **Isolamento Estrito por Psicólogo**:
1. Cada psicólogo tem acesso exclusivo aos seus dados cadastrais, suas configurações de integração, seus pacientes (consultantes) e seu histórico de prontuários/faturamento.
2. É terminantemente proibido que um psicólogo acesse qualquer recurso pertencente a outro profissional.
3. **Mecanismo de Segurança**: O token JWT emitido após a autenticação carrega o ID do psicólogo logado (`sub` claim). Toda requisição ao banco de dados que leia, crie ou altere dados de pacientes e sessões obrigatoriamente filtrará/definirá o ID do psicólogo utilizando o ID extraído do JWT (ex: `WHERE psychologist_id = ?`).

---

## 2. Modelagem Física de Dados

A seguir, a definição DDL completa para a estrutura do banco de dados. Estes tipos de dados são compatíveis com SQLite (DEV local) e RDS MySQL (PROD AWS).

```sql
-- Tabela 1: users (Psicólogos e Master Admin)
CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(36) PRIMARY KEY,              -- UUID v4
    email VARCHAR(255) UNIQUE NOT NULL,      -- E-mail de acesso
    password_hash VARCHAR(255) NOT NULL,     -- Hash SHA256 (com Salt) da senha
    role VARCHAR(20) NOT NULL,               -- 'admin' (Master) ou 'psicologo'
    name VARCHAR(255) NOT NULL,              -- Nome profissional completo
    crp VARCHAR(50),                         -- CRP (Registro Profissional)
    specialty VARCHAR(100),                  -- Especialidade do psicólogo
    phone VARCHAR(20),                       -- Celular principal
    base_fee DECIMAL(10,2) DEFAULT 0.00,       -- Valor padrão da consulta
    package_sessions INT DEFAULT 0,          -- Número de consultas em um pacote padrão
    package_fee DECIMAL(10,2) DEFAULT 0.00,  -- Valor padrão do pacote
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Tabela 2: patients (Cadastro de Pacientes/Consultantes)
CREATE TABLE IF NOT EXISTS patients (
    id VARCHAR(36) PRIMARY KEY,              -- UUID v4
    psychologist_id VARCHAR(36) NOT NULL,    -- FK para o psicólogo dono do registro
    name VARCHAR(255) NOT NULL,              -- Nome do paciente (obrigatório)
    phone VARCHAR(20) NOT NULL,              -- Telefone/WhatsApp (obrigatório)
    birthdate DATE,                          -- Opcional
    age INT,                                 -- Opcional
    profession VARCHAR(100),                 -- Opcional
    company VARCHAR(100),                    -- Opcional
    city VARCHAR(100),                       -- Opcional
    state VARCHAR(2),                        -- Opcional (UF de 2 letras)
    marital_status VARCHAR(50),              -- Opcional
    consultation_fee DECIMAL(10,2) NOT NULL, -- Valor cobrado especificamente deste paciente
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (psychologist_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Tabela 3: first_analysis (Anamnese/Primeira Análise do Paciente)
CREATE TABLE IF NOT EXISTS first_analysis (
    patient_id VARCHAR(36) PRIMARY KEY,      -- PK e FK referenciando o paciente (1-para-1)
    main_complaint TEXT,                     -- Queixa Principal
    symptom_diagnosis TEXT,                  -- Diagnóstico / Sintoma
    developmental_influence TEXT,            -- Influência do desenvolvimento
    situational_issues TEXT,                 -- Questões situacionais
    biological_factors TEXT,                 -- Fatores biológicos, genéticos e médicos
    strengths_resources TEXT,                -- Pontos fortes / Recursos
    addictions TEXT,                         -- Vícios
    stimuli TEXT,                            -- Estímulos
    thoughts TEXT,                           -- Pensamentos
    behaviors TEXT,                          -- Comportamentos
    affects TEXT,                            -- Afetos
    physiological TEXT,                      -- Fisiológico
    treatment_goals TEXT,                    -- Objetivos de tratamento
    treatment_plan TEXT,                     -- Plano de tratamento
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (patient_id) REFERENCES patients(id) ON DELETE CASCADE
);

-- Tabela 4: sessions (Histórico de Atendimentos, Notas Clínicas e Finanças)
CREATE TABLE IF NOT EXISTS sessions (
    id VARCHAR(36) PRIMARY KEY,              -- UUID v4
    patient_id VARCHAR(36) NOT NULL,         -- FK para a tabela de pacientes
    session_date DATETIME NOT NULL,          -- Data e hora do atendimento
    notes TEXT,                              -- Bloco de anotações da evolução do paciente
    status VARCHAR(20) NOT NULL,             -- 'pago' ou 'pendente'
    meet_link VARCHAR(255),                  -- Link gerado do Google Meet
    outlook_event_id VARCHAR(255),           -- ID do evento no Microsoft Outlook (para cancelamentos)
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (patient_id) REFERENCES patients(id) ON DELETE CASCADE
);

-- Tabela 5: oauth_tokens (Tokens de API salvos de forma encriptada)
CREATE TABLE IF NOT EXISTS oauth_tokens (
    user_id VARCHAR(36) NOT NULL,
    provider VARCHAR(20) NOT NULL,           -- 'google' ou 'outlook'
    access_token TEXT NOT NULL,              -- Token de acesso
    refresh_token TEXT,                      -- Token de atualização
    expiry DATETIME NOT NULL,                -- Validade do access_token
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, provider),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
```

---

## 3. Lógica das Regras de Negócio e Cálculos

### 3.1. Variação do Valor da Consulta do Paciente
Ao cadastrar ou atualizar o valor de consulta de um paciente (`patients.consultation_fee`), o sistema deve exibir visualmente a porcentagem de desvio em relação ao valor padrão do psicólogo (`users.base_fee`).
- **Fórmula**:
  $$\text{Variação \%} = \frac{\text{consultation\_fee} - \text{base\_fee}}{\text{base\_fee}} \times 100$$
- **Regras**:
  - Se $\text{Variação \%} > 0$, exibe "X% acima do valor padrão".
  - Se $\text{Variação \%} < 0$, exibe "X% abaixo do valor padrão" (em valor absoluto).
  - Se $\text{Variação \%} = 0$, exibe "Valor padrão".
  - Tratamento: Se `base_fee` do psicólogo for R$ 0,00, a variação será considerada 0% para evitar divisão por zero.

### 3.2. Primeiro Acesso
- O usuário `admin` (senha `admin`) é o administrador master pré-criado no sistema.
- Ao logar pela primeira vez com as credenciais padrões enviadas pelo administrador, o psicólogo deve ser direcionado para a tela de Boas-Vindas, onde é induzido a preencher seus dados profissionais na tela "Administração do Psicólogo" para definir sua tarifa base e configurar as integrações de calendário.

### 3.3. Exportação de Prontuário Completo (PDF)
O sistema deve gerar um documento PDF estruturado para o paciente contendo:
1. Dados Cadastrais básicos (Nome, Telefone, Profissão, Idade).
2. O histórico preenchido da **Primeira Análise** (seções clínicas).
3. Todas as **Sessões e Evoluções** anotadas, ordenadas de forma cronológica decrescente (da mais recente para a mais antiga), contendo Data, notas clínicas e status do atendimento.
- *Implementação Técnica Backend*: O endpoint `GET /api/patients/{id}/pdf` irá ler as tabelas `patients`, `first_analysis` e `sessions`, compilar a estrutura em HTML e renderizar em PDF usando o gerador PDF interno em Go (como a biblioteca `gofpdf`).

---

## 4. Integração com Calendários e Meetings

### 4.1. Fluxo de Agendamento Unificado
Ao criar uma sessão no sistema (Aba 3 da Ficha do Paciente), a aplicação realiza as seguintes operações integradas se as permissões OAuth estiverem ativas:

1. **Agenda do Outlook**:
   - O backend faz um POST para a API do Microsoft Graph `/me/events`.
   - Adiciona o compromisso com os detalhes do paciente e horário configurado.
   - Retorna o `id` do evento gerado no Outlook, que é salvo na coluna `sessions.outlook_event_id`.
2. **Google Meet**:
   - O backend cria um evento de calendário na conta do Google conectada usando a API do Google Calendar (`calendar.events.insert`).
   - Define a opção `conferenceDataVersion=1` com a solução `hangoutsMeet` para forçar a criação de uma sala do Google Meet.
   - Captura a URL retornada (`conferenceData.entryPoints[0].uri`) e a salva em `sessions.meet_link`.
3. **Fluxo de E-mail**:
   - O webservice envia um e-mail para o e-mail do consultante e/ou para o e-mail do psicólogo contendo o convite e o link de vídeo gerado para o atendimento.

### 4.2. Fluxo de Cancelamento/Remoção
Se uma sessão contendo integrações for deletada do sistema:
- O backend verifica se existe `outlook_event_id` e envia uma requisição `DELETE` para a API do Microsoft Graph `/me/events/{outlook_event_id}` para remover o compromisso automaticamente da agenda do Outlook do profissional.
- O mesmo é feito com o evento associado do Google Calendar, se houver.

---

## 5. Principais Queries SQL

Abaixo estão listadas as principais consultas utilizadas pelo backend Golang para garantir isolamento e integridade das regras:

### 5.1. Busca de Dados de Paciente com Isolamento
```sql
-- Busca dados pessoais de um paciente específico (garantindo que pertence ao psicólogo logado)
SELECT id, name, phone, birthdate, age, profession, company, city, state, marital_status, consultation_fee 
FROM patients 
WHERE id = ? AND psychologist_id = ?;
```

### 5.2. Inserção de Novo Paciente
```sql
-- Insere um novo paciente na base
INSERT INTO patients (id, psychologist_id, name, phone, birthdate, age, profession, company, city, state, marital_status, consultation_fee)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
```

### 5.3. Salvar Anotações de Atendimento (Upsert da Primeira Análise)
```sql
-- Executa o salvamento ou atualização da Primeira Análise do paciente
INSERT INTO first_analysis (
    patient_id, main_complaint, symptom_diagnosis, developmental_influence, situational_issues,
    biological_factors, strengths_resources, addictions, stimuli, thoughts, behaviors, affects, physiological,
    treatment_goals, treatment_plan, updated_at
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP)
ON CONFLICT(patient_id) DO UPDATE SET
    main_complaint = excluded.main_complaint,
    symptom_diagnosis = excluded.symptom_diagnosis,
    developmental_influence = excluded.developmental_influence,
    situational_issues = excluded.situational_issues,
    biological_factors = excluded.biological_factors,
    strengths_resources = excluded.strengths_resources,
    addictions = excluded.addictions,
    stimuli = excluded.stimuli,
    thoughts = excluded.thoughts,
    behaviors = excluded.behaviors,
    affects = excluded.affects,
    physiological = excluded.physiological,
    treatment_goals = excluded.treatment_goals,
    treatment_plan = excluded.treatment_plan,
    updated_at = CURRENT_TIMESTAMP;
```

### 5.4. Dashboard: Métricas Contábeis e de Atendimentos
```sql
-- Faturamento Total (Pago) por mês do Psicólogo logado
SELECT 
    strftime('%Y-%m', s.session_date) AS mes, 
    SUM(p.consultation_fee) AS total_recebido
FROM sessions s
JOIN patients p ON s.patient_id = p.id
WHERE p.psychologist_id = ? AND s.status = 'pago'
GROUP BY mes
ORDER BY mes DESC;

-- Total Pendente a Receber
SELECT 
    SUM(p.consultation_fee) AS total_pendente
FROM sessions s
JOIN patients p ON s.patient_id = p.id
WHERE p.psychologist_id = ? AND s.status = 'pendente';
```
*(Nota: As queries acima utilizam a sintaxe `strftime` do SQLite. Em produção no MySQL do RDS, a query de faturamento mensal utilizará `DATE_FORMAT(s.session_date, '%Y-%m')` ou uma função equivalente suportada pelo dialeto MySQL, tratada em tempo de execução de acordo com o driver configurado).*
