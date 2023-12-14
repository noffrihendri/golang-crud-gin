APP_NAME=crud-service

build: dep
	CGO_ENABLED=0 GOOS=linux go build -ldflags -a -installsuffix nocgo -o ./bin ./...

dep:
	@echo ">> Downloading Dependencies"
	@go mod download

run-api:
	@echo ">> Running API Server"
	@go run main.go serve-http

swag-init:
	@echo ">> Running swagger init"
	@swag init

run-consumer:
	@echo ">> Running Consumer"
	@go run main.go run-consumer

remock:
	#https://github.com/vektra/mockery

	@echo ">> Mock Domain"
	@mockery --all --recursive --dir ./internal/domain --output ./internal/domain/mocks_domain --outpkg mocks_domain

	@echo ">> Mock Interactor"
	@mockery --all --dir ./internal/usecases/interactor --output ./internal/usecases/interactor/mocks_interactor --outpkg mocks_interactor

	@echo ">> Mock Interfaces"
	@mockery --all --recursive --dir ./internal/interfaces --output ./internal/interfaces/mocks_interface --outpkg mocks_interface

	@echo ">> Mock Infrastructure"
	@mockery --all --recursive --dir ./internal/infrastructures --output ./internal/infrastructures/mocks_infra --outpkg mocks_infra

run-test: dep
	@echo ">> Running Test"
	@go test -v -cover -count=1 -failfast -covermode=atomic ./...