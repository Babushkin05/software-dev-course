package events

import (
	"fmt"

	"github.com/Babushkin05/software-dev-course/hw2/internal/domain"
)

type ConsoleEventPublisher struct{}

func NewConsoleEventPublisher() *ConsoleEventPublisher {
	return &ConsoleEventPublisher{}
}

func (p *ConsoleEventPublisher) Publish(event domain.DomainEvent) {
	fmt.Printf("Event published: %s => %+v\n", event.EventName(), event)
}
