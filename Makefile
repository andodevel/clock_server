SHELL := /bin/zsh

VERSION := $(shell cat ./VERSION).
LDFLAGS += -X "main.BuildTimestamp=$(shell date -u "+%Y-%m-%d %H:%M:%S")"
LDFLAGS += -X "main.Version=$(VERSION)$(shell git rev-parse --short HEAD)"

GO := GO111MODULE=on go

.PHONY: init
init:
	go get -u github.com/99designs/gqlgen
	go get -u golang.org/x/lint/golint
	go get -u golang.org/x/tools/cmd/goimports
	@echo "Install pre-commit hook"
	@chmod +x ./hack/check.sh
	@chmod +x $(shell pwd)/hooks/pre-commit
	@ln -sf $(shell pwd)/hooks/pre-commit $(shell pwd)/.git/hooks/pre-commit || true

.PHONY: setup
setup: init
	git init

.PHONY: check
check:
	@./hack/check.sh ${scope}

.PHONY: ci
ci: init
	@$(GO) mod tidy && $(GO) mod vendor

.PHONY: build
build: check
	$(GO) build -o ./tmp/go-echo-template -ldflags '$(LDFLAGS)' ./server/go-echo-template/debug

.PHONY: install
install: check
	@echo "Installing..."
	@$(GO) install -ldflags '$(LDFLAGS)' ./server/go-echo-template/debug

.PHONY: debug
debug: check
	GOOS=darwin GOARCH=amd64 $(GO) build -ldflags '$(LDFLAGS)' -tags debug -o bin/macos/debug/go-echo-template ./server/go-echo-template/debug
	GOOS=linux GOARCH=amd64 $(GO) build -ldflags '$(LDFLAGS)' -tags debug -o bin/linux/debug/go-echo-template ./server/go-echo-template/debug
	GOOS=windows GOARCH=amd64 $(GO) build -ldflags '$(LDFLAGS)' -tags debug -o bin/windows/debug/go-echo-template.exe ./server/go-echo-template/debug

.PHONY: release
release: check
	GOOS=darwin GOARCH=amd64 $(GO) build -ldflags '$(LDFLAGS)' -o bin/macos/release/go-echo-template ./server/go-echo-template/release
	GOOS=linux GOARCH=amd64 $(GO) build -ldflags '$(LDFLAGS)' -o bin/linux/release/go-echo-template ./server/go-echo-template/release
	GOOS=windows GOARCH=amd64 $(GO) build -ldflags '$(LDFLAGS)' -o bin/windows/release/go-echo-template.exe ./server/go-echo-template/release

.PHONY: docker-image
docker-image:
	docker build -t andodevel/go-echo-template:v1.10 -f ./Dockerfile .

.PHONY: clean
clean:
	@$(GO) clean ./server/go-echo-template

.PHONY: ide
ide:
	@export GO111MODULE=auto && $(GO) mod vendor

.PHONY: start
start:
	@$(GO) run ./server/go-echo-template/debug/main_debug.go

.PHONY: gql-init
gql-init:
	@[[ -d graphql ]] && cd graphql && ($(GO) run github.com/99designs/gqlgen init && rm -r server;)

.PHONY: gql
gql:
	@[[ -d graphql ]] && cd graphql && ($(GO) run github.com/99designs/gqlgen -v)
