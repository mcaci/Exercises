package motorbike

import "testing"
import "reflect"
import "math"

func TestPathToJumpSpeed(t *testing.T) {
	testMoto := newMotorbike(0, 0)
	testBridge := bridge{16, 4, 8}

	speedsToCover := getSpeedsToCover(testMoto, &testBridge)
	expectedSpeedsToCover := []int {1, 2, 3, 4, 5}
	if !reflect.DeepEqual(speedsToCover, expectedSpeedsToCover) {
		t.Fatalf("Expected speed of %d but got %d", expectedSpeedsToCover, speedsToCover)
	}
}

func TestPathToSlow(t *testing.T) {
	testMoto := newMotorbike(0, 6)
	testBridge := bridge{16, 4, 8}

	speedsToCover := getSpeedsToCover(testMoto, &testBridge)
	expectedSpeedsToCover := []int {1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(speedsToCover, expectedSpeedsToCover) {
		t.Fatalf("Expected speed of %d but got %d", expectedSpeedsToCover, speedsToCover)
	}
}

func getSpeedsToCover(m *motorbike, b *bridge) []int {
	baseSpeed := 1
	highSpeed := int(math.Max(float64(b.GapLength + 1), float64(m.Speed)))
	// maxSpeed := b.GapLength + 1

	speedsSize := highSpeed - baseSpeed + 1
	speeds := make([]int, speedsSize)
	for i := 0; i < speedsSize; i++ {
		speeds[i] = baseSpeed + i
	}
	return speeds
}
