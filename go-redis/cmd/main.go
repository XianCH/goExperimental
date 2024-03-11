package main

import goredis "github.com/x14n/goExperimental/go-redis"

func main() {
	goredis.RedisInit()
	goredis.ListDemo()
}
