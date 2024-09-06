package rearrangement

import (
	"bytes"
)

const keyCorrelation = 1

func NewCipher(key []int) *Cipher {
	return &Cipher{
		key: key,
	}
}

// Cipher constructs an object which
// can encrypt/decrypt data with rearrangement.
// Key should be an slice of integers with the rearrangement order.
// I.E: Text "KYIV" (1 - K, 2 - Y, 3 - I, 4 - V) with the key of []int{4,3,2,1} will result in
// "VIYK" (4 - V, 3 - I, 2 - Y, 1 - K).
type Cipher struct {
	key []int
}

// Encrypt method encrypts message reordering it's content
// by a given key. It will return an error only, if length
// of the message doesn't match length of the key.
func (c *Cipher) Encrypt(msg []byte) ([]byte, error) {
	msgRunes := bytes.Runes(msg)
	if c.keyLengthIsInvalid(msgRunes) {
		return nil, ErrInvalidKeyLength
	}

	encrypted := c.rearrange(msgRunes)
	return []byte(string(encrypted)), nil
}

// Decrypt method decrypts message reordering it's content
// by a given key. It will return an error only, if length
// of the message doesn't match length of the key.
func (wr *Cipher) Decrypt(msg []byte) ([]byte, error) {
	return wr.Encrypt(msg)
}

// Chunk return the length of the chunk of text
// it's capable to proccess. Chunk is equal to the
// length of the key.
func (c *Cipher) Chunk() int {
	return len(c.key)
}

func (wr *Cipher) rearrange(msg []rune) []rune {
	encryptedMsg := make([]rune, len(msg))
	for i, character := range msg {
		rearrangedIndex := wr.key[i] - keyCorrelation
		encryptedMsg[rearrangedIndex] = character
	}
	return encryptedMsg
}

func (c *Cipher) keyLengthIsInvalid(msg []rune) bool {
	return len(msg) != len(c.key)
}
