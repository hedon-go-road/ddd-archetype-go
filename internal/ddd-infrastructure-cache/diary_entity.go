package cache

import (
	"fmt"

	"github.com/redis/go-redis/v9"

	domain "github.com/hedon-go-road/ddd-archetype-go/internal/ddd-domain"
	"github.com/hedon-go-road/ddd-archetype-go/pkg/ddd"
)

// DiaryEntityCache is a cache for diary entity.
// TODO: it can be designed as an interface.
type DiaryEntityCache struct {
	*ddd.RedisCache[domain.DiaryID, domain.Diary]
}

func NewDiaryEntityCache(rdb redis.Cmdable) *DiaryEntityCache {
	return &DiaryEntityCache{
		RedisCache: ddd.NewRedisCache[domain.DiaryID, domain.Diary](rdb, func(id domain.DiaryID) string {
			return fmt.Sprintf("ddd:diary:%s", id.Value())
		}, 0),
	}
}
