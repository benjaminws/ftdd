package main

import (
	"github.com/benjaminws/ftdd/internal/server"
	"log"
)

func main() {
	if err := server.Server(":6969"); err != nil {
		log.Fatal(err)
	}
}
