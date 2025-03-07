package ddd

import "context"

// DomainRepository is a abstract repository for a domain entity.
type DomainRepository[ID EntityID[K], K comparable, T any] interface {
	// Load loads an entity by its ID.
	Load(ctx context.Context, id ID) (T, error)
	// Save saves an entity.
	Save(ctx context.Context, entity T) error
}
