package speed

import "testing"
import "reflect"
import "bridge/env"

func TestPathToJumpSpeeds(t *testing.T) {
	testEnv := env.Init(0, 0, 17, 4, 15)

	speeds := *BuildSpeedMap(testEnv.GetMotorbike(), testEnv.GetBridge())
	expectedSpeeds := make(Map)
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
	testEnv := env.Init(0, 6, 17, 4, 15)

	speeds := *BuildSpeedMap(testEnv.GetMotorbike(), testEnv.GetBridge())
	expectedSpeeds := make(Map)
	expectedSpeeds[5] = 2
	expectedSpeeds[6] = 1
	if !reflect.DeepEqual(speeds, expectedSpeeds) {
		t.Fatalf("Expected speed of %d but got %d", expectedSpeeds, speeds)
	}
}

func TestPathWithStartingSpeedMatchingJumpSpeedButNeedsToBeAdjustedDownwards(t *testing.T) {
	testEnv := env.Init(0, 4, 11, 3, 15)

	speeds := *BuildSpeedMap(testEnv.GetMotorbike(), testEnv.GetBridge())
	expectedSpeeds := make(Map)
	expectedSpeeds[3] = 2
	expectedSpeeds[4] = 1
	if !reflect.DeepEqual(speeds, expectedSpeeds) {
		t.Fatalf("Expected speed of %d but got %d", expectedSpeeds, speeds)
	}
}
