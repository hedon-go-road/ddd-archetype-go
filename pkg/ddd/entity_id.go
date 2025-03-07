package ddd

// EntityID is a value object that represents the unique identifier for a domain entity.
type EntityID[T comparable] interface {
	// Value returns the value of the entity ID.
	Value() T
}
