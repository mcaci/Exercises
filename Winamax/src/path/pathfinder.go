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
	} else if starts[0].X() == 0  && starts[0].Y() == 0 {
		step0 := Path{starts[0], &([]string{"v"})}
		step1 := Path{starts[1], &([]string{"v"})}
		paths = [](*Path){&step0, &step1}
	} else {
		step0 := Path{starts[0], &([]string{"<"})}
		step1 := Path{starts[1], &([]string{">"})}
		paths = [](*Path){&step0, &step1}
	}
	return BuildGolfCourseWithPaths(paths, emptyGolfMap)
}

func pathsFromBallToHole(balls [](*Ball), hole [](*Hole)) [](*Path) {
	var paths [](*Path)
	for i, ballAtStartPosition := range balls {
		ballAtCurrentPosition := *ballAtStartPosition
		var sequence []string	
		for !(ballAtCurrentPosition.X() == hole[i].X() && ballAtCurrentPosition.Y() == hole[i].Y()) {
			direction := getDirection(ballAtCurrentPosition, hole[i])
			sequence = append(sequence, direction)
			moveBall(&ballAtCurrentPosition, direction)
		}
		paths = append(paths, &Path{ballAtStartPosition, &sequence})
	}
	return paths
}

func moveBall(ball *Ball, direction string) {
	switch direction {
		case "v": ball.IncrX()
		case "^": ball.DecrX()
		case ">": ball.IncrY()
		case "<": ball.DecrY()
	}
}

func getDirection(ball, hole GolfElement) string {
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
