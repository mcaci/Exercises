package singlelinked

import "errors"

type SingleLinkedList struct {
	counter int
	first *SingleLinkedListCell
}

type SingleLinkedListCell struct {
	value int
	next *SingleLinkedListCell
}

func (sll *SingleLinkedList) Add(cell int) {
	sll.counter++;
	if sll.first == nil {
		sll.first = new(SingleLinkedListCell)
		sll.first.value = cell
	} else {
		node := sll.first
		for node.next != nil {
			node = node.next
		}
		node.next = new(SingleLinkedListCell)
		node.next.value = cell
	}
}

func (sll *SingleLinkedList) Get(position int) (int, error) {
	var err error
	var value int
	if sll.counter == 0 {
		err = errors.New("Cannot get from empty list")
	} else if position >= sll.counter {
		err = errors.New("Position requested is out of bounds")
	} else {
		node := sll.first
		for i := 0; i < position; i++ {
			node = node.next
		}
		value = node.value
	}
	return value, err	
}