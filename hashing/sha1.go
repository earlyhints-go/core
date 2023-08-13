package hashing

import (
	"crypto/sha1"
	"encoding/hex"
)

// SHA1Hasher ...
type SHA1Hasher struct {
	cfg Config
}

// NewSHA1Hasher creates new SHA1Hasher.
func NewSHA1Hasher(c Config) *SHA1Hasher {
	c.setDefault()

	return &SHA1Hasher{
		cfg: c,
	}
}

// GetHash from params.
func (h *SHA1Hasher) GetHash(params Params) string {
	hash := sha1.New()

	hash.Write([]byte(params.Method))
	hash.Write([]byte(params.Path))
	hash.Write([]byte(params.Path))

	return hex.EncodeToString(hash.Sum(nil))
}
