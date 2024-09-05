package app

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/hrvadl/security/internal/app/cli"
	"github.com/hrvadl/security/internal/app/iocrypto"
	"github.com/hrvadl/security/internal/domain/cipher"
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
	menu := cli.NewMenu()
	// opt := menu.GetAll()
	opt := cli.Options{
		InputPath: filepath.Join("./static", "in.txt"),
		KeyPath:   filepath.Join("./static", "key.txt"),
		OutPath:   filepath.Join("./static", "out.txt"),
	}

	key, err := getKey(opt.KeyPath)
	if err != nil {
		return err
	}

	outFile, err := recreateFile(opt.OutPath)
	if err != nil {
		return fmt.Errorf("failed to open the out file: %w", err)
	}
	defer func() {
		logIfError(outFile.Close())
	}()

	inputFile, err := os.Open(opt.InputPath)
	if err != nil {
		return fmt.Errorf("failed to open the input file: %w", err)
	}
	defer func() {
		logIfError(inputFile.Close())
	}()

	rearrangementCipher := cipher.NewWithRearranmegent(key)
	fw := bufio.NewWriter(outFile)
	enc := iocrypto.NewEncrypter(inputFile, fw, rearrangementCipher)
	defer func() {
		logIfError(fw.Flush())
	}()

	switch menu.GetIntention() {
	case cli.DecryptIntention:
		return enc.Decrypt()
	case cli.EncryptIntention:
		return enc.Encrypt()
	default:
		return errors.New("invalid operation")
	}
}

func recreateFile(path string) (*os.File, error) {
	if err := os.Remove(path); err != nil {
		return nil, err
	}
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
