package main

import (
	"fmt"
	"os"

	"ezyo/forum/http/route"
	_ "ezyo/forum/http/validation"

	"goyave.dev/goyave/v3"
	"goyave.dev/goyave/v3/config"
	_ "goyave.dev/goyave/v3/database/dialect/sqlite"
	// _ "database/model"
)

func main() {
	config.Load()
	// This is the entry point of your application.
	fmt.Printf("API Launched at %s://%s:%d", config.Get("server.protocol"), config.Get("server.host"), config.Get("server.port"))

	if err := goyave.Start(route.Register); err != nil {
		os.Exit(err.(*goyave.Error).ExitCode)
	}

}
