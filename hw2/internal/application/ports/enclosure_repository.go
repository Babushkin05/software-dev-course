package ports

import "github.com/Babushkin05/software-dev-course/hw2/internal/domain"

type EnclosureRepository interface {
	GetByID(id string) (*domain.Enclosure, error)
	Save(enclosure *domain.Enclosure) error
	Delete(id string) error
	List() ([]*domain.Enclosure, error)
}
