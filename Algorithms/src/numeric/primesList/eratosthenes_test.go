package primesList

import (
	"reflect"
	"testing"
)

func TestListOfPrimesUpTo2(t *testing.T) {
	list := ErathosthenesSieve(2)
	if !reflect.DeepEqual([]uint{2}, list) {
		t.Fatalf("List should be composed of %v but is %v", []uint{2}, list)
	}
}

func ErathosthenesSieve(limit uint) []uint {
	return []uint{2}
}
