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
	Repo db.AccountStorage
}

func NewAccountService(repo db.AccountStorage) *AccountService {
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

// Списание — используется при оплате заказа
func (s *AccountService) Withdraw(ctx context.Context, userID string, amount int64) (int64, error) {
	if amount <= 0 {
		return 0, fmt.Errorf("invalid amount")
	}
	return s.Repo.Withdraw(ctx, userID, amount)
}
