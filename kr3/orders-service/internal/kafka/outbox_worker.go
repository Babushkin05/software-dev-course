package kafka

import (
	"context"
	"log"
	"time"

	"github.com/Babushkin05/software-dev-course/kr3/orders-service/internal/db"
	segmentioKafka "github.com/segmentio/kafka-go"
)

type OutboxWorker struct {
	Repo   *db.OrderRepository
	Writer *segmentioKafka.Writer
}

func NewOutboxWorker(repo *db.OrderRepository, writer *segmentioKafka.Writer) *OutboxWorker {
	return &OutboxWorker{
		Repo:   repo,
		Writer: writer,
	}
}

func (w *OutboxWorker) Start(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(2 * time.Second)
		for {
			select {
			case <-ticker.C:
				if err := w.processBatch(ctx); err != nil {
					log.Printf("outbox worker error: %v", err)
				}
			case <-ctx.Done():
				log.Println("outbox worker stopped")
				return
			}
		}
	}()
}

func (w *OutboxWorker) processBatch(ctx context.Context) error {
	msgs, err := w.Repo.FetchUnsentOutboxMessages(ctx, 10)
	if err != nil {
		return err
	}

	for _, m := range msgs {
		kafkaMsg := segmentioKafka.Message{
			Topic: m.Topic,
			Value: []byte(m.Payload),
		}
		if m.Key != nil {
			kafkaMsg.Key = []byte(*m.Key)
		}

		if err := w.Writer.WriteMessages(ctx, kafkaMsg); err != nil {
			log.Printf("failed to write Kafka message: %v", err)
			continue
		}

		if err := w.Repo.MarkOutboxMessageSent(ctx, m.ID); err != nil {
			log.Printf("failed to mark message as sent: %v", err)
			continue
		}
	}

	return nil
}
