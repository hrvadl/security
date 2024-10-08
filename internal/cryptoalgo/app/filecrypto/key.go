package filecrypto

import (
	"bufio"
	"fmt"
	"os"

	"github.com/hrvadl/security/internal/cryptoalgo/app/iocrypto"
	"github.com/hrvadl/security/internal/cryptoalgo/domain/analysis"
)

func NewKeyDecryptor(enc string, out string, input string) *KeyDecryptor {
	return &KeyDecryptor{
		encryptPath: enc,
		outPath:     out,
		inputPath:   input,
	}
}

// KeyDecryptor struct is responsible for reading content
// from the given file then delegating the key guessing
// work to the underlying key decrypting algorithm.
type KeyDecryptor struct {
	encryptPath string
	outPath     string
	inputPath   string
}

func (d *KeyDecryptor) GetKey() error {
	encryptedFile, err := os.Open(d.encryptPath)
	if err != nil {
		return fmt.Errorf("failed to open the encrypted file: %w", err)
	}
	defer func() {
		logIfError(encryptedFile.Close())
	}()

	outFileFile, err := recreateFile(d.outPath)
	if err != nil {
		return fmt.Errorf("failed to open the out file: %w", err)
	}
	defer func() {
		logIfError(outFileFile.Close())
	}()

	inputFile, err := os.Open(d.inputPath)
	if err != nil {
		return fmt.Errorf("failed to open the input file: %w", err)
	}
	defer func() {
		logIfError(inputFile.Close())
	}()

	fw := bufio.NewWriter(outFileFile)
	ceasarKeyDec := analysis.NewCaesarKeyDecryptor()
	dec := iocrypto.NewKeyDecryptor(ceasarKeyDec, encryptedFile, inputFile, fw)
	if err := dec.GetKey(); err != nil {
		return fmt.Errorf("failed to decrypt: %w", err)
	}
	logIfError(fw.Flush())

	return nil
}
