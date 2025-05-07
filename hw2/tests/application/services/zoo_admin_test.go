package services_test

import (
	"errors"
	"testing"

	"github.com/Babushkin05/software-dev-course/hw2/internal/application/services"
	"github.com/Babushkin05/software-dev-course/hw2/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestZooAdminService_CreateAnimal(t *testing.T) {
	tests := []struct {
		name          string
		animal        *domain.Animal
		mockError     error
		expectedError error
	}{
		{
			"Success",
			&domain.Animal{ID: "1", Name: "Simba", Species: "Lion"},
			nil,
			nil,
		},
		{
			"RepositoryError",
			&domain.Animal{ID: "1", Name: "Simba", Species: "Lion"},
			errors.New("save error"),
			errors.New("save error"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			repo := new(MockAnimalRepository)
			repo.On("Save", tc.animal).Return(tc.mockError)

			service := services.NewZooAdminService(repo)

			// Act
			err := service.CreateAnimal(tc.animal)

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

func TestZooAdminService_DeleteAnimal(t *testing.T) {
	tests := []struct {
		name          string
		animalID      string
		mockError     error
		expectedError error
	}{
		{
			"Success",
			"1",
			nil,
			nil,
		},
		{
			"RepositoryError",
			"1",
			errors.New("delete error"),
			errors.New("delete error"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			repo := new(MockAnimalRepository)
			repo.On("Delete", tc.animalID).Return(tc.mockError)

			service := services.NewZooAdminService(repo)

			// Act
			err := service.DeleteAnimal(tc.animalID)

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

func TestZooAdminService_ListAnimals(t *testing.T) {
	tests := []struct {
		name           string
		mockAnimals    []*domain.Animal
		mockError      error
		expectedResult []*domain.Animal
		expectedError  error
	}{
		{
			"Success",
			[]*domain.Animal{
				{ID: "1", Name: "Simba", Species: "Lion"},
				{ID: "2", Name: "Zoe", Species: "Zebra"},
			},
			nil,
			[]*domain.Animal{
				{ID: "1", Name: "Simba", Species: "Lion"},
				{ID: "2", Name: "Zoe", Species: "Zebra"},
			},
			nil,
		},
		{
			"EmptyList",
			[]*domain.Animal{},
			nil,
			[]*domain.Animal{},
			nil,
		},
		{
			"RepositoryError",
			nil,
			errors.New("list error"),
			nil,
			errors.New("list error"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			repo := new(MockAnimalRepository)
			repo.On("List").Return(tc.mockAnimals, tc.mockError)

			service := services.NewZooAdminService(repo)

			// Act
			result, err := service.ListAnimals()

			// Assert
			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedResult, result)
			}
			repo.AssertExpectations(t)
		})
	}
}
