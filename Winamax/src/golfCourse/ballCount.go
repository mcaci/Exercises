package golfCourse

func CountBalls(golfCourseMap []string) int {
	ballCount := 0
	H := len(golfCourseMap)
	for i := 0; i < H; i++ {
		for j := 0; j < len(golfCourseMap[i]); j++ {
			if isPointABall(golfCourseMap[i][j]) {
				ballCount++
			}
		}
	}
	return ballCount
}
