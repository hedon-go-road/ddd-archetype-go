package thirdparty

import (
	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
)

func NewMiniRedis() *miniredis.Miniredis {
	miniRedisClient, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	return miniRedisClient
}

func NewMiniRedisClient() *redis.Client {
	server := NewMiniRedis()
	return redis.NewClient(&redis.Options{
		Addr: server.Addr(),
	})
}
