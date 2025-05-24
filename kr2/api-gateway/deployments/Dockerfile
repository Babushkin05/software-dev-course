FROM golang:1.24-alpine AS builder

RUN apk add --no-cache \
    protobuf \
    make \
    bash

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest


WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN make proto

RUN go build -o /api-gateway ./cmd/main.go

FROM gcr.io/distroless/base-debian11

WORKDIR /app

COPY --from=builder /api-gateway /app/api-gateway
COPY --from=builder /app/config/local.yaml /app/config/local.yaml

ENTRYPOINT ["/app/api-gateway", "--config=config/local.yaml"]
