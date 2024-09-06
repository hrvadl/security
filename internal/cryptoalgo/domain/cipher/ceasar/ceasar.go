package ceasar

import "bytes"

func NewCipher(shift int) *Cipher {
	return &Cipher{
		shifter: newShiftStrategy(shift),
	}
}

// Cipher struct is responsible for encrypting or
// decrypting given data with a Ceaser method with
// the configurable shift.
type Cipher struct {
	shifter *shiftStrategy
}

// Encrypt method will encrypt data by shifting
// unicode code point forward for each char with a given
// value. Note: it will shift values cyclically, meaning
// that if shifted value exceeds the maximum allowed code point
// for the latin alphabet, then it will start counting from the
// beginning.
func (c *Cipher) Encrypt(msg []byte) ([]byte, error) {
	msgRunes := bytes.Runes(msg)
	encrypted := make([]rune, 0, len(msgRunes))
	for _, r := range msgRunes {
		encrypted = append(encrypted, c.shifter.Forward(r))
	}
	return []byte(string(encrypted)), nil
}

// Encrypt method will decrypt data by shifting
// unicode code point backward for each char with a given
// value. Note: it will shift values cyclically, meaning
// that if shifted value lower than the min allowed code point
// for the latin alphabet, then it will start counting from the
// ending.
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
