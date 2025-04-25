package domain

import (
	"time"

	"github.com/google/uuid"
)

type FeedingSchedule struct {
	ID       string
	AnimalID string
	Time     time.Time
	FoodType string
	IsDone   bool
}

func NewFeedingSchedule(animalID string, t time.Time, foodType string) (*FeedingSchedule, error) {
	if animalID == "" || foodType == "" {
		return nil, ErrAnimalIdFoodRequire
	}
	if t.Before(time.Now()) {
		return nil, ErrFeedingTimeMustBeFuture
	}

	return &FeedingSchedule{
		ID:       uuid.New().String(),
		AnimalID: animalID,
		Time:     t,
		FoodType: foodType,
		IsDone:   false,
	}, nil
}

func (f *FeedingSchedule) MarkCompleted() {
	f.IsDone = true
}

func (f *FeedingSchedule) Reschedule(newTime time.Time) {
	f.Time = newTime
	f.IsDone = false
}
