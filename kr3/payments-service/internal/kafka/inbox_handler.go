package kafka

import (
	"github.com/Babushkin05/software-dev-course/kr3/payments-service/internal/db"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

func HandleInboxMessage(msg *kafka.Message, repo *db.AccountRepository) error {
	return repo.SaveInboxMessage(db.InboxMessage{
		MessageID: uuid.New().String(),
		Topic:     msg.Topic,
		Payload:   string(msg.Value),
	})
}
