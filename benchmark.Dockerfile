FROM inspirate789/dind-golang:1.21
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./cmd/benchmark/main.go ./
COPY ./internal ./internal
COPY ./pkg ./pkg

RUN go build -o /benchmark

EXPOSE 2112

ENTRYPOINT ["/benchmark"]
