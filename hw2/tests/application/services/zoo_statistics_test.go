package services_test

import (
	"testing"

	"github.com/Babushkin05/software-dev-course/hw2/internal/application/services"
	"github.com/Babushkin05/software-dev-course/hw2/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestZooStatisticsService_GetStatistics(t *testing.T) {
	tests := []struct {
		name           string
		mockAnimals    []*domain.Animal
		mockEnclosures []*domain.Enclosure
		expectedStats  *services.ZooStatistics
	}{
		{
			"EmptyZoo",
			[]*domain.Animal{},
			[]*domain.Enclosure{},
			&services.ZooStatistics{
				TotalAnimals:    0,
				FreeEnclosures:  0,
				TotalEnclosures: 0,
			},
		},
		{
			"WithAnimalsAndEnclosures",
			[]*domain.Animal{
				{ID: "1", Name: "Simba"},
				{ID: "2", Name: "Zoe"},
			},
			[]*domain.Enclosure{
				{ID: "e1", Capacity: 2, AnimalIDs: []string{"1"}},
				{ID: "e2", Capacity: 1, AnimalIDs: []string{"2"}},
				{ID: "e3", Capacity: 3, AnimalIDs: []string{}},
			},
			&services.ZooStatistics{
				TotalAnimals:    2,
				FreeEnclosures:  2, // e1 (1/2), e3 (0/3)
				TotalEnclosures: 3,
			},
		},
		{
			"FullEnclosures",
			[]*domain.Animal{
				{ID: "1"}, {ID: "2"}, {ID: "3"},
			},
			[]*domain.Enclosure{
				{ID: "e1", Capacity: 2, AnimalIDs: []string{"1", "2"}},
				{ID: "e2", Capacity: 1, AnimalIDs: []string{"3"}},
			},
			&services.ZooStatistics{
				TotalAnimals:    3,
				FreeEnclosures:  0,
				TotalEnclosures: 2,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			animalRepo := new(MockAnimalRepository)
			enclosureRepo := new(MockEnclosureRepository)

			animalRepo.On("List").Return(tc.mockAnimals, nil)
			enclosureRepo.On("List").Return(tc.mockEnclosures, nil)

			service := services.NewZooStatisticsService(animalRepo, enclosureRepo)

			// Act
			stats, err := service.GetStatistics()

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedStats, stats)
			animalRepo.AssertExpectations(t)
			enclosureRepo.AssertExpectations(t)
		})
	}
}
