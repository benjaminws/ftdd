package resolver

import (
	"bytes"
	"encoding/binary"
	"github.com/benjaminws/ftdd/internal/data"
	"log"
	"math"
)

func ResolveForzaDataForBuffer(buffer []byte, numBytes int) {

	fd := data.ForzaData{
		IsRaceOn:    computeInt32(getBytes(buffer, numBytes, 0, 4)),
		TimestampMS: binary.LittleEndian.Uint32(getBytes(buffer, numBytes, 4, 8)),
	}

	if fd.IsRaceOn == 1 {
		// SLED data
		fd.EngineMaxRpm = computeFloat32(getBytes(buffer, numBytes, 8, 12))
		fd.EngineIdleRpm = computeFloat32(getBytes(buffer, numBytes, 12, 16))
		fd.CurrentEngineRpm = computeFloat32(getBytes(buffer, numBytes, 16, 20))
		fd.AccelerationX = computeFloat32(getBytes(buffer, numBytes, 20, 24))
		fd.AccelerationY = computeFloat32(getBytes(buffer, numBytes, 24, 28))
		fd.AccelerationZ = computeFloat32(getBytes(buffer, numBytes, 28, 32))
		fd.VelocityX = computeFloat32(getBytes(buffer, numBytes, 32, 36))
		fd.VelocityY = computeFloat32(getBytes(buffer, numBytes, 36, 40))
		fd.VelocityZ = computeFloat32(getBytes(buffer, numBytes, 40, 44))
		fd.AngularVelocityX = computeFloat32(getBytes(buffer, numBytes, 44, 48))
		fd.AngularVelocityY = computeFloat32(getBytes(buffer, numBytes, 48, 52))
		fd.AngularVelocityZ = computeFloat32(getBytes(buffer, numBytes, 52, 56))
		fd.Yaw = computeFloat32(getBytes(buffer, numBytes, 56, 60))
		fd.Pitch = computeFloat32(getBytes(buffer, numBytes, 60, 64))
		fd.Roll = computeFloat32(getBytes(buffer, numBytes, 64, 68))
		fd.NormalizedSuspensionTravelFrontLeft = computeFloat32(getBytes(buffer, numBytes, 68, 72))
		fd.NormalizedSuspensionTravelFrontRight = computeFloat32(getBytes(buffer, numBytes, 72, 76))
		fd.NormalizedSuspensionTravelRearLeft = computeFloat32(getBytes(buffer, numBytes, 76, 80))
		fd.NormalizedSuspensionTravelRearRight = computeFloat32(getBytes(buffer, numBytes, 80, 84))
		fd.TireSlipRatioFrontLeft = computeFloat32(getBytes(buffer, numBytes, 84, 88))
		fd.TireSlipRatioFrontRight = computeFloat32(getBytes(buffer, numBytes, 88, 92))
		fd.TireSlipRatioRearLeft = computeFloat32(getBytes(buffer, numBytes, 92, 96))
		fd.TireSlipRatioRearRight = computeFloat32(getBytes(buffer, numBytes, 96, 100))
		fd.WheelRotationSpeedFrontLeft = computeFloat32(getBytes(buffer, numBytes, 100, 104))
		fd.WheelRotationSpeedFrontRight = computeFloat32(getBytes(buffer, numBytes, 104, 108))
		fd.WheelRotationSpeedRearLeft = computeFloat32(getBytes(buffer, numBytes, 108, 112))
		fd.WheelRotationSpeedRearRight = computeFloat32(getBytes(buffer, numBytes, 112, 116))
		fd.WheelOnRumbleStripFrontLeft = computeInt32(getBytes(buffer, numBytes, 116, 120))
		fd.WheelOnRumbleStripFrontRight = computeInt32(getBytes(buffer, numBytes, 120, 124))
		fd.WheelOnRumbleStripRearLeft = computeInt32(getBytes(buffer, numBytes, 124, 128))
		fd.WheelOnRumbleStripRearRight = computeInt32(getBytes(buffer, numBytes, 128, 132))
		fd.WheelInPuddleDepthFrontLeft = computeFloat32(getBytes(buffer, numBytes, 132, 136))
		fd.WheelInPuddleDepthFrontRight = computeFloat32(getBytes(buffer, numBytes, 136, 140))
		fd.WheelInPuddleDepthRearLeft = computeFloat32(getBytes(buffer, numBytes, 140, 144))
		fd.WheelInPuddleDepthRearRight = computeFloat32(getBytes(buffer, numBytes, 144, 148))
		fd.SurfaceRumbleFrontLeft = computeFloat32(getBytes(buffer, numBytes, 148, 152))
		fd.SurfaceRumbleFrontRight = computeFloat32(getBytes(buffer, numBytes, 152, 156))
		fd.SurfaceRumbleRearLeft = computeFloat32(getBytes(buffer, numBytes, 156, 160))
		fd.SurfaceRumbleRearRight = computeFloat32(getBytes(buffer, numBytes, 160, 164))
		fd.TireSlipAngleFrontLeft = computeFloat32(getBytes(buffer, numBytes, 164, 168))
		fd.TireSlipAngleFrontRight = computeFloat32(getBytes(buffer, numBytes, 168, 172))
		fd.TireSlipAngleRearLeft = computeFloat32(getBytes(buffer, numBytes, 172, 176))
		fd.TireSlipAngleRearRight = computeFloat32(getBytes(buffer, numBytes, 176, 180))
		fd.TireCombinedSlipFrontLeft = computeFloat32(getBytes(buffer, numBytes, 180, 184))
		fd.TireCombinedSlipFrontRight = computeFloat32(getBytes(buffer, numBytes, 184, 188))
		fd.TireCombinedSlipRearLeft = computeFloat32(getBytes(buffer, numBytes, 188, 192))
		fd.TireCombinedSlipRearRight = computeFloat32(getBytes(buffer, numBytes, 192, 196))
		fd.SuspensionTravelMetersFrontLeft = computeFloat32(getBytes(buffer, numBytes, 196, 200))
		fd.SuspensionTravelMetersFrontRight = computeFloat32(getBytes(buffer, numBytes, 200, 204))
		fd.SuspensionTravelMetersRearLeft = computeFloat32(getBytes(buffer, numBytes, 204, 208))
		fd.SuspensionTravelMetersRearRight = computeFloat32(getBytes(buffer, numBytes, 208, 212))

		// Car performance data
		fd.CarOrdinal = computeInt32(getBytes(buffer, numBytes, 212, 216))
		fd.CarClass = computeInt32(getBytes(buffer, numBytes, 216, 220))
		fd.CarPerformanceIndex = computeInt32(getBytes(buffer, numBytes, 220, 224))
		fd.DrivetrainType = computeInt32(getBytes(buffer, numBytes, 224, 228))
		fd.NumCylinders = computeInt32(getBytes(buffer, numBytes, 228, 232))

		// DASH Data
		fd.PositionX = computeFloat32(getBytes(buffer, numBytes, 232, 236))
		fd.PositionY = computeFloat32(getBytes(buffer, numBytes, 236, 240))
		fd.PositionZ = computeFloat32(getBytes(buffer, numBytes, 240, 244))
		fd.Speed = computeFloat32(getBytes(buffer, numBytes, 244, 248))
		// Computed MPH from meters per second
		fd.SpeedMPH = fd.Speed * 2.23694
		fd.Power = computeFloat32(getBytes(buffer, numBytes, 248, 252))
		// Computed Brake Horse Power from watts
		fd.BrakeHP = fd.Power / 745.699872
		fd.Torque = computeFloat32(getBytes(buffer, numBytes, 252, 256))
		fd.Gear = uint8(getBytes(buffer, numBytes, 307, 308)[0])

		// Convert slip values to ints as the precision of a float means a neutral state is rarely reporte
		totalSlipRear := int(fd.TireCombinedSlipRearLeft + fd.TireCombinedSlipRearRight)
		totalSlipFront := int(fd.TireCombinedSlipFrontLeft + fd.TireCombinedSlipFrontRight)

		carAttitude := carAttitude(totalSlipFront, totalSlipRear)
		fun := havingFun(carAttitude)

		log.Printf("RPM: %.0f \t Gear: %d \t BHP: %.0f \t Speed: %.0f \t Total slip: %.0f \t Attitude: %s \t Having Fun?: %t", fd.CurrentEngineRpm, fd.Gear, fd.BrakeHP, fd.SpeedMPH, (fd.TireCombinedSlipRearLeft + fd.TireCombinedSlipRearRight), carAttitude, fun)
	}
}

func getBytes(buffer []byte, numBytes, offset, length int) []byte {
	return buffer[:numBytes][offset:length]
}

func computeInt32(buffer []byte) int32 {
	buf := bytes.NewBuffer(buffer)
	var ret int32
	_ = binary.Read(buf, binary.LittleEndian, &ret)
	return ret
}

func computeFloat32(buffer []byte) float32 {
	bits := binary.LittleEndian.Uint32(buffer)
	return math.Float32frombits(bits)
}

func havingFun(attitude string) bool {
	switch attitude {
	case "Understeer":
		return false
	default:
		return true
	}
}

func carAttitude(totalSlipFront int, totalSlipRear int) string {
	// Check attitude of car by comparing front and rear slip levels
	// If front slip > rear slip, means car is understeering
	if totalSlipRear > totalSlipFront {
		return "Oversteer"
	} else if totalSlipFront > totalSlipRear {
		return "Understeer"
	} else {
		return "Neutral"
	}
}
