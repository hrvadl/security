package app

import (
	"path/filepath"

	"github.com/hrvadl/security/internal/cryptoalgo/app/cli"
	"github.com/hrvadl/security/internal/cryptoalgo/app/filecrypto"
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
		InputPath: filepath.Join("./static", "in.txt"),
		KeyPath:   filepath.Join("./static", "key.txt"),
		// GuessedKeyPath: filepath.Join("./static", "guessed_key.txt"),
		DecryptedFile: filepath.Join("./static", "decrypt.txt"),
		EncryptedFile: filepath.Join("./static", "encrypt.txt"),
	}

	fileEncDec := filecrypto.NewEncrypterDecrypter(
		opt.InputPath,
		opt.KeyPath,
		opt.EncryptedFile,
		opt.DecryptedFile,
	)

	return fileEncDec.EncryptAndDecrypt()

	// keyGuesser := filecrypto.NewKeyDecryptor(
	//
	//	opt.EncryptedFile,
	//	opt.GuessedKeyPath,
	//	opt.InputPath,
	//
	// )
	//
	// return keyGuesser.GetKey()
}
