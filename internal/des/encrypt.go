package des

import (
	"errors"
)

var ErrInvalidLength = errors.New("invalid length of the input")

func Encrypt(text string) (string, error) {
	if len(text) != 64 {
		return "", ErrInvalidLength
	}

	return "", nil
}
