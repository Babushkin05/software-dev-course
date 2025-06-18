package dto

type CreateAccountRequest struct {
	UserID string `json:"user_id" binding:"required"`
}

type PaymentRequest struct {
	UserID string `json:"user_id" binding:"required"`
	Amount int64  `json:"amount" binding:"required"`
}

type BalanceRequest struct {
	UserID string `json:"user_id" binding:"required"`
}
