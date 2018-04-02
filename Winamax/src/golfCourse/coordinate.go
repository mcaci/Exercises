package golfCourse 

type Coordinate interface {
	X() int
	Y() int
}

func AreInSamePosition(a, b Coordinate) bool {
	return a.X() == b.X() && a.Y() == b.Y()
}