package dto

type CreateOrderRequest struct {
	UserID      string `json:"user_id" binding:"required"`
	Amount      int64  `json:"amount" binding:"required"`
	Description string `json:"description"`
}
