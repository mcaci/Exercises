package lcd

import "../gcd"

func LCD(a, b int) int {
	gcd := gcd.GCD(a, b)
	return (a / gcd) * b
}
