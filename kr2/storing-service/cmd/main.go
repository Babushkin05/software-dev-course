package main

import (
	"fmt"
	"log"
	"net"

	filestoringpb "github.com/Babushkin05/software-dev-course/kr2/storing-service/api/gen"
	"github.com/Babushkin05/software-dev-course/kr2/storing-service/internal/application/services"
	"github.com/Babushkin05/software-dev-course/kr2/storing-service/internal/config"
	grpc_server "github.com/Babushkin05/software-dev-course/kr2/storing-service/internal/infrastructure/grpc"
	"github.com/Babushkin05/software-dev-course/kr2/storing-service/internal/infrastructure/postgres"
	"github.com/Babushkin05/software-dev-course/kr2/storing-service/internal/infrastructure/storage"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	// 1. Load environment variables from .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found — skipping...")
	}

	// 2. Load YAML config
	cfg := config.MustLoad()
	fmt.Printf("Loaded config: %+v\n", cfg)

	// 3. Init PostgreSQL
	dbConn, err := postgres.NewDB(cfg.PG)
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}
	defer dbConn.Close()

	if err := postgres.InitSchema(dbConn); err != nil {
		log.Fatalf("failed to initialize schema: %v", err)
	}

	db := postgres.NewFileRepo(dbConn)

	// 4. Init S3 storage
	s3Client, err := storage.NewS3Storage(cfg.S3)
	if err != nil {
		log.Fatalf("failed to init s3 storage: %v", err)
	}

	// 5. Init application layer
	service := services.NewFileService(s3Client, db)

	// 6. Init gRPC server
	grpcServer := grpc.NewServer()

	// 7. Register service implementation
	fileStoringServer := grpc_server.NewFileStoringServer(service)
	filestoringpb.RegisterFileStoringServiceServer(grpcServer, fileStoringServer)

	// 8. Start listening
	addr := fmt.Sprintf(":%d", cfg.GRPC.Port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Printf("gRPC server running on port %d...\n", cfg.GRPC.Port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve gRPC: %v", err)
	}
}
