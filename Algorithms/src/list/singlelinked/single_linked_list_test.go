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
	counter int
}

func (sll *SingleLinkedList) Add(cell int) {
	sll.counter++;
}

func (sll *SingleLinkedList) Get(position int) (int, error) {
	if sll.counter > 0 {
		return 0, nil
	} else {
		return 0, errors.New("Cannot get from empty list")
	}
}