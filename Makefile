.PHONY: help

VERSION ?= local-snapshot
SHELL := /bin/bash

help: ## Display this help page
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[33m%-30s\033[0m %s\n", $$1, $$2}'

run: ## Run funnel locally
	@go run main.go

build: ## Build the application
	@mkdir -p dist/
	@go build -o dist/funnel main.go

build-image: ## Build OCI image
	pack build ghcr.io/thymesave/funnel:$(VERSION) \
		--buildpack gcr.io/paketo-buildpacks/go \
		--builder paketobuildpacks/builder:tiny

push-image: ## Push OCR image using docker cli
	docker push ghcr.io/thymesave/funnel:$(VERSION)

test: ## Run tests
	@go test -v ./...

test-coverage: ## Run tests and measure coverage
	@go test -covermode=count -coverprofile=/tmp/count.out -v ./...

test-coverage-report: test-coverage ## Run test and display coverage report in browser
	@go tool cover -html=/tmp/count.out
