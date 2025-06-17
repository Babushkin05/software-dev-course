package db

import (
	"context"
	"database/sql"

	"github.com/Babushkin05/software-dev-course/kr3/orders-service/internal/model"
)

type OrderStorage interface {
	Create(ctx context.Context, order *model.Order) error
	GetByUser(ctx context.Context, userID string) ([]*model.Order, error)
	GetByID(ctx context.Context, orderID string) (*model.Order, error)
	UpdateStatus(ctx context.Context, orderID string, status model.OrderStatus) error
}

var _ OrderStorage = OrderRepository{}

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r OrderRepository) Create(ctx context.Context, order *model.Order) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO orders (id, user_id, amount, description, status, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`, order.ID, order.UserID, order.Amount, order.Description, order.Status, order.CreatedAt)
	return err
}

func (r OrderRepository) GetByUser(ctx context.Context, userID string) ([]*model.Order, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, user_id, amount, description, status, created_at
		FROM orders
		WHERE user_id = $1
		ORDER BY created_at DESC
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*model.Order
	for rows.Next() {
		var o model.Order
		if err := rows.Scan(&o.ID, &o.UserID, &o.Amount, &o.Description, &o.Status, &o.CreatedAt); err != nil {
			return nil, err
		}
		orders = append(orders, &o)
	}
	return orders, nil
}

func (r OrderRepository) GetByID(ctx context.Context, orderID string) (*model.Order, error) {
	row := r.db.QueryRowContext(ctx, `
		SELECT id, user_id, amount, description, status, created_at
		FROM orders
		WHERE id = $1
	`, orderID)

	var o model.Order
	err := row.Scan(&o.ID, &o.UserID, &o.Amount, &o.Description, &o.Status, &o.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &o, nil
}

func (r OrderRepository) UpdateStatus(ctx context.Context, orderID string, status model.OrderStatus) error {
	_, err := r.db.ExecContext(ctx, `
		UPDATE orders
		SET status = $1
		WHERE id = $2
	`, status, orderID)
	return err
}
