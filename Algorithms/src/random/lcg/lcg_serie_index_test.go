package lcg

import "testing"
import "math"

func TestIndexLCG1(t *testing.T) {
	testlcg := new(LCG)
	testlcg.Init(2, 1, 3, 10)
	indexTest(t, testlcg, 1)
}

func TestIndexLCG2(t *testing.T) {
	testlcg := new(LCG)
	testlcg.Init(1, 1, 4, 6)
	indexTest(t, testlcg, 0.5)
}

func TestIndexLCG3(t *testing.T) {
	testlcg := new(LCG)
	testlcg.Init(1, 10, 40, 1000)
	indexTest(t, testlcg, 0)
}

func TestIndexLCG4(t *testing.T) {
	testlcg := new(LCG)
	testlcg.Init(7, 23, 17, 550)
	indexTest(t, testlcg, 0.4)
}

func TestIndexLCG5(t *testing.T) {
	testlcg := new(LCG)
	testlcg.Init(10, 11, 17, 29)
	indexTest(t, testlcg, 0.96)
}

func indexTest(t *testing.T, lcg *LCG, expectedIndex float64) {
	index := serieIndex(lcg)
	t.Logf("Lcg serie index is %f", index)
	if !almostEqual(expectedIndex, index) {
		t.Fatalf("Lcg serie index is not %f", expectedIndex)
	}
}

const float64EqualityThreshold = 1e-2
func almostEqual(a, b float64) bool {
    return math.Abs(a - b) <= float64EqualityThreshold
}