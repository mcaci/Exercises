package prime

type Primes map[int]int

func Factor(numberToFactor uint) (*Primes, bool) {
	return new(Primes), numberToFactor < 2
}
