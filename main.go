package main

import (
	"context"
	"github.com/benjaminws/ftdd/internal/server"
	"log"
)

func main() {
	ctx := context.Background()
	if err := server.Server(ctx, ":6969"); err != nil {
		log.Fatal(err)
	}
}
