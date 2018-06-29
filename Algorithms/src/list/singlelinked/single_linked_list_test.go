package singlelinked

import "testing"
import "errors"

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

type SingleLinkedList struct {

}

func (s *SingleLinkedList) Add(cell int) {
	
}

func (s *SingleLinkedList) Get(position int) (int, error) {
	return 0, errors.New("Cannot get from empty list")
}