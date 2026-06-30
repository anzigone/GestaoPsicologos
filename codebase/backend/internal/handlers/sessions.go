package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/anzigone/GestaoPsicologos/backend/internal/auth"
	mw "github.com/anzigone/GestaoPsicologos/backend/internal/middleware"
	"github.com/anzigone/GestaoPsicologos/backend/internal/models"
	"github.com/go-chi/chi/v5"
)

// CreateSessionRequest holds data for creating/updating a session.
type CreateSessionRequest struct {
	SessionDate string `json:"session_date" example:"2026-07-15T14:00:00Z"`
	Notes       string `json:"notes,omitempty" example:"Sessão inicial de acolhimento"`
	Status      string `json:"status" example:"pendente" enums:"pago,pendente"`
}

const querySelectSession = `SELECT id, patient_id, session_date, COALESCE(notes,''), status,
	COALESCE(meet_link,''), COALESCE(outlook_event_id,''), created_at, updated_at FROM sessions`

func scanSessionRow(row *sql.Row) (models.Session, error) {
	var s models.Session
	err := row.Scan(&s.ID, &s.PatientID, &s.SessionDate, &s.Notes, &s.Status,
		&s.MeetLink, &s.OutlookEventID, &s.CreatedAt, &s.UpdatedAt)
	return s, err
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
func ListSessions(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		patientID := chi.URLParam(r, "id")
		if !patientBelongsTo(db, patientID, mw.UserIDFromContext(r)) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Paciente não encontrado"})
			return
		}

		rows, err := db.Query(querySelectSession+` WHERE patient_id=? ORDER BY session_date DESC`, patientID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Erro ao buscar sessões"})
			return
		}
		defer rows.Close()

		result := []models.Session{}
		for rows.Next() {
			var s models.Session
			rows.Scan(&s.ID, &s.PatientID, &s.SessionDate, &s.Notes, &s.Status,
				&s.MeetLink, &s.OutlookEventID, &s.CreatedAt, &s.UpdatedAt)
			result = append(result, s)
		}
		json.NewEncoder(w).Encode(result)
	}
}

// CreateSession godoc
// @Summary      Agendar sessão
// @Description  Agenda uma nova sessão para o paciente
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
func CreateSession(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		patientID := chi.URLParam(r, "id")
		if !patientBelongsTo(db, patientID, mw.UserIDFromContext(r)) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Paciente não encontrado"})
			return
		}

		var req CreateSessionRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.SessionDate == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "session_date é obrigatório"})
			return
		}
		if req.Status == "" {
			req.Status = "pendente"
		}

		id := auth.NewUUID()
		now := time.Now().UTC().Format(time.RFC3339)
		_, err := db.Exec(
			`INSERT INTO sessions (id, patient_id, session_date, notes, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)`,
			id, patientID, req.SessionDate, req.Notes, req.Status, now, now,
		)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Erro ao criar sessão"})
			return
		}

		s, err := scanSessionRow(db.QueryRow(querySelectSession+` WHERE id=?`, id))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(s)
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
func UpdateSession(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		patientID := chi.URLParam(r, "id")
		sessionID := chi.URLParam(r, "sid")
		if !patientBelongsTo(db, patientID, mw.UserIDFromContext(r)) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Paciente não encontrado"})
			return
		}

		var req CreateSessionRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Dados inválidos"})
			return
		}

		now := time.Now().UTC().Format(time.RFC3339)
		res, err := db.Exec(
			`UPDATE sessions SET notes=?, status=?, updated_at=? WHERE id=? AND patient_id=?`,
			req.Notes, req.Status, now, sessionID, patientID,
		)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Erro ao atualizar sessão"})
			return
		}
		if n, _ := res.RowsAffected(); n == 0 {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Sessão não encontrada"})
			return
		}

		s, _ := scanSessionRow(db.QueryRow(querySelectSession+` WHERE id=?`, sessionID))
		json.NewEncoder(w).Encode(s)
	}
}

// DeleteSession godoc
// @Summary      Remover sessão
// @Description  Remove uma sessão do paciente
// @Tags         Sessões
// @Security     BearerAuth
// @Param        id    path      string  true  "ID do paciente"
// @Param        sid   path      string  true  "ID da sessão"
// @Success      204  "Sessão removida com sucesso"
// @Failure      401  {object}  models.ErrorResponse
// @Failure      404  {object}  models.ErrorResponse
// @Router       /api/patients/{id}/sessions/{sid} [delete]
func DeleteSession(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		patientID := chi.URLParam(r, "id")
		sessionID := chi.URLParam(r, "sid")
		if !patientBelongsTo(db, patientID, mw.UserIDFromContext(r)) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		res, _ := db.Exec(`DELETE FROM sessions WHERE id=? AND patient_id=?`, sessionID, patientID)
		if n, _ := res.RowsAffected(); n == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
