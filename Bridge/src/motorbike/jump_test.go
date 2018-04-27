package motorbike

import "testing"
import "reflect"

func TestPathToJumpSpeed(t *testing.T) {
	testMoto := newMotorbike(0, 0)
	testBridge := bridge{16, 4, 8}

	speedsToCover := getSpeedsToCover(testMoto, &testBridge)
	expectedSpeedsToCover := []int{5, 4, 3, 2, 1}
	if !reflect.DeepEqual(speedsToCover, expectedSpeedsToCover) {
		t.Fatalf("Expected speed of %d but got %d", expectedSpeedsToCover, speedsToCover)
	}
}

func TestPathToSlow(t *testing.T) {
	testMoto := newMotorbike(0, 6)
	testBridge := bridge{16, 4, 8}

	speedsToCover := getSpeedsToCover(testMoto, &testBridge)
	expectedSpeedsToCover := []int{6, 5}
	if !reflect.DeepEqual(speedsToCover, expectedSpeedsToCover) {
		t.Fatalf("Expected speed of %d but got %d", expectedSpeedsToCover, speedsToCover)
	}
}

