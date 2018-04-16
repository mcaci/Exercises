package path

import "golfCourse"

func FindPath(golfCourseMap []string) []string {
	emptyGolfMap := golfCourse.CopyToEmptyGolfCourseMap(golfCourseMap)
	starts := golfCourse.FindBalls(golfCourseMap)
	ends := golfCourse.FindHoles(golfCourseMap)
	paths := pathsFromBallToHole(starts, ends)
	return BuildGolfCourseWithPaths(paths, emptyGolfMap)
}

func pathsFromBallToHole(balls, holes [](golfCourse.Coordinate)) [](*Path) {
	var paths [](*Path)
	for i, ball := range balls {
		ballAtCurrentPosition := *(ball.(*(golfCourse.Ball)))
		var sequence []string
		for !(golfCourse.AreInSamePosition(&ballAtCurrentPosition, holes[i])) {
			direction := getDirection(&ballAtCurrentPosition, holes[i])
			sequence = append(sequence, direction)
			ballAtCurrentPosition.Move(direction)
		}
		ballAtStartPosition := ball.(*(golfCourse.Ball))
		paths = append(paths, &Path{ballAtStartPosition, &sequence})
	}
	return paths
}

func getDirection(ball, hole golfCourse.Coordinate) string {
	var direction string
	if ball.X() < hole.X() {
		direction = "v"
	} else if ball.Y() < hole.Y() {
		direction = ">"
	} else if ball.X() > hole.X() {
		direction = "^"
	} else if ball.Y() > hole.Y() {
		direction = "<"
	} else {
		panic("no direction chosen")
	}
	return direction
}
