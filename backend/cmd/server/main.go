package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/artem4ik/yakutsk-finance-dept/internal/database"
	"github.com/artem4ik/yakutsk-finance-dept/internal/handlers"
	"github.com/artem4ik/yakutsk-finance-dept/internal/metrics"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment")
	}

	// Initialize database
	db, err := database.NewDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize metrics
	metrics.InitMetrics()

	// Setup router
	r := mux.NewRouter()

	// API routes
	api := r.PathPrefix("/api/v1").Subrouter()

	// Health check
	api.HandleFunc("/health", handlers.HealthCheck).Methods("GET")

	// News endpoints
	api.HandleFunc("/news", handlers.GetNews(db)).Methods("GET")
	api.HandleFunc("/news/{id}", handlers.GetNewsById(db)).Methods("GET")

	// Documents endpoints
	api.HandleFunc("/documents", handlers.GetDocuments(db)).Methods("GET")
	api.HandleFunc("/documents/{id}", handlers.GetDocumentById(db)).Methods("GET")

	// Budget endpoints
	api.HandleFunc("/budget", handlers.GetBudgetData(db)).Methods("GET")
	api.HandleFunc("/budget/reports", handlers.GetBudgetReports(db)).Methods("GET")

	// Department info
	api.HandleFunc("/department/info", handlers.GetDepartmentInfo).Methods("GET")
	api.HandleFunc("/department/contacts", handlers.GetContacts).Methods("GET")

	// Metrics endpoint for Prometheus
	r.Handle("/metrics", promhttp.Handler())

	// Middleware
	r.Use(metrics.MetricsMiddleware)
	r.Use(loggingMiddleware)

	// CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Handler:      handler,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Server starting on port %s", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
	})
}
