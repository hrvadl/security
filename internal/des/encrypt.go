package des

import (
	"errors"
)

const (
	bitSize  = 1
	byteSize = bitSize * 8
)

const (
	allowedTextLength = 8 * byteSize
	allowedKeyLength  = allowedTextLength
)

const totalRounds = 16

var (
	ErrFailedToEncrypt         = errors.New("failed to encrypt")
	ErrInvalidText             = errors.New("invalid text")
	ErrInvalidKey              = errors.New("invalid key")
	ErrInvalidLength           = errors.New("invalid length of the input")
	ErrFailedToConvertToBits   = errors.New("failed to convert to the bits")
	ErrFailedToConvertFromBits = errors.New("failed to convert from the bits")
)

func Encrypt(text, key string) (string, error) {
	if len(text) != allowedTextLength {
		return "", errors.Join(ErrInvalidText, ErrInvalidLength)
	}

	if len(key) != allowedKeyLength {
		return "", errors.Join(ErrInvalidKey, ErrInvalidLength)
	}

	bits, err := toBits(text)
	if err != nil {
		return "", errors.Join(ErrFailedToEncrypt, err)
	}

	permutated := ip(*(*[allowedTextLength]Bit)(bits))
	encrypted := encrypt(permutated)
	unpermutated := ipreverse(encrypted)

	encryptedStr, err := fromBits(unpermutated[:])
	if err != nil {
		return "", errors.Join(ErrFailedToEncrypt, err)
	}

	return encryptedStr, nil
}

// @TODO: implement
func encrypt(s [64]Bit) [64]Bit {
	encrypted := s
	for range totalRounds {
		encrypted = round(encrypted)
	}
	return encrypted
}

// @TODO: implement
func round(s [64]Bit) [64]Bit {
	return s
}
