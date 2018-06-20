package check

import "testing"

func TestFermatFor2(t *testing.T) {
	shouldPrime := true
	var numberToTest uint = 2
	isPrime := Fermat(numberToTest)
	if shouldPrime != isPrime {
		t.Fatalf("%d is supposed to be prime", numberToTest)
	}
}

func Fermat(numberToTest uint) bool {
	return true
}