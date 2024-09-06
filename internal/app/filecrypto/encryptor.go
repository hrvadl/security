package filecrypto

import (
	"bufio"
	"fmt"
	"os"

	"github.com/hrvadl/security/internal/app/iocrypto"
	"github.com/hrvadl/security/internal/domain/cipher/gamma"
)

func NewEncrypterDecrypter(
	input, key, encrypt, decrypt string,
) *EncrypterDecrypter {
	return &EncrypterDecrypter{
		inputPath:   input,
		keyPath:     key,
		encryptPath: encrypt,
		decryptPath: decrypt,
	}
}

type EncrypterDecrypter struct {
	inputPath   string
	keyPath     string
	encryptPath string
	decryptPath string
}

func (e *EncrypterDecrypter) EncryptAndDecrypt() error {
	encryptedFile, err := recreateFile(e.encryptPath)
	if err != nil {
		return fmt.Errorf("failed to open the encrypted file: %w", err)
	}
	defer func() {
		logIfError(encryptedFile.Close())
	}()

	decryptedFile, err := recreateFile(e.decryptPath)
	if err != nil {
		return fmt.Errorf("failed to open the decrypted file: %w", err)
	}
	defer func() {
		logIfError(decryptedFile.Close())
	}()

	inputFile, err := os.Open(e.inputPath)
	if err != nil {
		return fmt.Errorf("failed to open the input file: %w", err)
	}
	defer func() {
		logIfError(inputFile.Close())
	}()

	cipherSuite, err := gamma.NewCipher()
	if err != nil {
		return fmt.Errorf("failed to initialize cipher suite: %w", err)
	}

	fw := bufio.NewWriter(encryptedFile)
	enc := iocrypto.NewEncrypter(inputFile, fw, cipherSuite)

	if err := enc.Encrypt(); err != nil {
		return fmt.Errorf("failed to encrypt: %w", err)
	}
	logIfError(fw.Flush())

	encryptedFile, err = os.Open(e.encryptPath)
	if err != nil {
		return fmt.Errorf("failed to open the encrypted file: %w", err)
	}

	fw = bufio.NewWriter(decryptedFile)
	dec := iocrypto.NewDecrypter(encryptedFile, fw, cipherSuite)
	if err := dec.Decrypt(); err != nil {
		return fmt.Errorf("failed to decrypt: %w", err)
	}
	logIfError(fw.Flush())

	return nil
}

func recreateFile(path string) (*os.File, error) {
	_ = os.Remove(path)
	return os.Create(path)
}

func logIfError(err error) {
	if err != nil {
		fmt.Printf("got error: %v", err)
	}
}
