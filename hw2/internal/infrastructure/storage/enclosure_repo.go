package storage

import (
	"sync"

	"github.com/Babushkin05/software-dev-course/hw2/internal/domain"
)

type InMemoryEnclosureRepo struct {
	mu         sync.RWMutex
	enclosures map[string]*domain.Enclosure
}

func NewInMemoryEnclosureRepo() *InMemoryEnclosureRepo {
	return &InMemoryEnclosureRepo{
		enclosures: make(map[string]*domain.Enclosure),
	}
}

func (r *InMemoryEnclosureRepo) GetByID(id string) (*domain.Enclosure, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	enclosure, ok := r.enclosures[id]
	if !ok {
		return nil, domain.ErrEnclosureNotFound
	}
	return enclosure, nil
}

func (r *InMemoryEnclosureRepo) Save(enclosure *domain.Enclosure) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.enclosures[enclosure.ID] = enclosure
	return nil
}

func (r *InMemoryEnclosureRepo) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.enclosures, id)
	return nil
}

func (r *InMemoryEnclosureRepo) List() ([]*domain.Enclosure, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	list := make([]*domain.Enclosure, 0, len(r.enclosures))
	for _, e := range r.enclosures {
		list = append(list, e)
	}
	return list, nil
}
