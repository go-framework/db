package redis_test

import (
	"flag"
	"testing"

	"github.com/go-redis/redis/v8"
)

var (
	address = "localhost:6379"
)

func init() {
	testing.Init()
	flag.StringVar(&address, "addr", address, "redis address")
}

func newRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: address,
	})
	return rdb
}
