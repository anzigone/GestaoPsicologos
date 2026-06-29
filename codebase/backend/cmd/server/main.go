// @title           API Gestão Psicólogos
// @version         1.0
// @description     Webservice BFF para gestão de pacientes, atendimentos e financeiro para psicólogos.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Suporte GestaoPsi
// @contact.email  suporte@gestaopsi.com.br

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Informe o token JWT no formato: Bearer {token}

package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/anzigone/GestaoPsicologos/backend/docs"
	"github.com/anzigone/GestaoPsicologos/backend/internal/auth"
	"github.com/anzigone/GestaoPsicologos/backend/internal/config"
	"github.com/anzigone/GestaoPsicologos/backend/internal/database"
	"github.com/anzigone/GestaoPsicologos/backend/internal/handlers"
	"github.com/anzigone/GestaoPsicologos/backend/internal/logger"
	mw "github.com/anzigone/GestaoPsicologos/backend/internal/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func main() {
	cfg := config.Load()
	log := logger.New()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Error("Falha ao conectar ao banco de dados", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	if err := database.Migrate(db, cfg.DBDriver); err != nil {
		log.Error("Falha ao executar migrações", "error", err)
		os.Exit(1)
	}
	log.Info("Migrações executadas com sucesso", "driver", cfg.DBDriver)

	adminHash, err := auth.HashPassword("admin")
	if err != nil {
		log.Error("Falha ao gerar hash da senha padrão", "error", err)
		os.Exit(1)
	}
	if err := database.Seed(db, adminHash); err != nil {
		log.Error("Falha ao executar seed", "error", err)
		os.Exit(1)
	}
	log.Info("Seed executado com sucesso")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(corsMiddleware(cfg.FrontendURL))

	// Public routes
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))
	r.Get("/api/health", handlers.Health(db))
	r.Post("/api/auth/login", handlers.Login(db, cfg.JWTSecret))

	// Protected routes (JWT required)
	r.Group(func(r chi.Router) {
		r.Use(mw.JWTRequired(cfg.JWTSecret))

		r.Post("/api/auth/change-password", handlers.ChangePassword(db))

		// Psychologist profile
		r.Get("/api/psychologist", handlers.GetPsychologist(db))
		r.Put("/api/psychologist", handlers.UpdatePsychologist(db))

		// Admin only
		r.Route("/api/admin", func(r chi.Router) {
			r.Use(mw.AdminOnly)
			r.Get("/users", handlers.ListUsers(db))
			r.Post("/users", handlers.CreateUser(db))
			r.Delete("/users/{id}", handlers.DeleteUser(db))
		})

		// Patients (mock — real implementation in Sprint 6)
		r.Route("/api/patients", func(r chi.Router) {
			r.Get("/", handlers.ListPatients())
			r.Post("/", handlers.CreatePatient())
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", handlers.GetPatient())
				r.Put("/", handlers.UpdatePatient())
				r.Delete("/", handlers.DeletePatient())
				r.Get("/pdf", handlers.ExportPatientPDF())
				r.Get("/analysis", handlers.GetAnalysis())
				r.Put("/analysis", handlers.UpdateAnalysis())
				r.Route("/sessions", func(r chi.Router) {
					r.Get("/", handlers.ListSessions())
					r.Post("/", handlers.CreateSession())
					r.Put("/{sid}", handlers.UpdateSession())
					r.Delete("/{sid}", handlers.DeleteSession())
				})
			})
		})

		// Integrations (mock — real implementation in Sprint 8)
		r.Get("/api/integrations/google/connect", handlers.GoogleConnect())
		r.Get("/api/integrations/google/callback", handlers.GoogleCallback())
		r.Get("/api/integrations/outlook/connect", handlers.OutlookConnect())
		r.Get("/api/integrations/outlook/callback", handlers.OutlookCallback())
		r.Delete("/api/integrations/{provider}/disconnect", handlers.DisconnectIntegration())
	})

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Info("Servidor iniciado", "addr", addr)

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Error("Erro fatal no servidor", "error", err)
		os.Exit(1)
	}
}

func corsMiddleware(allowedOrigin string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
