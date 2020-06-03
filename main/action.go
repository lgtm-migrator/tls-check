package main

import (
	"log"
	"net/url"

	"github.com/urfave/cli"
)

var action = func(c *cli.Context) error {
	u, err := url.Parse(c.GlobalString("url"))
	if err != nil {
		return err
	}
	r, err := CheckTLS(u)
	if err != nil {
		return err
	}
	log.Printf("tls 1.0:\t%v", r.TLS10support)
	log.Printf("tls 1.1:\t%v", r.TLS11support)
	log.Printf("tls 1.2:\t%v", r.TLS12support)
	log.Printf("tls 1.3:\t%v", r.TLS13support)
	log.Printf("tls 3.0:\t%v", r.TLS30support)

	return nil
}
