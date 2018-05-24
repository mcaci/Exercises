package gcd

/**
GCD: gcd implementation with Euclid's Algorithm
*/
func GCD(a, b int) int {
	gcd := a
	for b != 0 {
		remainder := b % a
		gcd = GCD(b, remainder)
		a = b
		b = remainder
	}
	return gcd
}
