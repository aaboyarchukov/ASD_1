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
		if prev != nil && l.tail.next == prev {
			return true
		}
		tempNode = tempNode.next
		prev = tempNode.prev
	}

	return false
}

func (l *LinkedList2) GetCyclicList(item Node) {
	if l.head == nil {
		return
	}

	if l.Count() == 1 && l.head.value == item.value {
		l.tail.next = l.head
		return
	}

	tempNode := l.head

	for {
		if tempNode.value == item.value {
			l.tail.next = tempNode
			break
		}

		tempNode = tempNode.next
	}

}

// task 11
// Sort doubly linked list
// t = O(n^2), where n = len(list), mem = O(1)
func (l *LinkedList2) SortList() {
	if l.Count() <= 1 {
		return
	}

	for tempNode := l.head; tempNode != nil; tempNode = tempNode.next {
		for nextNode := l.head; nextNode != l.tail; nextNode = nextNode.next {
			if nextNode.value >= nextNode.next.value {
				nextNode.value, nextNode.next.value = nextNode.next.value, nextNode.value
			}
		}
	}
}

// task 12
// Sort merged doubly linked list
// t = O(n^2), where n = len(l1), mem = O(m), where m = len(l3)
func SortMergeList(l1, l2 *LinkedList2) *LinkedList2 {
	l1.SortList()
	l2.SortList()
	l3 := &LinkedList2{}

	tempNode1, tempNode2 := l1.head, l2.head

	for tempNode1 != nil && tempNode2 != nil {
		if tempNode1.value <= tempNode2.value {
			l3.AddInTail(*tempNode1)
			tempNode1 = tempNode1.next
		} else {
			l3.AddInTail(*tempNode2)
			tempNode2 = tempNode2.next
		}
	}

	for tempNode1 != nil {
		l3.AddInTail(*tempNode1)
		tempNode1 = tempNode1.next
	}

	for tempNode2 != nil {
		l3.AddInTail(*tempNode2)
		tempNode2 = tempNode2.next
	}

	return l3
}

// task 13
// Create dummy doubly linked list

type DummyNode struct {
	prev    *DummyNode
	next    *DummyNode
	value   int
	isDummy bool
}

type DummyLinkedList2 struct {
	head *DummyNode
	tail *DummyNode
}

func NewDLL2() DummyLinkedList2 {
	return DummyLinkedList2{
		head: &DummyNode{
			isDummy: true,
		},
		tail: &DummyNode{
			isDummy: true,
		},
	}
}

func GetDLL2(values []int) DummyLinkedList2 {
	return DummyLinkedList2{}
}

// t = O(1), mem = O(1)
func (dll *DummyLinkedList2) AddInTail(node DummyNode) {
	if node.isDummy {
		return
	}

	prev := dll.tail.prev

	prev.next = &node
	dll.tail.prev = &node

	node.prev = prev
	node.next = dll.tail
}

func EqualDLL2(dll2_1, dll2_2 DummyLinkedList2) bool {
	return true
}
