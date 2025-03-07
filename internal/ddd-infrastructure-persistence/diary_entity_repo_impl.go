package persistence

import (
	"context"

	domain "github.com/hedon-go-road/ddd-archetype-go/internal/ddd-domain"
	cache "github.com/hedon-go-road/ddd-archetype-go/internal/ddd-infrastructure-cache"
	persistencedb "github.com/hedon-go-road/ddd-archetype-go/internal/ddd-infrastructure-persistence/db"
	"github.com/hedon-go-road/ddd-archetype-go/internal/log"
)

type DiaryEntityRepoImpl struct {
	db    *persistencedb.DiaryDB
	cache *cache.DiaryEntityCache
}

func NewDiaryEntityRepoImpl(db *persistencedb.DiaryDB, cache *cache.DiaryEntityCache) *DiaryEntityRepoImpl {
	return &DiaryEntityRepoImpl{db: db, cache: cache}
}

func (r *DiaryEntityRepoImpl) Load(ctx context.Context, id domain.DiaryID) (domain.Diary, error) {
	entity, err := r.cache.Get(ctx, id)
	if err == nil {
		return entity, nil
	}
	entity, err = r.db.GetByID(ctx, id)
	if err != nil {
		return domain.Diary{}, err
	}
	err = r.cache.Set(ctx, id, entity)
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("failed to set diary entity to cache")
	}
	return entity, nil
}

func (r *DiaryEntityRepoImpl) Save(ctx context.Context, entity domain.Diary) error {
	err := r.db.Save(ctx, entity)
	if err != nil {
		return err
	}
	err = r.cache.Delete(ctx, entity.EntityID)
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("failed to delete diary entity from cache")
	}
	return nil
}
