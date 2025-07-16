package handlers

import (
	"net/http"
	"time"

	"github.com/Babushkin05/software-dev-course/hw2/internal/application/services"
	"github.com/gin-gonic/gin"
)

type FeedingScheduleHandler struct {
	Service services.IFeedingOrganizerService
}

func NewFeedingScheduleHandler(s services.IFeedingOrganizerService) FeedingScheduleHandler {
	return FeedingScheduleHandler{Service: s}
}

type CreateFeedingRequest struct {
	AnimalID string    `json:"animal_id"`
	Time     time.Time `json:"time"`
	Food     string    `json:"food"`
}

func (h *FeedingScheduleHandler) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/feedings", h.CreateFeeding)
	r.GET("/feedings", h.ListFeedings)
}

// CreateFeeding создает новое расписание кормления
// @Summary Создать расписание кормления
// @Description Добавляет новое расписание кормления животного
// @Tags feedings
// @Accept json
// @Produce json
// @Param schedule body CreateFeedingRequest true "Данные расписания кормления"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /feedings [post]
func (h *FeedingScheduleHandler) CreateFeeding(c *gin.Context) {
	var req CreateFeedingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	err := h.Service.AddFeeding(req.AnimalID, req.Time, req.Food)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "feeding scheduled"})
}

// ListFeedings возвращает список всех расписаний кормления
// @Summary Получить расписания кормления
// @Description Возвращает список всех расписаний кормления
// @Tags feedings
// @Produce json
// @Success 200 {array} domain.FeedingSchedule
// @Failure 500 {object} map[string]string
// @Router /feedings [get]
func (h *FeedingScheduleHandler) ListFeedings(c *gin.Context) {
	schedules, err := h.Service.GetAllSchedules()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list feedings"})
		return
	}
	c.JSON(http.StatusOK, schedules)
}
