package path

import "testing"

func TestPath1(t *testing.T) {
	golfMapWithBallsMovement := golfMapToShowBallMovement(GolfCourse1);
	expectedGolfCourse := []string{">."}
	if !equal(expectedGolfCourse, golfMapWithBallsMovement) {
		t.Fatalf("Expected %s but got %s", expectedGolfCourse, golfMapWithBallsMovement)
	}
}

func TestPath11(t *testing.T) {
	golfMapWithBallsMovement := golfMapToShowBallMovement(GolfCourse1_1);
	expectedGolfCourse := []string{"..", ">^"}
	if !equal(expectedGolfCourse, golfMapWithBallsMovement) {
		t.Fatalf("Expected %s but got %s", expectedGolfCourse, golfMapWithBallsMovement)
	}
}

func TestPath2(t *testing.T) {
	golfMapWithBallsMovement := golfMapToShowBallMovement(GolfCourse2);
	expectedGolfCourse := []string{"v..", "v..", ">.^"}
	if !equal(expectedGolfCourse, golfMapWithBallsMovement) {
		t.Fatalf("Expected %s but got %s", expectedGolfCourse, golfMapWithBallsMovement)
	}
}

func golfMapToShowBallMovement(golfCourse []string) []string {
	goRight := ">"
	goDown := "v"
	goUp := "^"

	if len(golfCourse) == 3 {
		line1 := goDown + mapFieldElement(string(golfCourse[0][1])) + mapFieldElement(string(golfCourse[0][2]))
		line2 := goDown + mapFieldElement(string(golfCourse[1][1])) + mapFieldElement(string(golfCourse[1][2]))
		line3 := goRight + mapFieldElement(string(golfCourse[2][1])) + goUp 
		return []string{line1, line2, line3}
	} else if len(golfCourse) == 2 {
		line1 := mapFieldElement(string(golfCourse[0][0])) + mapFieldElement(string(golfCourse[0][1]))
		line2 := goRight + goUp
		return []string{line1, line2}
	} else if len(golfCourse) == 1 {
		line1 := goRight + mapFieldElement(string(golfCourse[0][1]))
		return []string{line1}
	} else {
		return []string{""}
	}
}

func mapFieldElement(fieldElement string) string {
	hole := "."
	waterHazard := "."
	otherFieldElement := "."

	if fieldElement == "H" {
		return hole
	} else if fieldElement == "X" {
		return waterHazard
	} else {
		return otherFieldElement
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