package app

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/hrvadl/security/internal/sign/app/file"
	"github.com/hrvadl/security/internal/sign/domain/hash"
	"github.com/hrvadl/security/internal/sign/domain/keygen"
	"github.com/hrvadl/security/internal/sign/domain/sign"
	"github.com/hrvadl/security/internal/sign/domain/sign/contentsign"
)

const keySize = 4096

var (
	baseDir        = "./static"
	filePath       = filepath.Join(baseDir, "in.txt")
	privateKeyPath = filepath.Join(baseDir, "private.key")
	publicKeyPath  = filepath.Join(baseDir, "public.key")
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
	f, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer func() {
		logIfError(f.Close())
	}()

	content, err := io.ReadAll(f)
	if err != nil {
		return fmt.Errorf("failed to read file's content: %w", err)
	}

	hasher := hash.NewHasher()

	keyGenerator := keygen.New(keySize)
	key, err := keyGenerator.Generate()
	if err != nil {
		return fmt.Errorf("failed to generate key: %w", err)
	}

	signer := sign.NewSigner(key, hasher)
	signature, err := signer.SignToBase64(content)
	if err != nil {
		return fmt.Errorf("failed to sign data: %w", err)
	}

	appender := contentsign.NewAppender()
	signedContent := appender.AppendSign(content, signature)

	fileReplacer := file.NewReplacer()
	if err := fileReplacer.ReplaceOrCreate(filePath, signedContent); err != nil {
		return fmt.Errorf("failed to replace file: %w", err)
	}

	extracter := contentsign.NewExtracter()
	newContent, newSignature, err := extracter.ExtractSign(signedContent)
	if err != nil {
		return fmt.Errorf("failed to extract signature: %w", err)
	}

	if ok := signer.Verify(newSignature, newContent); !ok {
		return errors.New("signature is unverified")
	}

	replacer := file.NewReplacer()
	pemPrivateKey := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	})
	if err := replacer.ReplaceOrCreate(privateKeyPath, pemPrivateKey); err != nil {
		return fmt.Errorf("failed to save private key: %w", err)
	}

	pemPublicKey := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&key.PublicKey),
	})
	if err := replacer.ReplaceOrCreate(publicKeyPath, pemPublicKey); err != nil {
		return fmt.Errorf("failed to save public key: %w", err)
	}

	fmt.Println("Signature matched!")
	return nil
}

func logIfError(err error) {
	if err != nil {
		fmt.Printf("got error: %v", err)
	}
}
