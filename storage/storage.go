// Package storage defines Storage interface.
package storage

import (
	"context"
	"errors"
)

// ErrNotFound ...
var ErrNotFound = errors.New("hints by hash not found")

// Storage ...
type Storage interface {
	// GetHints from store.
	GetHints(ctx context.Context, hash string) ([]string, error)
	// SaveHints in store.
	SaveHints(ctx context.Context, hash string, hints []string) error
}
