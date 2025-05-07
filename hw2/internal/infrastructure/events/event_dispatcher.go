package events

import (
	"fmt"
	"io"
	"os"

	"github.com/Babushkin05/software-dev-course/hw2/internal/domain"
)

var Stdout io.Writer = io.Writer(os.Stdout)

type ConsoleEventPublisher struct{}

func NewConsoleEventPublisher() *ConsoleEventPublisher {
	return &ConsoleEventPublisher{}
}

func (p *ConsoleEventPublisher) Publish(event domain.DomainEvent) {
	fmt.Fprintf(Stdout, "Event published: %s => %+v\n", event.EventName(), event)
}
