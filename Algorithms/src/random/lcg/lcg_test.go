package lcg

import "testing"

func TestNextNumberWithA10B10M10StartingWith1(t *testing.T) {
	testlcg := new(LCG)
	testlcg.seed = 1
	testlcg.a = 10
	testlcg.b = 10
	testlcg.m = 10

	expectedNextNumber := 10
	actualNextNumber := testlcg.NextRandomNumber()
	if actualNextNumber != expectedNextNumber {
		t.Fatalf("Expected next number to be %d but %d was computed", expectedNextNumber, actualNextNumber)
	}
}
