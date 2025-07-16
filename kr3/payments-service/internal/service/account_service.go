package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/Babushkin05/software-dev-course/kr3/payments-service/internal/db"
	"github.com/Babushkin05/software-dev-course/kr3/payments-service/internal/model"
)

var (
	ErrAccountNotFound   = errors.New("account not found")
	ErrInsufficientFunds = errors.New("insufficient funds")
)

type AccountService struct {
	Repo AccountStorage
}

func NewAccountService(repo AccountStorage) *AccountService {
	return &AccountService{Repo: repo}
}

// Создание счёта
func (s *AccountService) CreateAccount(ctx context.Context, userID string) (*model.Account, error) {
	return s.Repo.Create(ctx, userID)
}

// Пополнение
func (s *AccountService) Deposit(ctx context.Context, userID string, amount int64) (int64, error) {
	if amount <= 0 {
		return 0, fmt.Errorf("invalid amount")
	}
	return s.Repo.AddBalance(ctx, userID, amount)
}

// Получение баланса
func (s *AccountService) GetBalance(ctx context.Context, userID string) (int64, error) {
	return s.Repo.GetBalance(ctx, userID)
}

func (s *AccountService) Withdraw(ctx context.Context, userID, orderID string, amount int64) (int64, error) {
	if amount <= 0 {
		_ = s.Repo.SaveOutboxMessage(ctx, db.OutboxMessage{
			Topic:   "order-events",
			Payload: buildOrderCanceledPayload(orderID),
		})
		return 0, fmt.Errorf("invalid amount")
	}

	tx, err := s.Repo.BeginTx(ctx)
	if err != nil {
		_ = s.Repo.SaveOutboxMessage(ctx, db.OutboxMessage{
			Topic:   "order-events",
			Payload: buildOrderCanceledPayload(orderID),
		})
		return 0, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer s.Repo.RollbackTx(tx)

	newBalance, err := s.Repo.WithdrawTx(ctx, tx, userID, amount)
	if err != nil {
		_ = s.Repo.SaveOutboxMessage(ctx, db.OutboxMessage{
			Topic:   "order-events",
			Payload: buildOrderCanceledPayload(orderID),
		})
		return 0, err
	}

	err = s.Repo.SaveOutboxMessageTx(ctx, tx, db.OutboxMessage{
		Topic:   "order-events",
		Payload: buildOrderFinishedPayload(orderID),
	})
	if err != nil {
		_ = s.Repo.SaveOutboxMessage(ctx, db.OutboxMessage{
			Topic:   "order-events",
			Payload: buildOrderCanceledPayload(orderID),
		})
		return 0, err
	}

	if err := s.Repo.CommitTx(tx); err != nil {
		_ = s.Repo.SaveOutboxMessage(ctx, db.OutboxMessage{
			Topic:   "order-events",
			Payload: buildOrderCanceledPayload(orderID),
		})
		return 0, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return newBalance, nil
}

func buildOrderFinishedPayload(orderID string) string {
	return fmt.Sprintf(`{"order_id":"%s","status":"finished"}`, orderID)
}

func buildOrderCanceledPayload(orderID string) string {
	return fmt.Sprintf(`{"order_id":"%s","status":"canceled"}`, orderID)
}
