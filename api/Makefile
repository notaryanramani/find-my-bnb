build:
	@go build -o bin/main

run: build
	@./bin/fs

test:
	@go clean -testcache && go test -v ./...