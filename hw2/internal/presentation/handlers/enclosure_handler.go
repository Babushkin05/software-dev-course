package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Babushkin05/software-dev-course/hw2/internal/application/ports"
	"github.com/Babushkin05/software-dev-course/hw2/internal/domain"
)

type EnclosureHandler struct {
	Repo ports.EnclosureRepository
}

type CreateEnclosureRequest struct {
	Type     string `json:"type"`
	Size     int    `json:"size"`
	Capacity int    `json:"capacity"`
}

func (h *EnclosureHandler) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/enclosures", h.Create)
	r.GET("/enclosures", h.List)
	r.DELETE("/enclosures/:id", h.Delete)
}

// Create создает новый вольер
// @Summary Создать вольер
// @Description Добавляет новый вольер в систему
// @Tags enclosures
// @Accept json
// @Produce json
// @Param enclosure body CreateEnclosureRequest true "Данные вольера"
// @Success 201 {object} domain.Enclosure
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /enclosures [post]
func (h *EnclosureHandler) Create(c *gin.Context) {
	var req CreateEnclosureRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	enclosureType := domain.EnclosureType(req.Type)
	enclosure, err := domain.NewEnclosure(enclosureType, req.Size, req.Capacity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.Repo.Save(enclosure)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save"})
		return
	}

	c.JSON(http.StatusCreated, enclosure)
}

// List возвращает список всех вольеров
// @Summary Получить список вольеров
// @Description Возвращает список всех вольеров в зоопарке
// @Tags enclosures
// @Produce json
// @Success 200 {array} domain.Enclosure
// @Failure 500 {object} map[string]string
// @Router /enclosures [get]
func (h *EnclosureHandler) List(c *gin.Context) {
	enclosures, err := h.Repo.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list"})
		return
	}
	c.JSON(http.StatusOK, enclosures)
}

// Delete удаляет вольер
// @Summary Удалить вольер
// @Description Удаляет вольер по ID
// @Tags enclosures
// @Produce json
// @Param id path string true "ID вольера"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /enclosures/{id} [delete]
func (h *EnclosureHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.Repo.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
