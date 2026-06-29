package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/anzigone/GestaoPsicologos/backend/internal/models"
)

// OAuthRedirectResponse holds the OAuth redirect URL.
type OAuthRedirectResponse struct {
	RedirectURL string `json:"redirect_url" example:"https://accounts.google.com/o/oauth2/auth?client_id=mock&redirect_uri=http://localhost:8080/api/integrations/google/callback&response_type=code&scope=https://www.googleapis.com/auth/calendar"`
}

// OAuthCallbackResponse holds the result of a successful OAuth callback.
type OAuthCallbackResponse struct {
	Provider string `json:"provider" example:"google"`
	Status   string `json:"status" example:"conectado"`
	Message  string `json:"message" example:"Integração com Google Calendar conectada com sucesso"`
}

// GoogleConnect godoc
// @Summary      Conectar Google Calendar
// @Description  Inicia o fluxo OAuth2 para conectar o Google Calendar do psicólogo. Redireciona para a tela de autorização do Google.
// @Tags         Integrações
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  OAuthRedirectResponse
// @Failure      401  {object}  models.ErrorResponse
// @Router       /api/integrations/google/connect [get]
func GoogleConnect() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(OAuthRedirectResponse{
			RedirectURL: "https://accounts.google.com/o/oauth2/auth?client_id=mock_client_id&redirect_uri=http://localhost:8080/api/integrations/google/callback&response_type=code&scope=https://www.googleapis.com/auth/calendar&access_type=offline",
		})
	}
}

// GoogleCallback godoc
// @Summary      Callback Google Calendar
// @Description  Recebe o código de autorização do Google OAuth2, troca pelo token de acesso e salva a integração
// @Tags         Integrações
// @Produce      json
// @Param        code   query     string  false  "Código de autorização retornado pelo Google"
// @Param        state  query     string  false  "State parameter para validação CSRF"
// @Success      200    {object}  OAuthCallbackResponse
// @Failure      400    {object}  models.ErrorResponse
// @Failure      401    {object}  models.ErrorResponse
// @Router       /api/integrations/google/callback [get]
func GoogleCallback() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(OAuthCallbackResponse{
			Provider: "google",
			Status:   "conectado",
			Message:  "Integração com Google Calendar conectada com sucesso",
		})
	}
}

// OutlookConnect godoc
// @Summary      Conectar Microsoft Outlook
// @Description  Inicia o fluxo OAuth2 para conectar o Outlook/Microsoft 365 do psicólogo. Redireciona para a tela de autorização da Microsoft.
// @Tags         Integrações
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  OAuthRedirectResponse
// @Failure      401  {object}  models.ErrorResponse
// @Router       /api/integrations/outlook/connect [get]
func OutlookConnect() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(OAuthRedirectResponse{
			RedirectURL: "https://login.microsoftonline.com/common/oauth2/v2.0/authorize?client_id=mock_client_id&redirect_uri=http://localhost:8080/api/integrations/outlook/callback&response_type=code&scope=offline_access%20Calendars.ReadWrite",
		})
	}
}

// OutlookCallback godoc
// @Summary      Callback Microsoft Outlook
// @Description  Recebe o código de autorização do Microsoft OAuth2, troca pelo token de acesso e salva a integração
// @Tags         Integrações
// @Produce      json
// @Param        code   query     string  false  "Código de autorização retornado pela Microsoft"
// @Param        state  query     string  false  "State parameter para validação CSRF"
// @Success      200    {object}  OAuthCallbackResponse
// @Failure      400    {object}  models.ErrorResponse
// @Failure      401    {object}  models.ErrorResponse
// @Router       /api/integrations/outlook/callback [get]
func OutlookCallback() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(OAuthCallbackResponse{
			Provider: "outlook",
			Status:   "conectado",
			Message:  "Integração com Microsoft Outlook conectada com sucesso",
		})
	}
}

// DisconnectIntegration godoc
// @Summary      Desconectar integração
// @Description  Remove os tokens OAuth salvos, desconectando a integração com o provedor de calendário
// @Tags         Integrações
// @Produce      json
// @Security     BearerAuth
// @Param        provider  path      string  true  "Provedor de calendário" Enums(google, outlook)
// @Success      200       {object}  models.MessageResponse
// @Failure      400       {object}  models.ErrorResponse
// @Failure      401       {object}  models.ErrorResponse
// @Router       /api/integrations/{provider}/disconnect [delete]
func DisconnectIntegration() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.MessageResponse{Message: "Integração desconectada com sucesso"})
	}
}
