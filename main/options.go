package main

import "github.com/urfave/cli"

var options = []cli.Flag{
	&cli.StringFlag{
		Name:   "url, u",
		EnvVar: "TLS_CHECK_URL",
		Usage:  "target test url",
		Value:  "https://github.com",
	},
}
