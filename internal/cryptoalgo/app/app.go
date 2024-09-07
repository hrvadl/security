package app

import (
	"fmt"
	"path/filepath"

	"github.com/hrvadl/security/internal/cryptoalgo/app/cli"
	"github.com/hrvadl/security/internal/cryptoalgo/app/filecrypto"
)

func New(cryptoType string) *App {
	return &App{
		cryptoType: cryptoType,
	}
}

type App struct {
	cryptoType string
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	basePath := filepath.Join("./static", "cipher", a.cryptoType)
	opt := cli.Options{
		InputPath:      filepath.Join(basePath, "in.txt"),
		KeyPath:        filepath.Join(basePath, "key.txt"),
		GuessedKeyPath: filepath.Join(basePath, "guessed.key.txt"),
		DecryptedFile:  filepath.Join(basePath, "decrypted.cipher.txt"),
		EncryptedFile:  filepath.Join(basePath, "encrypted.cipher.txt"),
	}

	fileEncDec := filecrypto.NewEncrypterDecrypter(
		opt.InputPath,
		opt.KeyPath,
		opt.EncryptedFile,
		opt.DecryptedFile,
		a.cryptoType,
	)

	if err := fileEncDec.EncryptAndDecrypt(); err != nil {
		return fmt.Errorf("failed to enc/dec: %w", err)
	}

	keyGuesser := filecrypto.NewKeyDecryptor(
		opt.EncryptedFile,
		opt.GuessedKeyPath,
		opt.InputPath,
	)

	if a.cryptoType == filecrypto.Caesar {
		return keyGuesser.GetKey()
	}

	return nil
}
