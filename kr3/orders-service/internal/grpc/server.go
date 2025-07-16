package grpc

import (
	"fmt"
	"log"
	"net"

	orderspb "github.com/Babushkin05/software-dev-course/kr3/orders-service/api/gen"
	"github.com/Babushkin05/software-dev-course/kr3/orders-service/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RunGRPCServer(svc *service.OrderService, port string) error {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	server := grpc.NewServer()
	orderspb.RegisterOrdersServiceServer(server, NewHandler(svc))

	reflection.Register(server)

	log.Printf("gRPC server started on port %s", port)
	return server.Serve(listener)
}
