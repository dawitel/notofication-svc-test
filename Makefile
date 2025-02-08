BINARY_NAME=bin/escrow_service
ENTRY_FILE=cmd/server/main.go

MIGRATE_CMD=migrate
MIGRATION_DIR=db/migrations
SCHEMA_DIR=db/schema

include .env
export $(shell sed 's/=.*//' .env)

.PHONY: build run tidy clean test fmt vet mg migrate-up migrate-down migrate-new sqlc-gen dump-schema

build: tidy fmt vet
	@mkdir -p bin
	@echo "Building the binary for escrow service..."
	@go build -o $(BINARY_NAME) $(ENTRY_FILE)
	@echo "Build completed: $(BINARY_NAME)"

run: build
	@echo "Running escrow service..."
	@./$(BINARY_NAME)

clean:
	@echo "Cleaning up..."
	@rm -rf bin
	@rm -rf log
	@rm -rf tests/unit/log
	@echo "Cleanup completed."

fmt:
	@echo "Formatting code..."
	@go fmt ./...

vet:
	@echo "Running go vet..."
	@go vet ./...

test:
	@echo "Running tests..."
	@go test ./tests/unit -cover
	@go test ./tests/integration -cover

tidy:
	@echo "Tidying up Go modules..."
	@go mod tidy

migrate-new:
ifndef name
	@echo "Error: Please provide a migration name. Usage: make migrate-new name=<name>"
	exit 1
endif
	@$(MIGRATE_CMD) create -ext sql -dir $(MIGRATION_DIR) -seq $(name)

migrate-up:
	@$(MIGRATE_CMD) -path $(MIGRATION_DIR) -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" up
	@make dump-schema

migrate-down:
	@$(MIGRATE_CMD) -path $(MIGRATION_DIR) -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" down
	@make dump-schema

dump-schema:
	@echo "Dumping schema..."
	@mkdir -p $(SCHEMA_DIR)
	@latest_version=$(shell ls $(SCHEMA_DIR) | grep -E 'dev_addispay-v[0-9]+\.sql' | sed -E 's/dev_addispay-v([0-9]+)\.sql/\1/' | sort -n | tail -1); \
	new_version=$$((latest_version + 1)); \
	schema_file="dev_addispay-v$$new_version.sql"; \
	echo "Generating schema file: $$schema_file"; \
	mysqldump --no-data \
			--skip-comments \
			--skip-add-locks \
			--skip-disable-keys \
			--user=$(DB_USER) \
			--password=$(DB_PASSWORD) \
			--host=$(DB_HOST) \
			--port=$(DB_PORT) \
			$(DB_NAME) > $(SCHEMA_DIR)/$$schema_file; \

	mv $(SCHEMA_DIR)/$$schema_file $(SCHEMA_DIR)/dev_addispay-v$$new_version.sql; \
	echo "Schema updated to version $$new_version in $(SCHEMA_DIR)/dev_addispay-v$$new_version.sql."


sqlc-gen:
	@echo "Generating SQLC..."
	@sqlc generate
