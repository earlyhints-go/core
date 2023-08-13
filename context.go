package core

import "context"

type contextKey struct{}

// NewContext returns ctx with Container value.
func NewContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, contextKey{}, &Container{})
}

// FromContext returns Container from ctx.
func FromContext(ctx context.Context) (*Container, bool) {
	c, ok := ctx.Value(contextKey{}).(*Container)
	return c, ok
}
