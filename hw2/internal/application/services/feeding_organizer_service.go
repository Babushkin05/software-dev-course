package services

import (
	"time"

	"github.com/Babushkin05/software-dev-course/hw2/internal/application/ports"
	"github.com/Babushkin05/software-dev-course/hw2/internal/domain"
)

type IFeedingOrganizerService interface {
	AddFeeding(animalID string, t time.Time, food string) error
	NotifyFeedingDue()
	GetAllSchedules() ([]*domain.FeedingSchedule, error)
}

type FeedingOrganizerService struct {
	schedules ports.FeedingScheduleRepository
	publisher ports.EventPublisher
}

func NewFeedingOrganizerService(s ports.FeedingScheduleRepository, p ports.EventPublisher) *FeedingOrganizerService {
	return &FeedingOrganizerService{s, p}
}

func (f *FeedingOrganizerService) AddFeeding(animalID string, t time.Time, food string) error {
	schedule, err := domain.NewFeedingSchedule(animalID, t, food)
	if err != nil {
		return err
	}
	return f.schedules.Add(schedule)
}

func (f *FeedingOrganizerService) NotifyFeedingDue() {
	all, _ := f.schedules.ListAll()
	now := time.Now()

	for _, s := range all {
		if !s.IsDone && now.After(s.Time) {
			event := domain.FeedingTimeEvent{
				AnimalID: s.AnimalID,
				Time:     s.Time.Format(time.RFC3339),
				Food:     s.FoodType,
			}
			f.publisher.Publish(event)
		}
	}
}

func (s *FeedingOrganizerService) GetAllSchedules() ([]*domain.FeedingSchedule, error) {
	return s.schedules.ListAll()
}
