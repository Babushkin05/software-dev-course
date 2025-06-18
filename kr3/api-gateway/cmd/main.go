package main

import (
	"log"

	"github.com/Babushkin05/software-dev-course/kr3/api-gateway/internal/client"
	"github.com/Babushkin05/software-dev-course/kr3/api-gateway/internal/config"
	"github.com/Babushkin05/software-dev-course/kr3/api-gateway/internal/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"google.golang.org/grpc"

	_ "github.com/Babushkin05/software-dev-course/kr3/api-gateway/docs"
)

// @title KR3 API Gateway
// @version 1.0
// @description Gateway to Orders and Payments microservices.
// @BasePath /api/v1
func main() {
	cfg := config.MustLoad()

	// --- gRPC Connections ---
	ordersConn, err := grpc.Dial(cfg.Services.Orders.Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial OrdersService: %v", err)
	}
	defer ordersConn.Close()

	paymentsConn, err := grpc.Dial(cfg.Services.Payments.Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial PaymentsService: %v", err)
	}
	defer paymentsConn.Close()

	// --- gRPC Clients ---
	ordersClient := client.NewOrdersClient(ordersConn)
	paymentsClient := client.NewPaymentsClient(paymentsConn)

	// --- Handler ---
	handler := handler.NewHandler(ordersClient, paymentsClient)

	// --- Gin Setup ---
	router := gin.Default()

	// --- Register API routes ---
	handler.RegisterRoutes(router)

	// --- Swagger ---
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// --- Start Server ---
	log.Printf("API Gateway running at http://localhost:%s", cfg.HTTP.Port)
	if err := router.Run(":" + cfg.HTTP.Port); err != nil {
		log.Fatalf("failed to run HTTP server: %v", err)
	}
}
