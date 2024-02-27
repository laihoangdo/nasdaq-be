ifndef VERBOSE
MAKEFLAGS += --no-print-directory
endif

include .env
export

DOCKER_COMPOSE_COMMAND = docker compose
MIGRATE_COMMAND = $(DOCKER_COMPOSE_COMMAND) --profile tools run --rm migrate

# ----------------------------
# Main
# ----------------------------
setup: init mysql-migrate generate

init:
	$(DOCKER_COMPOSE_COMMAND) up -d

run:
	go run ./cmd/api/main.go

build:
	go build ./cmd/api/main.go

generate:
	go generate ./...

test:
	go test -mod=vendor -coverprofile=c.out -failfast -timeout 5m ./...

down:
	$(DOCKER_COMPOSE_COMMAND) down -v && docker compose -f docker-compose.testing.yaml down -v

# ----------------------------
# Database
# ----------------------------
## Create a DB migration files e.g `make migrate-create name=test`
migrate-create:
	$(MIGRATE_COMMAND) create -ext sql -dir /migrations -seq $(name)

## Run migrations up
mysql-migrate:
	$(MIGRATE_COMMAND) up

## Rollback migrations
mysql-redo:
	$(MIGRATE_COMMAND) down

## Force migrations to a specific version e.g `make mysql-force version=1`
mysql-force:
	$(MIGRATE_COMMAND) force $(version)

go-test:
	go run test.go
