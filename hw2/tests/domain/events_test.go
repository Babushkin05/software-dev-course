package domain_test

import (
	"testing"

	"github.com/Babushkin05/software-dev-course/hw2/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestDomainEvents(t *testing.T) {
	t.Run("AnimalMovedEvent", func(t *testing.T) {
		// Arrange
		event := domain.AnimalMovedEvent{
			AnimalID:      "animal-123",
			FromEnclosure: "enc-old",
			ToEnclosure:   "enc-new",
		}

		// Act
		eventName := event.EventName()
		var domainEvent domain.DomainEvent = event // Проверяем реализацию интерфейса

		// Assert
		assert.Equal(t, "AnimalMoved", eventName)
		assert.Equal(t, "animal-123", event.AnimalID)
		assert.Equal(t, "enc-old", event.FromEnclosure)
		assert.Equal(t, "enc-new", event.ToEnclosure)
		assert.Implements(t, (*domain.DomainEvent)(nil), domainEvent)
	})

	t.Run("FeedingTimeEvent", func(t *testing.T) {
		// Arrange
		event := domain.FeedingTimeEvent{
			AnimalID: "animal-123",
			Time:     "12:00",
			Food:     "Meat",
		}

		// Act
		eventName := event.EventName()
		var domainEvent domain.DomainEvent = event // Проверяем реализацию интерфейса

		// Assert
		assert.Equal(t, "FeedingTime", eventName)
		assert.Equal(t, "animal-123", event.AnimalID)
		assert.Equal(t, "12:00", event.Time)
		assert.Equal(t, "Meat", event.Food)
		assert.Implements(t, (*domain.DomainEvent)(nil), domainEvent)
	})

	t.Run("DomainEventInterface", func(t *testing.T) {
		// Проверяем что интерфейс DomainEvent требует только EventName()
		var _ domain.DomainEvent = domain.AnimalMovedEvent{}
		var _ domain.DomainEvent = domain.FeedingTimeEvent{}
	})
}
