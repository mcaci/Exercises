package prime

type Primes []uint

var primeFactorsList = Primes{2, 3, 5}

func Factor(numberToFactor uint) (*Primes, bool) {
	var primes Primes
	var isSpecialCase = numberToFactor < 2
	var number = numberToFactor
	if !isSpecialCase {
		extractFactors(&primes, &numberToFactor, 2)
		for primeFactor := uint(3); primeFactor <= number; primeFactor += 2 {
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
