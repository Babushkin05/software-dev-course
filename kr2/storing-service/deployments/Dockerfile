# syntax=docker/dockerfile:1.4

FROM golang:1.24 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app

# Генерация gRPC 
RUN make proto

RUN go build -o storing-service ./cmd/main.go

# --- Minimal final image ---
FROM gcr.io/distroless/base-debian11

WORKDIR /app

COPY --from=builder /app/storing-service .
COPY --from=builder /app/config/local.yaml ./config/local.yaml
COPY --from=builder /app/.env .env

EXPOSE 50051

ENTRYPOINT ["./storing-service", "--config=config/local.yaml"]
