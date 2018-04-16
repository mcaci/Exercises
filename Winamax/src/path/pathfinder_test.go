package path

import "testing"

func TestFindPath1(t *testing.T) {
	testPath([]string {"1H"}, []string{">."}, t)
}

func TestFindReversePath1(t *testing.T) {
	testPath([]string {"H1"}, []string{".<"}, t)
}

func TestFindPath2(t *testing.T) {
	testPath([]string {"1","H"}, []string{"v","."}, t)
}

func TestFindReversePath2(t *testing.T) {
	testPath([]string {"H","1"}, []string{".","^"}, t)
}

func TestFindPath3(t *testing.T) {
	testPath([]string {".1",".H"}, []string{".v",".."}, t)
}

func TestFindPath4(t *testing.T) {
	testPath([]string {"11","HH"}, []string{"vv",".."}, t)
}

func TestFindPath5(t *testing.T) {
	testPath([]string {"H1","1H"}, []string{".<",">."}, t)
}

func TestFindPath6(t *testing.T) {
	testPath([]string {"2.",".H"}, []string{"v.",">."}, t)
}

func test2Paths(t *testing.T) {
	testPath([]string {"2.X", "..H", ".H1"}, []string{"v..", "v..", ">.^"}, t)
}

func testPath(golfCourseMap, expectedGolfCourse []string, t *testing.T) {
	golfMapWithBallsMovement := FindPath(golfCourseMap);
	if !equal(expectedGolfCourse, golfMapWithBallsMovement) {
		t.Log(golfCourseMap)
		t.Fatalf("Expected %s but got %s", expectedGolfCourse, golfMapWithBallsMovement)
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}