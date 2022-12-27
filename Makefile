SHELL := /bin/bash
BASEDIR = $(shell pwd)

APP_NAME=appboot
APP_VERSION=0.10.0
IMAGE_PREFIX=appboot/${APP_NAME}
IMAGE_NAME=${IMAGE_PREFIX}:v${APP_VERSION}
IMAGE_LATEST=${IMAGE_PREFIX}:latest

all: fmt imports mod lint test
install-pre-commit:
	brew install pre-commit
install-git-hooks:
	pre-commit install --hook-type commit-msg
	pre-commit install
run-pre-commit:
	pre-commit run --all-files
fmt:
	gofmt -w .
imports:
ifeq (, $(shell which goimports))
	go install golang.org/x/tools/cmd/goimports@latest
endif
	goimports -w .
mod:
	go mod tidy
lint: mod
ifeq (, $(shell which golangci-lint))
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
endif
	golangci-lint run -c .golangci.yml
.PHONY: test
test:
	sh scripts/test.sh
run:
	go run cmd/server/main.go
.PHONY: build
build: mod
	go build -o appboot cmd/appboot/main.go
	go build -o server cmd/server/main.go
build-docker:
	sh build/package/build.sh ${IMAGE_NAME}
push-docker:
	docker tag ${IMAGE_NAME} ${IMAGE_LATEST}
	docker push ${IMAGE_NAME}
	docker push ${IMAGE_LATEST}
.PHONY: web
web:
	cd web/appboot; \
	npm run serve
help:
	@echo "first - first time"
	@echo "fmt - gofmt"
	@echo "imports - goimports"
	@echo "mod - go mod tidy"
	@echo "lint - run golangci-lint"
	@echo "test - unit test"
	@echo "build - build binary"
	@echo "build-docker - build docker image"
	@echo "push-docker - push docker image"
	@echo "web - run web"
