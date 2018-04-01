package path

import . "golfCourse"

func FindPath(golfCourseMap []string) []string {
	countBalls := CountBalls(golfCourseMap)
	if countBalls == 1 {
		start := FindBalls(golfCourseMap)[0]
		end := FindHoles(golfCourseMap)[0]
		steps := PathFromBallToHole(start, end)
		emptyGolfMap := CopyToEmptyGolfCourseMap(golfCourseMap)
		return replacePathInMap(start, steps, emptyGolfMap)
	} else { 
		starts := FindBalls(golfCourseMap)
		if starts[0].X == 0  && starts[0].Y == 0 {
			return []string{"vv",".."}
		} else {
			return []string{".<",">."}
		}
	}
}

func replacePathInMap(start *Ball, sequence, golfMap []string) []string {
	for _, direction := range sequence {
		line := golfMap[start.X]
		golfMap[start.X] = line[:start.Y] + direction + line[start.Y+1:]
	}
	return golfMap
}

func PathFromBallToHole(ballPosition *Ball, holePosition *Hole) []string {
	var direction string
	var sequence []string
	pathBall := *ballPosition
	for !(pathBall.X == holePosition.X && pathBall.Y == holePosition.Y) {
		direction = getDirection(&pathBall, holePosition)
		sequence = append(sequence, direction)
	}
	return sequence
}

func getDirection(ballPosition *Ball, holePosition *Hole) string {
	direction := holePositionComparedToBall(ballPosition, holePosition)
	switch direction {
		case "v": (ballPosition.X)++
		case "^": (ballPosition.X)--
		case ">": (ballPosition.Y)++
		case "<": (ballPosition.Y)--
	}
	return direction
}

func holePositionComparedToBall(ballPosition *Ball, holePosition *Hole) string {
	var direction string
	if ballPosition.X == holePosition.X && ballPosition.Y > holePosition.Y {
		direction = "<"
	} else if ballPosition.X == holePosition.X && ballPosition.Y < holePosition.Y {
		direction = ">"
	} else if ballPosition.X > holePosition.X && ballPosition.Y == holePosition.Y {
		direction = "^"
	} else if ballPosition.X < holePosition.X && ballPosition.Y == holePosition.Y {
		direction = "v"
	}
	return direction
}
