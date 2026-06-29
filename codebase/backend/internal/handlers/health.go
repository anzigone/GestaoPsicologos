package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type healthResponse struct {
	Status   string `json:"status"`
	Database string `json:"database"`
}

func Health(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dbStatus := "connected"
		if err := db.Ping(); err != nil {
			dbStatus = "disconnected"
		}

		statusCode := http.StatusOK
		status := "OK"
		if dbStatus != "connected" {
			status = "ERROR"
			statusCode = http.StatusServiceUnavailable
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(healthResponse{Status: status, Database: dbStatus})
	}
}
