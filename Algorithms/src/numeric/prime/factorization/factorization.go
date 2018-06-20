package prime

type Primes []uint

func Factor(numberToFactor uint) (*Primes, bool) {
	var primes Primes
	var isSpecialCase = numberToFactor < 2
	if !isSpecialCase {
		var primeFactor uint = 2
		primes = *extractFactors(&numberToFactor, &primeFactor)
		primes = append(primes, (*extractLastFactor(&numberToFactor))...)
	}
	return &primes, isSpecialCase
}

func extractFactors(numberToFactorRef *uint, primeFactorRef *uint) *Primes {
	var primes = *extractFactor(numberToFactorRef, primeFactorRef)
	for *primeFactorRef = 3; (*primeFactorRef) * (*primeFactorRef) <= (*numberToFactorRef); (*primeFactorRef) += 2 {
		primes = append(primes, (*extractFactor(numberToFactorRef, primeFactorRef))...)
	}
	return &primes
}

func extractFactor(numberToFactorRef *uint, primeFactorRef *uint) *Primes {
	var primes Primes;
	for (*numberToFactorRef) % (*primeFactorRef) == 0 {
		primes = append(primes, *addFactor(primeFactorRef)...)
		updateNumberToFactor(numberToFactorRef, primeFactorRef)
	}
	return &primes
}

func extractLastFactor(primeFactorRef *uint) *Primes {
	var primes Primes;
	if (*primeFactorRef) > 1 {
		primes = *addFactor(primeFactorRef)
	}
	return &primes
}

func addFactor(primeFactorRef *uint) *Primes {
	return &Primes{*primeFactorRef}
}

func updateNumberToFactor(numberToFactorRef *uint, primeFactorRef *uint) {
	(*numberToFactorRef) /= (*primeFactorRef)
}