package lcg

type linearCongruentialGenerator struct {
	seed, a, b, m, current int
}

type LCG linearCongruentialGenerator

func (lcg *LCG) Next() int {
	lcg.current = lcg.current*lcg.a + lcg.b%lcg.m
	return lcg.current
}

func (lcg *LCG) Init(seed, a, b, m int) {
	lcg.seed = seed
	lcg.a = a
	lcg.b = b
	lcg.m = m
	lcg.current = seed
}
