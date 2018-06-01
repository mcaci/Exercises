package gcd

import "testing"

func TestGcdOf6And10(t *testing.T) {
	testGcd(t, 6, 10, 2)
}

func TestGcdOf10And15(t *testing.T) {
	testGcd(t, 10, 15, 5)
}

func TestGcdOf75And81(t *testing.T) {
	testGcd(t, 72, 81, 9)
}

func testGcd(t *testing.T, a, b, expectedGcd int) {
	gcd := GCD(a, b)
	if expectedGcd != gcd {
		t.Fatalf("Expected result to be %d but %d was computed", expectedGcd, gcd)
	}
}
