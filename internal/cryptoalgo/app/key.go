package app

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func getKey(keypath string) ([]int, error) {
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
