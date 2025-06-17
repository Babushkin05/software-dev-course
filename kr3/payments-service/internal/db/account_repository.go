package db

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/Babushkin05/software-dev-course/kr3/payments-service/internal/model"
	"github.com/google/uuid"
)

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

var ErrAccountAlreadyExists = errors.New("account already exists")
var ErrAccountNotFound = errors.New("account not found")
var ErrInsufficientFunds = errors.New("insufficient funds")

func (r *AccountRepository) Create(ctx context.Context, userID string) (*model.Account, error) {
	id := uuid.New().String()
	_, err := r.db.ExecContext(ctx, `
        INSERT INTO accounts(id, user_id, balance, created_at)
        VALUES ($1, $2, 0, now())
    `, id, userID)

	if err != nil {
		return nil, err
	}

	return &model.Account{
		ID:        id,
		UserID:    userID,
		Balance:   0,
		CreatedAt: time.Now(),
	}, nil
}

func (r *AccountRepository) AddBalance(ctx context.Context, userID string, amount int64) (int64, error) {
	var newBalance int64
	err := r.db.QueryRowContext(ctx, `
        UPDATE accounts
        SET balance = balance + $1
        WHERE user_id = $2
        RETURNING balance
    `, amount, userID).Scan(&newBalance)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrAccountNotFound
		}
		return 0, err
	}

	return newBalance, nil
}

func (r *AccountRepository) GetBalance(ctx context.Context, userID string) (int64, error) {
	var balance int64
	err := r.db.QueryRowContext(ctx, `
        SELECT balance FROM accounts WHERE user_id = $1
    `, userID).Scan(&balance)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrAccountNotFound
		}
		return 0, err
	}

	return balance, nil
}

func (r *AccountRepository) Withdraw(ctx context.Context, userID string, amount int64) (int64, error) {
	var newBalance int64

	err := r.db.QueryRowContext(ctx, `
        UPDATE accounts
        SET balance = balance - $1
        WHERE user_id = $2 AND balance >= $1
        RETURNING balance
    `, amount, userID).Scan(&newBalance)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// либо не найден пользователь, либо недостаточно средств
			exists, checkErr := r.exists(ctx, userID)
			if checkErr != nil {
				return 0, checkErr
			}
			if !exists {
				return 0, ErrAccountNotFound
			}
			return 0, ErrInsufficientFunds
		}
		return 0, err
	}

	return newBalance, nil
}

func (r *AccountRepository) exists(ctx context.Context, userID string) (bool, error) {
	var exists bool
	err := r.db.QueryRowContext(ctx, `
        SELECT EXISTS(SELECT 1 FROM accounts WHERE user_id = $1)
    `, userID).Scan(&exists)
	return exists, err
}
