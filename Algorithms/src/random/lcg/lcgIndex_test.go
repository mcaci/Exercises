package lcg

import "testing"

func TestIndexLCG(t *testing.T) {
	testlcg := new(LCG)
	seed := 2
	current := seed - 1
	m := 10
	testlcg.Init(seed, 1, 3, m)
	index := 0
	for ; seed != current && index < 100; index++ {
		current = testlcg.Next()
		t.Logf("Current number is %d", current)
	}
	t.Logf("Index is %f", float64(index)/float64(m))
}

func TestIndexLCG2(t *testing.T) {
	testlcg := new(LCG)
	seed := 1
	current := seed - 1
	m := 6
	testlcg.Init(seed, 1, 4, m)
	index := 0
	for ; seed != current && index < 100; index++ {
		current = testlcg.Next()
		t.Logf("Current number is %d", current)
	}
	t.Logf("Index is %f", float64(index)/float64(m))
}
