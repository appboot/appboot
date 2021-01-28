SHELL := /bin/bash
BASEDIR = $(shell pwd)

export GO111MODULE=on
export GOPROXY=https://goproxy.cn,direct
export GOSUMDB=off

APP_NAME=appboot
APP_VERSION=0.1.0
IMAGE_PREFIX=appboot/${APP_NAME}
IMAGE_NAME=${IMAGE_PREFIX}:v${APP_VERSION}
IMAGE_LATEST=${IMAGE_PREFIX}:latest

all: fmt imports mod lint test
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
.PHONY: build
build:
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
	@echo "fmt - gofmt"
	@echo "imports - goimports"
	@echo "mod - go mod tidy"
	@echo "lint - run golangci-lint"
	@echo "test - unit test"