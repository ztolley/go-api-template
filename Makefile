SWAG := $(shell which swag)

.PHONY: all swag init

build:
	@go build -o bin/api cmd/main.go

test:
	@go test -v ./...
	
run: build
	@./bin/api

swag:
	@echo "Generating Swagger documentation..."
	$(SWAG) init -g cmd/main.go -o ./internal/docs
	@echo "Renaming Swagger documentation files..."
	mv ./internal/docs/swagger.yaml ./internal/docs/openapi.yaml
	mv ./internal/docs/swagger.json ./internal/docs/openapi.json

init:
	@echo "Initializing project..."
	go mod tidy
	go install github.com/swaggo/swag/cmd/swag@latest

all: init swag build