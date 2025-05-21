package service_http

import (
	"github.com/Babushkin05/software-dev-course/kr2/api-gateway/internal/domain/ports/input"
	"github.com/Babushkin05/software-dev-course/kr2/api-gateway/internal/presentation/service_http/handlers"

	"github.com/gin-gonic/gin"
)

func NewRouter(usecase input.FileUsecase) *gin.Engine {
	router := gin.Default()

	fileHandler := handlers.NewFileHandler(usecase)
	fileHandler.RegisterRoutes(router)

	return router
}
