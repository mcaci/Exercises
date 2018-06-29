package singlelinked

import "errors"

type SingleLinkedList struct {
	counter int
}

func (sll *SingleLinkedList) Add(cell int) {
	sll.counter++;
}

func (sll *SingleLinkedList) Get(position int) (int, error) {
	cellValue := 0
	var err error
	if sll.counter == 0 {
		err = errors.New("Cannot get from empty list")
	} else if position >= sll.counter {
		err = errors.New("Position requested is out of bounds")
	}
	return cellValue, err	
}