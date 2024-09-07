package file

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func NewReplacer() *Replacer {
	return &Replacer{}
}

// Replacer is small helper structure
// which can be handy in case when you
// want to (re) create file with the given
// content.
type Replacer struct{}

// ReplaceOrCreate function firstly tries to remove existing file.
// Then it will create brand new one with the given content.
func (r *Replacer) ReplaceOrCreate(path string, content []byte) error {
	_ = os.Remove(path)
	return r.create(path, content)
}

func (r *Replacer) create(path string, content []byte) error {
	f, err := os.Create(filepath.Clean(path))
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
