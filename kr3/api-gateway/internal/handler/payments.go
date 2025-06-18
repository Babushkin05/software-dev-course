package handler

import (
	"net/http"

	paymentspb "github.com/Babushkin05/software-dev-course/kr3/api-gateway/api/gen/payments"
	"github.com/Babushkin05/software-dev-course/kr3/api-gateway/internal/dto"
	"github.com/gin-gonic/gin"
)

func (h *Handler) registerPaymentsRoutes(rg *gin.RouterGroup) {
	payments := rg.Group("/payments")
	{
		payments.POST("/account", h.CreateAccount)
		payments.POST("/deposit", h.Deposit)
		payments.POST("/withdraw", h.Withdraw)
		payments.GET("/balance", h.GetBalance)
	}
}

func (h *Handler) CreateAccount(c *gin.Context) {
	var req dto.CreateAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.paymentsClient.CreateAccount(c.Request.Context(), &paymentspb.CreateAccountRequest{
		UserId: req.UserID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *Handler) Deposit(c *gin.Context) {
	var req dto.PaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.paymentsClient.Deposit(c.Request.Context(), &paymentspb.DepositRequest{
		UserId: req.UserID,
		Amount: req.Amount,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) Withdraw(c *gin.Context) {
	var req dto.PaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.paymentsClient.Withdraw(c.Request.Context(), &paymentspb.WithdrawRequest{
		UserId: req.UserID,
		Amount: req.Amount,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetBalance(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	resp, err := h.paymentsClient.GetBalance(c.Request.Context(), &paymentspb.GetBalanceRequest{
		UserId: userID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
