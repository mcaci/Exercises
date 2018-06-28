package integration

import "math"
import "testing"

func verify(t *testing.T, area, expectedArea float64) {
	if !almostEqual(area, expectedArea) {
		t.Fatalf("Expected result to be %f but %f was computed", expectedArea, area)
	}
}

func testFunction(parameters ...float64) float64 {
	return 0.5 * parameters[0] + 1
}

const float64EqualityThreshold = 1e-2
func almostEqual(a, b float64) bool {
    return math.Abs(a - b) <= float64EqualityThreshold
}