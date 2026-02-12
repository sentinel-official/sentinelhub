# Build stage
FROM golang:1.26-alpine3.23 AS build

# Set working directory
WORKDIR /root

# Install build dependencies
RUN apk add --no-cache \
    build-base \
    ca-certificates \
    git \
    linux-headers \
    wget

# Cache Go modules
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# Copy source code
COPY . .

# Download and install CosmWasm static library
RUN unset GOTOOLCHAIN && \
    ARCH=$(uname -m) && \
    WASM_VERSION=$(go list -m all | grep github.com/CosmWasm/wasmvm | awk '{print $NF}') && \
    wget -q -O /usr/local/lib/libwasmvm_muslc.a \
        https://github.com/CosmWasm/wasmvm/releases/download/${WASM_VERSION}/libwasmvm_muslc.${ARCH}.a

# Build sentinelhub
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    STATIC=true make --jobs="$(nproc)" build

# Runtime stage
FROM alpine:3.23

# Copy the built binaries from build stage
COPY --from=build /root/bin/sentinelhub /usr/local/bin/sentinelhub

ENTRYPOINT ["sentinelhub"]
