package singlelinked

type SingleLinkedList struct {
	first *Cell
}
type SLL SingleLinkedList

type Cell struct {
	value int
	next *Cell
}