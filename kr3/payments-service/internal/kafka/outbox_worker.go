package kafka

import (
	"context"
	"log"
	"time"

	"github.com/Babushkin05/software-dev-course/kr3/payments-service/internal/db"
	"github.com/segmentio/kafka-go"
)

type OutboxWorker struct {
	Repo      db.OutboxStorage // интерфейс, реализуемый AccountRepository
	Writer    *kafka.Writer
	PollDelay time.Duration
	BatchSize int
}

// NewOutboxWorker создает воркера
func NewOutboxWorker(repo db.OutboxStorage, writer *kafka.Writer, pollDelay time.Duration, batchSize int) *OutboxWorker {
	return &OutboxWorker{
		Repo:      repo,
		Writer:    writer,
		PollDelay: pollDelay,
		BatchSize: batchSize,
	}
}

// Start запускает бесконечный цикл отправки сообщений из Outbox
func (w *OutboxWorker) Start(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("Outbox worker stopped")
				return
			default:
				w.processBatch(ctx)
				time.Sleep(w.PollDelay)
			}
		}
	}()
}

// processBatch обрабатывает партию сообщений
func (w *OutboxWorker) processBatch(ctx context.Context) {
	messages, err := w.Repo.FetchUnsentOutboxMessages(w.BatchSize)
	if err != nil {
		log.Printf("failed to fetch outbox messages: %v", err)
		return
	}

	for _, msg := range messages {
		kmsg := kafka.Message{
			Topic: msg.Topic,
			Key:   []byte(msg.ID),
			Value: []byte(msg.Payload),
		}
		if err := w.Writer.WriteMessages(ctx, kmsg); err != nil {
			log.Printf("failed to write message to kafka (id: %s): %v", msg.ID, err)
			continue
		}

		if err := w.Repo.MarkOutboxMessageSent(msg.ID); err != nil {
			log.Printf("failed to mark outbox message as sent (id: %s): %v", msg.ID, err)
		}
	}
}
