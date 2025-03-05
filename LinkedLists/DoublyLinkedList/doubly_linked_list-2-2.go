package main

// task 9
// Reverse doubly linked list
// t = O(n), where n = len(list), mem = O(1)
func (l *LinkedList2) ReverseList() {
	if l.head == nil || l.Count() == 1 {
		return
	}

	tempNode := l.tail

	for tempNode != l.head {
		tempNode.next, tempNode.prev = tempNode.prev, tempNode.next
		tempNode = tempNode.next
	}
	tempNode.next, tempNode.prev = tempNode.prev, tempNode.next
	l.tail, l.head = l.head, l.tail

}

// task 10
// Cyclic doubly list
// t = O(n), where n = len(list), mem = O(1)
func (l *LinkedList2) CyclicList() bool {
	if l.head == nil {
		return false
	}

	if l.Count() == 1 && l.tail.next == l.head {
		return true
	}

	tempNode := l.head
	prev := tempNode.prev
	for tempNode != l.tail {
		if prev != nil && l.tail.next == prev && prev.prev == l.tail {
			return true
		}
		tempNode = tempNode.next
		prev = tempNode.prev
	}

	return false
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
