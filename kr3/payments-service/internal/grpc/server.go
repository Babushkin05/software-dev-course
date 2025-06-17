package grpc

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/Babushkin05/software-dev-course/kr3/payments-service/api/gen"
	"github.com/Babushkin05/software-dev-course/kr3/payments-service/internal/service"
)

type Server struct {
	pb.UnimplementedPaymentsServiceServer
	svc *service.AccountService
}

func New(svc *service.AccountService) *Server {
	return &Server{svc: svc}
}

func (s *Server) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	acc, err := s.svc.CreateAccount(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.CreateAccountResponse{AccountId: acc.ID}, nil
}

func (s *Server) Deposit(ctx context.Context, req *pb.DepositRequest) (*pb.BalanceResponse, error) {
	balance, err := s.svc.Deposit(ctx, req.UserId, req.Amount)
	if err != nil {
		return nil, err
	}

	return &pb.BalanceResponse{Balance: balance}, nil
}

func (s *Server) GetBalance(ctx context.Context, req *pb.GetBalanceRequest) (*pb.BalanceResponse, error) {
	balance, err := s.svc.GetBalance(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.BalanceResponse{Balance: balance}, nil
}

func RunGRPCServer(svc *service.AccountService, addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPaymentsServiceServer(grpcServer, New(svc))

	reflection.Register(grpcServer)

	log.Printf("gRPC server listening at %s", addr)
	return grpcServer.Serve(lis)
}
