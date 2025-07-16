# Стейдж 1: билд Go-приложения
FROM golang:1.24 AS builder

WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь проект
COPY . .

# Сборка бинарника
RUN CGO_ENABLED=0 GOOS=linux go build -o api-gateway ./cmd/main.go

# Стейдж 2: минимальный runtime-образ
FROM alpine:latest

# Установка SSL-сертификатов (если нужно делать внешние HTTP-запросы)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируем бинарник из builder stage
COPY --from=builder /app/api-gateway .

# HTTP порт, можешь изменить при необходимости
EXPOSE 8080

# Запуск приложения
ENTRYPOINT ["./api-gateway", "--config=config/local.yaml"]
