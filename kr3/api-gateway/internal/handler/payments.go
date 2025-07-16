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

// CreateAccount godoc
// @Summary Создать платёжный аккаунт
// @Description Создает аккаунт для пользователя по user_id
// @Tags payments
// @Accept json
// @Produce json
// @Param account body dto.CreateAccountRequest true "Данные пользователя"
// @Success 201 {object} paymentspb.CreateAccountResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /payments/account [post]
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

// Deposit godoc
// @Summary Пополнить счёт
// @Description Зачисляет средства на счет пользователя
// @Tags payments
// @Accept json
// @Produce json
// @Param deposit body dto.PaymentRequest true "Платёжные данные"
// @Success 200 {object} paymentspb.BalanceResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /payments/deposit [post]
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

// Withdraw godoc
// @Summary Снять средства
// @Description Списывает средства со счёта пользователя
// @Tags payments
// @Accept json
// @Produce json
// @Param withdraw body dto.PaymentRequest true "Платёжные данные"
// @Success 200 {object} paymentspb.BalanceResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /payments/withdraw [post]
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

// GetBalance godoc
// @Summary Получить баланс
// @Description Возвращает текущий баланс пользователя
// @Tags payments
// @Produce json
// @Param user_id query string true "ID пользователя"
// @Success 200 {object} paymentspb.BalanceResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /payments/balance [get]
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
