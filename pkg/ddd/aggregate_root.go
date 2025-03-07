package ddd

// AggregateRoot is a struct that contains the metadata for a domain object.
type AggregateRoot struct {
	// DomainMask is the metadata for the domain object.
	DomainMask
	// DomainEvents is the events that have occurred on the domain object.
	DomainEvents []DomainEvent
}

// AddDomainEvent adds a domain event to the aggregate root.
func (ar *AggregateRoot) AddDomainEvent(events ...DomainEvent) {
	ar.DomainEvents = append(ar.DomainEvents, events...)
}
