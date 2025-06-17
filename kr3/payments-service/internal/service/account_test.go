package service

import (
	"context"
	"errors"
	"testing"

	"github.com/Babushkin05/software-dev-course/kr3/payments-service/internal/db"
	"github.com/Babushkin05/software-dev-course/kr3/payments-service/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Мок под интерфейс AccountStorage
type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) Create(ctx context.Context, userID string) (*model.Account, error) {
	args := m.Called(ctx, userID)
	account, _ := args.Get(0).(*model.Account)
	return account, args.Error(1)
}

func (m *MockStorage) AddBalance(ctx context.Context, userID string, amount int64) (int64, error) {
	args := m.Called(ctx, userID, amount)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockStorage) GetBalance(ctx context.Context, userID string) (int64, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockStorage) Withdraw(ctx context.Context, userID string, amount int64) (int64, error) {
	args := m.Called(ctx, userID, amount)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockStorage) SaveInboxMessage(msg db.InboxMessage) error {
	args := m.Called(msg)
	return args.Error(0)
}

func (m *MockStorage) FetchUnprocessedInboxMessages(limit int) ([]db.InboxMessage, error) {
	args := m.Called(limit)
	messages, _ := args.Get(0).([]db.InboxMessage)
	return messages, args.Error(1)
}

func (m *MockStorage) MarkInboxMessageProcessed(messageID string) error {
	args := m.Called(messageID)
	return args.Error(0)
}

func TestCreateAccount(t *testing.T) {
	mockStorage := new(MockStorage)
	svc := NewAccountService(mockStorage)

	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		account := &model.Account{ID: "acc123", UserID: "user1", Balance: 0}
		mockStorage.On("Create", ctx, "user1").Return(account, nil)

		_, err := svc.CreateAccount(ctx, "user1")
		assert.NoError(t, err)
		mockStorage.AssertCalled(t, "Create", ctx, "user1")
	})

	t.Run("failure", func(t *testing.T) {
		mockStorage.On("Create", ctx, "user2").Return(&model.Account{}, errors.New("already exists"))

		_, err := svc.CreateAccount(ctx, "user2")
		assert.Error(t, err)
		assert.EqualError(t, err, "already exists")
	})
}

func TestDeposit(t *testing.T) {
	mockStorage := new(MockStorage)
	svc := NewAccountService(mockStorage)

	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		mockStorage.On("AddBalance", ctx, "user1", int64(100)).Return(int64(150), nil)

		newBalance, err := svc.Deposit(ctx, "user1", 100)
		assert.NoError(t, err)
		assert.Equal(t, int64(150), newBalance)

		mockStorage.AssertCalled(t, "AddBalance", ctx, "user1", int64(100))
	})

	t.Run("failure", func(t *testing.T) {
		mockStorage.On("AddBalance", ctx, "user2", int64(200)).Return(int64(0), errors.New("db error"))

		_, err := svc.Deposit(ctx, "user2", 200)
		assert.Error(t, err)
		assert.EqualError(t, err, "db error")
	})
}

func TestWithdraw(t *testing.T) {
	mockStorage := new(MockStorage)
	svc := NewAccountService(mockStorage)

	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		mockStorage.On("Withdraw", ctx, "user1", int64(50)).Return(int64(50), nil)

		newBalance, err := svc.Withdraw(ctx, "user1", 50)
		assert.NoError(t, err)
		assert.Equal(t, int64(50), newBalance)

		mockStorage.AssertCalled(t, "Withdraw", ctx, "user1", int64(50))
	})

	t.Run("insufficient funds", func(t *testing.T) {
		mockStorage.On("Withdraw", ctx, "user2", int64(999)).Return(int64(0), db.ErrInsufficientFunds)

		_, err := svc.Withdraw(ctx, "user2", 999)
		assert.Error(t, err)
		assert.Equal(t, db.ErrInsufficientFunds, err)
	})
}

func TestGetBalance(t *testing.T) {
	mockStorage := new(MockStorage)
	svc := NewAccountService(mockStorage)

	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		mockStorage.On("GetBalance", ctx, "user1").Return(int64(200), nil)

		balance, err := svc.GetBalance(ctx, "user1")
		assert.NoError(t, err)
		assert.Equal(t, int64(200), balance)

		mockStorage.AssertCalled(t, "GetBalance", ctx, "user1")
	})

	t.Run("not found", func(t *testing.T) {
		mockStorage.On("GetBalance", ctx, "userX").Return(int64(0), db.ErrAccountNotFound)

		_, err := svc.GetBalance(ctx, "userX")
		assert.Error(t, err)
		assert.Equal(t, db.ErrAccountNotFound, err)
	})
}
