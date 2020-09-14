package redis

import (
	"context"
	"time"
)

type durationContextKey struct{}

func NewDurationContext(ctx context.Context, duration time.Duration) context.Context {
	return context.WithValue(ctx, durationContextKey{}, duration)
}

func GetDurationFromContext(ctx context.Context) (duration time.Duration, ok bool) {
	duration, ok = ctx.Value(durationContextKey{}).(time.Duration)
	return
}

func NewExpireContext(ctx context.Context, expire time.Duration) context.Context {
	return context.WithValue(ctx, durationContextKey{}, expire)
}

func GetExpireFromContext(ctx context.Context) (expire time.Duration, ok bool) {
	expire, ok = ctx.Value(durationContextKey{}).(time.Duration)
	return
}
