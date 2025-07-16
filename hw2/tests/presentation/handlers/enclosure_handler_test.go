package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Babushkin05/software-dev-course/hw2/internal/domain"
	"github.com/Babushkin05/software-dev-course/hw2/internal/presentation/handlers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockEnclosureRepository реализует ports.EnclosureRepository
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

func TestEnclosureHandler_Create(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    interface{}
		mockError      error
		expectedStatus int
	}{
		{
			"Success",
			handlers.CreateEnclosureRequest{
				Type:     "predator",
				Size:     100,
				Capacity: 5,
			},
			nil,
			http.StatusCreated,
		},
		{
			"InvalidType",
			handlers.CreateEnclosureRequest{
				Type:     "invalid",
				Size:     100,
				Capacity: 5,
			},
			nil,
			http.StatusInternalServerError,
		},
		{
			"InvalidRequest",
			"invalid",
			nil,
			http.StatusBadRequest,
		},
		{
			"RepositoryError",
			handlers.CreateEnclosureRequest{
				Type:     "predator",
				Size:     100,
				Capacity: 5,
			},
			domain.ErrCapacityMustBePositive,
			http.StatusBadRequest,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup
			repo := new(MockEnclosureRepository)
			handler := handlers.NewEnclosureHandler(repo)

			// Configure mock if needed
			if tc.mockError != nil || tc.expectedStatus == http.StatusCreated {
				repo.On("Save", mock.AnythingOfType("*domain.Enclosure")).Return(tc.mockError)
			}

			// Create test context
			router := gin.Default()
			router.POST("/enclosures", handler.Create)

			// Prepare request
			body, _ := json.Marshal(tc.requestBody)
			req, _ := http.NewRequest("POST", "/enclosures", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			// Execute
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			// Verify
			assert.Equal(t, tc.expectedStatus, w.Code)
			if tc.expectedStatus == http.StatusCreated {
				var response domain.Enclosure
				json.Unmarshal(w.Body.Bytes(), &response)
				assert.Equal(t, domain.Predator, response.Type)
			}

			repo.AssertExpectations(t)
		})
	}
}

func TestEnclosureHandler_List(t *testing.T) {
	tests := []struct {
		name           string
		mockEnclosures []*domain.Enclosure
		mockError      error
		expectedStatus int
	}{
		{
			"Success",
			[]*domain.Enclosure{
				{ID: "e1", Type: domain.Predator},
				{ID: "e2", Type: domain.Herbivore},
			},
			nil,
			http.StatusOK,
		},
		{
			"EmptyList",
			[]*domain.Enclosure{},
			nil,
			http.StatusOK,
		},
		{
			"RepositoryError",
			nil,
			domain.ErrEnclosureNotFound,
			http.StatusBadRequest,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup
			repo := new(MockEnclosureRepository)
			handler := handlers.NewEnclosureHandler(repo)

			// Configure mock
			repo.On("List").Return(tc.mockEnclosures, tc.mockError)

			// Create test context
			router := gin.Default()
			router.GET("/enclosures", handler.List)

			// Prepare request
			req, _ := http.NewRequest("GET", "/enclosures", nil)

			// Execute
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			// Verify
			assert.Equal(t, tc.expectedStatus, w.Code)
			if tc.expectedStatus == http.StatusOK {
				var response []*domain.Enclosure
				json.Unmarshal(w.Body.Bytes(), &response)
				assert.Equal(t, len(tc.mockEnclosures), len(response))
			}
			repo.AssertExpectations(t)
		})
	}
}

func TestEnclosureHandler_Delete(t *testing.T) {
	tests := []struct {
		name           string
		enclosureID    string
		mockError      error
		expectedStatus int
	}{
		{
			"Success",
			"e1",
			nil,
			http.StatusOK,
		},
		{
			"NotFound",
			"e99",
			domain.ErrEnclosureNotFound,
			http.StatusNotFound,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup
			repo := new(MockEnclosureRepository)
			handler := handlers.NewEnclosureHandler(repo)

			// Configure mock
			repo.On("Delete", tc.enclosureID).Return(tc.mockError)

			// Create test context
			router := gin.Default()
			router.DELETE("/enclosures/:id", handler.Delete)

			// Prepare request
			req, _ := http.NewRequest("DELETE", "/enclosures/"+tc.enclosureID, nil)

			// Execute
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			// Verify
			assert.Equal(t, tc.expectedStatus, w.Code)
			repo.AssertExpectations(t)
		})
	}
}

func NewEnclosureHandler(repo *MockEnclosureRepository) *handlers.EnclosureHandler {
	return &handlers.EnclosureHandler{
		Repo: repo,
	}
}
