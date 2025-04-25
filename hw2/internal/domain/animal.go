package domain

import "time"

type HealthStatus string
type Gender string

const (
	Healthy HealthStatus = "healthy"
	Sick    HealthStatus = "sick"

	Male   Gender = "male"
	Female Gender = "female"
)

type Animal struct {
	ID               string
	Name             string
	Species          string
	BirthDate        time.Time
	Gender           Gender
	FavoriteFood     string
	Health           HealthStatus
	CurrentEnclosure string
}

func (a *Animal) Feed() {
}

func (a *Animal) Heal() {
	a.Health = Healthy
}

func (a *Animal) MoveToEnclosure(enclosureID string) {
	a.CurrentEnclosure = enclosureID
}
