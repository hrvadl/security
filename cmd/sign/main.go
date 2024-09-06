package main

import "github.com/hrvadl/security/internal/sign/app"

func main() {
	app := app.New()
	app.MustRun()
}
