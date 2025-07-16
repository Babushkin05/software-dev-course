FROM golang:1.24 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o payments-service ./cmd/main.go

FROM gcr.io/distroless/base-debian11

WORKDIR /app

COPY --from=builder /app/payments-service /app/payments-service
COPY --from=builder /app/config/local.yaml /app/config/local.yaml
COPY --from=builder /app/internal/db/schema.sql /app/internal/db/schema.sql

EXPOSE 50051

ENTRYPOINT ["/app/payments-service", "--config=config/local.yaml"]
