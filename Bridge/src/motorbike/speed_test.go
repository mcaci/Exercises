package motorbike

import "testing"

func TestMotoMaxSpeedShouldBe1MoreThanGapSize(t *testing.T) {
	testMoto := newMotorbike(0, 0)
	testBridge := bridge{6, 4, 8}
	testMoto.SetMaxSpeed(testBridge.GapLength)
	expectedSpeed := 5
	if expectedSpeed != testMoto.GetMaxSpeed() {
		t.Fatalf("Expected speed of %d but got %d", expectedSpeed, testMoto.GetMaxSpeed())
	}
}


func TestMotoMaxSpeedShouldBe1MoreThanGapSize2(t *testing.T) {
	testMoto := newMotorbike(0, 0)
	testBridge := bridge{6, 2, 8}
	testMoto.SetMaxSpeed(testBridge.GapLength)
	expectedSpeed := 3
	if expectedSpeed != testMoto.GetMaxSpeed() {
		t.Fatalf("Expected speed of %d but got %d", expectedSpeed, testMoto.GetMaxSpeed())
	}
}