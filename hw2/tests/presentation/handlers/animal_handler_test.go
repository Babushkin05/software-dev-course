package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Babushkin05/software-dev-course/hw2/internal/domain"
	"github.com/Babushkin05/software-dev-course/hw2/internal/presentation/handlers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockAnimalTransferService реализует интерфейс AnimalTransferService
type MockAnimalTransferService struct {
	mock.Mock
}

func (m *MockAnimalTransferService) MoveAnimal(animalID, toEnclosureID string) error {
	args := m.Called(animalID, toEnclosureID)
	return args.Error(0)
}

// MockZooAdminService реализует интерфейс ZooAdminService
type MockZooAdminService struct {
	mock.Mock
}

func (m *MockZooAdminService) CreateAnimal(animal *domain.Animal) error {
	args := m.Called(animal)
	return args.Error(0)
}

func (m *MockZooAdminService) DeleteAnimal(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockZooAdminService) ListAnimals() ([]*domain.Animal, error) {
	args := m.Called()
	return args.Get(0).([]*domain.Animal), args.Error(1)
}

func TestAnimalHandler_CreateAnimal(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    interface{}
		mockError      error
		expectedStatus int
	}{
		{
			"Success",
			handlers.CreateAnimalRequest{
				Name:         "Simba",
				Species:      "Lion",
				BirthDate:    time.Now(),
				Gender:       "male",
				FavoriteFood: "Meat",
			},
			nil,
			http.StatusCreated,
		},
		{
			"InvalidGender",
			handlers.CreateAnimalRequest{
				Name:         "Simba",
				Species:      "Lion",
				BirthDate:    time.Now(),
				Gender:       "invalid",
				FavoriteFood: "Meat",
			},
			nil,
			http.StatusBadRequest,
		},
		{
			"InvalidRequest",
			"invalid",
			nil,
			http.StatusBadRequest,
		},
		{
			"ServiceError",
			handlers.CreateAnimalRequest{
				Name:         "Simba",
				Species:      "Lion",
				BirthDate:    time.Now(),
				Gender:       "male",
				FavoriteFood: "Meat",
			},
			domain.ErrNameAndSpeciesRequired,
			http.StatusBadRequest,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup
			transferService := new(MockAnimalTransferService)
			adminService := new(MockZooAdminService)
			handler := handlers.NewAnimalHandler(transferService, adminService)

			// Configure mock if needed
			if tc.mockError != nil || tc.expectedStatus == http.StatusCreated {
				adminService.On("CreateAnimal", mock.AnythingOfType("*domain.Animal")).Return(tc.mockError)
			}

			// Create test context
			router := gin.Default()
			router.POST("/animals", handler.CreateAnimal)

			// Prepare request
			body, _ := json.Marshal(tc.requestBody)
			req, _ := http.NewRequest("POST", "/animals", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			// Execute
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			// Verify
			assert.Equal(t, tc.expectedStatus, w.Code)
			if tc.expectedStatus == http.StatusCreated {
				var response domain.Animal
				json.Unmarshal(w.Body.Bytes(), &response)
				assert.Equal(t, tc.requestBody.(handlers.CreateAnimalRequest).Name, response.Name)
			}

			adminService.AssertExpectations(t)
		})
	}
}

func TestAnimalHandler_MoveAnimal(t *testing.T) {
	tests := []struct {
		name           string
		animalID       string
		requestBody    interface{}
		mockError      error
		expectedStatus int
	}{
		{
			"Success",
			"a1",
			handlers.MoveAnimalRequest{ToEnclosure: "e2"},
			nil,
			http.StatusOK,
		},
		{
			"InvalidRequest",
			"a1",
			"invalid",
			nil,
			http.StatusBadRequest,
		},
		{
			"ServiceError",
			"a1",
			handlers.MoveAnimalRequest{ToEnclosure: "e2"},
			domain.ErrAnimalNotFound,
			http.StatusBadRequest,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup
			transferService := new(MockAnimalTransferService)
			adminService := new(MockZooAdminService)
			handler := handlers.NewAnimalHandler(transferService, adminService)

			// Configure mock
			if tc.expectedStatus != http.StatusBadRequest || tc.mockError != nil {
				transferService.On("MoveAnimal", tc.animalID, mock.AnythingOfType("string")).Return(tc.mockError)
			}

			// Create test context
			router := gin.Default()
			router.POST("/animals/:id/move", handler.MoveAnimal)

			// Prepare request
			body, _ := json.Marshal(tc.requestBody)
			req, _ := http.NewRequest("POST", "/animals/"+tc.animalID+"/move", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			// Execute
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			// Verify
			assert.Equal(t, tc.expectedStatus, w.Code)
			transferService.AssertExpectations(t)
		})
	}
}

func TestAnimalHandler_ListAnimals(t *testing.T) {
	tests := []struct {
		name           string
		mockAnimals    []*domain.Animal
		mockError      error
		expectedStatus int
	}{
		{
			"Success",
			[]*domain.Animal{
				{ID: "a1", Name: "Simba"},
				{ID: "a2", Name: "Zoe"},
			},
			nil,
			http.StatusOK,
		},
		{
			"EmptyList",
			[]*domain.Animal{},
			nil,
			http.StatusOK,
		},
		{
			"ServiceError",
			nil,
			domain.ErrAnimalNotFound,
			http.StatusInternalServerError,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup
			transferService := new(MockAnimalTransferService)
			adminService := new(MockZooAdminService)
			handler := handlers.NewAnimalHandler(transferService, adminService)

			// Configure mock
			adminService.On("ListAnimals").Return(tc.mockAnimals, tc.mockError)

			// Create test context
			router := gin.Default()
			router.GET("/animals", handler.ListAnimals)

			// Prepare request
			req, _ := http.NewRequest("GET", "/animals", nil)

			// Execute
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			// Verify
			assert.Equal(t, tc.expectedStatus, w.Code)
			if tc.expectedStatus == http.StatusOK {
				var response []*domain.Animal
				json.Unmarshal(w.Body.Bytes(), &response)
				assert.Equal(t, len(tc.mockAnimals), len(response))
			}
			adminService.AssertExpectations(t)
		})
	}
}

func TestAnimalHandler_DeleteAnimal(t *testing.T) {
	tests := []struct {
		name           string
		animalID       string
		mockError      error
		expectedStatus int
	}{
		{
			"Success",
			"a1",
			nil,
			http.StatusOK,
		},
		{
			"NotFound",
			"a99",
			domain.ErrAnimalNotFound,
			http.StatusNotFound,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup
			transferService := new(MockAnimalTransferService)
			adminService := new(MockZooAdminService)
			handler := handlers.NewAnimalHandler(transferService, adminService)

			// Configure mock
			adminService.On("DeleteAnimal", tc.animalID).Return(tc.mockError)

			// Create test context
			router := gin.Default()
			router.DELETE("/animals/:id", handler.DeleteAnimal)

			// Prepare request
			req, _ := http.NewRequest("DELETE", "/animals/"+tc.animalID, nil)

			// Execute
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			// Verify
			assert.Equal(t, tc.expectedStatus, w.Code)
			adminService.AssertExpectations(t)
		})
	}
}

func NewAnimalHandler(transferService *MockAnimalTransferService, adminService *MockZooAdminService) *handlers.AnimalHandler {
	return &handlers.AnimalHandler{
		TransferService: transferService,
		AdminService:    adminService,
	}
}
