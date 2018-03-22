package path

import "testing"

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

func findPath(golfCourse []string) []string {
	goRight := ">"
	goLeft := "<"
	goUp := "^"
	goDown := "v"
	hole := "."

	start := "1"
	end := "H"

	startX, startY := findPoint(golfCourse, start)
	endX, endY := findPoint(golfCourse, end)

	var direction string
	X := startX;
	Y := startY;

	var sequence []string
	
	for X != endX || Y != endY {
		direction, X, Y = getDirection(X, Y, endX, endY)
		sequence = append(sequence, direction)
	}

	var golfMap []string
	mapElement := "."
	for i := 0; i < len(golfCourse); i++ {
		golfLine := ""
		for j := 0; j < len(golfCourse[i]); j++ {
			golfLine += mapElement
		}
		golfMap = append(golfMap, golfLine)
	}

	replacement := mapElement
	str[:index] + string(replacement) + str[index+1:]
	switch sequence[0] {
	case "N":
		replacement := 
		golfMap = []string{hole, goUp}
	case "S":
		golfMap = []string{goDown, hole}
	case "E":
		golfMap = []string{goRight + hole}
	case "W":
		golfMap = []string{hole + goLeft}
	default:
	}
	return golfMap
}

func getDirection(startX, startY, endX, endY int) (string, int, int) {
	var direction string
	if startX == endX && endY == startY + 1 {
		direction = "E"
		startY++
	} else if startX == endX && endY == startY - 1 {
		direction = "W"
		startY--
	} else if endX == startX + 1 && endY == startY {
		direction = "S"
		startX++
	} else if endX == startX - 1 && endY == startY {
		direction = "N"
		startX--
	}
	return direction, startX, startY
}

func findPoint(golfCourse []string, point string) (int, int) {
	H := len(golfCourse)
	for i := 0; i < H; i++ {
		for j := 0; j < len(golfCourse[i]); j++ {
			if point == string(golfCourse[i][j]) {
				return i, j
			}
		}
	}
	return -1, -1
}