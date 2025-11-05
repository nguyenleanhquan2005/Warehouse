# Load env file
ifneq (,$(wildcard .env))
    include .env
    export
endif

MIGRATE=migrate
MIGRATIONS_DIR=database/migrations
DB_DSN=mysql://$(DB_USER):$(DB_PASS)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)

# install dev tools
tool:
	@go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.3
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ${shell go env GOPATH}/bin v2.1.6

migrate-up:
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database '$(DB_DSN)' up

migrate-down:
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database '$(DB_DSN)' down 1


up:
	docker compose up

run:
	go run cmd/main.go

mod:
	@go mod tidy
	@go mod vendor

gen:
	@swag init -g internal/framework/route/route.go -o internal/docs  --exclude pkg,db,deployment,scripts,vendor

build:
	@go build -tags=jsoniter -o product-service cmd/main.go

lint: ## Run lint
	@./scripts/lint.sh

buildx:
	docker buildx create --use

build/multi:
	docker buildx build \
		--platform linux/amd64,linux/arm64 \
		-t ap-singapore-1.ocir.io/axfnrpyfvlpv/wws_product_service:latest \
		--push .
pull:
	docker pull ap-singapore-1.ocir.io/axfnrpyfvlpv/wws_product_service:latest
import_db:
	./scripts/importdb.sh
