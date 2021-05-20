package main

import (
	"os"

	"goyave.dev/template/http/route"
	_ "goyave.dev/template/http/validation"

	"goyave.dev/goyave/v3"
	_ "goyave.dev/goyave/v3/database/dialect/sqlite"
)

func main() {
	// This is the entry point of your application.
	if err := goyave.Start(route.Register); err != nil {
		os.Exit(err.(*goyave.Error).ExitCode)
	}
}
