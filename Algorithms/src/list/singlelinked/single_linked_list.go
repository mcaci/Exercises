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
		sll.first = createAndAddCell(cell)
	} else {
		node := sll.iterateUntil(func (index int, node *SingleLinkedListCell) bool { return node.next != nil })
		node.next = createAndAddCell(cell)
	}
}

func (sll *SingleLinkedList) Get(position int) (int, error) {
	err := errorCheck(sll.counter, position)
	var value int
	if err == nil {
		node := sll.iterateUntil(func (index int, node *SingleLinkedListCell) bool {return index < position})
		value = node.value
	}
	return value, err	
}

func errorCheck(listSize, position int) error {
	var err error
	if listSize == 0 {
		err = errors.New("Cannot get from empty list")
	} else if position >= listSize {
		err = errors.New("Position requested is out of bounds")
	}
	return err
}

func createAndAddCell(cell int) *SingleLinkedListCell {
	sllCell := new(SingleLinkedListCell)
	sllCell.value = cell
	return sllCell
}

func (sll *SingleLinkedList) iterateUntil(iteratingCondition func (int, *SingleLinkedListCell) bool) *SingleLinkedListCell {
	node := sll.first
	for i := 0; iteratingCondition(i, node); i++ {
		node = node.next
	}
	return node
}
