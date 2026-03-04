package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/artem4ik/yakutsk-finance-dept/internal/database"
	"github.com/artem4ik/yakutsk-finance-dept/internal/metrics"
	"github.com/artem4ik/yakutsk-finance-dept/internal/models"
	"github.com/gorilla/mux"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
		"time":   time.Now().Format(time.RFC3339),
	})
}

func GetNews(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		limit := 20
		if l := r.URL.Query().Get("limit"); l != "" {
			if parsed, err := strconv.Atoi(l); err == nil {
				limit = parsed
			}
		}

		rows, err := db.Query(`
			SELECT id, title, content, published_at, created_at, updated_at
			FROM news
			ORDER BY published_at DESC
			LIMIT $1
		`, limit)

		metrics.RecordDBQuery("select_news", statusFromError(err), time.Since(start))

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var newsList []models.News
		for rows.Next() {
			var n models.News
			if err := rows.Scan(&n.ID, &n.Title, &n.Content, &n.PublishedAt, &n.CreatedAt, &n.UpdatedAt); err != nil {
				continue
			}
			newsList = append(newsList, n)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newsList)
	}
}

func GetNewsById(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		vars := mux.Vars(r)
		id := vars["id"]

		var n models.News
		err := db.QueryRow(`
			SELECT id, title, content, published_at, created_at, updated_at
			FROM news
			WHERE id = $1
		`, id).Scan(&n.ID, &n.Title, &n.Content, &n.PublishedAt, &n.CreatedAt, &n.UpdatedAt)

		metrics.RecordDBQuery("select_news_by_id", statusFromError(err), time.Since(start))
		metrics.RecordNewsView()

		if err != nil {
			http.Error(w, "News not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(n)
	}
}

func GetDocuments(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		category := r.URL.Query().Get("category")
		query := `
			SELECT id, title, description, file_url, category, published_at, created_at
			FROM documents
			ORDER BY published_at DESC
		`

		var rows *sql.Rows
		var err error

		if category != "" {
			query += " WHERE category = $1"
			rows, err = db.Query(query, category)
		} else {
			rows, err = db.Query(query)
		}

		metrics.RecordDBQuery("select_documents", statusFromError(err), time.Since(start))

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var docs []models.Document
		for rows.Next() {
			var d models.Document
			if err := rows.Scan(&d.ID, &d.Title, &d.Description, &d.FileURL, &d.Category, &d.PublishedAt, &d.CreatedAt); err != nil {
				continue
			}
			docs = append(docs, d)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(docs)
	}
}

func GetDocumentById(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		vars := mux.Vars(r)
		id := vars["id"]

		var d models.Document
		err := db.QueryRow(`
			SELECT id, title, description, file_url, category, published_at, created_at
			FROM documents
			WHERE id = $1
		`, id).Scan(&d.ID, &d.Title, &d.Description, &d.FileURL, &d.Category, &d.PublishedAt, &d.CreatedAt)

		metrics.RecordDBQuery("select_document_by_id", statusFromError(err), time.Since(start))
		metrics.RecordDocumentDownload()

		if err != nil {
			http.Error(w, "Document not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(d)
	}
}

func GetBudgetData(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		year := r.URL.Query().Get("year")
		query := `
			SELECT id, year, quarter, title, revenue, expenses, balance, file_url, created_at
			FROM budget_reports
			ORDER BY year DESC, quarter DESC
		`

		var rows *sql.Rows
		var err error

		if year != "" {
			query = `
				SELECT id, year, quarter, title, revenue, expenses, balance, file_url, created_at
				FROM budget_reports
				WHERE year = $1
				ORDER BY quarter DESC
			`
			rows, err = db.Query(query, year)
		} else {
			rows, err = db.Query(query)
		}

		metrics.RecordDBQuery("select_budget", statusFromError(err), time.Since(start))

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var reports []models.BudgetReport
		for rows.Next() {
			var br models.BudgetReport
			if err := rows.Scan(&br.ID, &br.Year, &br.Quarter, &br.Title, &br.Revenue, &br.Expenses, &br.Balance, &br.FileURL, &br.CreatedAt); err != nil {
				continue
			}
			reports = append(reports, br)
		}

		metrics.RecordBudgetReportView()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(reports)
	}
}

func GetBudgetReports(db *database.DB) http.HandlerFunc {
	return GetBudgetData(db)
}

func GetDepartmentInfo(w http.ResponseWriter, r *http.Request) {
	info := models.DepartmentInfo{
		Name:        "Департамент финансов Окружной администрации города Якутска",
		Description: "Департамент финансов осуществляет функции по составлению и исполнению бюджета города, управлению муниципальным долгом, финансовому контролю и методологическому обеспечению бюджетного процесса.",
		Address:     "677000, Республика Саха (Якутия), г. Якутск, ул. Орджоникидзе, д. 24",
		Phone:       "+7 (4112) 42-26-26",
		Email:       "depfin@yakutsk.org",
		WorkHours:   "Пн-Пт: 9:00-18:00, обед 13:00-14:00",
		Head:        "Руководитель департамента",
		Departments: []string{
			"Отдел бюджетной политики",
			"Отдел исполнения бюджета",
			"Отдел муниципального долга",
			"Отдел финансового контроля",
			"Отдел методологии",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

func GetContacts(w http.ResponseWriter, r *http.Request) {
	contacts := []models.Contact{
		{
			Department: "Приемная",
			Head:       "Секретарь",
			Phone:      "+7 (4112) 42-26-26",
			Email:      "depfin@yakutsk.org",
		},
		{
			Department: "Отдел бюджетной политики",
			Head:       "Начальник отдела",
			Phone:      "+7 (4112) 42-26-27",
			Email:      "budget@yakutsk.org",
		},
		{
			Department: "Отдел исполнения бюджета",
			Head:       "Начальник отдела",
			Phone:      "+7 (4112) 42-26-28",
			Email:      "execution@yakutsk.org",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contacts)
}

func statusFromError(err error) string {
	if err != nil {
		return "error"
	}
	return "success"
}
