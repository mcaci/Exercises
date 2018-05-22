package lcg

import "testing"

func TestNextNumberWithA10B10M10StartingWith1(t *testing.T) {
	testlcg := new(LCG)
	testlcg.Init(1, 10, 10, 10)
	testNumberCorrectness(t, 10, testlcg)
}

func TestNextOfNextNumberWithA10B10M10StartingWith1(t *testing.T) {
	testlcg := new(LCG)
	testlcg.Init(1, 10, 10, 10)
	testlcg.Next()
	testNumberCorrectness(t, 100, testlcg)
}

func testNumberCorrectness(t *testing.T, expectedNextNumber int, lcg *LCG) {
	actualNextNumber := lcg.Next()
	if actualNextNumber != expectedNextNumber {
		t.Fatalf("Expected next number to be %d but %d was computed", expectedNextNumber, actualNextNumber)
	}
}
