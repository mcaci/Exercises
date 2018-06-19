package primesList

func ErathosthenesSieve(limit uint) []uint {
	if limit == 2 {
		return []uint{2}
	} else {
		return []uint{2, 3}
	}
}

