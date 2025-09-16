# syntax=docker/dockerfile:1

# ---- Build stage ----
FROM golang:1.24-alpine AS builder
WORKDIR /src

# Enable Go modules and better caching
ENV CGO_ENABLED=0 GO111MODULE=on

# Pre-cache dependencies
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download

# Copy source
COPY . .

# Build the binary
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -ldflags='-s -w' -o /app/server ./cmd

# ---- Runtime stage ----
FROM gcr.io/distroless/static-debian12:nonroot AS runtime
WORKDIR /app

# Copy binary and runtime assets (migrations, config if any)
COPY --from=builder /app/server /app/server
COPY --from=builder /src/migrations /app/migrations
COPY --from=builder /src/config/config.yml /app/config/config.yml

# Expose port
EXPOSE 8080

# Set non-root user by default from distroless image
USER nonroot:nonroot

# Entrypoint
ENTRYPOINT ["/app/server"]
