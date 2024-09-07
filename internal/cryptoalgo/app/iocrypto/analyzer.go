package iocrypto

import (
	"errors"
	"fmt"
	"io"
	"strconv"
)

func NewKeyDecryptor(
	decryptor KeyDecryptorSource,
	enc io.Reader,
	dec io.Reader,
	out io.Writer,
) *KeyDecryptor {
	return &KeyDecryptor{
		decryptor: decryptor,
		out:       out,
		encrypted: enc,
		decrypted: dec,
	}
}

type KeyDecryptorSource interface {
	GetKeyFor(source, encrypted []byte) (int, error)
}

// KeyDecryptor struct is responsible for reading content
// from the given io.Reader then delegating the key guessing
// work to the underlying key decrypting algorithm.
type KeyDecryptor struct {
	decryptor KeyDecryptorSource
	encrypted io.Reader
	decrypted io.Reader
	out       io.Writer
}

func (c *KeyDecryptor) GetKey() error {
	encryptedBuf, err := io.ReadAll(c.encrypted)
	if err != nil {
		return fmt.Errorf("failed to read encrypted buf: %w", err)
	}

	decryptedBuf, err := io.ReadAll(c.decrypted)
	if err != nil {
		return fmt.Errorf("failed to read decrypted buf: %w", err)
	}

	key, err := c.decryptor.GetKeyFor(decryptedBuf, encryptedBuf)
	if err != nil {
		return errors.New("failed to get key")
	}

	if _, err := c.out.Write([]byte(strconv.Itoa(key))); err != nil {
		return fmt.Errorf("failed to write key result: %w", err)
	}

	return nil
}
