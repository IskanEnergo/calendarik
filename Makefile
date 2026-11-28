include .env
export

export PROJECT_ROOT=$(subst /,\,$(CURDIR))

build: env-port-forward env-up migrate-up

env-up:
	docker compose up -d --wait calendarik-postgres

env-down:
	docker compose down calendarik-postgres

env-cleanup:
	docker compose down calendarik-postgres port-forwarder
	if exist $(PROJECT_ROOT)\out\pgdata rmdir /S $(PROJECT_ROOT)\out\pgdata

env-port-forward:
	docker compose up -d port-forwarder

env-port-close:
	docker compose down port-forwarder

migrate-create:
	docker compose run --rm calendarik-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-up:
	make migrate-action action=up


migrate-down:
	make migrate-action action=down

migrate-action:
	docker compose run --rm calendarik-postgres-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@calendarik-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		"$(action)"

calendarik-run:
	@set "LOGGER_FOLDER=$(PROJECT_ROOT)\out\logs" && \
	set "POSTGRES_HOST=localhost" && \
	go mod tidy && \
	go run $(PROJECT_ROOT)/cmd/calendarik/main.go