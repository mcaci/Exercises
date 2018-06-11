package prime

type Primes []uint

func Factor(numberToFactor uint) (*Primes, bool) {
	var primes Primes
	var isSpecialCase = numberToFactor < 2
	if !isSpecialCase {
		primes = append(primes, numberToFactor)
	}
	return &primes, isSpecialCase
}
