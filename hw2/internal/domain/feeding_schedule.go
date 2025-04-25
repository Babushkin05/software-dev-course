package domain

import (
	"time"

	"github.com/google/uuid"
)

type FeedingSchedule struct {
	ID       string    `json:"id" example:"f1g2h3i4"`
	AnimalID string    `json:"animal_id" example:"a1b2c3d4"`
	Time     time.Time `json:"time" example:"2023-05-15T14:30:00Z"`
	FoodType string    `json:"food_type" example:"Мясо"`
	IsDone   bool      `json:"is_done" example:"false"`
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
