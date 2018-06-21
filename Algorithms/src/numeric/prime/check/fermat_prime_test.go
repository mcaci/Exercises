package check

import "testing"

func TestFermatFor2(t *testing.T) {
	verifyPrimality(t, 2, true)
}

func TestFermatFor3(t *testing.T) {
	verifyPrimality(t, 3, true)
}

func TestFermatFor12547(t *testing.T) {
	verifyPrimality(t, 12547, true)
}

func TestFermatFor52009(t *testing.T) {
	verifyPrimality(t, 52009, true)
}

func TestFermatFor5200952009(t *testing.T) {
	verifyPrimality(t, 5200952009, true)
}

func verifyPrimality(t *testing.T, numberToTest uint, shouldPrime bool) {
	if shouldPrime != Fermat(numberToTest) {
		t.Fatalf("%d is supposed to be prime", numberToTest)
	}
}