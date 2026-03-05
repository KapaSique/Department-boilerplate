# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Full-stack web portal for Yakutsk Finance Department with integrated monitoring. Backend (Go), Frontend (Svelte), PostgreSQL database, Prometheus metrics, and Grafana dashboards.

## Development Commands

### Quick Start
```bash
# Automated setup and testing
./test.sh

# Manual start (requires Docker Desktop running)
make up
docker exec -i finance-dept-db psql -U postgres -d finance_dept < backend/init.sql
```

### Common Operations
```bash
make dev      # Start only postgres, prometheus, grafana (for local backend/frontend dev)
make build    # Rebuild Docker images
make down     # Stop all services
make logs     # Follow logs from all containers
make clean    # Remove volumes and prune Docker system
```

### Service URLs
- Frontend: http://localhost:3000
- Backend API: http://localhost:8080
- Prometheus: http://localhost:9090
- Grafana: http://localhost:3001 (admin/admin)
- PostgreSQL: localhost:5432

### Backend Development (without Docker)
```bash
cd backend
go mod download
go run cmd/server/main.go
```

### Frontend Development (without Docker)
```bash
cd frontend
npm install
npm run dev  # Runs on http://localhost:5173
```

### Database Operations
```bash
# Initialize/reset database
docker exec -i finance-dept-db psql -U postgres -d finance_dept < backend/init.sql

# Access PostgreSQL shell
docker exec -it finance-dept-db psql -U postgres -d finance_dept
```

## Architecture

### Backend Structure (Go)

**Entry Point**: `backend/cmd/server/main.go`
- Initializes database connection via `internal/database.NewDB()`
- Sets up Gorilla Mux router with `/api/v1` prefix
- Registers Prometheus metrics middleware globally
- Exposes `/metrics` endpoint for Prometheus scraping

**Handler Pattern**: All handlers in `internal/handlers/` follow dependency injection:
```go
func GetNews(db *database.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Handler implementation
    }
}
```

**Metrics Collection**: `internal/metrics/` provides:
- `MetricsMiddleware` - automatically tracks HTTP requests (count, duration, status)
- `RecordDBQuery()` - tracks database query performance
- Business metrics helpers: `RecordNewsView()`, `RecordDocumentDownload()`, etc.

All metrics are automatically exposed at `/metrics` for Prometheus.

**Database Layer**: `internal/database/database.go`
- Wraps `*sql.DB` with custom `DB` type
- Connection configured via environment variables (DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
- `InitSchema()` method creates tables and indexes

### Frontend Structure (Svelte)

**Entry Point**: `frontend/src/main.js` → `App.svelte`
- Uses `svelte-routing` for SPA navigation
- API client in `src/lib/api.js` makes requests to `/api/v1/*`

**Routing**: All routes defined in `App.svelte`:
- `/` - Home
- `/news` - News list
- `/news/:id` - News detail
- `/documents` - Documents
- `/budget` - Budget data
- `/about` - Department info
- `/contacts` - Contacts

**API Integration**: `src/lib/api.js` exports functions that call backend endpoints. In Docker, nginx proxies `/api` requests to backend:8080.

### Docker Architecture

**Service Dependencies** (defined in docker-compose.yml):
1. `postgres` - starts first with health check
2. `backend` - waits for postgres to be healthy
3. `frontend` - depends on backend
4. `prometheus` - scrapes backend metrics
5. `grafana` - visualizes prometheus data

**Health Checks**: PostgreSQL has `pg_isready` health check. Backend waits for this before starting to avoid connection errors.

**Networking**: All services on `finance-network` bridge network. Services reference each other by container name (e.g., backend connects to `postgres:5432`).

### Monitoring Flow

1. Backend handlers call `metrics.RecordDBQuery()` and other metric functions
2. `MetricsMiddleware` automatically tracks all HTTP requests
3. Prometheus scrapes `http://backend:8080/metrics` every 15s
4. Grafana queries Prometheus and displays dashboards
5. Pre-configured dashboard at `monitoring/grafana/dashboards/finance-dept-dashboard.json`

## Key Patterns

### Adding New API Endpoint

1. Define model in `backend/internal/models/models.go`
2. Create handler in `backend/internal/handlers/handlers.go`:
   ```go
   func GetNewEndpoint(db *database.DB) http.HandlerFunc {
       return func(w http.ResponseWriter, r *http.Request) {
           start := time.Now()
           // Query database
           metrics.RecordDBQuery("query_name", statusFromError(err), time.Since(start))
           // Return JSON
       }
   }
   ```
3. Register route in `backend/cmd/server/main.go`:
   ```go
   api.HandleFunc("/new-endpoint", handlers.GetNewEndpoint(db)).Methods("GET")
   ```
4. Add API function in `frontend/src/lib/api.js`
5. Metrics are automatically collected by middleware

### Adding New Prometheus Metric

1. Define metric in `backend/internal/metrics/metrics.go` using `promauto.NewCounter()` or similar
2. Create helper function to record metric
3. Call helper from handlers
4. Metric automatically exposed at `/metrics`

### Database Schema Changes

1. Modify `backend/internal/database/database.go` `InitSchema()` method
2. Update `backend/init.sql` with new schema
3. Reinitialize: `docker exec -i finance-dept-db psql -U postgres -d finance_dept < backend/init.sql`

## Environment Variables

Backend reads from environment (set in docker-compose.yml or .env):
- `PORT` - Server port (default: 8080)
- `DB_HOST` - PostgreSQL host (default: localhost)
- `DB_PORT` - PostgreSQL port (default: 5432)
- `DB_USER` - Database user (default: postgres)
- `DB_PASSWORD` - Database password (default: postgres)
- `DB_NAME` - Database name (default: finance_dept)

## Troubleshooting

**Backend won't start**: Check if PostgreSQL is healthy with `docker-compose ps`. Backend waits for postgres health check.

**Frontend shows 502**: Backend container may have crashed. Check logs: `docker-compose logs backend`

**Metrics not appearing**: Verify Prometheus is scraping backend by checking http://localhost:9090/targets

**Database connection refused**: Ensure postgres container is running and healthy before starting backend.
