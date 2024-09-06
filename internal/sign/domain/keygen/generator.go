package keygen

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"
)

func New(keySize int) *Generator {
	return &Generator{
		keySize: keySize,
	}
}

// Generator is a small facade which generates
// and RSA private key with the given key size.
type Generator struct {
	keySize int
}

func (g *Generator) Generate() (*rsa.PrivateKey, error) {
	key, err := rsa.GenerateKey(rand.Reader, g.keySize)
	if err != nil {
		return nil, errors.Join(ErrFailedToGenerateKey, err)
	}

	return key, nil
}
