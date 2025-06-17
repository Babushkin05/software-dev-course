package db

import (
	"context"
	"database/sql"
	"time"
)

type OutboxStorage interface {
	FetchUnsentOutboxMessages(limit int) ([]OutboxMessage, error)
	MarkOutboxMessageSent(id string) error
}

// OutboxMessage представляет сообщение, которое должно быть отправлено в Kafka.
type OutboxMessage struct {
	ID        string
	Topic     string
	Payload   string
	CreatedAt time.Time
	Sent      bool
}

// SaveOutboxMessage сохраняет новое сообщение в таблицу outbox.
func (r *AccountRepository) SaveOutboxMessage(ctx context.Context, msg OutboxMessage) error {
	_, err := r.db.Exec(`
		INSERT INTO outbox (id, topic, payload, sent, created_at)
		VALUES ($1, $2, $3, FALSE, now())
	`, msg.ID, msg.Topic, msg.Payload)
	return err
}

func (r *AccountRepository) SaveOutboxMessageTx(ctx context.Context, tx *sql.Tx, msg OutboxMessage) error {
	_, err := tx.ExecContext(ctx, `
        INSERT INTO outbox(message_id, topic, payload)
        VALUES (gen_random_uuid(), $1, $2)
    `, msg.Topic, msg.Payload)
	return err
}

// FetchUnsentOutboxMessages выбирает неотправленные сообщения из outbox (до лимита).
func (r *AccountRepository) FetchUnsentOutboxMessages(limit int) ([]OutboxMessage, error) {
	rows, err := r.db.Query(`
		SELECT id, topic, payload, created_at, sent
		FROM outbox
		WHERE sent = FALSE
		ORDER BY created_at ASC
		LIMIT $1
	`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var msgs []OutboxMessage
	for rows.Next() {
		var msg OutboxMessage
		if err := rows.Scan(&msg.ID, &msg.Topic, &msg.Payload, &msg.CreatedAt, &msg.Sent); err != nil {
			return nil, err
		}
		msgs = append(msgs, msg)
	}
	return msgs, nil
}

// MarkOutboxMessageSent отмечает сообщение как успешно отправленное.
func (r *AccountRepository) MarkOutboxMessageSent(id string) error {
	_, err := r.db.Exec(`UPDATE outbox SET sent = TRUE WHERE id = $1`, id)
	return err
}
