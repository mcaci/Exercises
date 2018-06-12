package prime

import (
	"reflect"
	"testing"
)

func TestFactorizationOf0isNil(t *testing.T) {
	verifyThatFactorsAreExpectedOnes(t, 0, nil)
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

func TestFactorizationOf5is5(t *testing.T) {
	verifyThatFactorsAreExpectedOnes(t, 5, []uint{5})
}

func TestFactorizationOf6is2times3(t *testing.T) {
	verifyThatFactorsAreExpectedOnes(t, 6, []uint{2, 3})
}

func verifyThatFactorsAreExpectedOnes(t *testing.T, value uint, expectedFactors []uint) {
	factors, _ := Factor(value)
	if !reflect.DeepEqual(*factors, Primes(expectedFactors)) {
		t.Fatalf("Expected factors are %v but got %v", expectedFactors, *factors)
	}
}
