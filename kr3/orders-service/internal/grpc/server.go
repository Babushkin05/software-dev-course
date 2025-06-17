package grpc

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/Babushkin05/software-dev-course/kr3/orders-service/api/gen"
	"github.com/Babushkin05/software-dev-course/kr3/orders-service/internal/config"
	"github.com/Babushkin05/software-dev-course/kr3/orders-service/internal/service"
)

func RunGRPCServer(orderSvc *service.OrderService, cfg config.GRPCConfig) error {
	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	pb.RegisterOrdersServiceServer(s, NewHandler(orderSvc))

	log.Printf("gRPC server listening on %s", cfg.Port)
	return s.Serve(lis)
}
