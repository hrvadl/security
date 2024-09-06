package ceasar

import "bytes"

func NewCipher(shifter Shifter) *Cipher {
	return &Cipher{
		shifter: shifter,
	}
}

type Cipher struct {
	shifter Shifter
}

func (c *Cipher) Encrypt(msg []byte) ([]byte, error) {
	msgRunes := bytes.Runes(msg)
	encrypted := make([]rune, 0, len(msgRunes))
	for _, r := range msgRunes {
		encrypted = append(encrypted, c.shifter.Forward(r))
	}
	return []byte(string(encrypted)), nil
}

func (c *Cipher) Decrypt(msg []byte) ([]byte, error) {
	msgRunes := bytes.Runes(msg)
	encrypted := make([]rune, 0, len(msgRunes))
	for _, r := range msgRunes {
		encrypted = append(encrypted, c.shifter.Backward(r))
	}
	return []byte(string(encrypted)), nil
}

func (c *Cipher) Chunk() int {
	return 10
}
