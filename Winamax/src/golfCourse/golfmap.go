package golfCourse

func CopyToEmptyGolfCourseMap(golfCourseMap []string) []string {
	var emptyGolfMap []string
	for i := 0; i < len(golfCourseMap); i++ {
		emptyGolfMap = append(emptyGolfMap, createMapLine(golfCourseMap[i]))
	}
	return emptyGolfMap
}

func createMapLine(golfCourseMapLine string) string {
	golfLine := ""
	mapElement := "."
	for i := 0; i < len(golfCourseMapLine); i++ {
		golfLine += mapElement
	}
	return golfLine
}