package prime

type Primes []uint

func Factor(numberToFactor uint) (*Primes, bool) {
	primes := Primes{2}
	return &primes, numberToFactor < 2
}
