# Этап сборки
FROM golang:1.24 AS builder

WORKDIR /app

# Установим зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходники
COPY . .

# Собираем бинарник
RUN CGO_ENABLED=0 GOOS=linux go build -o orders-service ./cmd/orders-service

# Этап исполнения
FROM gcr.io/distroless/base-debian11

WORKDIR /app

# Копируем бинарник из билдера
COPY --from=builder /app/orders-service /app/orders-service
COPY --from=builder /app/config/local.yaml /app/config/local.yaml
COPY --from=builder /app/internal/db/schema.sql /app/internal/db/schema.sql

EXPOSE 50052

# Запуск сервиса
CMD ["/app/orders-service", "--config=config/local.yaml"]
