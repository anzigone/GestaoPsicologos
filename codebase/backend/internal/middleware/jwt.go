package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/anzigone/GestaoPsicologos/backend/internal/auth"
)

type contextKey string

const (
	KeyUserID contextKey = "user_id"
	KeyRole   contextKey = "role"
)

// JWTRequired validates the Bearer token and injects user_id + role into context.
func JWTRequired(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			if header == "" || !strings.HasPrefix(header, "Bearer ") {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{"error": "Token não fornecido"})
				return
			}
			claims, err := auth.ParseClaims(strings.TrimPrefix(header, "Bearer "), secret)
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{"error": "Token inválido ou expirado"})
				return
			}
			ctx := context.WithValue(r.Context(), KeyUserID, claims["sub"])
			ctx = context.WithValue(ctx, KeyRole, claims["role"])
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// AdminOnly allows only users with role "admin". Must run after JWTRequired.
func AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role, _ := r.Context().Value(KeyRole).(string)
		if role != "admin" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(map[string]string{"error": "Acesso restrito ao administrador"})
			return
		}
		next.ServeHTTP(w, r)
	})
}

// UserIDFromContext extracts the authenticated user's ID from the request context.
func UserIDFromContext(r *http.Request) string {
	v, _ := r.Context().Value(KeyUserID).(string)
	return v
}
