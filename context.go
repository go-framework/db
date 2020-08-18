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

type tableContextKey struct{}

func NewTableContext(ctx context.Context, name string) context.Context {
	return context.WithValue(ctx, tableContextKey{}, name)
}

func GetTableFromContext(ctx context.Context) string {
	return ctx.Value(tableContextKey{}).(string)
}
