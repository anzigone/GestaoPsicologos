package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	mw "github.com/anzigone/GestaoPsicologos/backend/internal/middleware"
	"github.com/anzigone/GestaoPsicologos/backend/internal/models"
	"github.com/go-chi/chi/v5"
)

func patientBelongsTo(db *sql.DB, patientID, psychologistID string) bool {
	var count int
	db.QueryRow("SELECT COUNT(*) FROM patients WHERE id=? AND psychologist_id=?", patientID, psychologistID).Scan(&count)
	return count > 0
}

// GetAnalysis godoc
// @Summary      Buscar primeira análise
// @Description  Retorna a ficha de primeira análise (anamnese) preenchida para o paciente
// @Tags         Primeira Análise
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      string  true  "ID do paciente"
// @Success      200  {object}  models.FirstAnalysis
// @Failure      401  {object}  models.ErrorResponse
// @Failure      404  {object}  models.ErrorResponse
// @Router       /api/patients/{id}/analysis [get]
func GetAnalysis(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		patientID := chi.URLParam(r, "id")

		if !patientBelongsTo(db, patientID, mw.UserIDFromContext(r)) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Paciente não encontrado"})
			return
		}

		var a models.FirstAnalysis
		err := db.QueryRow(`
			SELECT patient_id,
			       COALESCE(main_complaint,''), COALESCE(symptom_diagnosis,''),
			       COALESCE(developmental_influence,''), COALESCE(situational_issues,''),
			       COALESCE(biological_factors,''), COALESCE(strengths_resources,''),
			       COALESCE(addictions,''), COALESCE(stimuli,''),
			       COALESCE(thoughts,''), COALESCE(behaviors,''),
			       COALESCE(affects,''), COALESCE(physiological,''),
			       COALESCE(treatment_goals,''), COALESCE(treatment_plan,''),
			       updated_at
			FROM first_analysis WHERE patient_id=?`, patientID).Scan(
			&a.PatientID,
			&a.MainComplaint, &a.SymptomDiagnosis,
			&a.DevelopmentalInfluence, &a.SituationalIssues,
			&a.BiologicalFactors, &a.StrengthsResources,
			&a.Addictions, &a.Stimuli,
			&a.Thoughts, &a.Behaviors,
			&a.Affects, &a.Physiological,
			&a.TreatmentGoals, &a.TreatmentPlan,
			&a.UpdatedAt,
		)
		if err == sql.ErrNoRows {
			json.NewEncoder(w).Encode(models.FirstAnalysis{PatientID: patientID})
			return
		}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Erro ao buscar análise"})
			return
		}
		json.NewEncoder(w).Encode(a)
	}
}

// UpdateAnalysis godoc
// @Summary      Salvar primeira análise
// @Description  Salva ou atualiza a ficha de primeira análise (anamnese) do paciente
// @Tags         Primeira Análise
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id    path      string               true  "ID do paciente"
// @Param        body  body      models.FirstAnalysis true  "Dados da primeira análise"
// @Success      200   {object}  models.FirstAnalysis
// @Failure      400   {object}  models.ErrorResponse
// @Failure      401   {object}  models.ErrorResponse
// @Failure      404   {object}  models.ErrorResponse
// @Router       /api/patients/{id}/analysis [put]
func UpdateAnalysis(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		patientID := chi.URLParam(r, "id")

		if !patientBelongsTo(db, patientID, mw.UserIDFromContext(r)) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Paciente não encontrado"})
			return
		}

		var req models.FirstAnalysis
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Dados inválidos"})
			return
		}

		now := time.Now().UTC().Format(time.RFC3339)
		_, err := db.Exec(`
			INSERT INTO first_analysis (patient_id, main_complaint, symptom_diagnosis, developmental_influence, situational_issues, biological_factors, strengths_resources, addictions, stimuli, thoughts, behaviors, affects, physiological, treatment_goals, treatment_plan, updated_at)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
			ON CONFLICT(patient_id) DO UPDATE SET
			    main_complaint=excluded.main_complaint,
			    symptom_diagnosis=excluded.symptom_diagnosis,
			    developmental_influence=excluded.developmental_influence,
			    situational_issues=excluded.situational_issues,
			    biological_factors=excluded.biological_factors,
			    strengths_resources=excluded.strengths_resources,
			    addictions=excluded.addictions,
			    stimuli=excluded.stimuli,
			    thoughts=excluded.thoughts,
			    behaviors=excluded.behaviors,
			    affects=excluded.affects,
			    physiological=excluded.physiological,
			    treatment_goals=excluded.treatment_goals,
			    treatment_plan=excluded.treatment_plan,
			    updated_at=excluded.updated_at`,
			patientID, req.MainComplaint, req.SymptomDiagnosis, req.DevelopmentalInfluence, req.SituationalIssues,
			req.BiologicalFactors, req.StrengthsResources, req.Addictions, req.Stimuli,
			req.Thoughts, req.Behaviors, req.Affects, req.Physiological,
			req.TreatmentGoals, req.TreatmentPlan, now,
		)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Erro ao salvar análise"})
			return
		}

		req.PatientID = patientID
		req.UpdatedAt = now
		json.NewEncoder(w).Encode(req)
	}
}
