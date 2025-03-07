package thirdparty

import "github.com/alicebob/miniredis/v2"

func NewMiniRedis() *miniredis.Miniredis {
	miniRedisClient, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	return miniRedisClient
}
