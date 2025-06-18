package db

import "context"

type InboxMessage struct {
	MessageID string
	Topic     string
	Payload   string
}

func (r *OrderRepository) SaveInboxMessage(ctx context.Context, msg InboxMessage) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO inbox (message_id, topic, payload)
		VALUES ($1, $2, $3)
		ON CONFLICT (message_id) DO NOTHING
	`, msg.MessageID, msg.Topic, msg.Payload)
	return err
}

func (r *OrderRepository) FetchUnprocessedInboxMessages(ctx context.Context, limit int) ([]InboxMessage, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT message_id, topic, payload
		FROM inbox
		WHERE processed = false
		ORDER BY created_at ASC
		LIMIT $1
	`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var msgs []InboxMessage
	for rows.Next() {
		var msg InboxMessage
		if err := rows.Scan(&msg.MessageID, &msg.Topic, &msg.Payload); err != nil {
			return nil, err
		}
		msgs = append(msgs, msg)
	}
	return msgs, nil
}

func (r *OrderRepository) MarkInboxMessageProcessed(ctx context.Context, messageID string) error {
	_, err := r.db.ExecContext(ctx, `UPDATE inbox SET processed = true WHERE message_id = $1`, messageID)
	return err
}
