package singlelinked

import "testing"

func TestGetFromEmptyListShouldReturnErrorMessage(t *testing.T) {
	verifyOutcome(t, createSingleLinkedList(0), 0, expectedErrorMessagePredicate)
}

func TestGetFirstElementFromNonEmptyListShouldNotReturnErrorMessage(t *testing.T) {
	verifyOutcome(t, createSingleLinkedList(1), 0, func(e error) bool {return e != nil})
}

func TestGet1stElementFromListWithSize1ShouldReturnErrorMessage(t *testing.T) {
	verifyOutcome(t, createSingleLinkedList(1), 1, expectedErrorMessagePredicate)
}

func TestGet5thElementFromListWithSize1ShouldReturnErrorMessage(t *testing.T) {
	verifyOutcome(t, createSingleLinkedList(1), 4, expectedErrorMessagePredicate)
}

func createSingleLinkedList(size int) *SLL {
	var v SLL
	if size == 1 { v.Add(10) }
	return &v
}

func expectedErrorMessagePredicate(err error) bool {return err == nil}

func verifyOutcome(t *testing.T, list *SLL, position int, comparisonPredicate func(error) bool) {
	_, err := list.Get(position)
	if comparisonPredicate(err) { t.Fatal() }
}