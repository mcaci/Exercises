package motorbike

import "math"

func speedMap(m *motorbike, b *bridge) map[int]int {
	speeds := make(map[int]int)
	distanceToCover := b.RoadLength - 1
	speedsToCover := getSpeedsToCover(m, b)
	for distanceToCover > 0 {

		for _, speed := range speedsToCover {
			distanceToCover -= speed
		}

		for _, speed := range speedsToCover {
			howManyTimesItKeepsTheSpeed := distanceToCover / speed
			distanceToCover = distanceToCover % speed
			speeds[speed] = 1 + howManyTimesItKeepsTheSpeed
		}

		if distanceToCover > 0 {
			lastItem := speedsToCover[len(speedsToCover)-1]
			speedsToCover = append(speedsToCover, lastItem-1)
			distanceToCover = b.RoadLength - 1
		}
	}
	return speeds
}

func getSpeedsToCover(m *motorbike, b *bridge) []int {
	baseSpeed := int(math.Min(float64(b.GapLength), float64(m.Speed))) + 1
	highSpeed := int(math.Max(float64(b.GapLength+1), float64(m.Speed)))
	speedsSize := highSpeed - baseSpeed + 1
	speeds := make([]int, speedsSize)
	for i := 0; i < speedsSize; i++ {
		speeds[i] = highSpeed - i
	}
	return speeds
}