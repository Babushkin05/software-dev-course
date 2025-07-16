package domain_test

import (
	"errors"
	"testing"

	"github.com/Babushkin05/software-dev-course/hw2/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewGender(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    domain.Gender
		expectedErr error
	}{
		{
			"Male",
			"male",
			domain.Male,
			nil,
		},
		{
			"Female",
			"female",
			domain.Female,
			nil,
		},
		{
			"Empty",
			"",
			"",
			errors.New("invalid gender"),
		},
		{
			"Invalid",
			"other",
			"",
			errors.New("invalid gender"),
		},
		{
			"CaseSensitive",
			"Male",
			"",
			errors.New("invalid gender"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			result, err := domain.NewGender(tc.input)

			// Assert
			assert.Equal(t, tc.expected, result)
			if tc.expectedErr != nil {
				assert.EqualError(t, err, tc.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestNewHealthStatus(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    domain.HealthStatus
		expectedErr error
	}{
		{
			"Healthy",
			"healthy",
			domain.Healthy,
			nil,
		},
		{
			"Sick",
			"sick",
			domain.Sick,
			nil,
		},
		{
			"Empty",
			"",
			"",
			errors.New("invalid health status"),
		},
		{
			"Invalid",
			"other",
			"",
			errors.New("invalid health status"),
		},
		{
			"CaseSensitive",
			"Healthy",
			"",
			errors.New("invalid health status"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			result, err := domain.NewHealthStatus(tc.input)

			// Assert
			assert.Equal(t, tc.expected, result)
			if tc.expectedErr != nil {
				assert.EqualError(t, err, tc.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
