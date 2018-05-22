package lcg

type linearCongruentialGenerator struct {
	seed, a, b, m, currentNumber int
}

type LCG linearCongruentialGenerator

func (lcg *LCG) NextRandomNumber() int {
	return lcg.seed*lcg.a + lcg.b%lcg.m
}
