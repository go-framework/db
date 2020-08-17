package db

import (
	"context"
)

type debugContextKey struct{}

func NewDebugContext(ctx context.Context, debug bool) context.Context {
	return context.WithValue(ctx, debugContextKey{}, debug)
}

func GetDebugFromContext(ctx context.Context) bool {
	return ctx.Value(debugContextKey{}).(bool)
}
