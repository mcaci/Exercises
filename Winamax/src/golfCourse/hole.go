package golfCourse

type Hole struct {
	x, y int 
}

func isPointAHole(point byte) bool {
	return point == 72 // letter "H"
}

func holeSupplier(x, y int) Coordinate {
	return &Hole{x, y}
}

func (hole *Hole) X() int {
	return hole.x
}

func (hole *Hole) Y() int {
	return hole.y
}