package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/Babushkin05/software-dev-course/kr3/payments-service/internal/config"
	"github.com/Babushkin05/software-dev-course/kr3/payments-service/internal/db"
	grpcsrv "github.com/Babushkin05/software-dev-course/kr3/payments-service/internal/grpc"
	"github.com/Babushkin05/software-dev-course/kr3/payments-service/internal/service"
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

}
