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

// Signer is a struct, responsible for the
// hashing given data, generating a signature
// and encoding in to the base64 format.
type Signer struct {
	key    *rsa.PrivateKey
	hasher Hasher
}

// SignToBase64 function is responsible for taking a hash
// of a given data, then signing it with the private key
// and returning signature in base64 encoded format.
func (s *Signer) SignToBase64(data []byte) ([]byte, error) {
	signed, err := s.sign(data)
	if err != nil {
		return nil, err
	}

	return s.toBase64(signed)
}

// Verify function verifies given signature. Signature should
// be passed encoded in the base64 format. Data should be in a row format.
// Under the hood it will take hash of the data then decode signature from
// the base64 format and compare signatures.
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

// sign function is responsible for generating a
// hash for the data and then signing that hash.
func (s *Signer) sign(data []byte) ([]byte, error) {
	hashed := s.hasher.Hash(data)
	signed, err := rsa.SignPKCS1v15(rand.Reader, s.key, crypto.SHA256, hashed)
	if err != nil {
		return nil, errors.Join(ErrFailedToSign, err)
	}

	return signed, nil
}

// toBase64 is a handy helper function responsible for
// encoding given data to the base64 format.
func (s *Signer) toBase64(data []byte) ([]byte, error) {
	encoded := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(encoded, data)
	return encoded, nil
}

// fromBase64 is a handy helper function responsible for
// decoding given data from the base64 format.
func (s *Signer) fromBase64(data []byte) ([]byte, error) {
	decoded := make(
		[]byte,
		base64.StdEncoding.DecodedLen(len(data)),
	)
	n, err := base64.StdEncoding.Decode(decoded, data)
	return decoded[:n], err
}
