# Start from the official Golang image for building
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install git for go mod if needed
RUN apk add --no-cache git

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Gin app
RUN go build -o main ./cmd/server

# Use a minimal image for running
FROM alpine:latest

WORKDIR /app

# Copy the built binary from builder
COPY ./config /config
COPY --from=builder /app/main .

# Expose port (default Gin port)
EXPOSE 8081

# Run the binary
ENTRYPOINT [ "/main","/config/local.yml" ]