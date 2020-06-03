package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

// Version string, in release version
// This variable will be overwrited by complier
var Version = "SNAPSHOT"

// AppName of this application
var AppName = "TLS Check Tool"

// AppUsage of this application
var AppUsage = "Check the TLS support situation of server"

func main() {
	app := cli.NewApp()
	app.Version = Version
	app.Name = AppName
	app.Usage = AppUsage
	app.Flags = options
	app.EnableBashCompletion = true

	app.Action = action

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
