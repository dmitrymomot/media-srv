# Build variables
LATEST_COMMIT := $$(git rev-parse HEAD)

.PHONY: help sqlc build docker up down

help: ## Show this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
%:
	@:

build: ## Build the app
	@go clean
	@CGO_ENABLED=0 \
	GOOS=linux \
	GOARCH=amd64 \
	go build \
	-a -installsuffix nocgo \
	-ldflags "-X main.buildTag=`date -u +%Y%m%d.%H%M%S`-$(LATEST_COMMIT)" \
	-o ./media-srv .

docker: ## Build docker image
	@docker build -t media-srv:latest .
	@go clean

up: ## Deploy pods to kubernetes
	@docker-compose up -d db \
	&& sleep 10 \
	&& docker-compose up -d api

down: ## Down pods
	@docker-compose down -v --rmi=local

test: ## Run all tests
	@go test ./handler
	@go test ./repository
	@go test ./resizer
	@go test ./storage
