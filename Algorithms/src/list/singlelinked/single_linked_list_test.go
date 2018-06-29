package singlelinked

import "testing"

func TestGetFromEmptyListShouldReturnErrorMessage(t *testing.T) {
	var v SingleLinkedList
	_, err := v.Get(0)
	if err == nil {
		t.Fatal()
	}
}

func TestGetFirstElementFromNonEmptyListShouldNotReturnErrorMessage(t *testing.T) {
	var v SingleLinkedList
	v.Add(10)
	_, err := v.Get(0)
	if err != nil {
		t.Fatalf("Error message was: %s", err.Error())
	}
}

func TestGet1stElementFromListWithSize1ShouldReturnErrorMessage(t *testing.T) {
	var v SingleLinkedList
	v.Add(10)
	_, err := v.Get(1)
	if err == nil {
		t.Fatal()
	}
}

func TestGet5thElementFromListWithSize1ShouldReturnErrorMessage(t *testing.T) {
	var v SingleLinkedList
	v.Add(10)
	_, err := v.Get(5)
	if err == nil {
		t.Fatal()
	}
}