package kafka

import (
	"context"
	"log"

	"github.com/Babushkin05/software-dev-course/kr3/payments-service/internal/config"
	"github.com/segmentio/kafka-go"
)

func StartKafkaConsumer(ctx context.Context, kafkaCfg config.KafkaConfig, handler func(msg *kafka.Message)) error {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaCfg.Broker},
		Topic:   kafkaCfg.Topic,
		GroupID: kafkaCfg.GroupID,
	})

	go func() {
		for {
			m, err := r.ReadMessage(ctx)
			if err != nil {
				log.Printf("error reading message: %v", err)
				continue
			}
			handler(&m)
		}
	}()
	return nil
}
