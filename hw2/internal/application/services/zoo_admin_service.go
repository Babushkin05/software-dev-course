package services

import (
	"github.com/Babushkin05/software-dev-course/hw2/internal/application/ports"
	"github.com/Babushkin05/software-dev-course/hw2/internal/domain"
)

type IZooAdminService interface {
	CreateAnimal(animal *domain.Animal) error
	DeleteAnimal(id string) error
	ListAnimals() ([]*domain.Animal, error)
}

type ZooAdminService struct {
	animalRepo ports.AnimalRepository
}

func NewZooAdminService(repo ports.AnimalRepository) *ZooAdminService {
	return &ZooAdminService{
		animalRepo: repo,
	}
}

func (s *ZooAdminService) CreateAnimal(animal *domain.Animal) error {
	return s.animalRepo.Save(animal)
}

func (s *ZooAdminService) DeleteAnimal(id string) error {
	return s.animalRepo.Delete(id)
}

func (s *ZooAdminService) ListAnimals() ([]*domain.Animal, error) {
	return s.animalRepo.List()
}
