package swap

import (
	"reflect"
	"testing"
)

func TestRandomizationAlgorithm(t *testing.T) {
	testArray := []int{1, 2, 3, 4, 5}
	shuffledArray := Rand(testArray)
	if reflect.DeepEqual(testArray, shuffledArray) {
		t.Fatalf("Array %v has not changed with respect to %v", testArray, shuffledArray)
	}
}
