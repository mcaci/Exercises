package singlelinked

import "testing"

func TestAddItemIsAddedCorrectly(t *testing.T) {
	var list SingleLinkedList
	addedItem := 10
	list.AddSequence(addedItem)
	item, _ := list.Get(0)
	if addedItem != item {
		 t.Fatalf("Expected item to be %d but %d was found", addedItem, item)
	}
}

func TestAddDifferentItemIsAddedCorrectly(t *testing.T) {
	var list SingleLinkedList
	addedItem := 5
	list.AddSequence(addedItem)
	item, _ := list.Get(0)
	if addedItem != item {
		 t.Fatalf("Expected item to be %d but %d was found", addedItem, item)
	}
}

func TestAddTwoElementsAndGetFirst(t *testing.T) {
	var list SingleLinkedList
	addedItem := 5
	list.AddSequence(addedItem, 10)
	item, _ := list.Get(0)
	if addedItem != item {
		 t.Fatalf("Expected item to be %d but %d was found", addedItem, item)
	}
}

func TestAddTwoElementsAndGetSecond(t *testing.T) {
	var list SingleLinkedList
	addedItem := 10
	list.AddSequence(5, addedItem)
	item, _ := list.Get(1)
	if addedItem != item {
		 t.Fatalf("Expected item to be %d but %d was found", addedItem, item)
	}
}

func TestAddThreeElementsAndGetThird(t *testing.T) {
	var list SingleLinkedList
	addedItem := 10
	list.AddSequence(5, 5, addedItem)
	item, _ := list.Get(2)
	if addedItem != item {
		 t.Fatalf("Expected item to be %d but %d was found", addedItem, item)
	}
}