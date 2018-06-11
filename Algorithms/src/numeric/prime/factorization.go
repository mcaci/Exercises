package prime

type Primes []uint

func Factor(numberToFactor uint) (*Primes, bool) {
	if numberToFactor < 1 {
		return nil, true
	} else {
		primes := Primes{2}
		return &primes, numberToFactor < 2
	}
}
