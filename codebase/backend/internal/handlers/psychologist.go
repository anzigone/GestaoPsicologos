package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	mw "github.com/anzigone/GestaoPsicologos/backend/internal/middleware"
	"github.com/anzigone/GestaoPsicologos/backend/internal/models"
)

// UpdatePsychologistRequest holds the updatable profile fields.
type UpdatePsychologistRequest struct {
	Name            string  `json:"name" example:"Dra. Ana Beatriz Santos"`
	CRP             string  `json:"crp" example:"06/123456"`
	Specialty       string  `json:"specialty" example:"Terapia Cognitivo Comportamental"`
	Phone           string  `json:"phone" example:"(11) 99999-9999"`
	BaseFee         float64 `json:"base_fee" example:"180.00"`
	PackageSessions int     `json:"package_sessions" example:"4"`
	PackageFee      float64 `json:"package_fee" example:"680.00"`
}

const querySelectUserByID = `
	SELECT id, email, role, name,
	       COALESCE(crp,''), COALESCE(specialty,''), COALESCE(phone,''),
	       COALESCE(base_fee,0), COALESCE(package_sessions,0), COALESCE(package_fee,0),
	       created_at, updated_at
	FROM users WHERE id = ?`

func scanUserRow(row *sql.Row) (models.User, error) {
	var u models.User
	err := row.Scan(
		&u.ID, &u.Email, &u.Role, &u.Name,
		&u.CRP, &u.Specialty, &u.Phone,
		&u.BaseFee, &u.PackageSessions, &u.PackageFee,
		&u.CreatedAt, &u.UpdatedAt,
	)
	return u, err
}

// GetPsychologist godoc
// @Summary      Obter perfil do psicólogo
// @Description  Retorna os dados cadastrais do psicólogo autenticado
// @Tags         Psicólogo
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  models.User
// @Failure      401  {object}  models.ErrorResponse
// @Failure      404  {object}  models.ErrorResponse
// @Router       /api/psychologist [get]
func GetPsychologist(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		u, err := scanUserRow(db.QueryRow(querySelectUserByID, mw.UserIDFromContext(r)))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Usuário não encontrado"})
			return
		}
		json.NewEncoder(w).Encode(u)
	}
}

// UpdatePsychologist godoc
// @Summary      Atualizar perfil do psicólogo
// @Description  Atualiza os dados cadastrais e valores de consulta do psicólogo autenticado
// @Tags         Psicólogo
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        body  body      UpdatePsychologistRequest  true  "Dados atualizados do psicólogo"
// @Success      200   {object}  models.User
// @Failure      400   {object}  models.ErrorResponse
// @Failure      401   {object}  models.ErrorResponse
// @Failure      500   {object}  models.ErrorResponse
// @Router       /api/psychologist [put]
func UpdatePsychologist(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		userID := mw.UserIDFromContext(r)

		var req UpdatePsychologistRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Name == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Nome é obrigatório"})
			return
		}

		now := time.Now().UTC().Format(time.RFC3339)
		_, err := db.Exec(
			`UPDATE users SET name=?, crp=?, specialty=?, phone=?, base_fee=?, package_sessions=?, package_fee=?, updated_at=? WHERE id=?`,
			req.Name, req.CRP, req.Specialty, req.Phone, req.BaseFee, req.PackageSessions, req.PackageFee, now, userID,
		)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Erro ao atualizar dados"})
			return
		}

		u, _ := scanUserRow(db.QueryRow(querySelectUserByID, userID))
		json.NewEncoder(w).Encode(u)
	}
}
