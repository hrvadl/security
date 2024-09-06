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

type Cipher struct {
	key []byte
}

func (c *Cipher) Encrypt(msg []byte) ([]byte, error) {
	c.key = newKey(len(msg))
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
	return int(m)
}

func newKey(length int) []byte {
	keyStream := make([]byte, length)
	T := T0
	for i := 0; i < length; i++ {
		T = (A*T + C) % m
		keyStream[i] = byte(T)
	}
	return keyStream
}
