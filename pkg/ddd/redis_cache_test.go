package ddd

import (
	"context"
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"

	"github.com/hedon-go-road/ddd-archetype-go/thirdparty"
)

func TestRedisCache(t *testing.T) {
	type Item struct {
		ID     int
		Name   string
		Addr   string
		Salary float64
	}

	rdb := thirdparty.NewMiniRedisClient()
	cache := NewRedisCache[int, Item](rdb, func(id int) string {
		return fmt.Sprintf("ddd:diary:%d", id)
	}, 0)

	t.Run("set and get should be equal", func(t *testing.T) {
		err := cache.Set(context.Background(), 1, Item{
			ID:     1,
			Name:   "test",
			Addr:   "test",
			Salary: 1002.01,
		})
		assert.NoError(t, err)

		get, err := cache.Get(context.Background(), 1)
		assert.NoError(t, err)
		assert.Equal(t, Item{
			ID:     1,
			Name:   "test",
			Addr:   "test",
			Salary: 1002.01,
		}, get)
	})

	t.Run("set, delete and get should be nil", func(t *testing.T) {
		err := cache.Set(context.Background(), 2, Item{
			ID:     2,
			Name:   "test2",
			Addr:   "test2",
			Salary: 1002.02,
		})
		assert.NoError(t, err)

		err = cache.Delete(context.Background(), 2)
		assert.NoError(t, err)

		get, err := cache.Get(context.Background(), 2)
		assert.Equal(t, redis.Nil, err)
		assert.Equal(t, Item{}, get)
	})
}
