// Package hashing defines Hasher interface and implementations.
package hashing

// Config of hasher.
type Config struct {
	UseMethod bool
	UseHost   bool
	UsePath   bool
}

func (c *Config) isEmpty() bool {
	return !c.UseMethod && !c.UseHost && !c.UsePath
}

func (c *Config) setDefault() {
	if c.isEmpty() {
		c.UseMethod = true
		c.UseHost = true
		c.UsePath = true
	}
}

// Params for hash.
type Params struct {
	Method string
	Host   string
	Path   string
}

// Hasher ...
type Hasher interface {
	GetHash(params Params) string
}
