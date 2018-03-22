package path

import ( 
	"testing"
	"golfCourse"
)

func TestFindPath1(t *testing.T) {
	testPath(GolfCourse1Base, []string{">."}, t)
}

func TestFindReversePath1(t *testing.T) {
	testPath(GolfCourse1Rev, []string{".<"}, t)
}

func TestFindPath2(t *testing.T) {
	testPath(GolfCourse2Base, []string{"v","."}, t)
}

func TestFindReversePath2(t *testing.T) {
	testPath(GolfCourse2Rev, []string{".","^"}, t)
}

func testFindPath3(t *testing.T) {
	testPath(GolfCourse2Base, []string{".v",".."}, t)
}

func testPath(golfCourse, expectedGolfCourse []string, t *testing.T) {
	golfMapWithBallsMovement := findPath(golfCourse);
	if !equal(expectedGolfCourse, golfMapWithBallsMovement) {
		t.Fatalf("Expected %s but got %s", expectedGolfCourse, golfMapWithBallsMovement)
	}
}

func findPath(golfCourseMap []string) []string {
	startX, startY, _ := golfCourse.FindStart(golfCourseMap)
	endX, endY, _ := golfCourse.FindEnd(golfCourseMap)
	sequence := golfCourse.PathFromBallToHole(startX, startY, endX, endY)
	golfMap := golfCourse.CopyToEmptyGolfCourseMap(golfCourseMap)
	return replacePathInMap(sequence, golfMap)
}

func replacePathInMap(sequence, golfMap []string) []string {
	// str[:index] + string(replacement) + str[index+1:]
	goRight := ">"
	goLeft := "<"
	goUp := "^"
	goDown := "v"
	hole := "."
	mapElement := "."
	for _, line := range sequence {
		switch line {
		case "N":
			mapElement = goUp
			golfMap = []string{hole, mapElement}
		case "S":
			mapElement = goDown
			golfMap = []string{mapElement, hole}
		case "E":
			mapElement = goRight
			golfMap = []string{mapElement + hole}
		case "W":
			mapElement = goLeft
			golfMap = []string{hole + mapElement}
		default:
		}
	}
	return golfMap
}