SHELL := /bin/bash
BASEDIR = $(shell pwd)

APP_NAME=backend
APP_VERSION=0.1.0
IMAGE_PREFIX=appboot/${APP_NAME}
IMAGE_NAME=${IMAGE_PREFIX}:v${APP_VERSION}
IMAGE_LATEST=${IMAGE_PREFIX}:latest

fmt:
	gofmt -w .
mod:
	go mod tidy
utest: mod
	go test -coverpkg=./... -coverprofile=coverage.data ./...
build: mod
	go build -o main ./main.go
build-docker:
	docker build -t ${IMAGE_NAME} -f Dockerfile .
push-docker:
	docker push ${IMAGE_NAME}
push-docker-latest:
	docker tag ${IMAGE_NAME} ${IMAGE_LATEST}
#	docker push ${IMAGE_LATEST}
help:
	@echo "fmt - format the source code"
	@echo "mod - mod tidy"
	@echo "utest - unit test"
	@echo "build - build binary"
	@echo "build-docker - build docker image"
	@echo "push-docker - push docker image"
	@echo "push-docker-latest - push latest docker image"