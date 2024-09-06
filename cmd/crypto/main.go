package main

import (
	"github.com/hrvadl/security/internal/app"
)

func main() {
	app := app.New()
	app.MustRun()
}
