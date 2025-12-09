.DEFAULT_GOAL := help

VERSION            := $(shell git describe --tags --always --dirty 2>/dev/null | sed 's/^v//')
COMMIT             := $(shell git log -1 --format='%H' 2>/dev/null || echo "unknown")
TENDERMINT_VERSION := $(shell go list -m github.com/cometbft/cometbft | sed 's/.* //')

comma      := ,
whitespace := $(empty) $(empty)

build_tags := netgo ledger
ld_flags   := -s -w \
	-X github.com/cosmos/cosmos-sdk/version.Name=sentinel \
	-X github.com/cosmos/cosmos-sdk/version.AppName=sentinelhub \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
	-X github.com/cometbft/cometbft/version.TMCoreSemVer=$(TENDERMINT_VERSION)

ifeq ($(STATIC),true)
	build_tags += muslc
	ld_flags += -linkmode=external -extldflags '-Wl,-z,muldefs -static'
endif

BUILD_TAGS := $(subst $(whitespace),$(comma),$(build_tags))
LD_FLAGS   := $(ld_flags) -X github.com/cosmos/cosmos-sdk/version.BuildTags=$(BUILD_TAGS)

build_flags = -ldflags="$(LD_FLAGS)" -mod=readonly -tags="$(BUILD_TAGS)" -trimpath

GOBIN ?= $(shell go env GOBIN)
ifeq ($(GOBIN),)
    GOBIN := $(shell go env GOPATH)/bin
endif

IMAGE ?= sentinelhub:latest

.PHONY: help
help: ## Show available targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "%-20s %s\n", $$1, $$2}'

.PHONY: build
build: ## Build the binary (./bin/sentinelhub)
	go build $(build_flags) -o ./bin/sentinelhub ./cmd/sentinelhub

.PHONY: install
install: ## Install the binary into $GOBIN
	go build $(build_flags) -o "$(GOBIN)/sentinelhub" ./cmd/sentinelhub

.PHONY: clean
clean: ## Remove build artifacts
	$(RM) -r ./bin ./vendor ./coverage.txt

.PHONY: test
test: ## Run tests
	go test -cover -mod=readonly -v ./...

.PHONY: test-coverage
test-coverage: ## Run tests with coverage and generate report
	go test -covermode=atomic -coverprofile=coverage.txt -mod=readonly -timeout 5m -v ./...

.PHONY: benchmark
benchmark: ## Run benchmarks
	go test -bench -mod=readonly -v ./...

.PHONY: go-lint
go-lint: ## Run golangci-lint with auto-fix
	golangci-lint run --fix

.PHONY: proto-gen
proto-gen: ## Generate protobuf code
	@scripts/proto-gen.sh

.PHONY: proto-lint
proto-lint: ## Lint protobuf definitions
	@find proto -name *.proto -exec buf format -w {} \;

.PHONY: build-image
build-image: ## Build Docker image
	docker build --compress --file Dockerfile --force-rm --tag $(IMAGE) .

.PHONY: tools
tools: ## Install development tools
	go install github.com/bufbuild/buf/cmd/buf@v1.57.0
	go install github.com/cosmos/gogoproto/protoc-gen-gocosmos@v1.7.2
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.5.0
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@v1.16.0
