package handlers

import (
	"net/http"
	"time"

	"github.com/Babushkin05/software-dev-course/hw2/internal/application/services"
	"github.com/gin-gonic/gin"
)

type FeedingScheduleHandler struct {
	Service *services.FeedingOrganizerService
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

func (h *FeedingScheduleHandler) ListFeedings(c *gin.Context) {
	schedules, err := h.Service.GetAllSchedules()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list feedings"})
		return
	}
	c.JSON(http.StatusOK, schedules)
}
