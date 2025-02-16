include .env

build:
	@go build -o bin/main

run: build
	@./bin/main

test:
	@go clean -testcache && go test -v ./...

migrate-up:
	@migrate -database $(DB_ADDR) -path db/migrations up
