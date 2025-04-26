package services

import (
	"github.com/Babushkin05/software-dev-course/hw2/internal/application/ports"
	"github.com/Babushkin05/software-dev-course/hw2/internal/domain"
)

type IAnimalTransferService interface {
	MoveAnimal(animalID, toEnclosureID string) error
}

type AnimalTransferService struct {
	animals    ports.AnimalRepository
	enclosures ports.EnclosureRepository
	publisher  ports.EventPublisher
}

func NewAnimalTransferService(a ports.AnimalRepository, e ports.EnclosureRepository, pub ports.EventPublisher) *AnimalTransferService {
	return &AnimalTransferService{a, e, pub}
}

func (s *AnimalTransferService) MoveAnimal(animalID, toEnclosureID string) error {
	animal, err := s.animals.GetByID(animalID)
	if err != nil {
		return err
	}

	fromEnclosureID := animal.CurrentEnclosure

	enclosure, err := s.enclosures.GetByID(toEnclosureID)
	if err != nil {
		return err
	}

	if !IsCompatible(animal.Species, enclosure.Type) {
		return domain.ErrInvalidEnclosureType
	}

	if len(enclosure.AnimalIDs) >= enclosure.Capacity {
		return domain.ErrEnclosureFull
	}

	enclosure.AddAnimal(animal.ID)
	animal.MoveToEnclosure(enclosure.ID)

	s.animals.Save(animal)
	s.enclosures.Save(enclosure)

	event := domain.AnimalMovedEvent{
		AnimalID:      animal.ID,
		FromEnclosure: fromEnclosureID,
		ToEnclosure:   enclosure.ID,
	}
	s.publisher.Publish(event)

	return nil
}
