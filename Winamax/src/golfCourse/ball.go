package golfCourse

import "strconv"

type Ball struct {
	x, y int 
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

func (ball *Ball) IncrX() {
	ball.x++
}

func (ball *Ball) DecrX() {
	ball.x--
}

func (ball *Ball) IncrY() {
	ball.y++
}

func (ball *Ball) DecrY() {
	ball.y--
}

func (ball Ball) X() int {
	return ball.x
}

func (ball Ball) Y() int {
	return ball.y
}