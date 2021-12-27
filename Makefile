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

deploy:
	@kubectl apply -f kubernetes/mysql-db-pv.yaml
	@kubectl apply -f kubernetes/mysql-db-pvc.yaml
	@kubectl apply -f kubernetes/mysql-db-deployment.yaml
	@kubectl apply -f kubernetes/mysql-db-service.yaml
	@echo "Mysql Successfully Initialized"
	@echo ""
	@kubectl apply -f kubernetes/redis-pv.yaml
	@kubectl apply -f kubernetes/redis-pvc.yaml
	@kubectl apply -f kubernetes/redis-deployment.yaml
	@kubectl apply -f kubernetes/redis-service.yaml
	@echo "Redis Successfully Initialized"
	@echo ""
	@kubectl apply -f kubernetes/app-mysql-deployment.yaml
	@kubectl apply -f kubernetes/app-mysql-service.yaml
	@echo "App Successfully Initialized"
	@echo ""


.PHONY: clean download unittest test run lint-prepare lint