package server

import (
	"context"
	"fmt"
	"github.com/benjaminws/ftdd/internal/resolver"
	"log"
	"net"
	"time"
)

const maxBufferSize = 1024

func Server(ctx context.Context, address string) (err error) {
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

	doneChan := make(chan error, 1)
	buffer := make([]byte, maxBufferSize)

	go func() {
		for {
			n, _, err := pc.ReadFrom(buffer)
			if err != nil {
				doneChan <- err
				return
			}

			//log.Printf("packet-received: bytes=%d from=%s\n",
			//	n, addr.String())

			resolver.ResolveForzaDataForBuffer(buffer, n)

			deadline := time.Now().Add(time.Duration(int64(30)))
			err = pc.SetWriteDeadline(deadline)
			if err != nil {
				doneChan <- err
				return
			}
		}
	}()

	select {
	case <-ctx.Done():
		fmt.Println("cancelled")
		err = ctx.Err()
	case err = <-doneChan:
	}

	return
}
