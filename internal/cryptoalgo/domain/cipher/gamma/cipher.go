package gamma

const (
	A  = uint32(1664525)
	C  = uint32(1013904223)
	m  = uint32(4294967295)
	T0 = uint32(123456789)
)

func NewCipher() (*Cipher, error) {
	return &Cipher{}, nil
}

// Cipher is a struct responsible for encoding
// given data with the Gamma method. It will generate
// key using formula from the instructions. Where A & C
// are constants, m is a maximum machine word.
type Cipher struct {
	key []byte
}

// Encrypt method will generate key for a given data
// using a special formula and then encrypt data with
// XOR operation and this key. NOTE: key will differ for
// messages with the different lengths.
func (c *Cipher) Encrypt(msg []byte) ([]byte, error) {
	c.key = newKey(len(msg))
	encrypted := make([]byte, 0, len(msg))
	lenKey := len(c.key)
	for i, val := range msg {
		encrypted = append(encrypted, val^c.key[i%lenKey])
	}
	return encrypted, nil
}

// Decrypt method is responsible for the decrypting data.
// Under the hood it calls encrypt method with the given
// encrypted data. By encrypting it again with XOR operation
// on encrypted data we can actually decrypt the data.
func (c *Cipher) Decrypt(msg []byte) ([]byte, error) {
	return c.Encrypt(msg)
}

func (c *Cipher) Chunk() int {
	return int(m)
}

func newKey(length int) []byte {
	keyStream := make([]byte, length)
	t := T0
	for i := range length {
		t = (A*t + C) % m
		keyStream[i] = byte(t)
	}
	return keyStream
}
