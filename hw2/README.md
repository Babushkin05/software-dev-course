# Второе домашнее задание по КПО

[условие](./КПО%20мини-ДЗ2.pdf)

## запуск

```sh
go run ./cmd/server  
```
swagger:
http://localhost:8080/swagger/index.html#/


## Функционал

Я реализовал все пункты из требуемого функционала

### Domain

[Animal](./internal/domain/animal.go)

[Enclosure](./internal/domain/enclosure.go)

[FeedingSchedule](./internal/domain/feeding_schedule.go)

### Services

[AnimalTransferService](./internal/application/services/animal_transfer_service.go)

[FeedingOrganizationService](./internal/application/services/feeding_organizer_service.go)

[ZooStatisticsService](./internal/application/services/zoo_statistics_service.go)

### Структура Проекта

* [Domain](./internal/domain/) (ядро, содержит наши модели)

* [Application](./internal/application/) (содержит сервисы, реализующие бизнес-логику приложения)

* [Infrastructure](./internal/infrastructure/) (внешние взаимодействия)

* [Presentation](./internal/presentation/) (контроллеры нашего веб-приложения)

## Примененные принципы

### DDD

* Богатая доменная модель, потому что структуры ```Animal```, ```Enclosure``` и ```ScheduleFeeding``` содержат не только данные, но и поведение

* Наличие ```ValueObjects```, такие как ```Gender``` ит```HealthStatus```

* Совпадение с языком доменной в области (как в настоящем зоопарке)

### Чистая Архитектура

* Слои с односторонней зависимостью

* Инверсия зависимостей - интерфейсы определены во внутренних слоях, а их реализации — во внешних

## ограничения

как можно увидеть смотря на файл [compability.go](./internal/application/services/compability.go)

```go
func IsCompatible(species string, enclosureType domain.EnclosureType) bool {
	switch species {
	case "lion", "tiger":
		return enclosureType == domain.Predator
	case "zebra", "deer":
		return enclosureType == domain.Herbivore
	case "parrot":
		return enclosureType == domain.BirdCage
	case "fish":
		return enclosureType == domain.Aquarium
	default:
		return false
	}
}
```
сейчас, программа поддерживает очень ограниченное количество животных

## тесты

```sh
make test-coverage 
```

итого 74% покрытие

```
github.com/Babushkin05/software-dev-course/hw2/cmd/server/main.go:26:						main				0.0%
github.com/Babushkin05/software-dev-course/hw2/docs/docs.go:665:						init				0.0%
github.com/Babushkin05/software-dev-course/hw2/internal/application/services/animal_transfer_service.go:18:	NewAnimalTransferService	100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/application/services/animal_transfer_service.go:22:	MoveAnimal			100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/application/services/compability.go:5:			IsCompatible			33.3%
github.com/Babushkin05/software-dev-course/hw2/internal/application/services/feeding_organizer_service.go:21:	NewFeedingOrganizerService	100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/application/services/feeding_organizer_service.go:25:	AddFeeding			100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/application/services/feeding_organizer_service.go:33:	NotifyFeedingDue		0.0%
github.com/Babushkin05/software-dev-course/hw2/internal/application/services/feeding_organizer_service.go:49:	GetAllSchedules			0.0%
github.com/Babushkin05/software-dev-course/hw2/internal/application/services/zoo_admin_service.go:18:		NewZooAdminService		100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/application/services/zoo_admin_service.go:24:		CreateAnimal			100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/application/services/zoo_admin_service.go:28:		DeleteAnimal			100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/application/services/zoo_admin_service.go:32:		ListAnimals			100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/application/services/zoo_statistics_service.go:16:	NewZooStatisticsService		100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/application/services/zoo_statistics_service.go:20:	GetStatistics			100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/domain/animal.go:31:					NewAnimal			100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/domain/animal.go:48:					Feed				0.0%
github.com/Babushkin05/software-dev-course/hw2/internal/domain/animal.go:51:					Heal				100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/domain/animal.go:55:					MoveToEnclosure			100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/domain/enclosure.go:24:					NewEnclosure			100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/domain/enclosure.go:36:					AddAnimal			100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/domain/enclosure.go:44:					RemoveAnimal			100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/domain/enclosure.go:53:					Clean				0.0%
github.com/Babushkin05/software-dev-course/hw2/internal/domain/events.go:13:					EventName			100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/domain/events.go:23:					EventName			100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/domain/feeding_schedule.go:17:				NewFeedingSchedule		100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/domain/feeding_schedule.go:34:				MarkCompleted			100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/domain/feeding_schedule.go:38:				Reschedule			100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/domain/value_objects.go:5:				NewGender			100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/domain/value_objects.go:16:				NewHealthStatus			100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/infrastructure/events/event_dispatcher.go:15:		NewConsoleEventPublisher	100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/infrastructure/events/event_dispatcher.go:19:		Publish				100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/infrastructure/storage/animal_repo.go:14:		NewInMemoryAnimalRepo		100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/infrastructure/storage/animal_repo.go:20:		GetByID				100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/infrastructure/storage/animal_repo.go:30:		Save				100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/infrastructure/storage/animal_repo.go:37:		Delete				100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/infrastructure/storage/animal_repo.go:44:		List				100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/infrastructure/storage/enclosure_repo.go:14:		NewInMemoryEnclosureRepo	100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/infrastructure/storage/enclosure_repo.go:20:		GetByID				100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/infrastructure/storage/enclosure_repo.go:30:		Save				100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/infrastructure/storage/enclosure_repo.go:37:		Delete				100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/infrastructure/storage/enclosure_repo.go:44:		List				100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/infrastructure/storage/schedule_repo.go:14:		NewInMemoryFeedingScheduleRepo	100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/infrastructure/storage/schedule_repo.go:20:		Add				100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/infrastructure/storage/schedule_repo.go:27:		ListByAnimal			100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/infrastructure/storage/schedule_repo.go:39:		MarkDone			100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/infrastructure/storage/schedule_repo.go:50:		ListAll				100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/presentation/handlers/animal_handler.go:18:		NewAnimalHandler		100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/presentation/handlers/animal_handler.go:30:		RegisterRoutes			0.0%
github.com/Babushkin05/software-dev-course/hw2/internal/presentation/handlers/animal_handler.go:48:		CreateAnimal			88.2%
github.com/Babushkin05/software-dev-course/hw2/internal/presentation/handlers/animal_handler.go:92:		MoveAnimal			100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/presentation/handlers/animal_handler.go:117:		ListAnimals			100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/presentation/handlers/animal_handler.go:135:		DeleteAnimal			100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/presentation/handlers/enclosure_handler.go:16:		NewEnclosureHandler		100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/presentation/handlers/enclosure_handler.go:26:		RegisterRoutes			0.0%
github.com/Babushkin05/software-dev-course/hw2/internal/presentation/handlers/enclosure_handler.go:43:		Create				85.7%
github.com/Babushkin05/software-dev-course/hw2/internal/presentation/handlers/enclosure_handler.go:74:		List				100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/presentation/handlers/enclosure_handler.go:92:		Delete				100.0%
github.com/Babushkin05/software-dev-course/hw2/internal/presentation/handlers/schedule_handler.go:15:		NewFeedingScheduleHandler	0.0%
github.com/Babushkin05/software-dev-course/hw2/internal/presentation/handlers/schedule_handler.go:25:		RegisterRoutes			0.0%
github.com/Babushkin05/software-dev-course/hw2/internal/presentation/handlers/schedule_handler.go:40:		CreateFeeding			0.0%
github.com/Babushkin05/software-dev-course/hw2/internal/presentation/handlers/schedule_handler.go:64:		ListFeedings			0.0%
github.com/Babushkin05/software-dev-course/hw2/internal/presentation/handlers/statistics_handler.go:14:		RegisterRoutes			0.0%
github.com/Babushkin05/software-dev-course/hw2/internal/presentation/handlers/statistics_handler.go:26:		GetStatistics			0.0%
github.com/Babushkin05/software-dev-course/hw2/internal/presentation/router.go:13:				SetupRouter			0.0%
total:														(statements)			74.9%
```


