.DEFAULT_GOAL := help

BIN_NAME := snakesneaks
BIN_DIR := ./bin
X_BIN_DIR := $(BIN_DIR)/gox
VERSION := "x.x.x"

GOBIN ?= $(shell go env GOPATH)/bin

.PHONY: all
all: build ##  ##  ## 

.PHONY: build
build: ##  ##  ## 
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BIN_NAME) main.go

.PHONY: x-build
x-build: $(GOBIN)/gox ##  ##  ## 
	mkdir -p $(X_BIN_DIR)
	gox -output="$(X_BIN_DIR)/$(BIN_NAME)-{{.OS}}-{{.Arch}}" -os="darwin linux windows" -osarch "!darwin/386 !darwin/arm" -arch="amd64 arm" .


.PHONY: clean
clean:  ##  ##  ## 
	rm -rf bin/*

$(GOBIN)/gox: ##  ##  ## 
	@go install github.com/mitchellh/gox@v1.0.1


.PHONY: help
help: ## show help ## make help ## a
	@echo "--- Makefile Help ---"
	@echo ""
	@echo "Usage: make SUB_COMMAND argument_name=argument_value"
	@echo ""
	@echo "Command list:"
	@printf "\033[36m%-30s\033[0m %-80s %s\n" "[Sub command]" "[Description]" "[Example]"
	@grep -E '^[/a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | perl -pe 's%^([/a-zA-Z_-]+):.*?(##)%$$1 $$2%' | awk -F " *?## *?" '{printf "\033[36m%-30s\033[0m %-80s %-30s\n", $$1, $$2, $$3}'
