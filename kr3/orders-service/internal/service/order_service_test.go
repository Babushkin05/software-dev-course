package service

import (
	"context"
	"errors"
	"testing"

	"github.com/Babushkin05/software-dev-course/kr3/orders-service/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// --- Mocks ---

type mockStorage struct {
	mock.Mock
}

func (m *mockStorage) Create(ctx context.Context, order *model.Order) error {
	args := m.Called(ctx, order)
	return args.Error(0)
}

func (m *mockStorage) GetByUser(ctx context.Context, userID string) ([]*model.Order, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]*model.Order), args.Error(1)
}

func (m *mockStorage) GetByID(ctx context.Context, orderID string) (*model.Order, error) {
	args := m.Called(ctx, orderID)
	return args.Get(0).(*model.Order), args.Error(1)
}

func (m *mockStorage) UpdateStatus(ctx context.Context, orderID string, status model.OrderStatus) error {
	args := m.Called(ctx, orderID, status)
	return args.Error(0)
}

func (m *mockStorage) SaveOutboxMessage(ctx context.Context, topic string, key *string, payload string) error {
	args := m.Called(ctx, topic, key, payload)
	return args.Error(0)
}

type mockKafkaWriter struct {
	mock.Mock
}

func (m *mockKafkaWriter) WriteOrder(ctx context.Context, order *model.Order) error {
	args := m.Called(ctx, order)
	return args.Error(0)
}

// --- Tests ---

func TestCreateOrder(t *testing.T) {
	storage := new(mockStorage)
	kafka := new(mockKafkaWriter)
	service := NewOrderService(storage, kafka)

	ctx := context.Background()
	userID := "user-123"
	amount := int64(500)
	desc := "Test order"

	// Any order passed to Create/WriteOrder will be checked by type, not exact values
	storage.On("Create", ctx, mock.AnythingOfType("*model.Order")).Return(nil)
	kafka.On("WriteOrder", ctx, mock.AnythingOfType("*model.Order")).Return(nil)

	order, err := service.CreateOrder(ctx, userID, amount, desc)
	assert.NoError(t, err)
	assert.Equal(t, userID, order.UserID)
	assert.Equal(t, amount, order.Amount)
	assert.Equal(t, model.OrderStatusNew, order.Status)

	storage.AssertExpectations(t)
	kafka.AssertExpectations(t)
}

func TestGetOrders(t *testing.T) {
	storage := new(mockStorage)
	service := NewOrderService(storage, nil)

	ctx := context.Background()
	userID := "user-1"
	expected := []*model.Order{
		{ID: "1", UserID: userID, Amount: 100},
		{ID: "2", UserID: userID, Amount: 200},
	}

	storage.On("GetByUser", ctx, userID).Return(expected, nil)

	orders, err := service.GetOrders(ctx, userID)
	assert.NoError(t, err)
	assert.Equal(t, expected, orders)
	storage.AssertExpectations(t)
}

func TestGetOrderStatus(t *testing.T) {
	storage := new(mockStorage)
	service := NewOrderService(storage, nil)

	ctx := context.Background()
	orderID := "order-1"
	order := &model.Order{ID: orderID, Status: model.OrderStatusFinished}

	storage.On("GetByID", ctx, orderID).Return(order, nil)

	status, err := service.GetOrderStatus(ctx, orderID)
	assert.NoError(t, err)
	assert.Equal(t, model.OrderStatusFinished, status)
	storage.AssertExpectations(t)
}

func TestMarkOrderFinished(t *testing.T) {
	storage := new(mockStorage)
	service := NewOrderService(storage, nil)

	ctx := context.Background()
	orderID := "order-xyz"

	storage.On("UpdateStatus", ctx, orderID, model.OrderStatusFinished).Return(nil)

	err := service.MarkOrderFinished(ctx, orderID)
	assert.NoError(t, err)
	storage.AssertExpectations(t)
}

func TestCreateOrder_Error(t *testing.T) {
	storage := new(mockStorage)
	kafka := new(mockKafkaWriter)
	service := NewOrderService(storage, kafka)

	ctx := context.Background()

	// Simулируем ошибку при сохранении
	storage.On("Create", ctx, mock.Anything).Return(errors.New("db error"))

	_, err := service.CreateOrder(ctx, "u", 100, "fail")
	assert.Error(t, err)
	storage.AssertExpectations(t)
}
