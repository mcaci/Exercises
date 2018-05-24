package speed

import "math"
import "bridge/env"

type Range []int

func MotorbikeSpeedRange(m *(env.Motorbike), b *(env.Bridge)) *Range {
	baseSpeed := int(math.Min(float64(b.GapLength), float64(m.Speed))) + 1
	highSpeed := int(math.Max(float64(b.GapLength+1), float64(m.Speed)))

	speedsSize := highSpeed - baseSpeed + 1
	speedRange := make(Range, speedsSize)
	for i := 0; i < speedsSize; i++ {
		speedRange[i] = highSpeed - i
	}
	return &speedRange
}
