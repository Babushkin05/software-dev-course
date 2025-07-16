package kafka

import (
	"context"
	"log"

	"github.com/Babushkin05/software-dev-course/kr3/orders-service/internal/db"
	"github.com/segmentio/kafka-go"
)

type InboxConsumer struct {
	reader *kafka.Reader
	repo   *db.OrderRepository
}

func NewInboxConsumer(broker, topic, groupID string, repo *db.OrderRepository) *InboxConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic:   topic,
		GroupID: groupID,
	})

	return &InboxConsumer{
		reader: reader,
		repo:   repo,
	}
}

func (c *InboxConsumer) Start(ctx context.Context) {
	go func() {
		for {
			m, err := c.reader.ReadMessage(ctx)
			if err != nil {
				log.Printf("inbox consumer read error: %v", err)
				continue
			}

			err = c.repo.SaveInboxMessage(ctx, string(m.Key), m.Topic, string(m.Value))
			if err != nil {
				log.Printf("failed to save inbox message: %v", err)
			}
		}
	}()
}
