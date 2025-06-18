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

// CreateOrder godoc
// @Summary Создание заказа
// @Description Создает новый заказ для пользователя
// @Tags orders
// @Accept json
// @Produce json
// @Param order body dto.CreateOrderRequest true "Данные для создания заказа"
// @Success 201 {object} orderspb.OrderResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /orders/ [post]
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

// GetOrders godoc
// @Summary Получение заказов пользователя
// @Description Возвращает список заказов по user_id
// @Tags orders
// @Produce json
// @Param user_id query string true "ID пользователя"
// @Success 200 {object} orderspb.OrdersList
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /orders/ [get]
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

// GetOrderStatus godoc
// @Summary Получение статуса заказа
// @Description Возвращает статус заказа по его ID
// @Tags orders
// @Produce json
// @Param id path string true "ID заказа"
// @Success 200 {object} orderspb.OrderStatusResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /orders/{id}/status [get]
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
