package domain

import (
	"github.com/google/uuid"
)

type EnclosureType string

const (
	Predator  EnclosureType = "predator"
	Herbivore EnclosureType = "herbivore"
	Aquarium  EnclosureType = "aquarium"
	BirdCage  EnclosureType = "bird"
)

type Enclosure struct {
	ID        string
	Type      EnclosureType
	Size      int
	Capacity  int
	AnimalIDs []string
}

func NewEnclosure(enclosureType EnclosureType, size, capacity int) (*Enclosure, error) {
	if capacity <= 0 {
		return nil, ErrCapacityMustBePositive
	}
	return &Enclosure{
		ID:       uuid.New().String(),
		Type:     enclosureType,
		Size:     size,
		Capacity: capacity,
	}, nil
}

func (e *Enclosure) AddAnimal(animalID string) error {
	if len(e.AnimalIDs) >= e.Capacity {
		return ErrEnclosureFull
	}
	e.AnimalIDs = append(e.AnimalIDs, animalID)
	return nil
}

func (e *Enclosure) RemoveAnimal(animalID string) {
	for i, id := range e.AnimalIDs {
		if id == animalID {
			e.AnimalIDs = append(e.AnimalIDs[:i], e.AnimalIDs[i+1:]...)
			break
		}
	}
}

func (e *Enclosure) Clean() {
}
