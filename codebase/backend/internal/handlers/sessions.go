package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/anzigone/GestaoPsicologos/backend/internal/models"
)

// CreateSessionRequest holds data for creating/updating a session.
type CreateSessionRequest struct {
	SessionDate string `json:"session_date" example:"2026-07-15T14:00:00Z"`
	Notes       string `json:"notes,omitempty" example:"Sessão inicial de acolhimento"`
	Status      string `json:"status" example:"pendente" enums:"pago,pendente"`
}

var mockSessions = []models.Session{
	{
		ID:          "770e8400-e29b-41d4-a716-446655440002",
		PatientID:   "660e8400-e29b-41d4-a716-446655440001",
		SessionDate: "2026-06-20T14:00:00Z",
		Notes:       "Paciente apresentou melhora significativa na regulação emocional. Trabalhou técnicas de respiração.",
		Status:      "pago",
		MeetLink:    "https://meet.google.com/abc-defg-hij",
		CreatedAt:   "2026-06-20T14:00:00Z",
		UpdatedAt:   "2026-06-20T15:00:00Z",
	},
	{
		ID:          "771e8400-e29b-41d4-a716-446655440003",
		PatientID:   "660e8400-e29b-41d4-a716-446655440001",
		SessionDate: "2026-06-06T14:00:00Z",
		Notes:       "Discussão sobre estratégias de enfrentamento no ambiente de trabalho.",
		Status:      "pago",
		MeetLink:    "https://meet.google.com/def-ghij-klm",
		CreatedAt:   "2026-06-06T14:00:00Z",
		UpdatedAt:   "2026-06-06T15:00:00Z",
	},
	{
		ID:          "772e8400-e29b-41d4-a716-446655440004",
		PatientID:   "660e8400-e29b-41d4-a716-446655440001",
		SessionDate: "2026-07-10T14:00:00Z",
		Notes:       "",
		Status:      "pendente",
		MeetLink:    "https://meet.google.com/nop-qrst-uvw",
		CreatedAt:   "2026-06-29T10:00:00Z",
		UpdatedAt:   "2026-06-29T10:00:00Z",
	},
}

// ListSessions godoc
// @Summary      Listar sessões do paciente
// @Description  Retorna a lista cronológica de sessões/atendimentos de um paciente
// @Tags         Sessões
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      string  true  "ID do paciente"
// @Success      200  {array}   models.Session
// @Failure      401  {object}  models.ErrorResponse
// @Failure      404  {object}  models.ErrorResponse
// @Router       /api/patients/{id}/sessions [get]
func ListSessions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mockSessions)
	}
}

// CreateSession godoc
// @Summary      Agendar sessão
// @Description  Agenda uma nova sessão para o paciente. Se as integrações OAuth estiverem ativas, cria evento no Google Calendar e no Outlook automaticamente.
// @Tags         Sessões
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id    path      string                true  "ID do paciente"
// @Param        body  body      CreateSessionRequest  true  "Dados da sessão"
// @Success      201   {object}  models.Session
// @Failure      400   {object}  models.ErrorResponse
// @Failure      401   {object}  models.ErrorResponse
// @Failure      404   {object}  models.ErrorResponse
// @Router       /api/patients/{id}/sessions [post]
func CreateSession() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(models.Session{
			ID:             "773e8400-e29b-41d4-a716-446655440099",
			PatientID:      "660e8400-e29b-41d4-a716-446655440001",
			SessionDate:    "2026-07-15T14:00:00Z",
			Notes:          "",
			Status:         "pendente",
			MeetLink:       "https://meet.google.com/xyz-abcd-efg",
			OutlookEventID: "AAMkADFmNTZhNmQ3LTk5ZGYtNDVmZC1iMjI2",
			CreatedAt:      "2026-06-29T10:00:00Z",
			UpdatedAt:      "2026-06-29T10:00:00Z",
		})
	}
}

// UpdateSession godoc
// @Summary      Atualizar sessão
// @Description  Atualiza as notas clínicas e o status de pagamento de uma sessão
// @Tags         Sessões
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id    path      string                true  "ID do paciente"
// @Param        sid   path      string                true  "ID da sessão"
// @Param        body  body      CreateSessionRequest  true  "Dados atualizados da sessão"
// @Success      200   {object}  models.Session
// @Failure      400   {object}  models.ErrorResponse
// @Failure      401   {object}  models.ErrorResponse
// @Failure      404   {object}  models.ErrorResponse
// @Router       /api/patients/{id}/sessions/{sid} [put]
func UpdateSession() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mockSessions[0])
	}
}

// DeleteSession godoc
// @Summary      Remover sessão
// @Description  Remove uma sessão e, se aplicável, cancela automaticamente os eventos associados no Google Calendar e Outlook
// @Tags         Sessões
// @Security     BearerAuth
// @Param        id    path      string  true  "ID do paciente"
// @Param        sid   path      string  true  "ID da sessão"
// @Success      204  "Sessão removida com sucesso"
// @Failure      401  {object}  models.ErrorResponse
// @Failure      404  {object}  models.ErrorResponse
// @Router       /api/patients/{id}/sessions/{sid} [delete]
func DeleteSession() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}
}
