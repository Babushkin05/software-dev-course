package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	fileanalysispb "github.com/Babushkin05/software-dev-course/kr2/api-gateway/api/gen/fileanalysispb"
	filestoringpb "github.com/Babushkin05/software-dev-course/kr2/api-gateway/api/gen/filestoringpb"
	"github.com/Babushkin05/software-dev-course/kr2/api-gateway/internal/application/services"
	"github.com/Babushkin05/software-dev-course/kr2/api-gateway/internal/config"
	grpcinfra "github.com/Babushkin05/software-dev-course/kr2/api-gateway/internal/infrastructure/service_grpc"
	"github.com/Babushkin05/software-dev-course/kr2/api-gateway/internal/presentation/service_http"
	"google.golang.org/grpc"
)

func main() {
	// 1. Load config
	cfg := config.MustLoad()
	fmt.Printf("Loaded config: %+v\n", cfg)

	// 2. Set up gRPC connections with timeout from config
	ctx, cancel := context.WithTimeout(context.Background(), cfg.GRPC.Timeout)
	defer cancel()

	fsConn, err := grpc.DialContext(
		ctx,
		cfg.Services.FileStoring,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("failed to connect to file storing service: %v", err)
	}
	defer fsConn.Close()

	faConn, err := grpc.DialContext(
		ctx,
		cfg.Services.FileAnalysis,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("failed to connect to file analysis service: %v", err)
	}
	defer faConn.Close()

	// 3. Create gRPC clients
	fsClient := filestoringpb.NewFileStoringServiceClient(fsConn)
	faClient := fileanalysispb.NewFileAnalysisServiceClient(faConn)

	// 4. Wrap them with infrastructure adapters
	fsAdapter := grpcinfra.NewFileStoringClient(fsClient)
	faAdapter := grpcinfra.NewFileAnalysisClient(faClient)

	// 5. Build usecase
	usecase := services.NewFileService(fsAdapter, faAdapter)

	// 6. Start HTTP server on port from config
	addr := fmt.Sprintf(":%d", cfg.HTTP.Port)
	fmt.Println("API Gateway running on", addr)
	router := service_http.NewRouter(usecase)

	// Optional: create http.Server with graceful shutdown if needed
	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("could not start server: %v", err)
	}
}
