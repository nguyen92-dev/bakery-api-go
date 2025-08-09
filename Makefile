.PHONY: migrate-up migrate-down run start

help:
	@echo "Available targets:"
	@echo "  migrate-up   - Run database migrations up"
	@echo "  migrate-down - Rollback database migrations down"
	@echo "  run          - Run the application"
	@echo "  start        - Start the application with all required services"

migrate-up:
	@echo "Running database migrations up..."
	@go run scripts/migrate.go up
	@echo "Done."
	@echo "Database migrations applied successfully."

migrate-down:
	@echo "Rolling back database migrations down..."
	@go run scripts/migrate.go down
	@echo "Done."
	@echo "Database migrations rolled back successfully."

run:
	@echo "Running server..."
	@go run main.go

start:
	@echo "Starting docker containers..."
	@docker compose up -d
	@echo "Docker containers started."
	@sleep 10
	@echo "Applying database migrations..."
	@make migrate-up
	@echo "Migrations applied."
	@make run
