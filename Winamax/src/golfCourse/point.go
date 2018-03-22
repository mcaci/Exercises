package golfCourse

import "strconv"

func FindStart(golfCourseMap []string) (int, int, bool) {
	return findPoint(golfCourseMap, isPointANumber)
}

func FindEnd(golfCourseMap []string) (int, int, bool) {
	return findPoint(golfCourseMap, isPointAHole)
}

func findPoint(golfCourseMap []string, pointPredicate func(string) bool) (int, int, bool) {
	H := len(golfCourseMap)
	for i := 0; i < H; i++ {
		for j := 0; j < len(golfCourseMap[i]); j++ {
			if pointPredicate(string(golfCourseMap[i][j])) {
				return i, j, true
			}
		}
	}
	return -1, -1, false
}

func isPointAHole(point string) bool {
	return point == "H"
}

func isPointANumber(point string) bool {
	_, err := strconv.Atoi(point)
	return err == nil
}