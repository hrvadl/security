package main

import (
	"github.com/hrvadl/security/internal/cryptoalgo/app"
)

func main() {
	app := app.New()
	app.MustRun()
}
