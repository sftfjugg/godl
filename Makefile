## help: Show makefile commands
.PHONY: help
help: Makefile
	@echo "---- Project: sftfjugg/godl ----"
	@echo " Usage: make COMMAND"
	@echo
	@echo " Management Commands:"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo


FILES_TO_FORMAT ?= $(shell find . -name '*.go')

GOFLAGS ?= -trimpath


## build: Build project
.PHONY: build
build:
	@go build $(GOFLAGS) -o ./bin/ ./cmd/godl


## deps: Ensures fresh go.mod and go.sum
.PHONY: deps
deps:
	@go mod tidy
	@go mod verify


## lint: Run golangci-lint check
.PHONY: lint
lint:
	@if [ ! where golangci-lint ]; then \
		curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s $(GOLANGCI_LINT_VERSION); \
	fi;
	@echo "golangci-lint checking..."
	@golangci-lint run --deadline=30m


## format: Formats Go code
.PHONY: format
format:
	@echo "formatting code..."
	@gofmt -s -e -w $(FILES_TO_FORMAT)
