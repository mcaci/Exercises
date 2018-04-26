package motorbike

import "testing"
import "reflect"

func TestPathToJumpSpeeds(t *testing.T) {
	testMoto := newMotorbike(0, 0)
	testBridge := bridge{17, 4, 15}

	// speedsToCover := getSpeedsToCover(testMoto, &testBridge)
	speeds := determineSpeeds(testMoto, &testBridge)
	expectedSpeeds := make(map[int]int)
	expectedSpeeds[1] = 2
	expectedSpeeds[2] = 1
	expectedSpeeds[3] = 1
	expectedSpeeds[4] = 1
	expectedSpeeds[5] = 1
	if !reflect.DeepEqual(speeds, expectedSpeeds) {
		t.Fatalf("Expected speed of %d but got %d", expectedSpeeds, speeds)
	}
}

func TestPathToSlows(t *testing.T) {
	testMoto := newMotorbike(0, 6)
	testBridge := bridge{17, 4, 15}

	// speedsToCover := getSpeedsToCover(testMoto, &testBridge)
	speeds := determineSpeeds(testMoto, &testBridge)
	expectedSpeeds := make(map[int]int)
	expectedSpeeds[5] = 2
	expectedSpeeds[6] = 1
	if !reflect.DeepEqual(speeds, expectedSpeeds) {
		t.Fatalf("Expected speed of %d but got %d", expectedSpeeds, speeds)
	}
}

func TestPath3(t *testing.T) {
	testMoto := newMotorbike(0, 4)
	testBridge := bridge{10, 3, 15}

	// speedsToCover := getSpeedsToCover(testMoto, &testBridge)
	speeds := determineSpeeds(testMoto, &testBridge)
	expectedSpeeds := make(map[int]int)
	expectedSpeeds[5] = 2
	expectedSpeeds[6] = 1
	if !reflect.DeepEqual(speeds, expectedSpeeds) {
		t.Fatalf("Expected speed of %d but got %d", expectedSpeeds, speeds)
	}
}

func determineSpeeds(m *motorbike, b *bridge) map[int]int {
	speeds := make(map[int]int)
	distanceToCover := b.RoadLength - 1
	speedsToCover := getSpeedsToCover(m, b)

	for _, speed := range speedsToCover {
		distanceToCover -= speed
	}

	for i, _ := range speedsToCover {
		reverseIndex := len(speedsToCover) - 1 - i
		speed := speedsToCover[reverseIndex]
		howManyTimesItKeepsTheSpeed := distanceToCover / speed
		distanceToCover = distanceToCover % speed
		speeds[speed] = 1 + howManyTimesItKeepsTheSpeed
	}
	return speeds
}
