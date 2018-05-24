package speed

import "bridge/env"

type Map map[int]int

func BuildSpeedMap(m *(env.Motorbike), b *(env.Bridge)) *Map {
	speedMap := make(Map)
	distanceToCover := b.RoadLength - 1
	speedsToCover := *MotorbikeSpeedRange(m, b)
	for distanceToCover > 0 {

		for _, speed := range speedsToCover {
			distanceToCover -= speed
		}

		for _, speed := range speedsToCover {
			howManyTimesItKeepsTheSpeed := distanceToCover / speed
			distanceToCover = distanceToCover % speed
			speedMap[speed] = 1 + howManyTimesItKeepsTheSpeed
		}

		if distanceToCover > 0 {
			lastItem := speedsToCover[len(speedsToCover)-1]
			speedsToCover = append(speedsToCover, lastItem-1)
			distanceToCover = b.RoadLength - 1
		}
	}
	return &speedMap
}
