package golfCourse

type Hole struct {
	X, Y int 
}

func FindHoles(golfCourseMap []string) ([](*Hole)) {
	var holes [](*Hole) 
	H := len(golfCourseMap)
	for i := 0; i < H; i++ {
		for j := 0; j < len(golfCourseMap[i]); j++ {
			if isPointAHole(string(golfCourseMap[i][j])) {
				holes = append(holes, &Hole{i, j})
			}
		}
	}
	return holes
}

func isPointAHole(point string) bool {
	return point == "H"
}
