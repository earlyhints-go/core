package core

import (
	"github.com/earlyhints-go/core/hashing"
	"github.com/earlyhints-go/core/storage"
	"github.com/earlyhints-go/core/storage/inmem"
)

// Options ...
type Options struct {
	Storage              storage.Storage
	Hasher               hashing.Hasher
	SaveHintsFromContext bool
}

func (o *Options) setDefaults() {
	if o.Storage == nil {
		o.Storage = inmem.New(0)
	}

	if o.Hasher == nil {
		o.Hasher = hashing.NewSHA1Hasher(hashing.Config{})
	}
}

// Option ...
type Option func(*Options)

// WithStorage ...
func WithStorage(s storage.Storage) Option {
	return func(o *Options) {
		o.Storage = s
	}
}

// WithHasher ...
func WithHasher(h hashing.Hasher) Option {
	return func(o *Options) {
		o.Hasher = h
	}
}

// WithSaveHintsFromContext ...
func WithSaveHintsFromContext(v bool) Option {
	return func(o *Options) {
		o.SaveHintsFromContext = v
	}
}

func getOptions(opts ...Option) *Options {
	var options Options

	for _, opt := range opts {
		opt(&options)
	}

	options.setDefaults()

	return &options
}
