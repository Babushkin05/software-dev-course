# syntax=docker/dockerfile:1.4

# Билдер
FROM --platform=linux/amd64 golang:1.24 as builder

RUN apt-get update && apt-get install -y protobuf-compiler

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
ENV PATH=$PATH:/go/bin

RUN make proto
RUN CGO_ENABLED=0 go build -o storing-service ./cmd/main.go

# Финальный образ на Debian 12
FROM debian:12-slim as runner

RUN apt-get update && apt-get install -y ca-certificates

WORKDIR /app
COPY --from=builder /app/storing-service .
COPY --from=builder /app/config/local.yaml ./config/local.yaml

ENTRYPOINT ["./storing-service", "--config=config/local.yaml"]