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

const querySelectPatient = `
	SELECT id, psychologist_id, name,
	       COALESCE(phone,''), COALESCE(birthdate,''), COALESCE(age,0),
	       COALESCE(profession,''), COALESCE(company,''), COALESCE(city,''), COALESCE(state,''),
	       COALESCE(marital_status,''), COALESCE(consultation_fee,0),
	       active, created_at, updated_at
	FROM patients WHERE id=? AND psychologist_id=?`

func scanPatient(row *sql.Row) (models.Patient, error) {
	var p models.Patient
	var active int
	err := row.Scan(
		&p.ID, &p.PsychologistID, &p.Name,
		&p.Phone, &p.Birthdate, &p.Age,
		&p.Profession, &p.Company, &p.City, &p.State,
		&p.MaritalStatus, &p.ConsultationFee,
		&active, &p.CreatedAt, &p.UpdatedAt,
	)
	p.Active = active == 1
	return p, err
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
func ListPatients(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		psychologistID := mw.UserIDFromContext(r)
		q := "%" + r.URL.Query().Get("q") + "%"

		includeInactive := r.URL.Query().Get("include_inactive") == "1"
		activeFilter := "AND active=1"
		if includeInactive {
			activeFilter = ""
		}

		rows, err := db.Query(`
			SELECT id, psychologist_id, name,
			       COALESCE(phone,''), COALESCE(birthdate,''), COALESCE(age,0),
			       COALESCE(profession,''), COALESCE(company,''), COALESCE(city,''), COALESCE(state,''),
			       COALESCE(marital_status,''), COALESCE(consultation_fee,0),
			       active, created_at, updated_at
			FROM patients WHERE psychologist_id=? AND name LIKE ? `+activeFilter+`
			ORDER BY name ASC`, psychologistID, q)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Erro ao listar pacientes"})
			return
		}
		defer rows.Close()

		patients := []models.Patient{}
		for rows.Next() {
			var p models.Patient
			var active int
			if err := rows.Scan(
				&p.ID, &p.PsychologistID, &p.Name,
				&p.Phone, &p.Birthdate, &p.Age,
				&p.Profession, &p.Company, &p.City, &p.State,
				&p.MaritalStatus, &p.ConsultationFee,
				&active, &p.CreatedAt, &p.UpdatedAt,
			); err == nil {
				p.Active = active == 1
				patients = append(patients, p)
			}
		}
		json.NewEncoder(w).Encode(patients)
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
func GetPatient(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		p, err := scanPatient(db.QueryRow(querySelectPatient, chi.URLParam(r, "id"), mw.UserIDFromContext(r)))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Paciente não encontrado"})
			return
		}
		json.NewEncoder(w).Encode(p)
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
func CreatePatient(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		psychologistID := mw.UserIDFromContext(r)

		var req CreatePatientRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Name == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Nome é obrigatório"})
			return
		}
		if req.ConsultationFee < 0 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Valor da consulta não pode ser negativo"})
			return
		}

		id := auth.NewUUID()
		now := time.Now().UTC().Format(time.RFC3339)
		_, err := db.Exec(`
			INSERT INTO patients (id, psychologist_id, name, phone, birthdate, age, profession, company, city, state, marital_status, consultation_fee, active, created_at, updated_at)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 1, ?, ?)`,
			id, psychologistID, req.Name, req.Phone, req.Birthdate, req.Age,
			req.Profession, req.Company, req.City, req.State, req.MaritalStatus, req.ConsultationFee,
			now, now,
		)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Erro ao cadastrar paciente"})
			return
		}

		p, _ := scanPatient(db.QueryRow(querySelectPatient, id, psychologistID))
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(p)
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
func UpdatePatient(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		psychologistID := mw.UserIDFromContext(r)
		id := chi.URLParam(r, "id")

		var req CreatePatientRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Name == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Nome é obrigatório"})
			return
		}

		now := time.Now().UTC().Format(time.RFC3339)
		res, err := db.Exec(`
			UPDATE patients SET name=?, phone=?, birthdate=?, age=?, profession=?, company=?, city=?, state=?, marital_status=?, consultation_fee=?, updated_at=?
			WHERE id=? AND psychologist_id=?`,
			req.Name, req.Phone, req.Birthdate, req.Age,
			req.Profession, req.Company, req.City, req.State, req.MaritalStatus, req.ConsultationFee,
			now, id, psychologistID,
		)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Erro ao atualizar paciente"})
			return
		}
		if n, _ := res.RowsAffected(); n == 0 {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Paciente não encontrado"})
			return
		}

		p, _ := scanPatient(db.QueryRow(querySelectPatient, id, psychologistID))
		json.NewEncoder(w).Encode(p)
	}
}

// DeletePatient godoc
// @Summary      Remover paciente
// @Description  Remove permanentemente um paciente do psicólogo autenticado
// @Tags         Pacientes
// @Security     BearerAuth
// @Param        id   path      string  true  "ID do paciente"
// @Success      204  "Paciente removido com sucesso"
// @Failure      401  {object}  models.ErrorResponse
// @Failure      404  {object}  models.ErrorResponse
// @Router       /api/patients/{id} [delete]
func DeletePatient(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := db.Exec(`DELETE FROM patients WHERE id=? AND psychologist_id=?`, chi.URLParam(r, "id"), mw.UserIDFromContext(r))
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Erro ao remover paciente"})
			return
		}
		if n, _ := res.RowsAffected(); n == 0 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Paciente não encontrado"})
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

// PatchPatientActive godoc
// @Summary      Ativar ou desativar paciente
// @Description  Altera o status ativo/inativo de um paciente sem remover seus dados
// @Tags         Pacientes
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id    path      string  true  "ID do paciente"
// @Param        body  body      object  true  "{'active': true|false}"
// @Success      200   {object}  models.Patient
// @Failure      400   {object}  models.ErrorResponse
// @Failure      401   {object}  models.ErrorResponse
// @Failure      404   {object}  models.ErrorResponse
// @Router       /api/patients/{id} [patch]
func PatchPatientActive(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		psychologistID := mw.UserIDFromContext(r)
		id := chi.URLParam(r, "id")

		var req struct {
			Active bool `json:"active"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Dados inválidos"})
			return
		}

		activeVal := 0
		if req.Active {
			activeVal = 1
		}
		now := time.Now().UTC().Format(time.RFC3339)
		res, err := db.Exec(`UPDATE patients SET active=?, updated_at=? WHERE id=? AND psychologist_id=?`,
			activeVal, now, id, psychologistID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Erro ao atualizar paciente"})
			return
		}
		if n, _ := res.RowsAffected(); n == 0 {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Paciente não encontrado"})
			return
		}

		p, _ := scanPatient(db.QueryRow(querySelectPatient, id, psychologistID))
		json.NewEncoder(w).Encode(p)
	}
}

// ExportPatientPDF godoc
// @Summary      Exportar prontuário em PDF
// @Description  Gera e retorna o prontuário completo do paciente em formato PDF (implementação real na Sprint 11)
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
		pdf := []byte("%PDF-1.4\n1 0 obj<</Type/Catalog/Pages 2 0 R>>endobj\n" +
			"2 0 obj<</Type/Pages/Kids[3 0 R]/Count 1>>endobj\n" +
			"3 0 obj<</Type/Page/MediaBox[0 0 612 792]/Parent 2 0 R>>endobj\n" +
			"xref\n0 4\n0000000000 65535 f\n0000000009 00000 n\n" +
			"0000000058 00000 n\n0000000115 00000 n\n" +
			"trailer<</Size 4/Root 1 0 R>>\nstartxref\n190\n%%EOF")
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", `attachment; filename="prontuario.pdf"`)
		w.Write(pdf)
	}
}
