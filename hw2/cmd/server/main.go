package main

import (
	_ "github.com/Babushkin05/software-dev-course/hw2/docs"
	"github.com/Babushkin05/software-dev-course/hw2/internal/application/services"
	"github.com/Babushkin05/software-dev-course/hw2/internal/infrastructure/events"
	"github.com/Babushkin05/software-dev-course/hw2/internal/infrastructure/storage"
	"github.com/Babushkin05/software-dev-course/hw2/internal/presentation"
)

// @title Zoo Management API
// @version 1.0
// @description API для управления зоопарком: животные, вольеры, расписание кормления и статистика.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://example.com/support
// @contact.email support@zoo.example.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api
// @schemes http
func main() {
	// Infrastructure (In-Memory + Console Publisher)
	animalRepo := storage.NewInMemoryAnimalRepo()
	enclosureRepo := storage.NewInMemoryEnclosureRepo()
	scheduleRepo := storage.NewInMemoryFeedingScheduleRepo()
	eventPublisher := events.NewConsoleEventPublisher()

	// Application Layer: Services
	animalTransferService := services.NewAnimalTransferService(animalRepo, enclosureRepo, eventPublisher)
	adminService := services.NewZooAdminService(animalRepo)
	scheduleService := services.NewFeedingOrganizerService(scheduleRepo, eventPublisher)
	statsService := services.NewZooStatisticsService(animalRepo, enclosureRepo)

	// Web API (Presentation Layer)
	r := presentation.SetupRouter(
		animalTransferService,
		adminService,
		enclosureRepo,
		scheduleService,
		statsService,
	)

	// Start Server
	r.Run(":8080")
}
