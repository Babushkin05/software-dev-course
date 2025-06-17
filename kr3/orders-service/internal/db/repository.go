package db

import (
	"context"

	"github.com/Babushkin05/software-dev-course/kr3/orders-service/internal/model"
)

type OrderStorage interface {
	Create(ctx context.Context, order *model.Order) error
	GetByUser(ctx context.Context, userID string) ([]*model.Order, error)
	GetByID(ctx context.Context, orderID string) (*model.Order, error)
	UpdateStatus(ctx context.Context, orderID string, status model.OrderStatus) error
}
