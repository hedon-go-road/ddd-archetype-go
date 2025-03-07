package ddd

import "time"

// DomainEvent is a struct that contains the metadata for a domain event.
type DomainEvent struct {
	// EventID is the unique identifier for the event.
	EventID int64
	// EntityID is the unique identifier for the entity that triggered the event.
	EntityID int64
	// EventType is the type of the event.
	EventType string
	// EventTime is the time the event occurred.
	EventTime time.Time
}
