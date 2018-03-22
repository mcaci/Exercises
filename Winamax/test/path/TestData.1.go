package path

var mapExample = map[string]string{
	"1" : "",
	"X" : ".",
}

var GolfCourseSize1L = 2
var GolfCourseSize1H = 1
var GolfCourse1 = []string {"1H"}

var GolfCourseSize1_1L = 2
var GolfCourseSize1_1H = 2
var GolfCourse1_1 = []string {".H", "2."}

var GolfCourseSize2L = 3
var GolfCourseSize2H = 3
var golfCourse2_1 = "2.X"
var golfCourse2_2 = "..H"
var golfCourse2_3 = ".H1"

var GolfCourse2 = []string {golfCourse2_1, golfCourse2_2, golfCourse2_3}

func mapIndex(x, y, mapWidth int) int {
	return x * mapWidth + y
}

func reverseMapIndex(mapIndex, mapWidth int) (x, y int) {
	x = mapIndex / mapWidth
	y = mapIndex % mapWidth
	return x, y 
}

// v..
// v..
// >.^