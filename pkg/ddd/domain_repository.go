package ddd

import "context"

// DomainRepo is a abstract repository for a domain entity.
type DomainRepo[ID EntityID[K], K comparable, T any] interface {
	// Load loads an entity by its ID.
	Load(ctx context.Context, id ID) (T, error)
	// Save saves an entity.
	Save(ctx context.Context, entity T) error
}
