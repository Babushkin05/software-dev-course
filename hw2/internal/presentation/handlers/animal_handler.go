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

// CreateAnimal создает новое животное
// @Summary Создать животное
// @Description Добавляет новое животное в систему
// @Tags animals
// @Accept json
// @Produce json
// @Param animal body CreateAnimalRequest true "Данные животного"
// @Success 201 {object} domain.Animal
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /animals [post]
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

// MoveAnimalRequest represents the request body for moving an animal to another enclosure
type MoveAnimalRequest struct {
	ToEnclosure string `json:"to_enclosure"`
}

// MoveAnimal перемещает животное в другой вольер
// @Summary Переместить животное
// @Description Перемещает животное в указанный вольер
// @Tags animals
// @Accept json
// @Produce json
// @Param id path string true "ID животного"
// @Param input body MoveAnimalRequest true "ID целевого вольера"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /animals/{id}/move [post]
func (h *AnimalHandler) MoveAnimal(c *gin.Context) {
	id := c.Param("id")
	var body MoveAnimalRequest
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

// ListAnimals возвращает список всех животных
// @Summary Получить список животных
// @Description Возвращает список всех животных в зоопарке
// @Tags animals
// @Produce json
// @Success 200 {array} domain.Animal
// @Failure 500 {object} map[string]string
// @Router /animals [get]
func (h *AnimalHandler) ListAnimals(c *gin.Context) {
	animals, err := h.AdminService.ListAnimals()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch animals"})
		return
	}
	c.JSON(http.StatusOK, animals)
}

// DeleteAnimal удаляет животное
// @Summary Удалить животное
// @Description Удаляет животное по ID
// @Tags animals
// @Produce json
// @Param id path string true "ID животного"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /animals/{id} [delete]
func (h *AnimalHandler) DeleteAnimal(c *gin.Context) {
	id := c.Param("id")
	err := h.AdminService.DeleteAnimal(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "animal not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "animal deleted"})
}
