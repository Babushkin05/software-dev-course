package kafka

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/Babushkin05/software-dev-course/kr3/payments-service/internal/service"
)

func StartInboxProcessor(ctx context.Context, svc *service.AccountService) {
	go func() {
		ticker := time.NewTicker(2 * time.Second)
		for {
			select {
			case <-ticker.C:
				processBatch(svc)
			case <-ctx.Done():
				return
			}
		}
	}()
}

func processBatch(svc *service.AccountService) {
	msgs, err := svc.Repo.FetchUnprocessedInboxMessages(10)
	if err != nil {
		log.Printf("error fetching inbox: %v", err)
		return
	}

	for _, msg := range msgs {
		var order struct {
			ID     string `json:"id"`
			UserID string `json:"user_id"`
			Amount int64  `json:"amount"`
		}
		if err := json.Unmarshal([]byte(msg.Payload), &order); err != nil {
			log.Printf("invalid payload: %v", err)
			continue
		}

		ctx := context.Background()

		if _, err := svc.Withdraw(ctx, order.UserID, order.Amount); err != nil {
			log.Printf("withdraw failed: %v", err)
			continue
		}

		_ = svc.Repo.MarkInboxMessageProcessed(msg.MessageID)
	}
}
