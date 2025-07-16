package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type OutboxMessage struct {
	ID        uuid.UUID
	Topic     string
	Key       *string
	Payload   string
	CreatedAt time.Time
	Sent      bool
}

func (r OrderRepository) SaveOutboxMessage(ctx context.Context, topic string, key *string, payload string) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO outbox (id, topic, key, payload, created_at, sent)
		VALUES ($1, $2, $3, $4, $5, false)
	`, uuid.New(), topic, key, payload, time.Now())
	return err
}

func (r *OrderRepository) FetchUnsentOutboxMessages(ctx context.Context, limit int) ([]OutboxMessage, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, topic, key, payload, created_at
		FROM outbox
		WHERE sent = false
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
		if err := rows.Scan(&msg.ID, &msg.Topic, &msg.Key, &msg.Payload, &msg.CreatedAt); err != nil {
			return nil, err
		}
		msgs = append(msgs, msg)
	}
	return msgs, nil
}

func (r *OrderRepository) MarkOutboxMessageSent(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.ExecContext(ctx, `UPDATE outbox SET sent = true WHERE id = $1`, id)
	return err
}
