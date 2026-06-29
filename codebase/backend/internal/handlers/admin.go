package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/anzigone/GestaoPsicologos/backend/internal/auth"
	"github.com/anzigone/GestaoPsicologos/backend/internal/models"
	"github.com/go-chi/chi/v5"
)

// CreateUserRequest holds data for creating a new psychologist.
type CreateUserRequest struct {
	Email     string  `json:"email" example:"dr.carlos@email.com"`
	Password  string  `json:"password" example:"senhaProvisoria123"`
	Name      string  `json:"name" example:"Dr. Carlos Eduardo Lima"`
	CRP       string  `json:"crp" example:"06/654321"`
	Specialty string  `json:"specialty" example:"Psicanálise"`
	Phone     string  `json:"phone" example:"(11) 97777-6666"`
	BaseFee   float64 `json:"base_fee" example:"180.00"`
}

// ListUsers godoc
// @Summary      Listar psicólogos
// @Description  Retorna a lista de todos os psicólogos cadastrados no sistema (acesso exclusivo do Admin Master)
// @Tags         Administração
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   models.User
// @Failure      401  {object}  models.ErrorResponse
// @Failure      403  {object}  models.ErrorResponse
// @Router       /api/admin/users [get]
func ListUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		rows, err := db.Query(
			`SELECT id, email, role, name,
			        COALESCE(crp,''), COALESCE(specialty,''), COALESCE(phone,''),
			        COALESCE(base_fee,0), COALESCE(package_sessions,0), COALESCE(package_fee,0),
			        created_at, updated_at
			 FROM users ORDER BY created_at ASC`,
		)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Erro ao listar usuários"})
			return
		}
		defer rows.Close()

		users := []models.User{}
		for rows.Next() {
			var u models.User
			if err := rows.Scan(
				&u.ID, &u.Email, &u.Role, &u.Name,
				&u.CRP, &u.Specialty, &u.Phone,
				&u.BaseFee, &u.PackageSessions, &u.PackageFee,
				&u.CreatedAt, &u.UpdatedAt,
			); err != nil {
				continue
			}
			users = append(users, u)
		}
		json.NewEncoder(w).Encode(users)
	}
}

// CreateUser godoc
// @Summary      Criar psicólogo
// @Description  Cadastra um novo psicólogo no sistema (acesso exclusivo do Admin Master)
// @Tags         Administração
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        body  body      CreateUserRequest  true  "Dados do novo psicólogo"
// @Success      201   {object}  models.User
// @Failure      400   {object}  models.ErrorResponse
// @Failure      401   {object}  models.ErrorResponse
// @Failure      403   {object}  models.ErrorResponse
// @Router       /api/admin/users [post]
func CreateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var req CreateUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Corpo inválido"})
			return
		}
		if req.Email == "" || req.Password == "" || req.Name == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Email, senha e nome são obrigatórios"})
			return
		}

		hash, err := auth.HashPassword(req.Password)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Erro ao processar senha"})
			return
		}

		id := auth.NewUUID()
		now := time.Now().UTC().Format(time.RFC3339)
		_, err = db.Exec(
			`INSERT INTO users (id, email, password_hash, role, name, crp, specialty, phone, base_fee, created_at, updated_at)
			 VALUES (?, ?, ?, 'psicologo', ?, ?, ?, ?, ?, ?, ?)`,
			id, req.Email, hash, req.Name, req.CRP, req.Specialty, req.Phone, req.BaseFee, now, now,
		)
		if err != nil {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "E-mail já cadastrado"})
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(models.User{
			ID: id, Email: req.Email, Role: "psicologo", Name: req.Name,
			CRP: req.CRP, Specialty: req.Specialty, Phone: req.Phone,
			BaseFee: req.BaseFee, CreatedAt: now, UpdatedAt: now,
		})
	}
}

// DeleteUser godoc
// @Summary      Remover psicólogo
// @Description  Remove um psicólogo do sistema (acesso exclusivo do Admin Master)
// @Tags         Administração
// @Security     BearerAuth
// @Param        id   path      string  true  "ID do psicólogo"
// @Success      204  "Psicólogo removido com sucesso"
// @Failure      401  {object}  models.ErrorResponse
// @Failure      403  {object}  models.ErrorResponse
// @Failure      404  {object}  models.ErrorResponse
// @Router       /api/admin/users/{id} [delete]
func DeleteUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		result, err := db.Exec("DELETE FROM users WHERE id = ? AND role != 'admin'", id)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Erro ao remover usuário"})
			return
		}
		if n, _ := result.RowsAffected(); n == 0 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Psicólogo não encontrado"})
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
