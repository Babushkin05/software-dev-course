package service

import (
	"context"
	"time"

	"github.com/Babushkin05/software-dev-course/kr3/orders-service/internal/db"
	"github.com/Babushkin05/software-dev-course/kr3/orders-service/internal/kafka"
	"github.com/Babushkin05/software-dev-course/kr3/orders-service/internal/model"
	"github.com/google/uuid"
)

type OrderService struct {
	storage     db.OrderStorage
	kafkaWriter kafka.OrderWriter
}

func NewOrderService(storage db.OrderStorage, kafkaWriter kafka.OrderWriter) *OrderService {
	return &OrderService{
		storage:     storage,
		kafkaWriter: kafkaWriter,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, userID string, amount int64, description string) (*model.Order, error) {
	order := &model.Order{
		ID:          uuid.New().String(),
		UserID:      userID,
		Amount:      amount,
		Description: description,
		Status:      model.OrderStatusNew,
		CreatedAt:   time.Now(),
	}

	if err := s.storage.Create(ctx, order); err != nil {
		return nil, err
	}

	// Отправка в Kafka
	if err := s.kafkaWriter.WriteOrder(ctx, order); err != nil {
		return nil, err
	}

	return order, nil
}

func (s *OrderService) GetOrders(ctx context.Context, userID string) ([]*model.Order, error) {
	return s.storage.GetByUser(ctx, userID)
}

func (s *OrderService) GetOrderStatus(ctx context.Context, orderID string) (model.OrderStatus, error) {
	order, err := s.storage.GetByID(ctx, orderID)
	if err != nil {
		return "", err
	}
	return order.Status, nil
}

func (s *OrderService) MarkOrderFinished(ctx context.Context, orderID string) error {
	return s.storage.UpdateStatus(ctx, orderID, model.OrderStatusFinished)
}

func (s *OrderService) GetByID(ctx context.Context, orderID string) (*model.Order, error) {
	return s.storage.GetByID(ctx, orderID)
}

func (s *OrderService) MarkFinished(ctx context.Context, orderID string) error {
	return s.storage.UpdateStatus(ctx, orderID, model.OrderStatusFinished)
}
