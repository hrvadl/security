package hash

import "crypto/sha256"

func NewHasher() *Hasher {
	return &Hasher{}
}

type Hasher struct{}

func (h *Hasher) Hash(data []byte) []byte {
	hash := sha256.New()
	hash.Write(data)
	return hash.Sum(nil)
}
