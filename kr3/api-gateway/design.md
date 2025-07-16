# API Gateway Design

## 📘 Описание

`api-gateway` — это шлюз, обеспечивающий взаимодействие между клиентами и микросервисами `orders-service` и `payments-service` через REST API. Он трансформирует HTTP-запросы в gRPC-вызовы и наоборот.

---

## 📦 Структура проекта

```
api-gateway/
├── api/
│   ├── orders/              # Protobuf-файлы для orders
│   ├── payments/            # Protobuf-файлы для payments
│   └── gen/                 # Сгенерированные gRPC клиенты
│
├── cmd/
│   └── main.go              # Точка входа
│
├── config/
│   └── local.yaml           # Конфигурация
│
├── deployments/             # Dockerfile, docker-compose и прочее
│
├── internal/
│   ├── client/              # gRPC клиенты
│   ├── config/              # Загрузка и структура конфига
│   ├── dto/                 # DTO структуры для входных/выходных данных
│   ├── handler/             # REST-хендлеры для Gin (orders, payments)
│   └── server/              # Инициализация и запуск Gin-сервера
│
├── docs/                    # Сгенерированная Swagger-документация
├── go.mod / go.sum
└── Makefile
```

---

## 🚪 Эндпоинты

### Orders

| Метод | URL                            | Описание                  |
|-------|--------------------------------|---------------------------|
| POST  | `/api/v1/orders/`              | Создание нового заказа    |
| GET   | `/api/v1/orders?user_id=...`   | Получение списка заказов  |
| GET   | `/api/v1/orders/:id/status`    | Получение статуса заказа  |

### Payments

| Метод | URL                              | Описание                    |
|-------|----------------------------------|-----------------------------|
| POST  | `/api/v1/payments/account`       | Создание платёжного аккаунта |
| POST  | `/api/v1/payments/deposit`       | Пополнение баланса          |
| POST  | `/api/v1/payments/withdraw`      | Списание с баланса          |
| GET   | `/api/v1/payments/balance`       | Получение текущего баланса  |

---

## 🧠 Технологии

- Go 1.24
- gRPC — связь между gateway и микросервисами
- Gin — HTTP сервер и маршрутизация
- Swagger (swaggo) — автогенерация документации
- Protobuf — IDL для gRPC
- Docker — контейнеризация

---

## 🔧 Конфигурация

Файл `config/local.yaml`:

```yaml
http:
  port: "8080"

services:
  orders:
    address: "localhost:50051"

grpc:
  paymentsAddr: "localhost:50052"
```

Загружается через `config.MustLoad()`.

---

## 🧪 Тестирование

- Swagger UI доступен по адресу:  
  [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)


---

## ⚙️ Разработка и запуск

```bash
# Сгенерировать gRPC код
make proto

# Сгенерировать Swagger
make swagger

# Запустить сервис
go run cmd/main.go
```

---

## 📌 Примечания

- Swagger требует DTO-структур для генерации документации.
- Прямые вызовы микросервисов (orders/payments) осуществляются через gRPC.
- Аутентификация и трейсинг пока не реализованы по условию.
