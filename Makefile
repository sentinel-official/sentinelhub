.DEFAULT_GOAL := help

VERSION   := $(shell git describe --tags --always --dirty 2>/dev/null | sed 's/^v//')
COMMIT    := $(shell git log -1 --format='%H' 2>/dev/null || echo "unknown")
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

IMAGE ?= sentinelhub:latest

.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build           Build the binary (./bin/sentinelhub)"
	@echo "  install         Install sentinelhub into \$$(GOBIN)"
	@echo "  clean           Remove build artifacts"
	@echo "  test            Run tests"
	@echo "  test-coverage   Run tests with coverage and generate report"
	@echo "  benchmark       Run benchmarks"
	@echo "  go-lint         Run golangci-lint with auto-fix"
	@echo "  proto-gen       Generate protobuf code"
	@echo "  proto-lint      Lint protobuf definitions"
	@echo "  build-image     Build Docker image"
	@echo "  tools           Install development tools"

.PHONY: build
build:
	go build $(build_flags) -o ./bin/sentinelhub ./cmd/sentinelhub

.PHONY: install
install:
	go build $(build_flags) -o $$(go env GOBIN)/sentinelhub ./cmd/sentinelhub

.PHONY: clean
clean:
	$(RM) -r ./bin ./vendor

.PHONY: test
test:
	go test -cover -mod=readonly -v ./...

.PHONY: test-coverage
test-coverage:
	go test -covermode=atomic -coverprofile=coverage.txt -mod=readonly -timeout 15m -v ./...

.PHONY: benchmark
benchmark:
	go test -bench -mod=readonly -v ./...

.PHONY: go-lint
go-lint:
	golangci-lint run --fix

.PHONY: proto-gen
proto-gen:
	@scripts/proto-gen.sh

.PHONY: proto-lint
proto-lint:
	find proto -name *.proto -exec buf format -w {} \;

.PHONY: build-image
build-image:
	docker build --compress --file Dockerfile --force-rm --tag $(IMAGE) .

.PHONY: tools
tools:
	go install github.com/bufbuild/buf/cmd/buf@v1.57.0
	go install github.com/cosmos/gogoproto/protoc-gen-gocosmos@v1.7.0
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.4.0
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@v1.16.0
