package prime

import (
	"testing"
)

func TestFactorizationOf0isNil(t *testing.T) {
	verifyThatFactorsAreExpectedOnes(t, 0, []uint{})
}

func TestFactorizationOf2is2(t *testing.T) {
	verifyThatFactorsAreExpectedOnes(t, 2, []uint{2})
}

func TestFactorizationOf3is3(t *testing.T) {
	verifyThatFactorsAreExpectedOnes(t, 3, []uint{3})
}

func TestFactorizationOf4is2times2(t *testing.T) {
	verifyThatFactorsAreExpectedOnes(t, 4, []uint{2, 2})
}

func verifyThatFactorsAreExpectedOnes(t *testing.T, value uint, expectedFactors []uint) {
	factors, _ := Factor(value)
	for i, factor := range *factors {
		if factor != expectedFactors[i] {
			t.Fatalf("Expected factors are %v but got %v", expectedFactors, *factors)
		}
	}
}
