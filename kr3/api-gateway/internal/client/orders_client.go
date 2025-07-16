package client

import (
	"context"

	orderspb "github.com/Babushkin05/software-dev-course/kr3/api-gateway/api/gen/orders"
	"google.golang.org/grpc"
)

type OrdersClient interface {
	CreateOrder(ctx context.Context, req *orderspb.CreateOrderRequest) (*orderspb.OrderResponse, error)
	GetOrders(ctx context.Context, req *orderspb.GetOrdersRequest) (*orderspb.OrdersList, error)
	GetOrderStatus(ctx context.Context, req *orderspb.GetOrderStatusRequest) (*orderspb.OrderStatusResponse, error)
}

type ordersClient struct {
	client orderspb.OrdersServiceClient
}

func NewOrdersClient(conn *grpc.ClientConn) OrdersClient {
	return &ordersClient{
		client: orderspb.NewOrdersServiceClient(conn),
	}
}

func (c *ordersClient) CreateOrder(ctx context.Context, req *orderspb.CreateOrderRequest) (*orderspb.OrderResponse, error) {
	return c.client.CreateOrder(ctx, req)
}

func (c *ordersClient) GetOrders(ctx context.Context, req *orderspb.GetOrdersRequest) (*orderspb.OrdersList, error) {
	return c.client.GetOrders(ctx, req)
}

func (c *ordersClient) GetOrderStatus(ctx context.Context, req *orderspb.GetOrderStatusRequest) (*orderspb.OrderStatusResponse, error) {
	return c.client.GetOrderStatus(ctx, req)
}
