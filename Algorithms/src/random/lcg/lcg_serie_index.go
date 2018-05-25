package lcg

func serieIndex(lcg *LCG) float64 {
	index := 0
	last := -1
	current := -1
	for ; lcg.seed != current; index++ {
		current = lcg.Next()
		if last == current {
			return 0.0
		}
		last = current
	}
	return float64(index)/float64(lcg.m)
}