.PHONY: mod
mod:
	go mod tidy -go=1.23 -compat=1.23
	go mod vendor

.PHONY: run
run: mod
	source .env && go run -mod=vendor main.go || true

.PHONY: run
help: mod
	source .env && go run -mod=vendor main.go help || true

.PHONY: wof
wof: mod
	source .env && go run -mod=vendor main.go wof || true

.PHONY: lint
lint:
	golangci-lint -v run ./...

.PHONY: lint-fix
lint-fix:
	goimports -local pupa -w .
	go fmt ./...
	golangci-lint run -v

.PHONY: generate
generate:
	go generate ./...

.PHONY: test-unit
test-unit:
	go clean -testcache
	go test -race -v -run Unit ./...

.PHONY: test-unit
test-integration:
	go clean -testcache
	go test -race -v -run Integration ./...

.PHONY: test
test: test-unit test-integration

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