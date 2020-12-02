package server

import (
	"fmt"
	"github.com/Datadog/datadog-go/statsd"
	"github.com/benjaminws/ftdd/internal/resolver"
	"log"
	"net"
	"sync"
)

const maxBufferSize = 1024

func Server(address string) (err error) {
	pc, err := net.ListenPacket("udp", address)
	if err != nil {
		return
	}

	defer func() {
		if err := pc.Close(); err != nil {
			err := fmt.Errorf("could not close udp socket: %w", err)
			log.Fatal(err)
		}
	}()

	ddstatsd, err := statsd.New("127.0.0.1:8125")
	if err != nil {
		log.Fatal(err)
	}

	buffer := make([]byte, maxBufferSize)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			n, _, err := pc.ReadFrom(buffer)
			if err != nil {
				return
			}

			if err := resolver.ResolveForzaDataForBuffer(ddstatsd, buffer, n); err != nil {
				return
			}
		}
	}()

	wg.Wait()
	return
}
