// Package inmem implements Storage interface.
package inmem

import (
	"context"
	"sync"

	"github.com/earlyhints-go/core/storage"
)

var _ storage.Storage = (*Storage)(nil)

// Storage ...
type Storage struct {
	data map[string][]string
	m    sync.RWMutex
}

// New creates new Storage.
func New(initialSize int) *Storage {
	return &Storage{
		data: make(map[string][]string, initialSize),
	}
}

// GetHints from store.
func (s *Storage) GetHints(_ context.Context, hash string) ([]string, error) {
	s.m.RLock()
	defer s.m.RUnlock()

	elem, found := s.data[hash]
	if !found {
		return nil, storage.ErrNotFound
	}

	return elem, nil
}

// SaveHints in store.
func (s *Storage) SaveHints(_ context.Context, hash string, hints []string) error {
	s.m.Lock()
	defer s.m.Unlock()

	s.data[hash] = hints

	return nil
}
