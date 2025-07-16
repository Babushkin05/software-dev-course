package main

import (
	"context"
	"database/sql"
	"log"
	"net"

	orderspb "github.com/Babushkin05/software-dev-course/kr3/orders-service/api/gen"
	"github.com/Babushkin05/software-dev-course/kr3/orders-service/internal/config"
	"github.com/Babushkin05/software-dev-course/kr3/orders-service/internal/db"
	"github.com/Babushkin05/software-dev-course/kr3/orders-service/internal/grpc"
	"github.com/Babushkin05/software-dev-course/kr3/orders-service/internal/kafka"
	"github.com/Babushkin05/software-dev-course/kr3/orders-service/internal/service"

	_ "github.com/lib/pq"
	segmentioKafka "github.com/segmentio/kafka-go"
	googleGrpc "google.golang.org/grpc"
)

func main() {
	ctx := context.Background()

	// --- Load Config ---
	cfg := config.MustLoad()
	log.Printf("Loaded config: %+v\n", cfg)

	// --- Connect to PostgreSQL ---
	dbConn, err := sql.Open("postgres", cfg.Postgres.DSN)
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}

	// --- Init Repository ---
	repo := db.NewOrderRepository(dbConn)

	// --- Kafka Writer (OutboxWorker will use it to publish) ---
	kafkaWriter := &segmentioKafka.Writer{
		Addr:     segmentioKafka.TCP(cfg.Kafka.Broker),
		Topic:    cfg.Kafka.Topic,
		Balancer: &segmentioKafka.LeastBytes{},
	}

	// --- Init Outbox Worker ---
	outboxWorker := kafka.NewOutboxWorker(repo, kafkaWriter)

	// --- Init OrderWriter (writes to outbox table only) ---
	orderWriter := kafka.NewWriter(repo, cfg.Kafka.Topic)

	// --- Init Order Service ---
	orderService := service.NewOrderService(repo, orderWriter)

	// --- Init gRPC Handler ---
	grpcHandler := grpc.NewHandler(orderService)

	// --- Init Inbox Consumer & Processor ---
	inboxConsumer := kafka.NewInboxConsumer(cfg.Kafka.Broker, cfg.Kafka.Topic, cfg.Kafka.GroupID, repo)
	inboxProcessor := kafka.NewInboxProcessor(repo)

	// --- Start Workers ---
	outboxWorker.Start(ctx)
	inboxConsumer.Start(ctx)
	inboxProcessor.Start(ctx)

	// --- Start gRPC Server ---
	listener, err := net.Listen("tcp", ":"+cfg.GRPC.Port)
	if err != nil {
		log.Fatalf("failed to listen on port %s: %v", cfg.GRPC.Port, err)
	}

	grpcServer := googleGrpc.NewServer()
	orderspb.RegisterOrdersServiceServer(grpcServer, grpcHandler)

	log.Printf("gRPC server is running on :%s", cfg.GRPC.Port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}
