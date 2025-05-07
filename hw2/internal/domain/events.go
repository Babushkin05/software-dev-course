package domain

type DomainEvent interface {
	EventName() string
}

type AnimalMovedEvent struct {
	AnimalID      string
	FromEnclosure string
	ToEnclosure   string
}

func (e AnimalMovedEvent) EventName() string {
	return "AnimalMoved"
}

type FeedingTimeEvent struct {
	AnimalID string
	Time     string
	Food     string
}

func (e FeedingTimeEvent) EventName() string {
	return "FeedingTime"
}
