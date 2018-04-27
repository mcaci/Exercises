package motorbike

import "testing"
import "reflect"

func TestPathToJumpSpeeds(t *testing.T) {
	testMoto := newMotorbike(0, 0)
	testBridge := bridge{17, 4, 15}

	// speedsToCover := getSpeedsToCover(testMoto, &testBridge)
	speeds := speedMap(testMoto, &testBridge)
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
	speeds := speedMap(testMoto, &testBridge)
	expectedSpeeds := make(map[int]int)
	expectedSpeeds[5] = 2
	expectedSpeeds[6] = 1
	if !reflect.DeepEqual(speeds, expectedSpeeds) {
		t.Fatalf("Expected speed of %d but got %d", expectedSpeeds, speeds)
	}
}

func TestPathWithStartingSpeedMatchingJumpSpeedButNeedsToBeAdjustedDownwards(t *testing.T) {
	testMoto := newMotorbike(0, 4)
	testBridge := bridge{11, 3, 15}

	speeds := speedMap(testMoto, &testBridge)
	expectedSpeeds := make(map[int]int)
	expectedSpeeds[3] = 2
	expectedSpeeds[4] = 1
	if !reflect.DeepEqual(speeds, expectedSpeeds) {
		t.Fatalf("Expected speed of %d but got %d", expectedSpeeds, speeds)
	}
}
