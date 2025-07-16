package model

import "time"

type OrderStatus string

const (
	OrderStatusNew      OrderStatus = "NEW"
	OrderStatusFinished OrderStatus = "FINISHED"
	OrderStatusCanceled OrderStatus = "CANCELED"
)

type Order struct {
	ID          string
	UserID      string
	Amount      int64
	Description string
	Status      OrderStatus
	CreatedAt   time.Time
}
