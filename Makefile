SHELL := /bin/bash
BASEDIR = $(shell pwd)

APP_NAME=appboot
APP_VERSION=0.1.2
IMAGE_PREFIX=appboot/${APP_NAME}
IMAGE_NAME=${IMAGE_PREFIX}:v${APP_VERSION}
IMAGE_LATEST=${IMAGE_PREFIX}:latest

all: first fmt imports mod lint test
first:
	go get golang.org/x/tools/cmd/goimports
	go get github.com/golangci/golangci-lint/cmd/golangci-lint
fmt:
	gofmt -w .
imports:
	goimports -w .
mod:
	go mod tidy
lint:
	golangci-lint run
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