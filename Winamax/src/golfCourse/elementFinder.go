package golfCourse

func FindElements(golfCourseMap []string, elementPredicate func (byte) bool, elementSupplier func (int, int) Coordinate) []Coordinate {
	var elements []Coordinate
	H := len(golfCourseMap)
	for i := 0; i < H; i++ {
		for j := 0; j < len(golfCourseMap[i]); j++ {
			if elementPredicate(golfCourseMap[i][j]) {
				elements = append(elements, elementSupplier(i, j))
			}
		}
	}
	return elements
}

func FindBalls(golfCourseMap []string) []Coordinate {
	return FindElements(golfCourseMap, isPointABall, ballSupplier)
}

func FindHoles(golfCourseMap []string) []Coordinate {
	return FindElements(golfCourseMap, isPointAHole, holeSupplier)
}