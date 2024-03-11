package goredis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var (
	rdc *redis.Client
	ctx = context.Background()
)

func RedisInit() {
	rdc = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "3953",
		DB:       0,
	})
}
