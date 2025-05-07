package services_test

import (
	"errors"
	"testing"

	"github.com/Babushkin05/software-dev-course/hw2/internal/application/services"
	"github.com/Babushkin05/software-dev-course/hw2/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Полный мок AnimalRepository с всеми методами интерфейса
type MockAnimalRepository struct {
	mock.Mock
}

func (m *MockAnimalRepository) GetByID(id string) (*domain.Animal, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Animal), args.Error(1)
}

func (m *MockAnimalRepository) Save(animal *domain.Animal) error {
	args := m.Called(animal)
	return args.Error(0)
}

func (m *MockAnimalRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockAnimalRepository) List() ([]*domain.Animal, error) {
	args := m.Called()
	return args.Get(0).([]*domain.Animal), args.Error(1)
}

// Полный мок EnclosureRepository с всеми методами интерфейса
type MockEnclosureRepository struct {
	mock.Mock
}

func (m *MockEnclosureRepository) GetByID(id string) (*domain.Enclosure, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Enclosure), args.Error(1)
}

func (m *MockEnclosureRepository) Save(enclosure *domain.Enclosure) error {
	args := m.Called(enclosure)
	return args.Error(0)
}

func (m *MockEnclosureRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockEnclosureRepository) List() ([]*domain.Enclosure, error) {
	args := m.Called()
	return args.Get(0).([]*domain.Enclosure), args.Error(1)
}

// Мок EventPublisher
type MockEventPublisher struct {
	mock.Mock
}

func (m *MockEventPublisher) Publish(event domain.DomainEvent) {
	m.Called(event)
}

func TestAnimalTransferService_MoveAnimal(t *testing.T) {
	// Общие тестовые данные
	animalID := "animal-1"
	fromEnclosureID := "enc-old"
	toEnclosureID := "enc-new"

	compatibleAnimal := &domain.Animal{
		ID:               animalID,
		Species:          "lion",
		CurrentEnclosure: fromEnclosureID,
	}

	compatibleEnclosure := &domain.Enclosure{
		ID:        toEnclosureID,
		Type:      domain.Predator,
		Capacity:  2,
		AnimalIDs: []string{},
	}

	tests := []struct {
		name          string
		setupMocks    func(*MockAnimalRepository, *MockEnclosureRepository, *MockEventPublisher)
		expectedError error
		expectedEvent bool
	}{
		{
			"Success",
			func(ar *MockAnimalRepository, er *MockEnclosureRepository, ep *MockEventPublisher) {
				ar.On("GetByID", animalID).Return(compatibleAnimal, nil)
				er.On("GetByID", toEnclosureID).Return(compatibleEnclosure, nil)
				ar.On("Save", compatibleAnimal).Return(nil)
				er.On("Save", compatibleEnclosure).Return(nil)
				ep.On("Publish", mock.AnythingOfType("domain.AnimalMovedEvent")).Once()
			},
			nil,
			true,
		},
		{
			"AnimalNotFound",
			func(ar *MockAnimalRepository, er *MockEnclosureRepository, ep *MockEventPublisher) {
				ar.On("GetByID", animalID).Return(&domain.Animal{}, errors.New("not found"))
			},
			errors.New("not found"),
			false,
		},
		{
			"EnclosureNotFound",
			func(ar *MockAnimalRepository, er *MockEnclosureRepository, ep *MockEventPublisher) {
				ar.On("GetByID", animalID).Return(compatibleAnimal, nil)
				er.On("GetByID", toEnclosureID).Return(&domain.Enclosure{}, errors.New("not found"))
			},
			errors.New("not found"),
			false,
		},
		{
			"IncompatibleEnclosure",
			func(ar *MockAnimalRepository, er *MockEnclosureRepository, ep *MockEventPublisher) {
				incompatibleEnclosure := &domain.Enclosure{
					ID:   toEnclosureID,
					Type: domain.Herbivore,
				}
				ar.On("GetByID", animalID).Return(compatibleAnimal, nil)
				er.On("GetByID", toEnclosureID).Return(incompatibleEnclosure, nil)
			},
			domain.ErrInvalidEnclosureType,
			false,
		},
		{
			"EnclosureFull",
			func(ar *MockAnimalRepository, er *MockEnclosureRepository, ep *MockEventPublisher) {
				fullEnclosure := &domain.Enclosure{
					ID:        toEnclosureID,
					Type:      domain.Predator,
					Capacity:  1,
					AnimalIDs: []string{"existing-animal"},
				}
				ar.On("GetByID", animalID).Return(compatibleAnimal, nil)
				er.On("GetByID", toEnclosureID).Return(fullEnclosure, nil)
			},
			domain.ErrEnclosureFull,
			false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			animalRepo := new(MockAnimalRepository)
			enclosureRepo := new(MockEnclosureRepository)
			eventPublisher := new(MockEventPublisher)
			tc.setupMocks(animalRepo, enclosureRepo, eventPublisher)

			service := services.NewAnimalTransferService(animalRepo, enclosureRepo, eventPublisher)

			// Act
			err := service.MoveAnimal(animalID, toEnclosureID)

			// Assert
			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}

			if tc.expectedEvent {
				eventPublisher.AssertCalled(t, "Publish", mock.AnythingOfType("domain.AnimalMovedEvent"))
			} else {
				eventPublisher.AssertNotCalled(t, "Publish", mock.Anything)
			}

			animalRepo.AssertExpectations(t)
			enclosureRepo.AssertExpectations(t)
			eventPublisher.AssertExpectations(t)
		})
	}
}

func TestListMethods(t *testing.T) {
	t.Run("AnimalRepository_List", func(t *testing.T) {
		repo := new(MockAnimalRepository)
		expected := []*domain.Animal{{ID: "1"}}
		repo.On("List").Return(expected, nil)

		result, err := repo.List()

		assert.NoError(t, err)
		assert.Equal(t, expected, result)
		repo.AssertExpectations(t)
	})

	t.Run("EnclosureRepository_List", func(t *testing.T) {
		repo := new(MockEnclosureRepository)
		expected := []*domain.Enclosure{{ID: "1"}}
		repo.On("List").Return(expected, nil)

		result, err := repo.List()

		assert.NoError(t, err)
		assert.Equal(t, expected, result)
		repo.AssertExpectations(t)
	})
}
