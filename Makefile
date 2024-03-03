APP_NAME := wfi

GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin
DOCKER_TEST_CONTAINER_NAME = mariatest
DOCKER_TEST_IMAGE = mariadb

.PHONY: all build clean install test test-dependency

all: build

build:
	@go mod tidy
	@GOBIN=$(GOBIN) go build -o $(GOBIN)/$(APP_NAME) $(GOBASE)/cmd

test-dependency:
	@scripts/start-mariadb-docker.sh

test:
	@go test ./...

install:
	@cp $(GOBIN)/$(APP_NAME) ~/bin/
	
clean:
	@docker kill $(DOCKER_TEST_CONTAINER_NAME)
	@docker rm $(DOCKER_TEST_CONTAINER_NAME)
	@docker rmi $(DOCKER_TEST_IMAGE)
	@rm -f $(GOBASE)/bin/*
