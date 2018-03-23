package golfCourse

func PathFromBallToHole(ballPosition, holePosition *Coordinate) []string {
	var direction string
	var sequence []string
	pathCoordinate := *ballPosition
	for pathCoordinate != *holePosition {
		direction = getDirection(&pathCoordinate, holePosition)
		sequence = append(sequence, direction)
	}
	return sequence
}

func getDirection(ballPosition, holePosition *Coordinate) string {
	direction := holePositionComparedToBall(ballPosition, holePosition)
	switch direction {
		case "v": (ballPosition.X)++
		case "^": (ballPosition.X)--
		case ">": (ballPosition.Y)++
		case "<": (ballPosition.Y)--
	}
	return direction
}

func holePositionComparedToBall(ballPosition, holePosition *Coordinate) string {
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
