package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/anzigone/GestaoPsicologos/backend/internal/models"
)

// CreatePatientRequest holds data for creating/updating a patient.
type CreatePatientRequest struct {
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
}

var mockPatients = []models.Patient{
	{
		ID:              "660e8400-e29b-41d4-a716-446655440001",
		PsychologistID:  "550e8400-e29b-41d4-a716-446655440000",
		Name:            "Ana Souza",
		Phone:           "(11) 98888-7777",
		Birthdate:       "1990-05-12",
		Age:             36,
		Profession:      "Professora",
		Company:         "Escola Estadual Dom Pedro II",
		City:            "São Paulo",
		State:           "SP",
		MaritalStatus:   "Casada",
		ConsultationFee: 200.00,
		CreatedAt:       "2026-03-01T09:00:00Z",
		UpdatedAt:       "2026-06-01T09:00:00Z",
	},
	{
		ID:              "661e8400-e29b-41d4-a716-446655440002",
		PsychologistID:  "550e8400-e29b-41d4-a716-446655440000",
		Name:            "Carlos Lima",
		Phone:           "(21) 97777-5555",
		Birthdate:       "1985-11-23",
		Age:             40,
		Profession:      "Engenheiro",
		Company:         "Construtora Lima & Associados",
		City:            "Rio de Janeiro",
		State:           "RJ",
		MaritalStatus:   "Divorciado",
		ConsultationFee: 180.00,
		CreatedAt:       "2026-04-10T11:00:00Z",
		UpdatedAt:       "2026-06-10T11:00:00Z",
	},
	{
		ID:              "662e8400-e29b-41d4-a716-446655440003",
		PsychologistID:  "550e8400-e29b-41d4-a716-446655440000",
		Name:            "Roberta Silva",
		Phone:           "(31) 96666-4444",
		Birthdate:       "1998-03-08",
		Age:             28,
		Profession:      "Designer",
		Company:         "Agência Criativa",
		City:            "Belo Horizonte",
		State:           "MG",
		MaritalStatus:   "Solteira",
		ConsultationFee: 220.00,
		CreatedAt:       "2026-05-20T14:00:00Z",
		UpdatedAt:       "2026-06-20T14:00:00Z",
	},
}

// ListPatients godoc
// @Summary      Listar pacientes
// @Description  Retorna a lista de pacientes do psicólogo autenticado, com filtro opcional por nome
// @Tags         Pacientes
// @Produce      json
// @Security     BearerAuth
// @Param        q    query     string  false  "Filtro por nome do paciente"
// @Success      200  {array}   models.Patient
// @Failure      401  {object}  models.ErrorResponse
// @Router       /api/patients [get]
func ListPatients() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("q")
		result := mockPatients
		if q != "" {
			filtered := []models.Patient{}
			for _, p := range mockPatients {
				if len(p.Name) >= len(q) {
					for i := 0; i <= len(p.Name)-len(q); i++ {
						if p.Name[i:i+len(q)] == q {
							filtered = append(filtered, p)
							break
						}
					}
				}
			}
			result = filtered
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}

// GetPatient godoc
// @Summary      Buscar paciente
// @Description  Retorna os dados completos de um paciente específico do psicólogo autenticado
// @Tags         Pacientes
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      string  true  "ID do paciente"
// @Success      200  {object}  models.Patient
// @Failure      401  {object}  models.ErrorResponse
// @Failure      404  {object}  models.ErrorResponse
// @Router       /api/patients/{id} [get]
func GetPatient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mockPatients[0])
	}
}

// CreatePatient godoc
// @Summary      Cadastrar paciente
// @Description  Cadastra um novo paciente para o psicólogo autenticado
// @Tags         Pacientes
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        body  body      CreatePatientRequest  true  "Dados do paciente"
// @Success      201   {object}  models.Patient
// @Failure      400   {object}  models.ErrorResponse
// @Failure      401   {object}  models.ErrorResponse
// @Router       /api/patients [post]
func CreatePatient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(models.Patient{
			ID:              "663e8400-e29b-41d4-a716-446655440099",
			PsychologistID:  "550e8400-e29b-41d4-a716-446655440000",
			Name:            "Novo Paciente",
			Phone:           "(11) 90000-0000",
			ConsultationFee: 200.00,
			CreatedAt:       "2026-06-29T10:00:00Z",
			UpdatedAt:       "2026-06-29T10:00:00Z",
		})
	}
}

// UpdatePatient godoc
// @Summary      Atualizar paciente
// @Description  Atualiza os dados de um paciente do psicólogo autenticado
// @Tags         Pacientes
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id    path      string                true  "ID do paciente"
// @Param        body  body      CreatePatientRequest  true  "Dados atualizados do paciente"
// @Success      200   {object}  models.Patient
// @Failure      400   {object}  models.ErrorResponse
// @Failure      401   {object}  models.ErrorResponse
// @Failure      404   {object}  models.ErrorResponse
// @Router       /api/patients/{id} [put]
func UpdatePatient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mockPatients[0])
	}
}

// DeletePatient godoc
// @Summary      Remover paciente
// @Description  Remove um paciente do psicólogo autenticado
// @Tags         Pacientes
// @Security     BearerAuth
// @Param        id   path      string  true  "ID do paciente"
// @Success      204  "Paciente removido com sucesso"
// @Failure      401  {object}  models.ErrorResponse
// @Failure      404  {object}  models.ErrorResponse
// @Router       /api/patients/{id} [delete]
func DeletePatient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}
}

// ExportPatientPDF godoc
// @Summary      Exportar prontuário em PDF
// @Description  Gera e retorna o prontuário completo do paciente (dados cadastrais, primeira análise e sessões) em formato PDF
// @Tags         Pacientes
// @Produce      application/pdf
// @Security     BearerAuth
// @Param        id   path      string  true  "ID do paciente"
// @Success      200  {file}    binary
// @Failure      401  {object}  models.ErrorResponse
// @Failure      404  {object}  models.ErrorResponse
// @Router       /api/patients/{id}/pdf [get]
func ExportPatientPDF() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Minimal valid PDF for mock purposes
		pdf := []byte("%PDF-1.4\n1 0 obj<</Type/Catalog/Pages 2 0 R>>endobj\n" +
			"2 0 obj<</Type/Pages/Kids[3 0 R]/Count 1>>endobj\n" +
			"3 0 obj<</Type/Page/MediaBox[0 0 612 792]/Parent 2 0 R>>endobj\n" +
			"xref\n0 4\n0000000000 65535 f\n0000000009 00000 n\n" +
			"0000000058 00000 n\n0000000115 00000 n\n" +
			"trailer<</Size 4/Root 1 0 R>>\nstartxref\n190\n%%EOF")
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", `attachment; filename="prontuario_mock.pdf"`)
		w.Write(pdf)
	}
}
