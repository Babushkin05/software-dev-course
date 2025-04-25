package handlers

import (
	"net/http"

	"github.com/Babushkin05/software-dev-course/hw2/internal/application/services"
	"github.com/gin-gonic/gin"
)

type StatisticsHandler struct {
	Service *services.ZooStatisticsService
}

func (h *StatisticsHandler) RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/statistics", h.GetStatistics)
}

func (h *StatisticsHandler) GetStatistics(c *gin.Context) {
	stats, err := h.Service.GetStatistics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get statistics"})
		return
	}
	c.JSON(http.StatusOK, stats)
}
