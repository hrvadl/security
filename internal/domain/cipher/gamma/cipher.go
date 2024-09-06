package gamma

import (
	"crypto/rand"
	"fmt"
)

func NewCipher() (*Cipher, error) {
	token := make([]byte, 4)
	if _, err := rand.Read(token); err != nil {
		return nil, fmt.Errorf("failed to generate random key: %w", err)
	}

	return &Cipher{
		key: token,
	}, nil
}

type Cipher struct {
	key []byte
}

func (c *Cipher) Encrypt(msg []byte) ([]byte, error) {
	encrypted := make([]byte, 0, len(msg))
	lenKey := len(c.key)
	for i, val := range msg {
		encrypted = append(encrypted, val^c.key[i%lenKey])
	}
	return encrypted, nil
}

func (c *Cipher) Decrypt(msg []byte) ([]byte, error) {
	return c.Encrypt(msg)
}

func (c *Cipher) Chunk() int {
	return len(c.key)
}
