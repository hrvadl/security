package main

import (
	"flag"

	"github.com/hrvadl/security/internal/cryptoalgo/app"
)

var cipherSuite = flag.String(
	"cipher",
	"",
	"Pass the cipher suite you'd like to use. Could be rearrangement/caesar/gamma",
)

func main() {
	flag.Parse()

	app := app.New(*cipherSuite)
	app.MustRun()
}
