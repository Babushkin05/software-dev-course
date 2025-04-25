package services

import "github.com/Babushkin05/software-dev-course/hw2/internal/domain"

func IsCompatible(species string, enclosureType domain.EnclosureType) bool {
	switch species {
	case "lion", "tiger":
		return enclosureType == domain.Predator
	case "zebra", "deer":
		return enclosureType == domain.Herbivore
	case "parrot":
		return enclosureType == domain.BirdCage
	case "fish":
		return enclosureType == domain.Aquarium
	default:
		return false
	}
}
