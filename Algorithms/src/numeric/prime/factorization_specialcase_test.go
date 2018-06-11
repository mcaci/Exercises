package prime

import (
	"testing"
)

func TestFactorizationOf0isASpecialCase(t *testing.T) {
	verifyIsValueIsASpecialCaseOfFactorization(t, 0, func(b bool) bool { return !b })
}

func TestFactorizationOf2isNotASpecialCase(t *testing.T) {
	verifyIsValueIsASpecialCaseOfFactorization(t, 2, func(b bool) bool { return b })
}

func verifyIsValueIsASpecialCaseOfFactorization(t *testing.T, value uint, predicate func(bool) bool) {
	_, specialCase := Factor(value)
	if predicate(specialCase) {
		t.Fatalf("%d is not a special case of factorization", value)
	}
}
