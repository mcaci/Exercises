package golfCourse 

type GolfElement interface {
	X() int
	Y() int
}

func AreInSamePosition(a, b GolfElement) bool {
	return a.X() == b.X() && a.Y() == b.Y()
}