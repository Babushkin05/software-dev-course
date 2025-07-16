# 🧩 `payments-service` Design Document

## 📌 Назначение

Сервис отвечает за управление пользовательскими аккаунтами и обработку платежей. Основная задача — **списание средств с аккаунта при создании заказа**. В случае успеха или неудачи — отправка финального статуса заказа в `orders-service` через Kafka.

---

## 📁 Структура проекта

```
payments-service/
│
├── api/                  # gRPC protobuf definitions (compiled)
├── internal/
│   ├── config/           # Загрузка конфигурации из env
│   ├── db/               # Работа с Postgres (AccountRepository, inbox, outbox)
│   ├── grpc/             # gRPC сервер
│   ├── kafka/            # Kafka consumer/producer
│   └── service/          # Бизнес-логика (AccountService)
│
├── cmd/
│   └── main.go           # Точка входа
```

---

## 🔧 Основные компоненты

### ✅ gRPC API

Методы:
- `CreateAccount`
- `AddBalance`
- `GetBalance`

### 📦 Kafka

- 📥 `order_created` (inbox)
- 📤 `order_finished`, `order_canceled` (outbox)

### 🧠 Бизнес-логика (`AccountService`)
- `Withdraw(ctx, userID, orderID, amount)` — списание средств
  - В случае успеха → `order_finished`
  - В случае ошибки (например, недостаточно средств) → `order_canceled`
- Сохраняет сообщения в `outbox`, не отправляет напрямую в Kafka

---

## 🗄️ Хранилище

PostgreSQL, 2 ключевые таблицы:

### `accounts`
| id | user_id | balance | created_at |

### `inbox`
| message_id | topic | payload | processed | created_at |

### `outbox`
| id | topic | key | payload | dispatched | created_at |

---

## 🔁 Паттерны надёжности

### ✅ Inbox Pattern
- Гарантирует, что Kafka-сообщение будет обработано точно один раз

### ✅ Outbox Pattern
- Гарантирует надёжную доставку события об изменении статуса заказа

---

## ⚙️ Воркеры

### `InboxProcessor`
- Каждые 2 секунды:
  - Читает `inbox` (необработанные)
  - Парсит заказ
  - Вызывает `Withdraw`
  - Помечает сообщение как `processed`

### `OutboxDispatcher`
- Каждые 2 секунды:
  - Читает `outbox` (неотправленные)
  - Публикует в Kafka
  - Помечает как `dispatched = true`

---

## 🧪 Тестирование

- Покрыта бизнес-логика
- Покрытие с mock-интерфейсами (`AccountStorage`, `KafkaWriter`)
- Проверка:
  - Успешного списания
  - Ошибки недостатка средств
  - Записи в `outbox`

`coverage: 17.9% of statements`
