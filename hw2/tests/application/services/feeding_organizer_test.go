package services_test

import (
	"errors"
	"testing"
	"time"

	"github.com/Babushkin05/software-dev-course/hw2/internal/application/services"
	"github.com/Babushkin05/software-dev-course/hw2/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Полный мок FeedingScheduleRepository
type MockFeedingScheduleRepository struct {
	mock.Mock
}

func (m *MockFeedingScheduleRepository) Add(schedule *domain.FeedingSchedule) error {
	args := m.Called(schedule)
	return args.Error(0)
}

func (m *MockFeedingScheduleRepository) ListByAnimal(animalID string) ([]*domain.FeedingSchedule, error) {
	args := m.Called(animalID)
	return args.Get(0).([]*domain.FeedingSchedule), args.Error(1)
}

func (m *MockFeedingScheduleRepository) MarkDone(scheduleID string) error {
	args := m.Called(scheduleID)
	return args.Error(0)
}

func (m *MockFeedingScheduleRepository) ListAll() ([]*domain.FeedingSchedule, error) {
	args := m.Called()
	return args.Get(0).([]*domain.FeedingSchedule), args.Error(1)
}

func TestFeedingOrganizerService_AddFeeding(t *testing.T) {
	tests := []struct {
		name          string
		animalID      string
		time          time.Time
		food          string
		setupMock     func(*MockFeedingScheduleRepository)
		expectedError error
	}{
		{
			"Success",
			"animal-1",
			time.Now().Add(1 * time.Hour),
			"Meat",
			func(m *MockFeedingScheduleRepository) {
				m.On("Add", mock.AnythingOfType("*domain.FeedingSchedule")).Return(nil)
			},
			nil,
		},
		{
			"RepositoryError",
			"animal-1",
			time.Now().Add(1 * time.Hour),
			"Meat",
			func(m *MockFeedingScheduleRepository) {
				m.On("Add", mock.AnythingOfType("*domain.FeedingSchedule")).Return(errors.New("repository error"))
			},
			errors.New("repository error"),
		},
		{
			"InvalidTime",
			"animal-1",
			time.Now().Add(-1 * time.Hour),
			"Meat",
			func(m *MockFeedingScheduleRepository) {
				// Не ожидаем вызова Add, так как будет ошибка валидации
			},
			domain.ErrFeedingTimeMustBeFuture,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			repo := new(MockFeedingScheduleRepository)
			publisher := new(MockEventPublisher)
			tc.setupMock(repo)

			service := services.NewFeedingOrganizerService(repo, publisher)

			// Act
			err := service.AddFeeding(tc.animalID, tc.time, tc.food)

			// Assert
			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}

			repo.AssertExpectations(t)
		})
	}
}
