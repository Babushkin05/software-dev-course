package storage

import (
	"sync"

	"github.com/Babushkin05/software-dev-course/hw2/internal/domain"
)

type InMemoryAnimalRepo struct {
	mu      sync.RWMutex
	animals map[string]*domain.Animal
}

func NewInMemoryAnimalRepo() *InMemoryAnimalRepo {
	return &InMemoryAnimalRepo{
		animals: make(map[string]*domain.Animal),
	}
}

func (r *InMemoryAnimalRepo) GetByID(id string) (*domain.Animal, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	animal, ok := r.animals[id]
	if !ok {
		return nil, domain.ErrAnimalNotFound
	}
	return animal, nil
}

func (r *InMemoryAnimalRepo) Save(animal *domain.Animal) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.animals[animal.ID] = animal
	return nil
}

func (r *InMemoryAnimalRepo) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.animals, id)
	return nil
}

func (r *InMemoryAnimalRepo) List() ([]*domain.Animal, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	list := make([]*domain.Animal, 0, len(r.animals))
	for _, a := range r.animals {
		list = append(list, a)
	}
	return list, nil
}
