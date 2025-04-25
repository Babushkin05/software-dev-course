package ports

import "github.com/Babushkin05/software-dev-course/hw2/internal/domain"

type FeedingScheduleRepository interface {
	Add(schedule *domain.FeedingSchedule) error
	ListByAnimal(animalID string) ([]*domain.FeedingSchedule, error)
	MarkDone(scheduleID string) error
	ListAll() ([]*domain.FeedingSchedule, error)
}
