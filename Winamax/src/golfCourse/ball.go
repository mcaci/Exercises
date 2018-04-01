package golfCourse

import "strconv"

type Ball struct {
	X, Y int 
}

func FindBalls(golfCourseMap []string) ([](*Ball)) {
	var balls [](*Ball) 
	H := len(golfCourseMap)
	for i := 0; i < H; i++ {
		for j := 0; j < len(golfCourseMap[i]); j++ {
			if isPointABall(string(golfCourseMap[i][j])) {
				balls = append(balls, &Ball{i, j})
			}
		}
	}
	return balls
}

func isPointABall(point string) bool {
	_, err := strconv.Atoi(point)
	return err == nil
}
