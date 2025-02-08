.PHONY: run
run:
	go mod tidy -go=1.23 -compat=1.23
	go mod vendor
	source .env && go run -mod=vendor main.go || true

.PHONY: lint
lint:
	golangci-lint -v run ./...

.PHONY: lint-fix
lint-fix:
	goimports -local github.com/thrownew/go-middlewares -w .
	go fmt ./...
	golangci-lint run -v

.PHONY: generate
generate:
	go generate ./...

.PHONY: build
build:
	docker build -t pupa .

.PHONY: up
up:
	docker compose up -d

.PHONY: down
down:
	docker compose down

cleanup: down
	docker volume rm pupa_mysql_data