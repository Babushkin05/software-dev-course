package client

import (
	"time"

	"google.golang.org/grpc"
)

type Clients struct {
	OrdersClient   OrdersClient
	PaymentsClient PaymentsClient
}

func NewClients(ordersAddr, paymentsAddr string) (*Clients, error) {
	ordersConn, err := grpc.Dial(ordersAddr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		return nil, err
	}

	paymentsConn, err := grpc.Dial(paymentsAddr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		return nil, err
	}

	return &Clients{
		OrdersClient:   NewOrdersClient(ordersConn),
		PaymentsClient: NewPaymentsClient(paymentsConn),
	}, nil
}

func (c *Clients) Close() {
	if closer, ok := c.OrdersClient.(interface{ Close() }); ok {
		closer.Close()
	}
	if closer, ok := c.PaymentsClient.(interface{ Close() }); ok {
		closer.Close()
	}
}
