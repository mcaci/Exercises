package singlelinked

import "testing"

func TestAddItemIsAddedCorrectly(t *testing.T) {
	verify(t, 10, []int{10}, 0)
}

func TestAddDifferentItemIsAddedCorrectly(t *testing.T) {
	verify(t, 5, []int{5}, 0)
}

func TestAddTwoElementsAndGetFirst(t *testing.T) {
	verify(t, 5, []int{5, 10}, 0)
}

func TestAddTwoElementsAndGetSecond(t *testing.T) {
	verify(t, 10, []int{5, 10}, 1)
}

func TestAddThreeElementsAndGetThird(t *testing.T) {
	verify(t, 10, []int{5, 5, 10}, 2)
}

func verify(t *testing.T, addedItem int, addedSequence []int, indexForGet int) {
	var list SLL
	list.AddSequence(addedSequence...)
	item, _ := list.Get(indexForGet)
	if addedItem != item {
		 t.Fatalf("Expected item to be %d but %d was found", addedItem, item)
	}
}