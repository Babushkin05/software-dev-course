package domain

import "errors"

var (
	ErrEnclosureFull           = errors.New("enclosure is full")
	ErrInvalidEnclosureType    = errors.New("invalid enclosure type for animal")
	ErrAnimalNotFound          = errors.New("animal not found")
	ErrFeedingAlreadyDone      = errors.New("feeding already marked as done")
	ErrNameAndSpeciesRequired  = errors.New("name and species are required")
	ErrCapacityMustBePositive  = errors.New("capacity must be positive")
	ErrAnimalIdFoodRequire     = errors.New("animal ID and food type are required")
	ErrFeedingTimeMustBeFuture = errors.New("feeding time must be in the future")
)
