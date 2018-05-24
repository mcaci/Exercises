package lcg

import "testing"

func TestNextNumberWithA10B10M10StartingWith1(t *testing.T) {
	testlcg := new(LCG)
	testlcg.Init(1, 10, 10, 10)
	testNumberCorrectness(t, 0, testlcg)
}

func TestNextNumberWithA3B5M7StartingWith1(t *testing.T) {
	testlcg := new(LCG)
	testlcg.Init(1, 3, 5, 7)
	testNumberCorrectness(t, 1, testlcg)
}

func TestNextOfNextNumberWithA10B10M10StartingWith1(t *testing.T) {
	testlcg := new(LCG)
	testlcg.Init(1, 10, 10, 10)
	testlcg.Next()
	testNumberCorrectness(t, 0, testlcg)
}

func testNumberCorrectness(t *testing.T, expectedNextNumber int, lcg *LCG) {
	actualNextNumber := lcg.Next()
	if actualNextNumber != expectedNextNumber {
		t.Fatalf("Expected next number to be %d but %d was computed", expectedNextNumber, actualNextNumber)
	}
}
