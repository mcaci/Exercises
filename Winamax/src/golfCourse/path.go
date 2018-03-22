package golfCourse

func PathFromBallToHole(ballX, ballY, holeX, holeY int) []string {
	var direction string
	var sequence []string
	for ballX != holeX || ballY != holeY {
		direction, ballX, ballY = getDirection(ballX, ballY, holeX, holeY)
		sequence = append(sequence, direction)
	}
	return sequence
}

func getDirection(startX, startY, endX, endY int) (string, int, int) {
	var direction string
	if startX == endX && endY == startY + 1 {
		direction = "E"
		startY++
	} else if startX == endX && endY == startY - 1 {
		direction = "W"
		startY--
	} else if endX == startX + 1 && endY == startY {
		direction = "S"
		startX++
	} else if endX == startX - 1 && endY == startY {
		direction = "N"
		startX--
	}
	return direction, startX, startY
}

