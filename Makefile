.PHONY: help dev dev-simple up down clean logs migrate-up migrate-down seed

help:
	@echo "Available commands:"
	@echo "  make dev          - Start development environment with hot reload"
	@echo "  make dev-simple   - Start development environment (no hot reload)"
	@echo "  make up           - Start all services"
	@echo "  make down         - Stop all services"
	@echo "  make clean        - Clean all containers and volumes"
	@echo "  make logs         - Show logs"
	@echo "  make migrate-up   - Run database migrations"
	@echo "  make migrate-down - Rollback migrations"
	@echo "  make seed         - Seed database with sample data"

dev:
	docker compose up

dev-build:
	@echo "Building for local development..."
	docker compose build --no-cache

dev-simple:
	docker compose -f docker-compose.simple.yml up

up:
	docker compose up -d

down:
	docker compose down

clean:
	docker compose down -v
	rm -rf backend/tmp/*
	rm -rf frontend/dist/*

logs:
	docker compose logs -f

migrate-up:
	docker compose exec backend go run cmd/migrate/main.go up

migrate-down:
	docker compose exec backend go run cmd/migrate/main.go down

seed:
	docker compose exec backend go run cmd/seed/main.go
