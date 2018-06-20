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

func TestListOfPrimesUpTo4(t *testing.T) {
	verify(t, ErathosthenesSieve(4), []uint{2, 3})
}

func TestListOfPrimesUpTo5(t *testing.T) {
	verify(t, ErathosthenesSieve(5), []uint{2, 3, 5})
}

func TestListOfPrimesUpTo100(t *testing.T) {
	verify(t, ErathosthenesSieve(100), []uint{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97})
}

func verify(t *testing.T, list *[]uint, expectedList []uint) {
	if !reflect.DeepEqual(*list, expectedList) {
		t.Fatalf("List should be composed of %v but is %v", expectedList, list)
	}
}
