# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2026-03-04

### Added

#### Backend
- REST API server with Go and Gorilla Mux
- PostgreSQL database integration
- Prometheus metrics collection:
  - HTTP request metrics (count, duration, status codes)
  - Business metrics (page views, news views, document downloads)
  - Database query metrics (count, duration)
- CORS configuration for frontend integration
- Health check endpoint
- API endpoints for news, documents, budget reports, and department info

#### Frontend
- Svelte SPA with modern minimalist design
- Responsive layout for mobile, tablet, and desktop
- Pages: Home, News, Documents, Budget, About, Contacts
- Professional design system:
  - Lexend + Source Sans 3 typography
  - Cyan/green government-friendly color palette
  - WCAG AAA accessibility compliance
  - Smooth animations with reduced-motion support
  - Visible focus states for keyboard navigation
- API integration with backend

#### Monitoring
- Prometheus configuration for metrics collection
- Grafana dashboards with:
  - HTTP metrics visualization
  - Business metrics tracking
  - Database performance monitoring
- Pre-configured datasources and dashboards

#### Infrastructure
- Docker Compose orchestration for all services
- Dockerfiles for backend and frontend
- Nginx configuration for frontend
- Database initialization script with sample data
- Makefile with convenient development commands

#### Documentation
- Comprehensive README with badges and banners
- Deployment guide (DEPLOYMENT.md)
- Contributing guidelines (CONTRIBUTING.md)
- Security policy (SECURITY.md)
- MIT License

#### CI/CD
- GitHub Actions workflow for:
  - Backend tests and linting
  - Frontend build verification
  - Docker image builds
  - Code quality checks

### Security
- Environment variable configuration
- Secure database connection handling
- CORS protection
- Input validation on API endpoints

---

## [Unreleased]

### Planned
- User authentication system
- Admin panel for content management
- Search functionality
- Email notifications
- PDF report generation
- Multi-language support (Russian/Yakut)
- Mobile app (React Native)

---

[1.0.0]: https://github.com/KapaSique/Department-boilerplate/releases/tag/v1.0.0
