package ddd

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache[K comparable, T any] struct {
	rdb     redis.Cmdable
	keyFunc func(K) string
	ex      time.Duration
}

func NewRedisCache[K comparable, T any](rdb redis.Cmdable, kf func(K) string, ex time.Duration) *RedisCache[K, T] {
	return &RedisCache[K, T]{rdb: rdb, keyFunc: kf, ex: ex}
}

func (c *RedisCache[K, T]) Get(ctx context.Context, id K) (T, error) {
	var v T
	res, err := c.rdb.Get(ctx, c.keyFunc(id)).Result()
	if err != nil {
		return v, err
	}
	if err := json.Unmarshal([]byte(res), &v); err != nil {
		return v, err
	}
	return v, nil
}

func (c *RedisCache[K, T]) Set(ctx context.Context, id K, value T) error {
	bs, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.rdb.Set(ctx, c.keyFunc(id), bs, c.ex).Err()
}

func (c *RedisCache[K, T]) Delete(ctx context.Context, id K) error {
	return c.rdb.Del(ctx, c.keyFunc(id)).Err()
}
