package speed

import (
	"bridge/env"
	"reflect"
	"testing"
)

func TestPathToJumpSpeed(t *testing.T) {
	expectedSpeedsToCover := Range{5, 4, 3, 2, 1}
	testEnv := env.Init(0, 0, 16, 4, 8)

	speedsToCover := *MotorbikeSpeedRange(testEnv.GetMotorbike(), testEnv.GetBridge())
	if !reflect.DeepEqual(speedsToCover, expectedSpeedsToCover) {
		t.Fatalf("Expected speed of %d but got %d", expectedSpeedsToCover, speedsToCover)
	}
}

func TestPathToSlow(t *testing.T) {
	expectedSpeedsToCover := Range{6, 5}
	testEnv := env.Init(0, 6, 16, 4, 8)

	speedsToCover := *MotorbikeSpeedRange(testEnv.GetMotorbike(), testEnv.GetBridge())
	if !reflect.DeepEqual(speedsToCover, expectedSpeedsToCover) {
		t.Fatalf("Expected speed of %d but got %d", expectedSpeedsToCover, speedsToCover)
	}
}

func TestMotoMaxSpeedShouldBe1MoreThanGapSize(t *testing.T) {
	testEnv := env.Init(0, 0, 6, 4, 8)
	testEnv.GetMotorbike().SetMaxSpeed(testEnv.GetBridge().GapLength)
	expectedSpeed := 5
	if expectedSpeed != testEnv.GetMotorbike().GetMaxSpeed() {
		t.Fatalf("Expected speed of %d but got %d", expectedSpeed, testEnv.GetMotorbike().GetMaxSpeed())
	}
}

func TestMotoMaxSpeedShouldBe1MoreThanGapSize2(t *testing.T) {
	testEnv := env.Init(0, 0, 6, 2, 8)
	testEnv.GetMotorbike().SetMaxSpeed(testEnv.GetBridge().GapLength)
	expectedSpeed := 3
	if expectedSpeed != testEnv.GetMotorbike().GetMaxSpeed() {
		t.Fatalf("Expected speed of %d but got %d", expectedSpeed, testEnv.GetMotorbike().GetMaxSpeed())
	}
}
