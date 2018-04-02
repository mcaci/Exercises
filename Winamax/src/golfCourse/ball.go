package golfCourse

type Ball struct {
	x, y int 
}

func isPointABall(point byte) bool {
	return point <= 57 && point >= 48 // between 0 and 9
}

func ballSupplier(x, y int) Coordinate {
	return Ball{x, y}
}

func (ball *Ball) Move(direction string) {
	switch direction {
		case "v": ball.IncrX()
		case "^": ball.DecrX()
		case ">": ball.IncrY()
		case "<": ball.DecrY()
	}
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