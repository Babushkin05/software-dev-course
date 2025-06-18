package handler

import (
	"net/http"

	orderspb "github.com/Babushkin05/software-dev-course/kr3/api-gateway/api/gen/orders"
	"github.com/Babushkin05/software-dev-course/kr3/api-gateway/internal/dto"
	"github.com/gin-gonic/gin"
)

func (h *Handler) registerOrdersRoutes(rg *gin.RouterGroup) {
	orders := rg.Group("/orders")
	{
		orders.POST("/", h.CreateOrder)
		orders.GET("/", h.GetOrders)
		orders.GET("/:id/status", h.GetOrderStatus)
	}
}

// POST /api/orders
func (h *Handler) CreateOrder(c *gin.Context) {
	var req dto.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grpcReq := &orderspb.CreateOrderRequest{
		UserId:      req.UserID,
		Amount:      req.Amount,
		Description: req.Description,
	}

	resp, err := h.ordersClient.CreateOrder(c.Request.Context(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GET /api/orders?user_id=123
func (h *Handler) GetOrders(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	grpcReq := &orderspb.GetOrdersRequest{
		UserId: userID,
	}

	resp, err := h.ordersClient.GetOrders(c.Request.Context(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GET /api/orders/:id/status
func (h *Handler) GetOrderStatus(c *gin.Context) {
	orderID := c.Param("id")
	if orderID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "order_id is required"})
		return
	}

	grpcReq := &orderspb.GetOrderStatusRequest{
		OrderId: orderID,
	}

	resp, err := h.ordersClient.GetOrderStatus(c.Request.Context(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
