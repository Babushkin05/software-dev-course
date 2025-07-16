package kafka

import (
	"context"
	"encoding/json"

	"github.com/Babushkin05/software-dev-course/kr3/orders-service/internal/db"
	"github.com/Babushkin05/software-dev-course/kr3/orders-service/internal/model"
)

type OrderWriter interface {
	WriteOrder(ctx context.Context, order *model.Order) error
}

type Writer struct {
	storage db.OrderStorage
	topic   string
}

func NewWriter(storage db.OrderStorage, topic string) *Writer {
	return &Writer{
		storage: storage,
		topic:   topic,
	}
}

func (w *Writer) WriteOrder(ctx context.Context, order *model.Order) error {
	msg, err := json.Marshal(order)
	if err != nil {
		return err
	}

	return w.storage.SaveOutboxMessage(ctx, w.topic, &order.ID, string(msg))
}
