package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type ScanKeysFunc func(ctx context.Context, rdb *redis.Client, startCursor uint64, match string, count int64) (keys []string, nextCursor uint64, err error)
