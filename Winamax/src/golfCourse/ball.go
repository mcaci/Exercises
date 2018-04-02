package golfCourse

type Ball struct {
	x, y int 
}

func FindBalls(golfCourseMap []string) ([](*Ball)) {
	var balls [](*Ball) 
	H := len(golfCourseMap)
	for i := 0; i < H; i++ {
		for j := 0; j < len(golfCourseMap[i]); j++ {
			if isPointABall(golfCourseMap[i][j]) {
				balls = append(balls, ballSupplier(i, j))
			}
		}
	}
	return balls
}

func isPointABall(point byte) bool {
	return point <= 57 && point >= 48 // between 0 and 9
}

func ballSupplier(x, y int) *Ball {
	return &Ball{x, y}
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