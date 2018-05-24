package lcg

import "testing"

func TestIndexLCG1(t *testing.T) {
	testlcg := new(LCG)
	testlcg.Init(2, 1, 3, 10)
	t.Logf("Lcg serie index is %f", serieIndex(testlcg))
	if 1 != serieIndex(testlcg) {
		t.Fatalf("Lcg serie index is not 1")
	}
}

func TestIndexLCG2(t *testing.T) {
	testlcg := new(LCG)
	testlcg.Init(1, 1, 4, 6)
	t.Logf("Lcg serie index is %f", serieIndex(testlcg))
	if 0.5 != serieIndex(testlcg) {
		t.Fatalf("Lcg serie index is not 0.5")
	}
}

func TestIndexLCG3(t *testing.T) {
	testlcg := new(LCG)
	testlcg.Init(1, 10, 40, 1000)
	t.Logf("Lcg serie index is %f", serieIndex(testlcg))
	if 0 != serieIndex(testlcg) {
		t.Fatalf("Lcg serie index is not 0")
	}
}