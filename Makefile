.PHONY: up
up:
	docker-compose up -d

.PHONY: down
down:
	docker-compose down

.PHONY: run
run:
	go run cmd/api/*.go

.PHONY: createmigration
create_migration:
	migrate create -ext=sql -dir=sql/migrations -seq init

.PHONY: migrate
migrate:
	migrate -path=sql/migrations -database "postgres://postgres:postgres@localhost:5432/mcc?sslmode=disable" -verbose up

.PHONY: migrate_down
migrate_down:
	migrate -path=sql/migrations -database "postgres://postgres:postgres@localhost:5432/mcc?sslmode=disable" -verbose down

.PHONY: generate
generate:
	sqlc generate

.PHONY: wire
wire:
	wire cmd/server/wire.go
