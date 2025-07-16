package kafka

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/Babushkin05/software-dev-course/kr3/orders-service/internal/db"
	"github.com/Babushkin05/software-dev-course/kr3/orders-service/internal/model"
)

type InboxProcessor struct {
	repo *db.OrderRepository
}

func NewInboxProcessor(repo *db.OrderRepository) *InboxProcessor {
	return &InboxProcessor{repo: repo}
}

func (p *InboxProcessor) Start(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(2 * time.Second)
		for {
			select {
			case <-ticker.C:
				if err := p.processBatch(ctx); err != nil {
					log.Printf("inbox processor error: %v", err)
				}
			case <-ctx.Done():
				log.Println("inbox processor stopped")
				return
			}
		}
	}()
}

func (p *InboxProcessor) processBatch(ctx context.Context) error {
	msgs, err := p.repo.FetchUnprocessedInboxMessages(ctx, 10)
	if err != nil {
		return err
	}

	for _, msg := range msgs {
		var event struct {
			OrderID string `json:"order_id"`
			Status  string `json:"status"`
		}

		if err := json.Unmarshal([]byte(msg.Payload), &event); err != nil {
			log.Printf("invalid inbox payload: %v", err)
			continue
		}

		var newStatus model.OrderStatus
		switch event.Status {
		case "FINISHED":
			newStatus = model.OrderStatusFinished
		case "CANCELED":
			newStatus = model.OrderStatusCanceled
		default:
			log.Printf("unknown status in inbox message: %s", event.Status)
			continue
		}

		if err := p.repo.UpdateStatus(ctx, event.OrderID, newStatus); err != nil {
			log.Printf("failed to update order status: %v", err)
			continue
		}

		if err := p.repo.MarkInboxMessageProcessed(ctx, msg.ID); err != nil {
			log.Printf("failed to mark inbox message processed: %v", err)
		}
	}

	return nil
}
