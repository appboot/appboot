SHELL := /bin/bash
BASEDIR = $(shell pwd)

APP_NAME=backend
APP_VERSION=0.1.0
IMAGE_PREFIX=appboot/${APP_NAME}
IMAGE_NAME=${IMAGE_PREFIX}:v${APP_VERSION}
IMAGE_JENKINSAPI=${IMAGE_PREFIX}:jenkinsapi-v${APP_VERSION}
IMAGE_LATEST=${IMAGE_PREFIX}:latest
IMAGE_JENKINSAPI_LATEST=${IMAGE_PREFIX}:jenkinsapi-latest

fmt:
	gofmt -w .
mod:
	go mod tidy
utest: mod
	go test -coverpkg=./... -coverprofile=coverage.data ./...
build: mod
	go build -o main ./main.go
build-docker:
	docker build -t ${IMAGE_NAME} -f docker/Dockerfile .
	docker tag ${IMAGE_NAME} ${IMAGE_LATEST}
build-docker-jenkinsapi:
	docker build -t ${IMAGE_JENKINSAPI} -f docker/Dockerfile-jenkinsapi .
	docker tag ${IMAGE_JENKINSAPI} ${IMAGE_JENKINSAPI_LATEST}
push-docker:
	docker push ${IMAGE_NAME}
push-docker-jenkinsapi:
	docker push ${IMAGE_JENKINSAPI}	
push-docker-latest:
	docker push ${IMAGE_LATEST}
push-docker-jenkinsapi-latest:
	docker push ${IMAGE_JENKINSAPI_LATEST}
help:
	@echo "fmt - format the source code"
	@echo "mod - mod tidy"
	@echo "utest - unit test"
	@echo "build - build binary"
	@echo "build-docker - build docker image"
	@echo "push-docker - push docker image"
	@echo "push-docker-latest - push latest docker image"