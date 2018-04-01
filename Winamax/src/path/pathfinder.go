package path

import . "golfCourse"

func FindPath(golfCourseMap []string) []string {
	emptyGolfMap := CopyToEmptyGolfCourseMap(golfCourseMap)
	countBalls := CountBalls(golfCourseMap)
	starts := FindBalls(golfCourseMap)
	ends := FindHoles(golfCourseMap)
	var steps [](*Path)
	if countBalls == 1 {
		steps = PathFromBallToHole(starts, ends)
	} else if starts[0].X == 0  && starts[0].Y == 0 {
		step0 := Path{starts[0], &([]string{"v"})}
		step1 := Path{starts[1], &([]string{"v"})}
		steps = [](*Path){&step0, &step1}
	} else {
		step0 := Path{starts[0], &([]string{"<"})}
		step1 := Path{starts[1], &([]string{">"})}
		steps = [](*Path){&step0, &step1}
	}
	return replacePathInMap(steps, emptyGolfMap)
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

func PathFromBallToHole(balls [](*Ball), holePosition [](*Hole)) [](*Path) {
	var paths [](*Path)
	for i, ballPosition := range balls {
		pathBall := *ballPosition
		var sequence []string	
		for !(pathBall.X == holePosition[i].X && pathBall.Y == holePosition[i].Y) {
			direction := getDirection(&pathBall, holePosition[i])
			sequence = append(sequence, direction)
		}
		paths = append(paths, &Path{ballPosition, &sequence})
	}
	return paths
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
