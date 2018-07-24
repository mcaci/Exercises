package singlelinked

import "testing"

func TestSizeShrinksWhenItemIsRemovedCorrectly(t *testing.T) {
	var list SLL
	list.AddSequence([]int {1, 2, 3}...)
	list.Remove(0)
	size := list.counter
	expectedSize := 2
	if expectedSize != size {
		 t.Fatalf("Expected item to be %d but %d was found", expectedSize, size)
	}
}

// TODO: Failing test
func TestItemIsRemovedCorrectly(t *testing.T) {
	var list SLL
	list.AddSequence([]int {1, 2, 3}...)
	list.Remove(0)
	item, _ := list.Get(0)
	expectedItem := 2
	if expectedItem != item {
		 t.Fatalf("Expected item to be %d but %d was found", expectedItem, item)
	}
}

func (sll *SLL) Remove(position int) {
	var prevNode *SLLCell
	var node *SLLCell
	if sll.counter > 0 {
		if position > 1 {
			prevNode = sll.iterateUntil(func (index int, node *SLLCell) bool {return index == position - 1})
			node = sll.iterateUntil(func (index int, node *SLLCell) bool {return index == position})
		} else {
			prevNode = sll.first
			node = sll.first.next
		}
		prevNode.next = node.next
		sll.counter--
	}
}