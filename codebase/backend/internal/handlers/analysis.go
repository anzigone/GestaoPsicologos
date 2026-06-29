package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/anzigone/GestaoPsicologos/backend/internal/models"
)

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
func GetAnalysis() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.FirstAnalysis{
			PatientID:              "660e8400-e29b-41d4-a716-446655440001",
			MainComplaint:          "Ansiedade generalizada e dificuldades no trabalho",
			SymptomDiagnosis:       "TAG - Transtorno de Ansiedade Generalizada",
			DevelopmentalInfluence: "Infância com cobranças excessivas dos pais",
			SituationalIssues:      "Alta pressão no ambiente de trabalho atual",
			BiologicalFactors:      "Histórico familiar de ansiedade, sem uso de medicação",
			StrengthsResources:     "Alta resiliência, rede de suporte familiar sólida",
			Addictions:             "Nenhuma",
			Stimuli:                "Ambientes com muitas pessoas e prazos curtos",
			Thoughts:               "Pensamentos catastrofizantes e ruminação",
			Behaviors:              "Evitação de situações sociais e procrastinação",
			Affects:                "Medo intenso, preocupação excessiva e irritabilidade",
			Physiological:          "Taquicardia, sudorese e insônia",
			TreatmentGoals:         "Redução da ansiedade e melhora do funcionamento diário",
			TreatmentPlan:          "TCC com foco em reestruturação cognitiva, 12 sessões quinzenais",
			UpdatedAt:              "2026-06-15T14:30:00Z",
		})
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
func UpdateAnalysis() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.FirstAnalysis{
			PatientID:              "660e8400-e29b-41d4-a716-446655440001",
			MainComplaint:          "Ansiedade generalizada e dificuldades no trabalho",
			SymptomDiagnosis:       "TAG - Transtorno de Ansiedade Generalizada",
			DevelopmentalInfluence: "Infância com cobranças excessivas dos pais",
			SituationalIssues:      "Alta pressão no ambiente de trabalho atual",
			BiologicalFactors:      "Histórico familiar de ansiedade, sem uso de medicação",
			StrengthsResources:     "Alta resiliência, rede de suporte familiar sólida",
			Addictions:             "Nenhuma",
			Stimuli:                "Ambientes com muitas pessoas e prazos curtos",
			Thoughts:               "Pensamentos catastrofizantes e ruminação",
			Behaviors:              "Evitação de situações sociais e procrastinação",
			Affects:                "Medo intenso, preocupação excessiva e irritabilidade",
			Physiological:          "Taquicardia, sudorese e insônia",
			TreatmentGoals:         "Redução da ansiedade e melhora do funcionamento diário",
			TreatmentPlan:          "TCC com foco em reestruturação cognitiva, 12 sessões quinzenais",
			UpdatedAt:              "2026-06-29T10:00:00Z",
		})
	}
}
