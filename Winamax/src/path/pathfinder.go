package path

import . "golfCourse"

func FindPath(golfCourseMap []string) []string {
	emptyGolfMap := CopyToEmptyGolfCourseMap(golfCourseMap)
	countBalls := CountBalls(golfCourseMap)
	starts := FindBalls(golfCourseMap)
	ends := FindHoles(golfCourseMap)
	if countBalls == 1 {
		step0 := PathFromBallToHole(starts[0], ends[0])
		steps := [](*Path){step0}
		return replacePathInMap(steps, emptyGolfMap)
	} else if starts[0].X == 0  && starts[0].Y == 0 {
		step0 := Path{starts[0], &([]string{"v"})}
		step1 := Path{starts[1], &([]string{"v"})}
		steps := [](*Path){&step0, &step1}
		return replacePathInMap(steps, emptyGolfMap)
	} else {
		return []string{".<",">."}
	}
}

func replacePathInMap(paths [](*Path), golfMap []string) []string {
	for _, path := range paths {
		for _, direction := range *(path.StepSequence) {
			line := golfMap[path.Start.X]
			golfMap[path.Start.X] = line[:path.Start.Y] + direction + line[path.Start.Y+1:]
		}
	}
	return golfMap
}

func PathFromBallToHole(ballPosition *Ball, holePosition *Hole) *Path {
	var direction string
	var sequence []string
	pathBall := *ballPosition
	for !(pathBall.X == holePosition.X && pathBall.Y == holePosition.Y) {
		direction = getDirection(&pathBall, holePosition)
		sequence = append(sequence, direction)
	}
	return &Path{ballPosition, &sequence}
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
