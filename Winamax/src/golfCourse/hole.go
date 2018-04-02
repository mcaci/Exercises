package golfCourse

type Hole struct {
	x, y int 
}

func FindHoles(golfCourseMap []string) ([](*Hole)) {
	var holes [](*Hole) 
	H := len(golfCourseMap)
	for i := 0; i < H; i++ {
		for j := 0; j < len(golfCourseMap[i]); j++ {
			if isPointAHole(golfCourseMap[i][j]) {
				holes = append(holes, holeSupplier(i, j))
			}
		}
	}
	return holes
}

func isPointAHole(point byte) bool {
	return point == 72 // letter "H"
}

func holeSupplier(x, y int) *Hole {
	return &Hole{x, y}
}

func (hole Hole) X() int {
	return hole.x
}

func (hole Hole) Y() int {
	return hole.y
}