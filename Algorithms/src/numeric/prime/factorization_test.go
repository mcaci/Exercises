package prime

import "testing"

func TestFactorizationOf0(t *testing.T) {
	value := 0
	_, specialCase := Factor(0)
	if !specialCase {
		t.Fatalf("%d is a special case", value)
	}
}

