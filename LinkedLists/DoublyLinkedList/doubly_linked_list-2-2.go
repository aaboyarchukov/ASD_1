package main

// task 9
// Reverse doubly linked list
func (l *LinkedList2) ReverseList() {

}

// task 10
// Cyclic doubly list
func (l *LinkedList2) CyclicList() {

}

// task 11
// Sort doubly linked list
func (l *LinkedList2) SortList() {

}

// task 12
// Sort merged doubly linked list
func (l *LinkedList2) SortMergeList() {

}

// task 13
// Create dummy doubly linked list

type DummyNode struct {
	value int
	next  *DummyNode
	prev  *DummyNode
}

type DummyDoublyList struct {
	head *DummyNode
	tail *DummyNode
}

func GetDummyList(values []int) *DummyDoublyList {
	return nil
}
