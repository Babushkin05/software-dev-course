package storage

import (
	"sync"

	"github.com/Babushkin05/software-dev-course/hw2/internal/domain"
)

type InMemoryFeedingScheduleRepo struct {
	mu        sync.RWMutex
	schedules map[string]*domain.FeedingSchedule
}

func NewInMemoryFeedingScheduleRepo() *InMemoryFeedingScheduleRepo {
	return &InMemoryFeedingScheduleRepo{
		schedules: make(map[string]*domain.FeedingSchedule),
	}
}

func (r *InMemoryFeedingScheduleRepo) Add(schedule *domain.FeedingSchedule) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.schedules[schedule.ID] = schedule
	return nil
}

func (r *InMemoryFeedingScheduleRepo) ListByAnimal(animalID string) ([]*domain.FeedingSchedule, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var result []*domain.FeedingSchedule
	for _, s := range r.schedules {
		if s.AnimalID == animalID {
			result = append(result, s)
		}
	}
	return result, nil
}

func (r *InMemoryFeedingScheduleRepo) MarkDone(scheduleID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	s, ok := r.schedules[scheduleID]
	if !ok {
		return domain.ErrScheduleNotFound
	}
	s.MarkCompleted()
	return nil
}

func (r *InMemoryFeedingScheduleRepo) ListAll() ([]*domain.FeedingSchedule, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	result := make([]*domain.FeedingSchedule, 0, len(r.schedules))
	for _, s := range r.schedules {
		result = append(result, s)
	}
	return result, nil
}
