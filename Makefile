DB_URL=postgres://postgres:password@localhost:5432/dbname

SERVICE?=order

MIGRATIONS_DIR=$(SERVICE)/migrations

GOOSE=goose -dir $(MIGRATIONS_DIR) postgres $(DB_URL)

.PHONY: run migrate-up migrate-down migrate-status migrate-create migrate-reset

run:
	go run ./$(SERVICE)/cmd

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