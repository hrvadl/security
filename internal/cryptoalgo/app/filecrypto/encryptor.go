package filecrypto

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/hrvadl/security/internal/cryptoalgo/app/iocrypto"
)

func NewEncrypterDecrypter(
	input, key, encrypt, decrypt, cryptoType string,
) *EncrypterDecrypter {
	return &EncrypterDecrypter{
		inputPath:   input,
		keyPath:     key,
		encryptPath: encrypt,
		decryptPath: decrypt,
		cryptoType:  cryptoType,
	}
}

// EncrypterDecrypter struct is responsible for reading content
// from the given file and then delegating all the
// work to the iocrypto encrypter.
type EncrypterDecrypter struct {
	inputPath   string
	keyPath     string
	encryptPath string
	decryptPath string
	cryptoType  string
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

	cipherSuite, err := newCipherFactory(e.keyPath).Create(e.cryptoType)
	if err != nil {
		return fmt.Errorf("failed to initialize cipher suite: %w", err)
	}

	fw := bufio.NewWriter(encryptedFile)
	enc := iocrypto.NewEncrypter(inputFile, fw, cipherSuite)

	if err = enc.Encrypt(); err != nil {
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
	path = filepath.Clean(path)
	_ = os.Remove(path)
	return os.Create(path)
}

func logIfError(err error) {
	if err != nil {
		slog.Error("got error", slog.Any("err", err))
	}
}
