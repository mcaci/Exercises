package prime

type Primes []uint

func Factor(numberToFactor uint) (*Primes, bool) {
	var primes Primes
	isSpecialCase := numberToFactor < 2
	if !isSpecialCase {
		primes = append(primes, 2)
	}
	return &primes, isSpecialCase
}
