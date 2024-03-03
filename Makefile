APP_NAME := wfi

GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin

.PHONY: all build clean install

all: build

build:
	@GOBIN=$(GOBIN) go build -o $(GOBIN)/$(APP_NAME) $(GOBASE)/cmd

install:
	@cp $(GOBIN)/$(APP_NAME) ~/bin/
	
clean:
	rm -f $(GOBASE)/bin/*
