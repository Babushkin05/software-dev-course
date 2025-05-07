package ports

import "github.com/Babushkin05/software-dev-course/hw2/internal/domain"

type AnimalRepository interface {
	GetByID(id string) (*domain.Animal, error)
	Save(animal *domain.Animal) error
	Delete(id string) error
	List() ([]*domain.Animal, error)
}
