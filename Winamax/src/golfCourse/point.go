package golfCourse

import "strconv"

type Coordinate struct {
	X, Y int 
}

func FindStart(golfCourseMap []string) (*Coordinate, bool) {
	return findPoint(golfCourseMap, isPointANumber)
}

func FindEnd(golfCourseMap []string) (*Coordinate, bool) {
	return findPoint(golfCourseMap, isPointAHole)
}

func findPoint(golfCourseMap []string, pointPredicate func(string) bool) (*Coordinate, bool) {
	H := len(golfCourseMap)
	for i := 0; i < H; i++ {
		for j := 0; j < len(golfCourseMap[i]); j++ {
			if pointPredicate(string(golfCourseMap[i][j])) {
				return &Coordinate{i, j}, true
			}
		}
	}
	return &Coordinate{-1, -1}, false
}

func isPointAHole(point string) bool {
	return point == "H"
}

func isPointANumber(point string) bool {
	_, err := strconv.Atoi(point)
	return err == nil
}