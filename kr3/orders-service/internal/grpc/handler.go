package grpc

import (
	"context"

	pb "github.com/Babushkin05/software-dev-course/kr3/orders-service/api/gen"
	"github.com/Babushkin05/software-dev-course/kr3/orders-service/internal/service"
)

type handler struct {
	pb.UnimplementedOrdersServiceServer
	svc *service.OrderService
}

func NewHandler(svc *service.OrderService) pb.OrdersServiceServer {
	return &handler{svc: svc}
}

func (h *handler) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	order, err := h.svc.CreateOrder(ctx, req.UserId, req.Amount, req.Description)
	if err != nil {
		return nil, err
	}

	return &pb.OrderResponse{
		Order: &pb.Order{
			Id:          order.ID,
			UserId:      order.UserID,
			Amount:      order.Amount,
			Description: order.Description,
			Status:      string(order.Status),
		},
	}, nil
}

func (h *handler) ListOrders(ctx context.Context, req *pb.GetOrdersRequest) (*pb.ListOrdersResponse, error) {
	orders, err := h.svc.GetOrders(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	var pbOrders []*pb.Order
	for _, o := range orders {
		pbOrders = append(pbOrders, &pb.Order{
			Id:          o.ID,
			UserId:      o.UserID,
			Amount:      o.Amount,
			Description: o.Description,
			Status:      string(o.Status),
		})
	}

	return &pb.OrdersList{Orders: pbOrders}, nil
}

func (h *handler) GetOrderStatus(ctx context.Context, req *pb.GetOrderStatusRequest) (*pb.GetOrderResponse, error) {
	order, err := h.svc.GetByID(ctx, req.OrderId)
	if err != nil {
		return nil, err
	}

	return &pb.OrderStatusResponse{
		Status: string(order.Status),
	}, nil
}
