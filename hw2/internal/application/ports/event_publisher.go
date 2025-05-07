package ports

import "github.com/Babushkin05/software-dev-course/hw2/internal/domain"

type EventPublisher interface {
	Publish(event domain.DomainEvent)
}
