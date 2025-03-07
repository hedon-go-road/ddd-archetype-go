package persistence

import (
	"context"

	domain "github.com/hedon-go-road/ddd-archetype-go/internal/ddd-domain"
	cache "github.com/hedon-go-road/ddd-archetype-go/internal/ddd-infrastructure-cache"
	"github.com/hedon-go-road/ddd-archetype-go/internal/ddd-infrastructure-persistence/db"
	"github.com/hedon-go-road/ddd-archetype-go/internal/log"
)

type StickyNoteEntityRepoImpl struct {
	db    *db.StickyNoteDB
	cache *cache.StickyNoteEntityCache
}

func NewStickyNoteEntityRepoImpl(db *db.StickyNoteDB, cache *cache.StickyNoteEntityCache) *StickyNoteEntityRepoImpl {
	return &StickyNoteEntityRepoImpl{db: db, cache: cache}
}

func (r *StickyNoteEntityRepoImpl) Load(ctx context.Context, id domain.StickyNoteID) (domain.StickyNote, error) {
	entity, err := r.cache.Get(ctx, id)
	if err == nil {
		return entity, nil
	}
	entity, err = r.db.GetByEntityID(ctx, id)
	if err != nil {
		return domain.StickyNote{}, err
	}
	err = r.cache.Set(ctx, id, entity)
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("failed to set sticky note entity to cache")
	}
	return entity, nil
}
func (r *StickyNoteEntityRepoImpl) Save(ctx context.Context, entity domain.StickyNote) error {
	err := r.db.Save(ctx, entity)
	if err != nil {
		return err
	}
	err = r.cache.Delete(ctx, entity.EntityID)
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("failed to delete sticky note entity from cache")
	}
	return nil
}
