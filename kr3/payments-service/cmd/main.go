package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"

	"github.com/Babushkin05/software-dev-course/kr3/payments-service/internal/config"
	"github.com/Babushkin05/software-dev-course/kr3/payments-service/internal/db"
	grpcsrv "github.com/Babushkin05/software-dev-course/kr3/payments-service/internal/grpc"
	"github.com/Babushkin05/software-dev-course/kr3/payments-service/internal/kafka"
	"github.com/Babushkin05/software-dev-course/kr3/payments-service/internal/service"
	segmentioKafka "github.com/segmentio/kafka-go"
)

func main() {
	cfg := config.MustLoad()
	log.Printf("Loaded config: %+v\n", cfg)

	dbConn, err := sql.Open("postgres", cfg.Postgres.DSN)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.InitSchemaFromFile(dbConn); err != nil {
		log.Fatalf("failed to init schema: %v", err)
	}

	repo := db.NewAccountRepository(dbConn)
	svc := service.NewAccountService(repo)

	if err := grpcsrv.RunGRPCServer(svc, cfg.GRPC.Port); err != nil {
		log.Fatalf("failed to run gRPC server: %v", err)
	}

	ctx := context.Background()

	kafka.StartInboxConsumer(ctx, cfg.Kafka, func(msg *segmentioKafka.Message) {
		_ = kafka.HandleInboxMessage(msg, repo)
	})

	kafka.StartInboxProcessor(ctx, svc)

	writer := &segmentioKafka.Writer{
		Addr:     segmentioKafka.TCP(cfg.Kafka.Broker),
		Balancer: &segmentioKafka.LeastBytes{},
	}

	outboxWorker := kafka.NewOutboxWorker(repo, writer, 2*time.Second, 10)
	outboxWorker.Start(ctx)

}
