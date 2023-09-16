.DEFAULT_GOAL := help

BIN_NAME := snakesneaks
BIN_DIR := ./bin
X_BIN_DIR := $(BIN_DIR)/goxz
VERSION := "x.x.x"

GOBIN ?= $(shell go env GOPATH)/bin

.PHONY: all
all: build ##  ##  ## 

.PHONY: build
build: ##  ##  ## 
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BIN_NAME) main.go

.PHONY: x-build
x-build: $(GOBIN)/goxz ##  ##  ## 
	goxz -d $(X_BIN_DIR) -n $(BIN_NAME) .

$(GOBIN)/goxz: ##  ##  ## 
	@go install github.com/Songmu/goxz/cmd/goxz@v0.9.1


$(GOBIN)/ghr: ##  ##  ## 
   @go install github.com/tcnksm/ghr@latest

$(GOBIN)/gobump: ##  ##  ## 
   @go install github.com/x-motemen/gobump/cmd/gobump@master

.PHONY: help
help: ## show help ## make help ## a
	@echo "--- Makefile Help ---"
	@echo ""
	@echo "Usage: make SUB_COMMAND argument_name=argument_value"
	@echo ""
	@echo "Command list:"
	@printf "\033[36m%-30s\033[0m %-80s %s\n" "[Sub command]" "[Description]" "[Example]"
	@grep -E '^[/a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | perl -pe 's%^([/a-zA-Z_-]+):.*?(##)%$$1 $$2%' | awk -F " *?## *?" '{printf "\033[36m%-30s\033[0m %-80s %-30s\n", $$1, $$2, $$3}'
