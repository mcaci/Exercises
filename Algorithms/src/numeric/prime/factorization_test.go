package prime

import "testing"

func TestFactorizationOf0isASpecialCase(t *testing.T) {
	value := uint(0)
	_, specialCase := Factor(value)
	if !specialCase {
		t.Fatalf("%d is a special case of factorization", value)
	}
}

func TestFactorizationOf2isNotASpecialCase(t *testing.T) {
	value := uint(2)
	_, specialCase := Factor(value)
	if specialCase {
		t.Fatalf("%d is not a special case of factorization", value)
	}
}
