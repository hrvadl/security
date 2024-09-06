package hash

import "crypto/sha256"

func NewHasher() *Hasher {
	return &Hasher{}
}

// Hasher is a facade, which will generate
// hash for the given data. It's main purpose is
// to encapsulate hashing algorithm and simplify the
// signature.
type Hasher struct{}

func (h *Hasher) Hash(data []byte) []byte {
	hash := sha256.New()
	hash.Write(data)
	return hash.Sum(nil)
}
