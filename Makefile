.PHONY: migrate-up migrate-down

help:
	@echo "Available targets:"
	@echo "  migrate-up   - Run database migrations up"
	@echo "  migrate-down - Rollback database migrations down"

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