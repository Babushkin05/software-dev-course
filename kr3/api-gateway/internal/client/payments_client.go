package client

import (
	"context"

	paymentspb "github.com/Babushkin05/software-dev-course/kr3/api-gateway/api/gen/payments"
	"google.golang.org/grpc"
)

type PaymentsClient interface {
	CreateAccount(ctx context.Context, req *paymentspb.CreateAccountRequest) (*paymentspb.CreateAccountResponse, error)
	Deposit(ctx context.Context, req *paymentspb.DepositRequest) (*paymentspb.BalanceResponse, error)
	GetBalance(ctx context.Context, req *paymentspb.GetBalanceRequest) (*paymentspb.BalanceResponse, error)
	Withdraw(ctx context.Context, req *paymentspb.WithdrawRequest) (*paymentspb.BalanceResponse, error)
}

type paymentsClient struct {
	client paymentspb.PaymentsServiceClient
}

func NewPaymentsClient(conn *grpc.ClientConn) PaymentsClient {
	return &paymentsClient{
		client: paymentspb.NewPaymentsServiceClient(conn),
	}
}

func (c *paymentsClient) CreateAccount(ctx context.Context, req *paymentspb.CreateAccountRequest) (*paymentspb.CreateAccountResponse, error) {
	return c.client.CreateAccount(ctx, req)
}

func (c *paymentsClient) Deposit(ctx context.Context, req *paymentspb.DepositRequest) (*paymentspb.BalanceResponse, error) {
	return c.client.Deposit(ctx, req)
}

func (c *paymentsClient) GetBalance(ctx context.Context, req *paymentspb.GetBalanceRequest) (*paymentspb.BalanceResponse, error) {
	return c.client.GetBalance(ctx, req)
}

func (c *paymentsClient) Withdraw(ctx context.Context, req *paymentspb.WithdrawRequest) (*paymentspb.BalanceResponse, error) {
	return c.client.Withdraw(ctx, req)
}
