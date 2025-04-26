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
	ddl := DummyLinkedList2{
		head: &DummyNode{
			isDummy: true,
		},
		tail: &DummyNode{
			isDummy: true,
		},
	}

	ddl.head.next = ddl.tail
	ddl.tail.prev = ddl.head

	return ddl
}

// t = O(n), mem = O(n), where n = len(values), because we are using result
// with struct of DLL2
func GetDLL2(values []int) DummyLinkedList2 {
	var result DummyLinkedList2 = NewDLL2()
	for _, item := range values {
		result.AddInTail(DummyNode{
			value: item,
		})
	}

	return result
}

// head -> [dummy] -> node1 -> ... -> nodeN -> [dummy] <- tail
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

// t = O(n), where n = len(dll2), mem = O(1)
func (dll *DummyLinkedList2) Insert(after *DummyNode, add DummyNode) {
	if dll.Count() == 0 {
		dll.AddInTail(add)
		return
	}

	if add.isDummy {
		return
	}

	temp := dll.head

	for temp != nil {
		if !temp.isDummy && temp.value == after.value {
			break
		}
		temp = temp.next
	}

	if temp == nil {
		return
	}

	next := temp.next
	temp.next = &add
	next.prev = &add
	add.next = next
	add.prev = temp

}

func (dll DummyLinkedList2) InsertFirst(add DummyNode) {
	if add.isDummy {
		return
	}

	next := dll.head.next
	dll.head.next = &add
	next.prev = &add
	add.next = next
	add.prev = dll.head

}

// t = O(n), where n = len(dll2), mem = O(1)
func (dll *DummyLinkedList2) Delete(n int, all bool) {
	if dll.Count() == 0 {
		return
	}

	temp := dll.head

	for temp != nil {
		prev := temp.prev
		next := temp.next

		if temp.value == n {
			prev.next = next
			next.prev = prev

			if !all {
				return
			}
		}
		temp = temp.next
	}
}

// t = O(n), where n = size of dll, mem = O(1)
func (dll *DummyLinkedList2) Count() int {
	count := 0

	tempNode := dll.head

	for tempNode != nil {
		if !tempNode.isDummy {
			count++
		}
		tempNode = tempNode.next
	}
	return count
}

// t = O(1), mem = O(1)
func EqualDLL2(dll2_1, dll2_2 DummyLinkedList2) bool {
	size_1, size_2 := dll2_1.Count(), dll2_2.Count()
	if size_1 != size_2 {
		return false
	}

	temp_1, temp_2 := dll2_1.head, dll2_2.head

	for temp_1 != dll2_1.tail {
		if temp_1.value != temp_2.value {
			return false
		}
		temp_1 = temp_1.next
		temp_2 = temp_2.next
	}

	return true
}
