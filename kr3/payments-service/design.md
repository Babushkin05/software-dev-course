# Design — Payments Service

## 📦 Назначение сервиса

`payments-service` отвечает за управление пользовательскими счетами, включая:
- создание счёта,
- пополнение баланса,
- просмотр текущего баланса,
- списание средств по заказу (асинхронно через Kafka).

## 🏗️ Архитектура

### Общая структура проекта

```
/payments-service
├── Makefile
├── api/                  # Протобаф-схемы
│   ├── gen/              # Сгенерированный код
│   └── payments.proto
├── cmd/
│   └── main.go           # Точка входа
├── config/
│   └── local.yaml        # Конфигурация
├── deployments/
│   └── Dockerfile        # Сборка образа
├── internal/
│   ├── config/           # Загрузка конфигурации
│   ├── db/               # Работа с БД и schema.sql
│   ├── grpc/             # gRPC-сервер
│   ├── kafka/            # Kafka consumer
│   ├── model/            # Бизнес-модели
│   └── service/          # Бизнес-логика
```

---

## ⚙️ Бизнес-функциональность

### 📌 Открытые методы для пользователя (через API Gateway → gRPC)

- `CreateAccount(user_id)` — создать счёт (не более одного на пользователя)
- `AddBalance(user_id, amount)` — пополнить баланс
- `GetBalance(user_id)` — получить текущий баланс

### 📌 Обработка Kafka-сообщений (`orders-topic`)

Асинхронно обрабатываются заказы:
- При получении сообщения из Kafka производится попытка списания средств с учётом transactional inbox.
- Используется семантика **at-most-once**, чтобы избежать двойного списания.

---

## 🛢️ Работа с базой данных

### 💾 Таблицы
- `accounts(id, user_id, balance, created_at)`
- `inbox(message_id, topic, payload, processed, created_at)`

### 📄 `schema.sql`
- При запуске автоматически удаляет старые данные и создаёт таблицы с нуля.
- Загружается при старте из `internal/db/schema.sql`.

---

## 🧠 Бизнес-логика

### Интерфейс `AccountStorage`

```go
type AccountStorage interface {
	Create(ctx context.Context, userID string) (*model.Account, error)
	AddBalance(ctx context.Context, userID string, amount int64) (int64, error)
	GetBalance(ctx context.Context, userID string) (int64, error)
	Withdraw(ctx context.Context, userID string, amount int64) (int64, error)

	// Inbox паттерн
	SaveInboxMessage(msg InboxMessage) error
	FetchUnprocessedInboxMessages(limit int) ([]InboxMessage, error)
	MarkInboxMessageProcessed(messageID string) error
}
```

### Реализация
- Все методы работают через `sql.DB` (PostgreSQL).
- Списания защищены от гонок через `WHERE balance >= amount`.

---

## 📬 Kafka: Consumer + Inbox Pattern

- Реализован Kafka consumer (`StartKafkaConsumer`) через `segmentio/kafka-go`.
- Каждое входящее сообщение записывается в `inbox`, обрабатывается, и помечается как `processed`.
- Гарантия идемпотентности обеспечена `ON CONFLICT (message_id) DO NOTHING`.

---

## 🧪 Тесты

- Юнит-тесты для бизнес-логики (через мок-репозиторий).
- Используется `testify/mock` для создания `MockStorage`.
- Тесты покрывают: создание, пополнение, просмотр баланса, списание средств и ошибки (например, недостаток средств).

---

## ⚙️ Devtools

### Makefile
Поддерживает:
- запуск приложения: `make run`
- запуск через docker: `make docker-run`
- генерацию protobuf: `make proto`

---

## 🔧 Конфигурация

```yaml
grpc:
  port: ":50051"

database:
  dsn: "postgres://user:password@localhost:5432/payments?sslmode=disable"

kafka:
  broker: "localhost:9092"
  topic: "orders-topic"
  group_id: "payments-group"
```

---

## 🧩 Протобаф

- Сервисный контракт находится в `api/payments.proto`
- Используется gRPC для внутренних вызовов.

---

