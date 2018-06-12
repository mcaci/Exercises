package prime

type Primes []uint

var primeFactorsList = Primes{2, 3, 5}

func Factor(numberToFactor uint) (*Primes, bool) {
	var primes Primes
	var isSpecialCase = numberToFactor < 2
	if !isSpecialCase {
		for _, primeFactor := range primeFactorsList {
			extractFactors(&primes, &numberToFactor, primeFactor)
		}
	}
	return &primes, isSpecialCase
}

func extractFactors(primes *Primes, numberToFactorRef *uint, primeFactor uint) {
	numberToFactor := *numberToFactorRef
	for numberToFactor%primeFactor == 0 {
		*primes = append(*primes, primeFactor)
		numberToFactor /= primeFactor
	}
}
