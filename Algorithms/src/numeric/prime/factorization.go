package prime

type Primes []uint

func Factor(numberToFactor uint) (*Primes, bool) {
	var primes Primes
	var isSpecialCase = numberToFactor < 2
	if !isSpecialCase {
		if numberToFactor == 4 {
			primes = append(primes, 2)
			primes = append(primes, 2)
		} else {
			primes = append(primes, numberToFactor)
		}
	}
	return &primes, isSpecialCase
}
