package domain

import "errors"

func NewGender(value string) (Gender, error) {
	switch value {
	case "male":
		return Male, nil
	case "female":
		return Female, nil
	default:
		return "", errors.New("invalid gender")
	}
}

func NewHealthStatus(value string) (HealthStatus, error) {
	switch value {
	case "healthy":
		return Healthy, nil
	case "sick":
		return Sick, nil
	default:
		return "", errors.New("invalid health status")
	}
}
