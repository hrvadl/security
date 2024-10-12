package main

import (
	"crypto/des"
	"encoding/base64"
	"errors"
	"fmt"
)

const (
	allowedKeyLength  = 8
	allowedTextLength = allowedKeyLength
)

func main() {
	var (
		key  string
		text string
	)

	fmt.Printf("Provide encryption key (8 bytes):\n")
	if _, err := fmt.Scanln(&key); err != nil {
		panic(fmt.Sprintf("failed to scan key: %v", err))
	}

	fmt.Printf("\nProvide text to encrypt (8 bytes):\n")
	if _, err := fmt.Scanln(&text); err != nil {
		panic(fmt.Sprintf("failed to scan text: %v", err))
	}

	encrypted, err := encrypt([]byte(key), []byte(text))
	if err != nil {
		panic(fmt.Sprintf("failed to encrypt text: %v", err))
	}

	fmt.Printf("\nYour encrypted text:\n%s\n", encrypted)
}

func encrypt(key, text []byte) ([]byte, error) {
	if len(key) != allowedKeyLength {
		return nil, errors.New("key should be 64 bytes length")
	}

	if len(text) != allowedTextLength {
		return nil, errors.New("text should be 64 bytes length")
	}

	block, err := des.NewCipher([]byte(key))
	if err != nil {
		return nil, fmt.Errorf("failed to get DES block: %w", err)
	}

	encrypted := make([]byte, block.BlockSize())
	block.Encrypt(encrypted, text)

	encoded := make([]byte, base64.StdEncoding.EncodedLen(len(encrypted)))
	base64.StdEncoding.Encode(encoded, encrypted)

	return encoded, nil
}
