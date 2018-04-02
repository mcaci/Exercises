package path

import . "golfCourse"

func FindPath(golfCourseMap []string) []string {
	emptyGolfMap := CopyToEmptyGolfCourseMap(golfCourseMap)
	countBalls := CountBalls(golfCourseMap)
	starts := FindBalls(golfCourseMap)
	ends := FindHoles(golfCourseMap)
	var paths [](*Path)
	if countBalls == 1 {
		paths = pathsFromBallToHole(starts, ends)
	} else if starts[0].X() == 0 && starts[0].Y() == 0 {
		ball0 := starts[0].(Ball)
		step0 := Path{&ball0, &([]string{"v"})}
		ball1 := starts[1].(Ball)
		step1 := Path{&ball1, &([]string{"v"})}
		paths = [](*Path){&step0, &step1}
	} else {
		ball0 := starts[0].(Ball)
		step0 := Path{&ball0, &([]string{"<"})}
		ball1 := starts[1].(Ball)
		step1 := Path{&ball1, &([]string{">"})}
		paths = [](*Path){&step0, &step1}
	}
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
