package domain

import (
	"time"

	"github.com/google/uuid"
)

type HealthStatus string
type Gender string

const (
	Healthy HealthStatus = "healthy"
	Sick    HealthStatus = "sick"

	Male   Gender = "male"
	Female Gender = "female"
)

type Animal struct {
	ID               string       `json:"id" example:"a1b2c3d4"`
	Name             string       `json:"name" example:"Симба"`
	Species          string       `json:"species" example:"Lion"`
	BirthDate        time.Time    `json:"birth_date" example:"2020-01-15T00:00:00Z"`
	Gender           Gender       `json:"gender" example:"male"`
	FavoriteFood     string       `json:"favorite_food" example:"Мясо"`
	Health           HealthStatus `json:"health" example:"healty"`
	CurrentEnclosure string       `json:"current_enclosure,omitempty" example:"e1f2g3h4"`
}

func NewAnimal(name, species string, birthDate time.Time, gender Gender, favoriteFood string) (*Animal, error) {
	if name == "" || species == "" {
		return nil, ErrNameAndSpeciesRequired
	}

	animal := &Animal{
		ID:           uuid.New().String(),
		Name:         name,
		Species:      species,
		BirthDate:    birthDate,
		Gender:       gender,
		FavoriteFood: favoriteFood,
		Health:       Healthy,
	}
	return animal, nil
}

func (a *Animal) Feed() {
}

func (a *Animal) Heal() {
	a.Health = Healthy
}

func (a *Animal) MoveToEnclosure(enclosureID string) {
	a.CurrentEnclosure = enclosureID
}
