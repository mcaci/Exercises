package singlelinked

import "errors"

type SingleLinkedListTDD struct {
	counter int
	first *SLLCell
}
type SLLTDD SingleLinkedListTDD

type SingleLinkedListCellTDD struct {
	value int
	next *SLLCell
}
type SLLCellTDD SingleLinkedListCellTDD

func (sll *SLL) Add(cell int) {
	sll.counter++;
	listCell := createCell(cell)
	if sll.first == nil {
		sll.first = listCell
	} else {
		node := sll.iterateUntil(func (index int, node *SLLCell) bool { return node.next != nil })
		node.next = listCell
	}
}

func (sll *SLL) AddSequence(sequence ...int) {
	for _, element := range sequence { sll.Add(element) }
}

func (sll *SLL) Get(position int) (int, error) {
	err := errorCheck(sll.counter, position)
	var value int
	if err == nil {
		node := sll.iterateUntil(func (index int, node *SLLCell) bool {return index < position})
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

func createCell(cell int) *SLLCell {
	sllCell := new(SLLCell)
	sllCell.value = cell
	return sllCell
}

func (sll *SLL) iterateUntil(iteratingCondition func (int, *SLLCell) bool) *SLLCell {
	node := sll.first
	for i := 0; iteratingCondition(i, node); i++ {
		node = node.next
	}
	return node
}