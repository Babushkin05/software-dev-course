package domain_test

import (
	"testing"
	"time"

	"github.com/Babushkin05/software-dev-course/hw2/internal/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewAnimal_Success(t *testing.T) {
	// Arrange
	name := "Симба"
	species := "Lion"
	birthDate := time.Date(2020, 1, 15, 0, 0, 0, 0, time.UTC)
	gender := domain.Male
	food := "Мясо"

	// Act
	animal, err := domain.NewAnimal(name, species, birthDate, gender, food)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, animal)
	assert.Equal(t, name, animal.Name)
	assert.Equal(t, species, animal.Species)
	assert.Equal(t, birthDate, animal.BirthDate)
	assert.Equal(t, gender, animal.Gender)
	assert.Equal(t, food, animal.FavoriteFood)
	assert.Equal(t, domain.Healthy, animal.Health)
	assert.NotEmpty(t, animal.ID)
	_, err = uuid.Parse(animal.ID)
	assert.NoError(t, err, "ID should be valid UUID")
}

func TestNewAnimal_ValidationErrors(t *testing.T) {
	tests := []struct {
		name     string
		species  string
		expected error
	}{
		{"", "Lion", domain.ErrNameAndSpeciesRequired},
		{"Симба", "", domain.ErrNameAndSpeciesRequired},
		{"", "", domain.ErrNameAndSpeciesRequired},
	}

	for _, tc := range tests {
		t.Run(tc.name+"_"+tc.species, func(t *testing.T) {
			// Arrange
			birthDate := time.Now()
			gender := domain.Male
			food := "Мясо"

			// Act
			animal, err := domain.NewAnimal(tc.name, tc.species, birthDate, gender, food)

			// Assert
			assert.Nil(t, animal)
			assert.ErrorIs(t, err, tc.expected)
		})
	}
}

func TestAnimal_Methods(t *testing.T) {
	// Arrange
	animal := &domain.Animal{
		ID:           uuid.New().String(),
		Name:         "Симба",
		Species:      "Lion",
		BirthDate:    time.Now(),
		Gender:       domain.Male,
		FavoriteFood: "Мясо",
		Health:       domain.Healthy,
	}

	t.Run("Feed", func(t *testing.T) {
		// Act
		animal.Feed()

		// Assert
		// Можно добавить логирование или другие проверки, если метод Feed будет расширен
	})

	t.Run("Heal", func(t *testing.T) {
		// Arrange
		animal.Health = domain.Sick

		// Act
		animal.Heal()

		// Assert
		assert.Equal(t, domain.Healthy, animal.Health)
	})

	t.Run("MoveToEnclosure", func(t *testing.T) {
		// Arrange
		enclosureID := "enclosure-123"

		// Act
		animal.MoveToEnclosure(enclosureID)

		// Assert
		assert.Equal(t, enclosureID, animal.CurrentEnclosure)
	})
}

func TestAnimal_HealthStatus(t *testing.T) {
	tests := []struct {
		status   domain.HealthStatus
		expected string
	}{
		{domain.Healthy, "healthy"},
		{domain.Sick, "sick"},
	}

	for _, tc := range tests {
		t.Run(string(tc.status), func(t *testing.T) {
			assert.Equal(t, tc.expected, string(tc.status))
		})
	}
}

func TestAnimal_Gender(t *testing.T) {
	tests := []struct {
		gender   domain.Gender
		expected string
	}{
		{domain.Male, "male"},
		{domain.Female, "female"},
	}

	for _, tc := range tests {
		t.Run(string(tc.gender), func(t *testing.T) {
			assert.Equal(t, tc.expected, string(tc.gender))
		})
	}
}
