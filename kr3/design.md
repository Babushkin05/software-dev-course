# 💳 Архитектура сервиса оплаты заказов

## 🧭 Общий обзор

Система состоит из двух микросервисов и собственного API Gateway. Коммуникация между сервисами — через Kafka (асинхронно) и gRPC (синхронно). Внешние клиенты взаимодействуют с системой через REST-интерфейс API Gateway.

---

## 📦 Компоненты

| Компонент         | Описание |
|------------------|----------|
| **API Gateway**  | REST-интерфейс, конвертирует HTTP-запросы в gRPC и маршрутизирует их на сервисы. Также отдает Swagger-документацию. |
| **Orders Service** | Создание заказов и получение их статусов |
| **Payments Service** | Управление счетами пользователей, списание и пополнение средств |
| **Kafka**         | Асинхронный обмен событиями |
| **PostgreSQL**    | База данных для каждого сервиса (своя) |
| **Schema Registry** | Хранилище схем Kafka-сообщений (Protobuf/Avro) |
| **Swagger UI**    | Отдельный компонент для отображения документации REST API |

---

## 🧰 Технологии

- Язык: **Go**
- REST-фреймворк: `gin`, `echo` или `fiber`
- gRPC: `google.golang.org/grpc`
- Kafka: `segmentio/kafka-go` или `confluent-kafka-go`
- БД: PostgreSQL
- Swagger: `swaggo/swag`
- Схемы сообщений: Protobuf или Avro

---

## ⚙️ Доступный функционал

### 📄 Payments Service (через API Gateway)
1. **Создание счета**
   - Один счёт на пользователя. Повторный запрос — ошибка.
2. **Пополнение счета**
   - Зачисление средств с логированием транзакции.
3. **Просмотр баланса**
   - Текущий баланс счёта пользователя.

### 📄 Orders Service (через API Gateway)
1. **Создание заказа**
   - Асинхронный процесс: создается заказ, публикуется событие оплаты.
2. **Просмотр списка заказов**
   - Получение всех заказов конкретного пользователя.
3. **Просмотр статуса заказа**
   - Получение полной информации по конкретному заказу (по ID).

---

## 📐 Поток создания заказа

1. Клиент отправляет `POST /orders` в API Gateway
2. Gateway вызывает gRPC-метод `OrdersService.CreateOrder`
3. Сервис сохраняет заказ со статусом `NEW` и записывает событие `OrderCreated` в Outbox
4. Kafka Publisher читает Outbox и публикует в Kafka
5. Payments Service через Inbox читает `OrderCreated`, проверяет баланс, списывает деньги
6. Если успешно — публикует `PaymentSucceeded`
7. Orders Service получает это сообщение и обновляет заказ в БД на `FINISHED`

---

## 🧾 Структура БД

### Orders Service

```sql
orders (
  id UUID PRIMARY KEY,
  user_id UUID,
  amount BIGINT,
  description TEXT,
  status TEXT CHECK (status IN ('NEW', 'FINISHED', 'CANCELED')),
  created_at TIMESTAMP
)
```

### Payments Service

```sql
accounts (
  id UUID PRIMARY KEY,
  user_id UUID UNIQUE,
  balance BIGINT,
  updated_at TIMESTAMP
)

transactions (
  id UUID PRIMARY KEY,
  account_id UUID,
  amount BIGINT,
  type TEXT CHECK (type IN ('DEBIT', 'CREDIT')),
  reference_order_id UUID,
  created_at TIMESTAMP
)
```

---

## 📨 Kafka события

### `OrderCreated`
```json
{
  "event_id": "uuid",
  "event_type": "OrderCreated",
  "occurred_at": "timestamp",
  "data": {
    "order_id": "uuid",
    "user_id": "uuid",
    "amount": 1000,
    "description": "тестовая покупка"
  }
}
```

### `PaymentSucceeded`
```json
{
  "event_id": "uuid",
  "event_type": "PaymentSucceeded",
  "occurred_at": "timestamp",
  "data": {
    "order_id": "uuid",
    "user_id": "uuid"
  }
}
```

---

## 📁 Структура каталогов

```
/api-gateway
  /docs             # Swagger
  /handlers         # REST endpoints
  /grpcclients      # gRPC stubs
  /main.go

/services/
  /orders
    /cmd
    /internal
      /grpc
      /kafka
      /db
      /outbox
  /payments
    /cmd
    /internal
      /grpc
      /kafka
      /db
      /inbox

/api/
  /proto/           # protobuf
  /schemas/         # Kafka schemas

```

---

## 🧰 Inbox / Outbox

- **Outbox**: события сохраняются в БД как часть основной транзакции, затем публикуются в Kafka отдельным воркером.
- **Inbox**: входящие события сохраняются перед обработкой, повторная обработка блокируется по `event_id`.

---

## 📚 Swagger

- Используется `swaggo/swag`
- OpenAPI спецификация доступна по `GET /swagger/doc.json`
- Swagger UI разворачивается внутри API Gateway (`/docs`)

---

## ✅ Особенности

- gRPC используется для внутренней коммуникации
- REST используется только на внешнем уровне
- Вся финансовая логика (списание, CAS, баланс) — строго идемпотентна
- Kafka события проходят схемавалидацию
- Метрик и мониторинга пока нет
