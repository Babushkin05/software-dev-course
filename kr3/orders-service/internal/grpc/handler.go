package grpc

import (
	"context"

	orderspb "github.com/Babushkin05/software-dev-course/kr3/orders-service/api/gen"
	"github.com/Babushkin05/software-dev-course/kr3/orders-service/internal/service"
)

type handler struct {
	orderspb.UnimplementedOrdersServiceServer
	svc *service.OrderService
}

func NewHandler(svc *service.OrderService) orderspb.OrdersServiceServer {
	return &handler{svc: svc}
}

func (h *handler) CreateOrder(ctx context.Context, req *orderspb.CreateOrderRequest) (*orderspb.OrderResponse, error) {
	order, err := h.svc.CreateOrder(ctx, req.UserId, req.Amount, req.Description)
	if err != nil {
		return nil, err
	}

	return &orderspb.OrderResponse{
		Order: &orderspb.Order{
			Id:          order.ID,
			UserId:      order.UserID,
			Amount:      order.Amount,
			Description: order.Description,
			Status:      string(order.Status),
			CreatedAt:   order.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		},
	}, nil
}

func (h *handler) GetOrders(ctx context.Context, req *orderspb.GetOrdersRequest) (*orderspb.OrdersList, error) {
	orders, err := h.svc.GetOrders(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	var pbOrders []*orderspb.Order
	for _, o := range orders {
		pbOrders = append(pbOrders, &orderspb.Order{
			Id:          o.ID,
			UserId:      o.UserID,
			Amount:      o.Amount,
			Description: o.Description,
			Status:      string(o.Status),
			CreatedAt:   o.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	return &orderspb.OrdersList{Orders: pbOrders}, nil
}

func (h *handler) GetOrderStatus(ctx context.Context, req *orderspb.GetOrderStatusRequest) (*orderspb.OrderStatusResponse, error) {
	order, err := h.svc.GetByID(ctx, req.OrderId)
	if err != nil {
		return nil, err
	}

	return &orderspb.OrderStatusResponse{
		Status: string(order.Status),
	}, nil
}

func (h *handler) MarkOrderFinished(ctx context.Context, req *orderspb.MarkOrderFinishedRequest) (*orderspb.Empty, error) {
	if err := h.svc.MarkFinished(ctx, req.OrderId); err != nil {
		return nil, err
	}
	return &orderspb.Empty{}, nil
}
