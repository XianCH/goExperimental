package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "3953",
		DB:       0,
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()

	if err != nil {
		panic(err)
	}

	value, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", value)

	value2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 not exist!" + err.Error())
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", value2)
	}
}
