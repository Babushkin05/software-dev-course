package events_test

import (
	"bytes"
	"testing"

	"github.com/Babushkin05/software-dev-course/hw2/internal/domain"
	"github.com/Babushkin05/software-dev-course/hw2/internal/infrastructure/events"
	"github.com/stretchr/testify/assert"
)

func TestConsoleEventPublisher_Publish(t *testing.T) {
	tests := []struct {
		name     string
		event    domain.DomainEvent
		expected string
	}{
		{
			name: "AnimalMovedEvent",
			event: domain.AnimalMovedEvent{
				AnimalID:      "123",
				FromEnclosure: "old",
				ToEnclosure:   "new",
			},
			expected: "Event published: AnimalMoved => {AnimalID:123 FromEnclosure:old ToEnclosure:new}\n",
		},
		{
			name: "FeedingTimeEvent",
			event: domain.FeedingTimeEvent{
				AnimalID: "456",
				Time:     "12:00",
				Food:     "Meat",
			},
			expected: "Event published: FeedingTime => {AnimalID:456 Time:12:00 Food:Meat}\n",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			var buf bytes.Buffer
			publisher := events.NewConsoleEventPublisher()

			// Redirect stdout to buffer
			old := setOutput(&buf)
			defer func() { setOutput(old) }()

			// Act
			publisher.Publish(tc.event)

			// Assert
			assert.Equal(t, tc.expected, buf.String())
		})
	}
}

// setOutput redirects fmt output and returns the previous writer
func setOutput(w *bytes.Buffer) *bytes.Buffer {
	old := new(bytes.Buffer)
	old, w = w, old
	return old
}
