SHELL := /bin/bash

VERSION := $(SHELL cat ./VERSION).
LDFLAGS += -X "main.BuildTimestamp=$(SHELL date -u "+%Y-%m-%d %H:%M:%S")"
LDFLAGS += -X "main.Version=$(VERSION)$(SHELL git rev-parse --short HEAD)"

GO := GO111MODULE=on go

# FIXME: Correct working dir and cd to this.
WORK_DIR=$(SHELL pwd)

# Initialize project. All init(s) target should be run only once to create the project.
.PHONY: init
init: dev-tools
	go get -u github.com/skip2/go-qrcode/
	go get -u github.com/99designs/gqlgen
	@echo "Install pre-commit hooks"
	@chmod +x ./hack/check.sh
	@chmod +x ./hooks/pre-commit
	@ln -sf ./hooks/pre-commit ./.git/hooks/pre-commit || true

.PHONY: git-init
git-init: 
	git init

.PHONY: gql-init
gql-init:
	@[[ -d graphql ]] && cd graphql && ($(GO) run github.com/99designs/gqlgen init && rm -r server;)

# Developement
.PHONY: dev-tools
dev-tools:
	go get -u -mod=readonly golang.org/x/lint/golint
	go get -u -mod=readonly golang.org/x/tools/cmd/goimports

.PHONY: gql
gql:
	@[[ -d graphql ]] && cd graphql && ($(GO) run github.com/99designs/gqlgen -v)

.PHONY: check
check:
	@./hack/check.sh ${scope}

.PHONY: docker-image
docker-image:
	docker build -t andodevel/clock_server:v0.0.1 -f ./Dockerfile .

# Deprecated. Please use language server.
.PHONY: ide
ide:
	@export GO111MODULE=auto && $(GO) mod vendor

# Build and run
.PHONY: build
build: check
	$(GO) build -o ./tmp/clock_server -ldflags '$(LDFLAGS)' ./server/clock_server/debug

.PHONY: test
test:
	@$(GO) test -v ./...

.PHONY: start
start:
	@$(GO) run ./server/clock_server/debug/main_debug.go

.PHONY: install
install: check
	@echo "Installing..."
	@$(GO) install -ldflags '$(LDFLAGS)' ./server/clock_server/release

.PHONY: debug
debug: check
	GOOS=darwin GOARCH=amd64 $(GO) build -ldflags '$(LDFLAGS)' -tags debug -o bin/macos/debug/clock_server ./server/clock_server/debug
	GOOS=linux GOARCH=amd64 $(GO) build -ldflags '$(LDFLAGS)' -tags debug -o bin/linux/debug/clock_server ./server/clock_server/debug
	GOOS=windows GOARCH=amd64 $(GO) build -ldflags '$(LDFLAGS)' -tags debug -o bin/windows/debug/clock_server.exe ./server/clock_server/debug

.PHONY: release
release: check
	GOOS=darwin GOARCH=amd64 $(GO) build -ldflags '$(LDFLAGS)' -o bin/macos/release/clock_server ./server/clock_server/release
	GOOS=linux GOARCH=amd64 $(GO) build -ldflags '$(LDFLAGS)' -o bin/linux/release/clock_server ./server/clock_server/release
	GOOS=windows GOARCH=amd64 $(GO) build -ldflags '$(LDFLAGS)' -o bin/windows/release/clock_server.exe ./server/clock_server/release

.PHONY: clean
clean:
	@$(GO) clean ./server/clock_server

# CI/CD
.PHONY: preci
preci: dev-tools
	@$(GO) mod tidy
	@$(GO) mod vendor
	@export PATH=$PATH:${GOPATH}\bin
	@echo "GOPATH was set: ${GOPATH}"
	@echo "Add GOPATH to OS PATH variable: ${PATH}"

.PHONY: ci
ci: preci build test
