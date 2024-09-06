package iocrypto

import (
	"bytes"
	"io"
	"slices"
)

func NewDecrypter(in io.Reader, out io.Writer, cipher CipherEncryptSource) *Encrypter {
	return NewEncrypter(in, out, cipher)
}

func NewEncrypter(in io.Reader, out io.Writer, cipher CipherEncryptSource) *Encrypter {
	return &Encrypter{
		in:     in,
		out:    out,
		cipher: cipher,
	}
}

type CipherEncryptSource interface {
	Encrypt(msg []byte) ([]byte, error)
	Decrypt(msg []byte) ([]byte, error)
	Chunk() int
}

// Encrypter struct is responsible for reading content
// from the given io.Reader then splitting the content
// into the chunks, so underlying cipher algorithm could
// understand it and then writing the output to the given
// io.Writer.
type Encrypter struct {
	in     io.Reader
	out    io.Writer
	cipher CipherEncryptSource
}

func (e *Encrypter) Encrypt() error {
	buf, err := io.ReadAll(e.in)
	if err != nil {
		return err
	}

	buf = bytes.Trim(buf, "\n")
	encrypted, err := e.encryptChunks(buf)
	if err != nil {
		return err
	}

	_, err = e.out.Write(encrypted)
	return err
}

func (e *Encrypter) Decrypt() error {
	buf, err := io.ReadAll(e.in)
	if err != nil {
		return err
	}

	decrypted, err := e.cipher.Decrypt(buf)
	if err != nil {
		return err
	}

	_, err = e.out.Write(decrypted)
	return err
}

func (e *Encrypter) encryptChunks(buf []byte) ([]byte, error) {
	encrypted := make([]byte, 0, len(buf))
	bufRunes := bytes.Runes(buf)

	for chunk := range slices.Chunk(bufRunes, e.cipher.Chunk()) {
		encryptedChunk, err := e.cipher.Encrypt([]byte(string(chunk)))
		if err != nil {
			return nil, err
		}
		encrypted = append(encrypted, encryptedChunk...)
	}

	return encrypted, nil
}
