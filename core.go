package core

import (
	"context"

	"github.com/earlyhints-go/core/hashing"
	"github.com/earlyhints-go/core/storage"
)

// Core ...
type Core struct {
	storage storage.Storage
	hasher  hashing.Hasher

	saveHintsFromContext bool
}

// New creates new Core.
func New(opts ...Option) *Core {
	options := getOptions(opts...)

	return &Core{
		storage: options.Storage,
		hasher:  options.Hasher,

		saveHintsFromContext: options.SaveHintsFromContext,
	}
}

// CalculateHash for params.
func (c *Core) CalculateHash(params hashing.Params) string {
	return c.hasher.GetHash(params)
}

// GetHints for hash.
func (c *Core) GetHints(ctx context.Context, hash string) ([]string, error) {
	return c.storage.GetHints(ctx, hash)
}

// SaveHints for hash.
func (c *Core) SaveHints(ctx context.Context, hash string, hints []string) error {
	if c.saveHintsFromContext {
		if container, exist := FromContext(ctx); exist {
			hints = append(hints, container.GetHints()...)
		}
	}

	return c.storage.SaveHints(ctx, hash, hints)
}
