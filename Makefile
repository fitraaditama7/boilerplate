lint-prepare:
	@echo "Installing golangci-lint" 
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

lint:
	./bin/golangci-lint run ./...

test: 
	@go test -v -cover -covermode=atomic ./...

unittest: 
	@go test -short  ./...

migrate: 
	@go run ./migrations/migrate.go

run: 
	@go run main.go

download:
	@go mod download

.PHONY: clean download unittest test run lint-prepare lint