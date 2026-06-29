package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/anzigone/GestaoPsicologos/backend/internal/auth"
	"github.com/anzigone/GestaoPsicologos/backend/internal/middleware"
	"github.com/anzigone/GestaoPsicologos/backend/internal/models"
)

// LoginRequest holds login credentials.
type LoginRequest struct {
	Email    string `json:"email" example:"admin@admin.com.br"`
	Password string `json:"password" example:"admin"`
}

// LoginResponse holds the JWT token and user data.
type LoginResponse struct {
	Token string      `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	User  models.User `json:"user"`
}

// ChangePasswordRequest holds the password change payload.
type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password" example:"senhaAtual123"`
	NewPassword     string `json:"new_password" example:"novaSenha456"`
}

// Login godoc
// @Summary      Login do usuário
// @Description  Autentica um usuário (Psicólogo ou Admin) e retorna um token JWT com os dados do perfil
// @Tags         Autenticação
// @Accept       json
// @Produce      json
// @Param        body  body      LoginRequest   true  "Credenciais de acesso"
// @Success      200   {object}  LoginResponse
// @Failure      400   {object}  models.ErrorResponse
// @Failure      401   {object}  models.ErrorResponse
// @Router       /api/auth/login [post]
func Login(db *sql.DB, jwtSecret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var req LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Corpo da requisição inválido"})
			return
		}
		if req.Email == "" || req.Password == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "E-mail e senha são obrigatórios"})
			return
		}

		var user models.User
		var passwordHash string
		row := db.QueryRow(
			`SELECT id, email, password_hash, role, name,
			        COALESCE(crp,''), COALESCE(specialty,''), COALESCE(phone,''),
			        COALESCE(base_fee,0), COALESCE(package_sessions,0), COALESCE(package_fee,0),
			        created_at, updated_at
			 FROM users WHERE email = ?`, req.Email,
		)
		err := row.Scan(
			&user.ID, &user.Email, &passwordHash, &user.Role, &user.Name,
			&user.CRP, &user.Specialty, &user.Phone,
			&user.BaseFee, &user.PackageSessions, &user.PackageFee,
			&user.CreatedAt, &user.UpdatedAt,
		)
		if err == sql.ErrNoRows || !auth.CheckPassword(req.Password, passwordHash) {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Credenciais inválidas"})
			return
		}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Erro interno"})
			return
		}

		token, err := auth.GenerateToken(user.ID, user.Role, jwtSecret)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Erro ao gerar token"})
			return
		}

		json.NewEncoder(w).Encode(LoginResponse{Token: token, User: user})
	}
}

// ChangePassword godoc
// @Summary      Alterar senha
// @Description  Altera a senha do usuário autenticado validando a senha atual
// @Tags         Autenticação
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        body  body      ChangePasswordRequest  true  "Dados para troca de senha"
// @Success      200   {object}  models.MessageResponse
// @Failure      400   {object}  models.ErrorResponse
// @Failure      401   {object}  models.ErrorResponse
// @Router       /api/auth/change-password [post]
func ChangePassword(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		userID := middleware.UserIDFromContext(r)
		var req ChangePasswordRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.CurrentPassword == "" || req.NewPassword == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Dados inválidos"})
			return
		}

		var storedHash string
		if err := db.QueryRow("SELECT password_hash FROM users WHERE id = ?", userID).Scan(&storedHash); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Usuário não encontrado"})
			return
		}
		if !auth.CheckPassword(req.CurrentPassword, storedHash) {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Senha atual incorreta"})
			return
		}

		newHash, err := auth.HashPassword(req.NewPassword)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Erro ao processar senha"})
			return
		}

		now := time.Now().UTC().Format(time.RFC3339)
		if _, err := db.Exec("UPDATE users SET password_hash = ?, updated_at = ? WHERE id = ?", newHash, now, userID); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Erro ao atualizar senha"})
			return
		}

		json.NewEncoder(w).Encode(models.MessageResponse{Message: "Senha alterada com sucesso"})
	}
}
