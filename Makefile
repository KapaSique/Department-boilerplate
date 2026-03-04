# Makefile for Finance Department Project

.PHONY: help dev build up down logs clean

help:
	@echo "Available commands:"
	@echo "  make dev      - Start development environment"
	@echo "  make build    - Build all Docker images"
	@echo "  make up       - Start all services"
	@echo "  make down     - Stop all services"
	@echo "  make logs     - View logs"
	@echo "  make clean    - Clean up volumes and images"

dev:
	docker-compose up -d postgres prometheus grafana
	@echo "Development services started"
	@echo "Postgres: localhost:5432"
	@echo "Prometheus: http://localhost:9090"
	@echo "Grafana: http://localhost:3001 (admin/admin)"

build:
	docker-compose build

up:
	docker-compose up -d
	@echo "All services started"
	@echo "Frontend: http://localhost:3000"
	@echo "Backend: http://localhost:8080"
	@echo "Prometheus: http://localhost:9090"
	@echo "Grafana: http://localhost:3001"

down:
	docker-compose down

logs:
	docker-compose logs -f

clean:
	docker-compose down -v
	docker system prune -f
