package cache

import (
	"fmt"

	"github.com/redis/go-redis/v9"

	domain "github.com/hedon-go-road/ddd-archetype-go/internal/ddd-domain"
	"github.com/hedon-go-road/ddd-archetype-go/pkg/ddd"
)

// StickyNoteEntityCache is a cache for sticky note entity.
// TODO: it can be designed as an interface.
type StickyNoteEntityCache struct {
	*ddd.RedisCache[domain.StickyNoteID, domain.StickyNote]
}

func NewStickyNoteEntityCache(rdb redis.Cmdable) *StickyNoteEntityCache {
	return &StickyNoteEntityCache{
		RedisCache: ddd.NewRedisCache[domain.StickyNoteID, domain.StickyNote](rdb, func(id domain.StickyNoteID) string {
			return fmt.Sprintf("ddd:sticky_note:%s", id.Value())
		}, 0),
	}
}
