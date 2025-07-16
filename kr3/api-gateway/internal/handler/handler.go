package handler

import (
	"github.com/Babushkin05/software-dev-course/kr3/api-gateway/internal/client"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	ordersClient   client.OrdersClient
	paymentsClient client.PaymentsClient
}

func NewHandler(orders client.OrdersClient, payments client.PaymentsClient) *Handler {
	return &Handler{
		ordersClient:   orders,
		paymentsClient: payments,
	}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		h.registerOrdersRoutes(api)
		h.registerPaymentsRoutes(api)
	}
}
