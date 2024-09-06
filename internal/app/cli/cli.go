package cli

import (
	"bufio"
	"fmt"
	"os"
)

func NewMenu() *Menu {
	return &Menu{
		scanner: bufio.NewScanner(os.Stdin),
	}
}

type Menu struct {
	scanner *bufio.Scanner
}

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
	fmt.Println("Where's your input file located?")
	return m.getText()
}

func (m *Menu) GetKeyFilepath() string {
	fmt.Println("Where's your key file located?")
	return m.getText()
}

func (m *Menu) GetEncryptedFilepath() string {
	fmt.Println("Where would you like to save encrypted data?")
	return m.getText()
}

func (m *Menu) GetDecryptedFilepath() string {
	fmt.Println("Where would you like to save encrypted data?")
	return m.getText()
}

func (m *Menu) getText() string {
	_ = m.scanner.Scan()
	return m.scanner.Text()
}
