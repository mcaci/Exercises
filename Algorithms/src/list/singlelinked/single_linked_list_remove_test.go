package singlelinked

import "testing"

func TestItemIsRemovedCorrectly(t *testing.T) {
	var list SLL
	list.AddSequence([]int {1}...)
	list.Remove(0)
	size := list.counter
	expectedSize := 0
	if expectedSize != size {
		 t.Fatalf("Expected item to be %d but %d was found", expectedSize, size)
	}
}

func (sll *SLL) Remove(position int) {

}