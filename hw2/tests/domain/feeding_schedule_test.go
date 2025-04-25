package domain_test

import (
	"testing"
	"time"

	"github.com/Babushkin05/software-dev-course/hw2/internal/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewFeedingSchedule_Success(t *testing.T) {
	// Arrange
	animalID := "a1b2c3d4"
	foodType := "Мясо"
	futureTime := time.Now().Add(2 * time.Hour)

	// Act
	schedule, err := domain.NewFeedingSchedule(animalID, futureTime, foodType)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, schedule)
	assert.Equal(t, animalID, schedule.AnimalID)
	assert.Equal(t, foodType, schedule.FoodType)
	assert.Equal(t, futureTime, schedule.Time)
	assert.False(t, schedule.IsDone)
	assert.NotEmpty(t, schedule.ID)
	_, err = uuid.Parse(schedule.ID)
	assert.NoError(t, err, "ID should be valid UUID")
}

func TestNewFeedingSchedule_ValidationErrors(t *testing.T) {
	tests := []struct {
		name      string
		animalID  string
		foodType  string
		time      time.Time
		expectErr error
	}{
		{
			"EmptyAnimalID",
			"",
			"Мясо",
			time.Now().Add(1 * time.Hour),
			domain.ErrAnimalIdFoodRequire,
		},
		{
			"EmptyFoodType",
			"a1b2c3d4",
			"",
			time.Now().Add(1 * time.Hour),
			domain.ErrAnimalIdFoodRequire,
		},
		{
			"PastTime",
			"a1b2c3d4",
			"Мясо",
			time.Now().Add(-1 * time.Hour),
			domain.ErrFeedingTimeMustBeFuture,
		},
		{
			"AllInvalid",
			"",
			"",
			time.Now().Add(-1 * time.Hour),
			domain.ErrAnimalIdFoodRequire,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			schedule, err := domain.NewFeedingSchedule(tc.animalID, tc.time, tc.foodType)

			// Assert
			assert.Nil(t, schedule)
			assert.ErrorIs(t, err, tc.expectErr)
		})
	}
}

func TestFeedingSchedule_Methods(t *testing.T) {
	// Arrange
	animalID := "a1b2c3d4"
	foodType := "Мясо"
	futureTime := time.Now().Add(2 * time.Hour)
	schedule, _ := domain.NewFeedingSchedule(animalID, futureTime, foodType)

	t.Run("MarkCompleted", func(t *testing.T) {
		// Act
		schedule.MarkCompleted()

		// Assert
		assert.True(t, schedule.IsDone)
	})

	t.Run("Reschedule", func(t *testing.T) {
		// Arrange
		newTime := time.Now().Add(4 * time.Hour)
		schedule.IsDone = true // Помечаем как выполненное

		// Act
		schedule.Reschedule(newTime)

		// Assert
		assert.Equal(t, newTime, schedule.Time)
		assert.False(t, schedule.IsDone)
	})
}

func TestFeedingSchedule_EdgeCases(t *testing.T) {
	t.Run("RescheduleToPast", func(t *testing.T) {
		// Arrange
		animalID := "a1b2c3d4"
		foodType := "Мясо"
		futureTime := time.Now().Add(2 * time.Hour)
		schedule, _ := domain.NewFeedingSchedule(animalID, futureTime, foodType)
		pastTime := time.Now().Add(-1 * time.Hour)

		// Act
		schedule.Reschedule(pastTime)

		// Assert
		assert.Equal(t, pastTime, schedule.Time) // Reschedule позволяет установить любое время
		assert.False(t, schedule.IsDone)
	})

	t.Run("MultipleMarkCompleted", func(t *testing.T) {
		// Arrange
		animalID := "a1b2c3d4"
		foodType := "Мясо"
		futureTime := time.Now().Add(2 * time.Hour)
		schedule, _ := domain.NewFeedingSchedule(animalID, futureTime, foodType)

		// Act
		schedule.MarkCompleted()
		schedule.MarkCompleted() // Повторный вызов

		// Assert
		assert.True(t, schedule.IsDone) // Остается true
	})
}
