package singlelinked

import "errors"

type SingleLinkedList struct {
	counter int
}

func (sll *SingleLinkedList) Add(cell int) {
	sll.counter++;
}

func (sll *SingleLinkedList) Get(position int) (int, error) {
	if sll.counter == 0 {
		return 0, errors.New("Cannot get from empty list")
	} else if position >= sll.counter {
		return 0, errors.New("Position requested is out of bounds")
	} else {
		return 0, nil
	}
}