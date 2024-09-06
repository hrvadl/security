package file

import (
	"bufio"
	"fmt"
	"os"
)

func NewReplacer() *Replacer {
	return &Replacer{}
}

type Replacer struct{}

func (r *Replacer) ReplaceOrCreate(path string, content []byte) error {
	_ = os.Remove(path)
	return r.create(path, content)
}

func (r *Replacer) create(path string, content []byte) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	fw := bufio.NewWriter(f)
	if _, err := fw.Write(content); err != nil {
		return fmt.Errorf("failed to write content to the file: %w", err)
	}

	if err := fw.Flush(); err != nil {
		return fmt.Errorf("failed to flush bytes to the file: %w", err)
	}

	return nil
}
