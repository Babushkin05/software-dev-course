package kafka

import (
	"context"
	"encoding/json"

	"github.com/Babushkin05/software-dev-course/kr3/orders-service/internal/model"
	"github.com/segmentio/kafka-go"
)

type Writer struct {
	writer *kafka.Writer
}

func NewWriter(broker, topic string) *Writer {
	return &Writer{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(broker),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		},
	}
}

func (w *Writer) WriteOrder(ctx context.Context, order *model.Order) error {
	msg, err := json.Marshal(order)
	if err != nil {
		return err
	}

	return w.writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(order.ID),
		Value: msg,
	})
}
