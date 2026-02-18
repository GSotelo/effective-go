.PHONY: help run build test test-variadic test-interfaces clean

.DEFAULT_GOAL := help

help: ## Show this help message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

run: ## Run all snippets
	go run .

build: ## Build the binary
	go build -o bin/snippets .

vet: ## Static AST code analysis 
	go vet ./...

test: ## Test all packages
	go test ./... -v

test-variadic: ## Test variadic package
	go test ./variadic -v

test-interfaces: ## Test interfaces package
	go test ./interfaces -v

test-maps: ## Test maps package
	go test ./maps -v

clean: ## Clean build artifacts
	rm -rf bin/