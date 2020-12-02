package server

import (
	"context"
	"fmt"
	"github.com/Datadog/datadog-go/statsd"
	"github.com/benjaminws/ftdd/internal/data"
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

	ddstatsd, err := statsd.New("127.0.0.1:8125")
	if err != nil {
		log.Fatal(err)
	}
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

			forzaData := resolver.ResolveForzaDataForBuffer(buffer, n)
			if forzaData.IsRaceOn == 1 {
				totalSlipRear := int(forzaData.TireCombinedSlipRearLeft + forzaData.TireCombinedSlipRearRight)
				totalSlipFront := int(forzaData.TireCombinedSlipFrontLeft + forzaData.TireCombinedSlipFrontRight)
				carAttitude := resolver.CarAttitude(totalSlipFront, totalSlipRear)
				fun := resolver.HavingFun(carAttitude)

				log.Printf("forzaData: %+v", forzaData)
				tags := []string{
					fmt.Sprintf("car_ordinal:%d", forzaData.CarOrdinal),
					fmt.Sprintf("car_class:%s", data.CarClass(forzaData.CarClass)),
					fmt.Sprintf("car_drivetrain_type:%s", data.DrivetrainType(forzaData.DrivetrainType)),
					fmt.Sprintf("car_attitude:%s", carAttitude),
					fmt.Sprintf("driver_having_fun:%t", fun),
				}
				log.Printf("sending tags: %+v", tags)
				if err := ddstatsd.Gauge("car.current_engine_rpm", float64(forzaData.CurrentEngineRpm), tags, 1); err != nil {
					return
				}
				if err := ddstatsd.Gauge("car.speed_mph", float64(forzaData.SpeedMPH), tags, 1); err != nil {
					return
				}
				if err := ddstatsd.Gauge("car.gear", float64(forzaData.Gear), tags, 1); err != nil {
					return
				}
				if err := ddstatsd.Gauge("car.bhp", float64(forzaData.BrakeHP), tags, 1); err != nil {
					return
				}
				if err := ddstatsd.Gauge("car.torque_ft_lbs", float64(forzaData.TorqueFtLbs), tags, 1); err != nil {
					return
				}
			}

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
