package lcg

import "testing"

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

func indexTest(t *testing.T, lcg *LCG, expectedIndex float64) {
	index := serieIndex(lcg)
	t.Logf("Lcg serie index is %f", index)
	if expectedIndex != index {
		t.Fatalf("Lcg serie index is not %f", expectedIndex)
	}
}