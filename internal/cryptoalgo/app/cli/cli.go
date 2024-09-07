package cli

import (
	"bufio"
	"log/slog"
	"os"
)

func NewMenu() *Menu {
	return &Menu{
		scanner: bufio.NewScanner(os.Stdin),
	}
}

// Menu is a structure, responsible for filling
// Options struct with the values provided by user
// from the CLI.
type Menu struct {
	scanner *bufio.Scanner
}

// Options is a main configuration options availablel
// for the user. All paths should be a valid paths
// provided in the UNIX compatible format.
type Options struct {
	InputPath      string
	EncryptedFile  string
	DecryptedFile  string
	KeyPath        string
	GuessedKeyPath string
}

func (m *Menu) GetAll() Options {
	return Options{
		KeyPath:       m.GetKeyFilepath(),
		InputPath:     m.GetInputFilepath(),
		DecryptedFile: m.GetEncryptedFilepath(),
		EncryptedFile: m.GetEncryptedFilepath(),
	}
}

func (m *Menu) GetInputFilepath() string {
	slog.Info("Where's your input file located?")
	return m.getText()
}

func (m *Menu) GetKeyFilepath() string {
	slog.Info("Where's your key file located?")
	return m.getText()
}

func (m *Menu) GetEncryptedFilepath() string {
	slog.Info("Where would you like to save encrypted data?")
	return m.getText()
}

func (m *Menu) GetDecryptedFilepath() string {
	slog.Info("Where would you like to save encrypted data?")
	return m.getText()
}

func (m *Menu) getText() string {
	_ = m.scanner.Scan()
	return m.scanner.Text()
}
