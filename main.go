package main

import (
	"os"

	"ezyo/forum/http/route"
	_ "ezyo/forum/http/validation"

	"goyave.dev/goyave/v3"
	"goyave.dev/goyave/v3/config"
	_ "goyave.dev/goyave/v3/database/dialect/sqlite"
)

func main() {
	// Load configuration
	config.Load()

	goyave.Logger.Println("Starting...")
	goyave.RegisterStartupHook(func() {
		goyave.Logger.Printf("API Started. Listening at %s", goyave.BaseURL())
	})

	if err := goyave.Start(route.Register); err != nil {
		os.Exit(err.(*goyave.Error).ExitCode)
	}

}
