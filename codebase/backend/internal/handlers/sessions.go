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

// ListSessions godoc
// @Summary      Listar sessões do paciente
// @Description  Retorna a lista cronológica de sessões/atendimentos de um paciente (implementação real na Sprint 7)
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
		json.NewEncoder(w).Encode([]models.Session{})
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
		w.WriteHeader(http.StatusNotImplemented)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Implementação na Sprint 7"})
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
		w.WriteHeader(http.StatusNotImplemented)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Implementação na Sprint 7"})
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
