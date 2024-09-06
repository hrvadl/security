package sign

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"errors"
)

func NewSigner(key *rsa.PrivateKey, hasher Hasher) *Signer {
	return &Signer{
		key:    key,
		hasher: hasher,
	}
}

type Hasher interface {
	Hash(data []byte) []byte
}

type Signer struct {
	key    *rsa.PrivateKey
	hasher Hasher
}

func (s *Signer) SignToBase64(data []byte) ([]byte, error) {
	signed, err := s.sign(data)
	if err != nil {
		return nil, err
	}

	return s.toBase64(signed)
}

func (s *Signer) Verify(encodedSignature []byte, data []byte) bool {
	hashed := s.hasher.Hash(data)
	signature, err := s.fromBase64(encodedSignature)
	if err != nil {
		return false
	}

	err = rsa.VerifyPKCS1v15(
		&s.key.PublicKey,
		crypto.SHA256,
		hashed,
		signature,
	)

	return err == nil
}

func (s *Signer) sign(data []byte) ([]byte, error) {
	hashed := s.hasher.Hash(data)
	signed, err := rsa.SignPKCS1v15(rand.Reader, s.key, crypto.SHA256, hashed)
	if err != nil {
		return nil, errors.Join(ErrFailedToSign, err)
	}

	return signed, nil
}

func (s *Signer) toBase64(data []byte) ([]byte, error) {
	encoded := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(encoded, data)
	return encoded, nil
}

func (s *Signer) fromBase64(data []byte) ([]byte, error) {
	decoded := make(
		[]byte,
		base64.StdEncoding.DecodedLen(len(data)),
	)
	n, err := base64.StdEncoding.Decode(decoded, data)
	return decoded[:n], err
}
