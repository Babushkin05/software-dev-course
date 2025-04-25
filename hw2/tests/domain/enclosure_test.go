package domain_test

import (
	"testing"

	"github.com/Babushkin05/software-dev-course/hw2/internal/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewEnclosure_Success(t *testing.T) {
	tests := []struct {
		name     string
		encType  domain.EnclosureType
		size     int
		capacity int
	}{
		{"Predator", domain.Predator, 100, 5},
		{"Herbivore", domain.Herbivore, 200, 10},
		{"Aquarium", domain.Aquarium, 50, 20},
		{"BirdCage", domain.BirdCage, 30, 15},
	}

	for _, tc := range tests {
		t.Run(string(tc.encType), func(t *testing.T) {
			// Act
			enc, err := domain.NewEnclosure(tc.encType, tc.size, tc.capacity)

			// Assert
			assert.NoError(t, err)
			assert.NotNil(t, enc)
			assert.Equal(t, tc.encType, enc.Type)
			assert.Equal(t, tc.size, enc.Size)
			assert.Equal(t, tc.capacity, enc.Capacity)
			assert.Empty(t, enc.AnimalIDs)
			assert.NotEmpty(t, enc.ID)
			_, err = uuid.Parse(enc.ID)
			assert.NoError(t, err, "ID should be valid UUID")
		})
	}
}

func TestNewEnclosure_InvalidCapacity(t *testing.T) {
	// Arrange
	encType := domain.Predator
	size := 100
	invalidCapacities := []int{0, -1, -100}

	for _, cap := range invalidCapacities {
		t.Run(string(rune(cap)), func(t *testing.T) {
			// Act
			enc, err := domain.NewEnclosure(encType, size, cap)

			// Assert
			assert.Nil(t, enc)
			assert.ErrorIs(t, err, domain.ErrCapacityMustBePositive)
		})
	}
}

func TestEnclosure_AddAnimal(t *testing.T) {
	t.Run("AddAnimal_Success", func(t *testing.T) {
		// Arrange
		enc, _ := domain.NewEnclosure(domain.Predator, 100, 2)
		animalID1 := "animal-1"
		animalID2 := "animal-2"

		// Act & Assert
		err := enc.AddAnimal(animalID1)
		assert.NoError(t, err)
		assert.Contains(t, enc.AnimalIDs, animalID1)
		assert.Len(t, enc.AnimalIDs, 1)

		err = enc.AddAnimal(animalID2)
		assert.NoError(t, err)
		assert.Contains(t, enc.AnimalIDs, animalID2)
		assert.Len(t, enc.AnimalIDs, 2)
	})

	t.Run("AddAnimal_EnclosureFull", func(t *testing.T) {
		// Arrange
		enc, _ := domain.NewEnclosure(domain.Predator, 100, 1)
		animalID1 := "animal-1"
		animalID2 := "animal-2"

		// Add first animal
		err := enc.AddAnimal(animalID1)
		assert.NoError(t, err)

		// Try to add second animal
		err = enc.AddAnimal(animalID2)
		assert.ErrorIs(t, err, domain.ErrEnclosureFull)
		assert.NotContains(t, enc.AnimalIDs, animalID2)
		assert.Len(t, enc.AnimalIDs, 1)
	})
}

func TestEnclosure_RemoveAnimal(t *testing.T) {
	tests := []struct {
		name           string
		initialAnimals []string
		removeID       string
		expectedLeft   int
	}{
		{"RemoveOnly", []string{"animal-1"}, "animal-1", 0},
		{"RemoveFirst", []string{"animal-1", "animal-2"}, "animal-1", 1},
		{"RemoveLast", []string{"animal-1", "animal-2"}, "animal-2", 1},
		{"RemoveMiddle", []string{"animal-1", "animal-2", "animal-3"}, "animal-2", 2},
		{"RemoveNonExistent", []string{"animal-1"}, "animal-99", 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			enc, _ := domain.NewEnclosure(domain.Predator, 100, 10)
			enc.AnimalIDs = tc.initialAnimals

			// Act
			enc.RemoveAnimal(tc.removeID)

			// Assert
			assert.Len(t, enc.AnimalIDs, tc.expectedLeft)
			assert.NotContains(t, enc.AnimalIDs, tc.removeID)
		})
	}
}

func TestEnclosureType_Values(t *testing.T) {
	tests := []struct {
		encType  domain.EnclosureType
		expected string
	}{
		{domain.Predator, "predator"},
		{domain.Herbivore, "herbivore"},
		{domain.Aquarium, "aquarium"},
		{domain.BirdCage, "bird"},
	}

	for _, tc := range tests {
		t.Run(string(tc.encType), func(t *testing.T) {
			assert.Equal(t, tc.expected, string(tc.encType))
		})
	}
}
