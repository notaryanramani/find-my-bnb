include .env

build:
	@go build -o bin/main

run: build
	@./bin/main

test:
	@docker-compose up -d
	@go test -v ./...
	@docker-compose down

migrate-up:
	@migrate -database $(DB_ADDR) -path db/migrations up
