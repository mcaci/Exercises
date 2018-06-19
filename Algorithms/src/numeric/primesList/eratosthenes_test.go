package primesList

import (
	"reflect"
	"testing"
)

func TestListOfPrimesUpTo2(t *testing.T) {
	verify(t, ErathosthenesSieve(2), []uint{2})
}

func TestListOfPrimesUpTo3(t *testing.T) {
	verify(t, ErathosthenesSieve(3), []uint{2, 3})
}

// func TestListOfPrimesUpTo4(t *testing.T) {
// 	verify(t, ErathosthenesSieve(4), []uint{2, 3})
// }

func TestListOfPrimesUpTo5(t *testing.T) {
	verify(t, ErathosthenesSieve(5), []uint{2, 3, 5})
}

func verify(t *testing.T, list, expectedList []uint) {
	if !reflect.DeepEqual(list, expectedList) {
		t.Fatalf("List should be composed of %v but is %v", expectedList, list)
	}
}
