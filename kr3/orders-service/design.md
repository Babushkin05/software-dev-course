
# Orders Service — Design Document

## Описание

`orders-service` — микросервис, отвечающий за приём, хранение и отслеживание заказов пользователей.  
Он реализует взаимодействие с другими сервисами через Kafka и предоставляет gRPC API для вызова из внешних систем.

---

## Основные компоненты

### 1. OrderService (бизнес-логика)

Методы:
- `CreateOrder(userID, amount, description)`  
  Создаёт заказ и регистрирует событие в outbox.
  
- `GetOrders(userID)`  
  Возвращает список заказов пользователя.

- `GetOrderStatus(orderID)`  
  Показывает текущий статус заказа.

- `MarkOrderFinished(orderID)`  
  Устанавливает статус заказа как "завершён".

---

### 2. Хранилище (PostgreSQL)

#### Таблицы:

- `orders`
  - `id`, `user_id`, `amount`, `description`, `status`, `created_at`

- `outbox`
  - `id`, `topic`, `key`, `payload`, `created_at`, `sent`

- `inbox`
  - `message_id`, `topic`, `payload`, `processed`, `created_at`

---

### 3. Kafka-интеграция

#### Outbox паттерн

- `OrderWriter.WriteOrder(...)`  
  Сериализует заказ и сохраняет событие в таблицу `outbox`.

- `OutboxWorker`  
  Периодически читает unsent-сообщения из `outbox`, отправляет их в Kafka и помечает как отправленные.

#### Inbox паттерн

- `InboxConsumer`  
  Kafka consumer, который сохраняет входящее сообщение в таблицу `inbox` (если не было обработано ранее).

- `InboxProcessor`  
  Периодически обрабатывает сообщения из `inbox`, обновляя статусы заказов (например, `finished`/`cancelled`).

---

### 4. gRPC API

Интерфейс:
- `CreateOrder`
- `GetOrders`
- `GetOrderStatus`
- `MarkOrderFinished`

Сгенерировано из `.proto` файлов, используется пакет `orderspb`.

---

## Потоки данных

### ➕ Новый заказ
1. Внешняя система вызывает `CreateOrder`
2. Сервис сохраняет заказ в таблицу `orders`
3. Пишет событие в `outbox`
4. `OutboxWorker` отправляет сообщение в Kafka

### ✅ Завершение заказа
1. `payments-service` публикует сообщение в Kafka
2. `InboxConsumer` сохраняет его в `inbox`
3. `InboxProcessor` читает, обновляет статус заказа

---

## Точки расширения

- Добавить валидацию и аутентификацию
- Расширить типы событий Kafka
- Хранить в Kafka отдельные типы (OrderCreated, OrderFinished, и т.д.)
- Обработка dead-letter или retries для Kafka

---

## Зависимости

- PostgreSQL
- Kafka
- gRPC
- Go modules

---

## Переменные окружения (config)

- `POSTGRES_DSN`
- `KAFKA_BROKER`
- `KAFKA_TOPIC`
- `KAFKA_GROUP_ID`
- `GRPC_PORT`


