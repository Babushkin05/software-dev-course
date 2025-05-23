package main

import (
	"fmt"
	"log"
	"net"

	fileanalysispb "github.com/Babushkin05/software-dev-course/kr2/analise-service/api/gen/fileanalisys"
	"github.com/Babushkin05/software-dev-course/kr2/analise-service/internal/application/services"
	"github.com/Babushkin05/software-dev-course/kr2/analise-service/internal/config"
	analysis "github.com/Babushkin05/software-dev-course/kr2/analise-service/internal/infrastructure/analisys"
	grpcinfra "github.com/Babushkin05/software-dev-course/kr2/analise-service/internal/infrastructure/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfg := config.MustLoad()
	fmt.Printf("Loaded config: %+v\n", cfg)

	// Connect to storing-service
	storingConn, err := grpc.Dial(cfg.StoringService.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to storing service: %v", err)
	}
	defer storingConn.Close()

	// Init dependencies
	storingClient := grpcinfra.NewStoringClient(storingConn)
	sentimentAnalyzer := analysis.NewVaderAdapter()
	wordCloudGenerator := analysis.NewWordCloudAdapter()

	// Init application service
	analysisService := services.NewAnalysisService(storingClient, sentimentAnalyzer, wordCloudGenerator)

	// Init gRPC server
	listener, err := net.Listen("tcp", cfg.GRPC.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	fileanalysispb.RegisterFileAnalysisServiceServer(server, grpcinfra.NewGRPCServer(analysisService))

	log.Printf("Analysis service is running on %s", cfg.GRPC.Port)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
