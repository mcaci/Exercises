package path

import . "golfCourse"

func FindPath(golfCourseMap []string) []string {
	emptyGolfMap := CopyToEmptyGolfCourseMap(golfCourseMap)
	starts := FindBalls(golfCourseMap)
	ends := FindHoles(golfCourseMap)
	paths := pathsFromBallToHole(starts, ends)
	return BuildGolfCourseWithPaths(paths, emptyGolfMap)
}

func pathsFromBallToHole(balls, holes []Coordinate) [](*Path) {
	var paths [](*Path)
	for i, ball := range balls {
		ballAtCurrentPosition := ball.(Ball)
		var sequence []string	
		for !AreInSamePosition(ballAtCurrentPosition, holes[i]) {
			direction := getDirection(ballAtCurrentPosition, holes[i])
			sequence = append(sequence, direction)
			ballAtCurrentPosition.Move(direction)
		}
		ballAtStartPosition := ball.(Ball)
		paths = append(paths, &Path{&ballAtStartPosition, &sequence})
	}
	return paths
}

func getDirection(ball, hole Coordinate) string {
	var direction string
	if ball.X() < hole.X() && ball.Y() == hole.Y() {
		direction = "v"
	} else if ball.X() == hole.X() && ball.Y() < hole.Y() {
		direction = ">"
	} else if ball.X() > hole.X() && ball.Y() == hole.Y() {
		direction = "^"
	} else if ball.X() == hole.X() && ball.Y() > hole.Y() {
		direction = "<"
	}	
	return direction
}
