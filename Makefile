include .env
.PHONY:
.SILENT:
.DEFAULT_GOAL := run
OUTPUT_NAME=

## deps: download all deps to local vendor
deps:
	go mod tidy && go mod vendor

## build: build application
build: deps
	CGO_ENABLED=0 GOOS=$(go env GOOS) GOARCH=$(go env GOARCH) go build -o ./bin/debug/app ./cmd/service/main.go

## run: run application
run: build
	docker compose up -d --remove-orphans app

## debug: run application in debug mode
debug: build
	docker compose up --remove-orphans debug

## release: build application
release: deps
	CGO_ENABLED=0 GOOS=$(RELEASE_OS) GOARCH=$(RELEASE_ARCH) go build -o ./bin/release/app ./cmd/service/main.go

## test: run tests
test:
	go test --short -coverprofile=cover.out -v ./...
	make test.coverage

## clean: clean bin directory
clean:
	go clean && rm -rf bin

help: Makefile
	@echo "Choose a command run in "$(APP)":"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'