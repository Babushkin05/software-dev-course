package presentation

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/Babushkin05/software-dev-course/hw2/internal/application/ports"
	"github.com/Babushkin05/software-dev-course/hw2/internal/application/services"
	"github.com/Babushkin05/software-dev-course/hw2/internal/presentation/handlers"
)

func SetupRouter(
	transferService *services.AnimalTransferService,
	adminService *services.ZooAdminService,
	enclosureRepo ports.EnclosureRepository,
	scheduleService *services.FeedingOrganizerService,
	statsService *services.ZooStatisticsService,
) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	animalHandler := &handlers.AnimalHandler{
		TransferService: transferService,
		AdminService:    adminService,
	}
	animalHandler.RegisterRoutes(api)

	enclosureHandler := &handlers.EnclosureHandler{
		Repo: enclosureRepo,
	}
	enclosureHandler.RegisterRoutes(api)

	scheduleHandler := &handlers.FeedingScheduleHandler{
		Service: scheduleService,
	}
	scheduleHandler.RegisterRoutes(api)

	statsHandler := &handlers.StatisticsHandler{
		Service: statsService,
	}
	statsHandler.RegisterRoutes(api)

	return r
}
