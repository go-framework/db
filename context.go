package db

import (
	"context"
)

type debugContextKey struct{}

func NewDebugContext(ctx context.Context, debug bool) context.Context {
	return context.WithValue(ctx, debugContextKey{}, debug)
}

func GetDebugFromContext(ctx context.Context) bool {
	value, _ := ctx.Value(debugContextKey{}).(bool)
	return value
}

type tableContextKey struct{}

func NewTableContext(ctx context.Context, name string) context.Context {
	return context.WithValue(ctx, tableContextKey{}, name)
}

func GetTableFromContext(ctx context.Context) string {
	value, _ := ctx.Value(tableContextKey{}).(string)
	return value
}
