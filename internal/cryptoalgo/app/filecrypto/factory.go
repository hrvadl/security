package filecrypto

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/hrvadl/security/internal/cryptoalgo/app/iocrypto"
	"github.com/hrvadl/security/internal/cryptoalgo/domain/cipher/caesar"
	"github.com/hrvadl/security/internal/cryptoalgo/domain/cipher/gamma"
	"github.com/hrvadl/security/internal/cryptoalgo/domain/cipher/rearrangement"
)

const (
	Gamma         = "gamma"
	Caesar        = "caesar"
	Rearrangement = "rearrangement"
)

func newCipherFactory(keypath string) *CipherFactory {
	return &CipherFactory{
		keypath: keypath,
	}
}

type CipherFactory struct {
	keypath string
}

func (cf *CipherFactory) Create(cryptoType string) (iocrypto.CipherEncryptSource, error) {
	switch cryptoType {
	case Gamma:
		return newGammaCipher()
	case Caesar:
		return newCaesarCipher(cf.keypath)
	case Rearrangement:
		return newRearrangementCipher(cf.keypath)
	default:
		return nil, fmt.Errorf("unknown crypto type: %s", cryptoType)
	}
}

func newGammaCipher() (*gamma.Cipher, error) {
	return gamma.NewCipher()
}

func newRearrangementCipher(keypath string) (*rearrangement.Cipher, error) {
	key, err := getKey(keypath)
	if err != nil {
		return nil, fmt.Errorf("failed to read key: %w", err)
	}

	order, err := splitStringToIntSlice(string(key))
	if err != nil {
		return nil, fmt.Errorf("failed to convert key to slice of int: %w", err)
	}

	return rearrangement.NewCipher(order), nil
}

func newCaesarCipher(keypath string) (*caesar.Cipher, error) {
	key, err := getKey(keypath)
	if err != nil {
		return nil, fmt.Errorf("failed to read key: %w", err)
	}

	shift, err := strconv.Atoi(string(key))
	if err != nil {
		return nil, fmt.Errorf("failed to convert shift to int value: %w", err)
	}

	return caesar.NewCipher(shift), nil
}

func getKey(keypath string) ([]byte, error) {
	keyFile, err := os.Open(filepath.Clean(keypath))
	if err != nil {
		return nil, fmt.Errorf("failed to open the key file: %w", err)
	}
	defer func() {
		if err = keyFile.Close(); err != nil {
			slog.Error("Got error", slog.Any("err", err))
		}
	}()

	content, err := io.ReadAll(keyFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read key file: %w", err)
	}

	return bytes.Trim(content, "\n"), nil
}

func splitStringToIntSlice(content string) ([]int, error) {
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
