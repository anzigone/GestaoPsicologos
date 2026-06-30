package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	mw "github.com/anzigone/GestaoPsicologos/backend/internal/middleware"
	"github.com/anzigone/GestaoPsicologos/backend/internal/models"
)

var ptBRMonths = [12]string{"Jan", "Fev", "Mar", "Abr", "Mai", "Jun", "Jul", "Ago", "Set", "Out", "Nov", "Dez"}

// GetDashboardStats godoc
// @Summary      KPIs do Dashboard
// @Description  Retorna os indicadores financeiros: faturamento total, sessões, pacientes ativos e pendências
// @Tags         Dashboard
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  models.DashboardStats
// @Failure      401  {object}  models.ErrorResponse
// @Router       /api/dashboard/stats [get]
func GetDashboardStats(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		psychologistID := mw.UserIDFromContext(r)

		var stats models.DashboardStats
		db.QueryRow(`SELECT COALESCE(SUM(p.consultation_fee),0) FROM sessions s JOIN patients p ON p.id=s.patient_id WHERE p.psychologist_id=? AND s.status='pago'`, psychologistID).Scan(&stats.TotalRevenue)
		db.QueryRow(`SELECT COUNT(*) FROM sessions s JOIN patients p ON p.id=s.patient_id WHERE p.psychologist_id=?`, psychologistID).Scan(&stats.TotalSessions)
		db.QueryRow(`SELECT COUNT(*) FROM patients WHERE psychologist_id=? AND active=1`, psychologistID).Scan(&stats.ActivePatients)
		db.QueryRow(`SELECT COALESCE(SUM(p.consultation_fee),0) FROM sessions s JOIN patients p ON p.id=s.patient_id WHERE p.psychologist_id=? AND s.status='pendente'`, psychologistID).Scan(&stats.PendingAmount)

		json.NewEncoder(w).Encode(stats)
	}
}

// GetDashboardCharts godoc
// @Summary      Dados de gráfico do Dashboard
// @Description  Retorna o faturamento mensal (sessões pagas) dos últimos 6 meses
// @Tags         Dashboard
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   models.ChartPoint
// @Failure      401  {object}  models.ErrorResponse
// @Router       /api/dashboard/charts [get]
func GetDashboardCharts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		psychologistID := mw.UserIDFromContext(r)

		now := time.Now()
		points := make([]models.ChartPoint, 6)
		keys := make([]string, 6)
		for i := 0; i < 6; i++ {
			t := now.AddDate(0, -(5 - i), 0)
			points[i] = models.ChartPoint{Month: ptBRMonths[t.Month()-1]}
			keys[i] = t.Format("2006-01")
		}

		rows, err := db.Query(`
			SELECT strftime('%Y-%m', s.session_date) as ym, COALESCE(SUM(p.consultation_fee),0)
			FROM sessions s JOIN patients p ON p.id=s.patient_id
			WHERE p.psychologist_id=? AND s.status='pago'
			  AND s.session_date >= date('now','-6 months')
			GROUP BY ym ORDER BY ym`, psychologistID)
		if err == nil {
			defer rows.Close()
			dataMap := make(map[string]float64)
			for rows.Next() {
				var ym string
				var fat float64
				rows.Scan(&ym, &fat)
				dataMap[ym] = fat
			}
			for i, key := range keys {
				points[i].Faturamento = dataMap[key]
			}
		}

		json.NewEncoder(w).Encode(points)
	}
}

// GetDashboardTransactions godoc
// @Summary      Transações do Dashboard
// @Description  Lista as sessões como transações financeiras com filtro opcional por status
// @Tags         Dashboard
// @Produce      json
// @Security     BearerAuth
// @Param        status  query     string  false  "Filtrar por status (pago/pendente)"
// @Success      200     {array}   models.Transaction
// @Failure      401     {object}  models.ErrorResponse
// @Router       /api/dashboard/transactions [get]
func GetDashboardTransactions(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		psychologistID := mw.UserIDFromContext(r)
		statusFilter := r.URL.Query().Get("status")

		query := `SELECT s.id, s.session_date, p.name, p.consultation_fee, s.status
			FROM sessions s JOIN patients p ON p.id=s.patient_id
			WHERE p.psychologist_id=?`
		args := []any{psychologistID}

		if statusFilter == "pago" || statusFilter == "pendente" {
			query += " AND s.status=?"
			args = append(args, statusFilter)
		}
		query += " ORDER BY s.session_date DESC"

		rows, err := db.Query(query, args...)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Erro ao buscar transações"})
			return
		}
		defer rows.Close()

		result := []models.Transaction{}
		for rows.Next() {
			var tx models.Transaction
			var rawDate string
			rows.Scan(&tx.ID, &rawDate, &tx.PatientName, &tx.Value, &tx.Status)
			if t, err := time.Parse(time.RFC3339, rawDate); err == nil {
				tx.Date = t.Format("02/01/2006")
			} else {
				tx.Date = rawDate
			}
			result = append(result, tx)
		}
		json.NewEncoder(w).Encode(result)
	}
}
