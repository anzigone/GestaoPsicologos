package models

// User represents a psychologist or admin user.
// @Description Usuário do sistema (Psicólogo ou Administrador Master)
type User struct {
	ID              string  `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Email           string  `json:"email" example:"dr.ana@email.com"`
	Role            string  `json:"role" example:"psicologo" enums:"admin,psicologo"`
	Name            string  `json:"name" example:"Dra. Ana Beatriz Santos"`
	CRP             string  `json:"crp" example:"06/123456"`
	Specialty       string  `json:"specialty" example:"Terapia Cognitivo-Comportamental"`
	Phone           string  `json:"phone" example:"(11) 99999-9999"`
	BaseFee         float64 `json:"base_fee" example:"200.00"`
	PackageSessions int     `json:"package_sessions" example:"4"`
	PackageFee      float64 `json:"package_fee" example:"720.00"`
	CreatedAt       string  `json:"created_at" example:"2026-01-15T10:00:00Z"`
	UpdatedAt       string  `json:"updated_at" example:"2026-06-01T10:00:00Z"`
}

// Patient represents a patient/consultante.
// @Description Paciente (Consultante) cadastrado no sistema
type Patient struct {
	ID              string  `json:"id" example:"660e8400-e29b-41d4-a716-446655440001"`
	PsychologistID  string  `json:"psychologist_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Name            string  `json:"name" example:"Ana Souza"`
	Phone           string  `json:"phone" example:"(11) 98888-7777"`
	Birthdate       string  `json:"birthdate,omitempty" example:"1990-05-12"`
	Age             int     `json:"age,omitempty" example:"36"`
	Profession      string  `json:"profession,omitempty" example:"Professora"`
	Company         string  `json:"company,omitempty" example:"Escola Estadual"`
	City            string  `json:"city,omitempty" example:"São Paulo"`
	State           string  `json:"state,omitempty" example:"SP"`
	MaritalStatus   string  `json:"marital_status,omitempty" example:"Casada"`
	ConsultationFee float64 `json:"consultation_fee" example:"200.00"`
	CreatedAt       string  `json:"created_at" example:"2026-03-01T09:00:00Z"`
	UpdatedAt       string  `json:"updated_at" example:"2026-06-01T09:00:00Z"`
}

// FirstAnalysis represents the anamnesis/initial analysis for a patient.
// @Description Ficha de Primeira Análise (Anamnese) do Paciente
type FirstAnalysis struct {
	PatientID              string `json:"patient_id" example:"660e8400-e29b-41d4-a716-446655440001"`
	MainComplaint          string `json:"main_complaint" example:"Ansiedade generalizada e dificuldades no trabalho"`
	SymptomDiagnosis       string `json:"symptom_diagnosis" example:"TAG - Transtorno de Ansiedade Generalizada"`
	DevelopmentalInfluence string `json:"developmental_influence" example:"Infância com cobranças excessivas dos pais"`
	SituationalIssues      string `json:"situational_issues" example:"Alta pressão no ambiente de trabalho atual"`
	BiologicalFactors      string `json:"biological_factors" example:"Histórico familiar de ansiedade, sem uso de medicação"`
	StrengthsResources     string `json:"strengths_resources" example:"Alta resiliência, rede de suporte familiar sólida"`
	Addictions             string `json:"addictions" example:"Nenhuma"`
	Stimuli                string `json:"stimuli" example:"Ambientes com muitas pessoas e prazos curtos"`
	Thoughts               string `json:"thoughts" example:"Pensamentos catastrofizantes e ruminação"`
	Behaviors              string `json:"behaviors" example:"Evitação de situações sociais e procrastinação"`
	Affects                string `json:"affects" example:"Medo intenso, preocupação excessiva e irritabilidade"`
	Physiological          string `json:"physiological" example:"Taquicardia, sudorese e insônia"`
	TreatmentGoals         string `json:"treatment_goals" example:"Redução da ansiedade e melhora do funcionamento diário"`
	TreatmentPlan          string `json:"treatment_plan" example:"TCC com foco em reestruturação cognitiva, 12 sessões quinzenais"`
	UpdatedAt              string `json:"updated_at" example:"2026-06-15T14:30:00Z"`
}

// Session represents a therapy session/appointment.
// @Description Sessão de Atendimento (Prontuário)
type Session struct {
	ID             string `json:"id" example:"770e8400-e29b-41d4-a716-446655440002"`
	PatientID      string `json:"patient_id" example:"660e8400-e29b-41d4-a716-446655440001"`
	SessionDate    string `json:"session_date" example:"2026-07-10T14:00:00Z"`
	Notes          string `json:"notes" example:"Paciente apresentou melhora significativa na regulação emocional"`
	Status         string `json:"status" example:"pago" enums:"pago,pendente"`
	MeetLink       string `json:"meet_link,omitempty" example:"https://meet.google.com/abc-defg-hij"`
	OutlookEventID string `json:"outlook_event_id,omitempty" example:"AAMkADFmNTZhNmQ3"`
	CreatedAt      string `json:"created_at" example:"2026-07-01T10:00:00Z"`
	UpdatedAt      string `json:"updated_at" example:"2026-07-10T15:00:00Z"`
}

// ErrorResponse represents an error response.
// @Description Resposta de erro da API
type ErrorResponse struct {
	Error string `json:"error" example:"Credenciais inválidas"`
}

// MessageResponse represents a success message response.
// @Description Resposta de mensagem de sucesso
type MessageResponse struct {
	Message string `json:"message" example:"Operação realizada com sucesso"`
}
