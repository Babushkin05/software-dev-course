package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Babushkin05/software-dev-course/hw2/internal/application/services"
	"github.com/Babushkin05/software-dev-course/hw2/internal/domain"
)

type AnimalHandler struct {
	TransferService *services.AnimalTransferService
	AdminService    *services.ZooAdminService
}

type CreateAnimalRequest struct {
	Name         string    `json:"name"`
	Species      string    `json:"species"`
	BirthDate    time.Time `json:"birth_date"`
	Gender       string    `json:"gender"`
	FavoriteFood string    `json:"favorite_food"`
}

func (h *AnimalHandler) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/animals", h.CreateAnimal)
	r.GET("/animals", h.ListAnimals)
	r.DELETE("/animals/:id", h.DeleteAnimal)
	r.POST("/animals/:id/move", h.MoveAnimal)
}

func (h *AnimalHandler) CreateAnimal(c *gin.Context) {
	var req CreateAnimalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	gender, err := domain.NewGender(req.Gender)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid gender"})
		return
	}

	animal, err := domain.NewAnimal(req.Name, req.Species, req.BirthDate, gender, req.FavoriteFood)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.AdminService.CreateAnimal(animal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save"})
		return
	}

	c.JSON(http.StatusCreated, animal)
}

func (h *AnimalHandler) MoveAnimal(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		ToEnclosure string `json:"to_enclosure"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	err := h.TransferService.MoveAnimal(id, body.ToEnclosure)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "animal moved"})
}

func (h *AnimalHandler) ListAnimals(c *gin.Context) {
	animals, err := h.AdminService.ListAnimals()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch animals"})
		return
	}
	c.JSON(http.StatusOK, animals)
}

func (h *AnimalHandler) DeleteAnimal(c *gin.Context) {
	id := c.Param("id")
	err := h.AdminService.DeleteAnimal(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "animal not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "animal deleted"})
}
