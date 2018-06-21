package check

import "testing"

func TestFermatFor4(t *testing.T) {
	verifyNonPrimality(t, 4, false)
}

func TestFermatFor603(t *testing.T) {
	verifyNonPrimality(t, 603, false)
}

func TestFermatFor12543(t *testing.T) {
	verifyPrimality(t, 12547, true)
}

func TestFermatFor52001(t *testing.T) {
	verifyPrimality(t, 52009, true)
}

func TestFermatFor5200110025(t *testing.T) {
	verifyPrimality(t, 5200110025, true)
}

func verifyNonPrimality(t *testing.T, numberToTest uint, shouldPrime bool) {
	if shouldPrime == Fermat(numberToTest) {
		t.Fatalf("%d is not supposed to be prime", numberToTest)
	}
}