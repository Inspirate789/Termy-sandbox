# syntax=docker/dockerfile:1

## Build
FROM golang:1.21-alpine3.17 AS build
WORKDIR /app

# Copy source code & Install dependencies
RUN go mod init tracer
COPY ./cmd/tracer/main.go ./
RUN go mod tidy

# Run
ENTRYPOINT ["go", "run", "main.go"]
