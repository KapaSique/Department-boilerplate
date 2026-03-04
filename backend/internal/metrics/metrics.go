package metrics

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// HTTP metrics
	httpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint", "status"},
	)

	httpRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)

	// Business metrics
	pageViews = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "page_views_total",
			Help: "Total number of page views",
		},
		[]string{"page"},
	)

	newsViews = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "news_views_total",
			Help: "Total number of news article views",
		},
	)

	documentDownloads = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "document_downloads_total",
			Help: "Total number of document downloads",
		},
	)

	budgetReportViews = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "budget_report_views_total",
			Help: "Total number of budget report views",
		},
	)

	// Database metrics
	dbQueriesTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "db_queries_total",
			Help: "Total number of database queries",
		},
		[]string{"query_type", "status"},
	)

	dbQueryDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "db_query_duration_seconds",
			Help:    "Database query duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"query_type"},
	)
)

func InitMetrics() {
	// Metrics are automatically registered via promauto
}

// MetricsMiddleware tracks HTTP request metrics
func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Wrap response writer to capture status code
		wrapped := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(wrapped, r)

		duration := time.Since(start).Seconds()

		route := r.URL.Path
		if routeMatch := mux.CurrentRoute(r); routeMatch != nil {
			if pathTemplate, err := routeMatch.GetPathTemplate(); err == nil {
				route = pathTemplate
			}
		}

		httpRequestsTotal.WithLabelValues(
			r.Method,
			route,
			strconv.Itoa(wrapped.statusCode),
		).Inc()

		httpRequestDuration.WithLabelValues(
			r.Method,
			route,
		).Observe(duration)
	})
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// Business metric helpers
func RecordPageView(page string) {
	pageViews.WithLabelValues(page).Inc()
}

func RecordNewsView() {
	newsViews.Inc()
}

func RecordDocumentDownload() {
	documentDownloads.Inc()
}

func RecordBudgetReportView() {
	budgetReportViews.Inc()
}

// Database metric helpers
func RecordDBQuery(queryType, status string, duration time.Duration) {
	dbQueriesTotal.WithLabelValues(queryType, status).Inc()
	dbQueryDuration.WithLabelValues(queryType).Observe(duration.Seconds())
}
