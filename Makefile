DB_URL=postgres://postgres:password@localhost:5432/databaseName
MIGRATIONS_DIR=auth/migrations

GOOSE=goose -dir $(MIGRATIONS_DIR) postgres $(DB_URL)

.PHONY: run migrate-up migrate-down migrate-status migrate-create migrate-reset

run:
	go run ./auth/cmd

migrate-up:
	$(GOOSE) up

migrate-down:
	$(GOOSE) down

migrate-status:
	$(GOOSE) status

migrate-reset:
	$(GOOSE) reset

migrate-create:
	goose -dir $(MIGRATIONS_DIR) create $(name) sql