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

type Generator struct {
	keySize int
}

func (g *Generator) Generate() (*rsa.PrivateKey, error) {
	privKey, err := rsa.GenerateKey(rand.Reader, g.keySize)
	if err != nil {
		return nil, errors.Join(ErrFailedToGenerateKey, err)
	}

	return privKey, nil
}
