.PHONY: help

VERSION ?= local-snapshot
SHELL := /bin/bash
NOW=$(shell date +'%y-%m-%d_%H:%M:%S')
GO_BUILD_ARGS=-ldflags "-X github.com/thymesave/funnel/pkg/buildinfo.GitSha=$(shell git rev-parse --short HEAD) -X github.com/thymesave/funnel/pkg/buildinfo.Version=$(VERSION) -X github.com/thymesave/funnel/pkg/buildinfo.BuildTime=$(NOW)"
GIT_REF=$(shell git rev-parse --short HEAD)

help: ## Display this help page
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[33m%-30s\033[0m %s\n", $$1, $$2}'

run: ## Run funnel locally
	go run main.go $(GO_BUILD_ARGS)

build: build-binaries create-checksums ## Build the application

build-binaries: ## Build all binaries for the project
	@rm -rf dist/ || true && mkdir -p dist/
	@GOOS=linux GOARCH=amd64 go build -o dist/funnel_linux-amd64-$(VERSION) $(GO_BUILD_ARGS)
	@GOOS=linux GOARCH=arm64 go build -o dist/funnel_linux-arm64-$(VERSION) $(GO_BUILD_ARGS)
	@GOOS=linux GOARCH=386 go build -o dist/funnel_linux-i386-$(VERSION) $(GO_BUILD_ARGS)

build-image: ## Build OCI image
	pack build ghcr.io/thymesave/funnel:$(VERSION) \
		--buildpack gcr.io/paketo-buildpacks/go \
		--buildpack  gcr.io/paketo-buildpacks/image-labels \
		--builder paketobuildpacks/builder:tiny \
		--env "BP_OCI_SOURCE=https://github.com/thymesave/funnel" \
		--env "BP_OCI_AUTHORS=https://github.com/thymesave/funnel/graphs/contributors" \
		--env "BP_OCI_URL=https://github.com/thymesave/funnel" \
		--env "BP_OCI_DOCUMENTATION=https://github.com/thymesave/funnel" \
		--env "BP_OCI_VENDOR=ThymeSave" \
		--env "BP_OCI_TITLE=Funnel" \
		--env "BP_OCI_REF_NAME=thymesave/funnel" \
		--env "BP_OCI_DESCRIPTION=Funnel is a core part of ThymeSave and does what the name suggests: It filters all backend requests. It includes routing, cors and authentication." \
		--env "BP_OCI_LICENSES=GPLv3" \
		--env "BP_OCI_VERSION=$(VERSION)" \
		--env "BP_OCI_REVISION=$(GIT_REF)" \
		--env "BP_OCI_CREATED=$(shell date --rfc-3339=seconds)"

push-image: ## Push OCR image using docker cli
	docker push ghcr.io/thymesave/funnel:$(VERSION)

test: ## Run tests
	@go test -v ./... $(GO_BUILD_ARGS)

test-coverage: ## Run tests and measure coverage
	@go test -covermode=count -coverprofile=/tmp/count.out -v ./... $(GO_BUILD_ARGS)

test-coverage-report: test-coverage ## Run test and display coverage report in browser
	@go tool cover -html=/tmp/count.out

create-checksums: ## Create checksums for binaries
	@find ./dist -type f -exec sh -c 'sha256sum {} | cut -d " " -f 1 > {}.sha256' {} \;
	@find ./dist ! -name '*.sha256' -type f -exec sh -c 'sha1sum {} | cut -d " " -f 1 > {}.sha1' {} \;
