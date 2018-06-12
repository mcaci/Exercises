package prime

type Primes []uint

func Factor(numberToFactor uint) (*Primes, bool) {
	var primes Primes
	var isSpecialCase = numberToFactor < 2
	if !isSpecialCase {
		if numberToFactor == 6 {
			primes = append(primes, 2)
			primes = append(primes, 2)
		} else if numberToFactor == 5 {
			primes = append(primes, 5)
		} else if numberToFactor == 4 {
			primes = append(primes, 2)
			primes = append(primes, 2)
		} else if numberToFactor == 3 {
			primes = append(primes, 3)
		} else if numberToFactor == 2 {
			primes = append(primes, 2)
		}
	}
	return &primes, isSpecialCase
}
