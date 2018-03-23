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
	// t.Log(golfCourse)
	if !equal(expectedGolfCourse, golfMapWithBallsMovement) {
		t.Fatalf("Expected %s but got %s", expectedGolfCourse, golfMapWithBallsMovement)
	}
}

func findPath(golfCourseMap []string) []string {
	start, _ := golfCourse.FindStart(golfCourseMap)
	end, _ := golfCourse.FindEnd(golfCourseMap)
	steps := golfCourse.PathFromBallToHole(start, end)
	emptyGolfMap := golfCourse.CopyToEmptyGolfCourseMap(golfCourseMap)
	return replacePathInMap(start, steps, emptyGolfMap)
}

func replacePathInMap(start *golfCourse.Coordinate, sequence, golfMap []string) []string {
	for _, direction := range sequence {
		line := golfMap[start.X]
		golfMap[start.X] = line[:start.Y] + direction + line[start.Y+1:]
	}
	return golfMap
}