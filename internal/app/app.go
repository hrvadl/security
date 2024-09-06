package app

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/hrvadl/security/internal/app/cli"
	"github.com/hrvadl/security/internal/app/iocrypto"
	"github.com/hrvadl/security/internal/domain/cipher/rearrangement"
)

func New() *App {
	return &App{}
}

type App struct{}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	// menu := cli.NewMenu()
	// opt := menu.GetAll()
	opt := cli.Options{
		InputPath:     filepath.Join("./static", "in.txt"),
		KeyPath:       filepath.Join("./static", "key.txt"),
		DecryptedFile: filepath.Join("./static", "decrypt.txt"),
		EncryptedFile: filepath.Join("./static", "encrypt.txt"),
	}

	encryptedFile, err := recreateFile(opt.EncryptedFile)
	if err != nil {
		return fmt.Errorf("failed to open the encrypted file: %w", err)
	}
	defer func() {
		logIfError(encryptedFile.Close())
	}()

	decryptedFile, err := recreateFile(opt.DecryptedFile)
	if err != nil {
		return fmt.Errorf("failed to open the decrypted file: %w", err)
	}
	defer func() {
		logIfError(decryptedFile.Close())
	}()

	inputFile, err := os.Open(opt.InputPath)
	if err != nil {
		return fmt.Errorf("failed to open the input file: %w", err)
	}
	defer func() {
		logIfError(inputFile.Close())
	}()

	cipherSuite := rearrangement.NewCipher([]int{4, 3, 2, 1})

	fw := bufio.NewWriter(encryptedFile)
	enc := iocrypto.NewEncrypter(inputFile, fw, cipherSuite)

	if err := enc.Encrypt(); err != nil {
		return fmt.Errorf("failed to encrypt: %w", err)
	}
	logIfError(fw.Flush())

	encryptedFile, err = os.Open(opt.EncryptedFile)
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

func getKey(key string) ([]int, error) {
	keyFile, err := os.Open(key)
	if err != nil {
		return nil, fmt.Errorf("failed to open the key file: %w", err)
	}

	defer keyFile.Close()

	content, err := io.ReadAll(keyFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read key file: %w", err)
	}

	return convertKeyFromStrToInt(string(content))
}

func convertKeyFromStrToInt(content string) ([]int, error) {
	splits := strings.Split(strings.TrimSuffix(content, "\n"), ",")
	res := make([]int, 0, len(splits))
	for _, str := range splits {
		order, err := strconv.Atoi(str)
		if err != nil {
			return nil, fmt.Errorf("failed to convert key order to int: %w", err)
		}
		res = append(res, order)
	}

	return res, nil
}

func logIfError(err error) {
	if err != nil {
		fmt.Printf("got error: %v", err)
	}
}
