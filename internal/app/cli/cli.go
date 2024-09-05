package cli

import (
	"bufio"
	"fmt"
	"os"
)

const (
	EncryptIntention = "encrypt"
	DecryptIntention = "decrypt"
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
	InputPath string
	OutPath   string
	KeyPath   string
}

func (m *Menu) GetAll() Options {
	return Options{
		KeyPath:   m.GetKeyFilepath(),
		InputPath: m.GetInputFilepath(),
		OutPath:   m.GetOutputFilepath(),
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

func (m *Menu) GetOutputFilepath() string {
	fmt.Println("Where would you like to save output data?")
	return m.getText()
}

func (m *Menu) GetIntention() string {
	fmt.Println("What would you like to do with it? (encrypt/decrypt)")
	return m.getText()
}

func (m *Menu) getText() string {
	_ = m.scanner.Scan()
	return m.scanner.Text()
}
