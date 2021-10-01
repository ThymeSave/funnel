package main

import (
	"log"

	"github.com/thymesave/funnel/cmd"
)

func main() {
	// TODO Make port configurable
	if err := cmd.RunHTTPServer(3000); err != nil {
		log.Fatalf("HTTP-Server crashed: %e", err)
	}
}
