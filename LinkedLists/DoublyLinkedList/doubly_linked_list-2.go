package main

import (
	"errors"
	_ "os"
	_ "reflect"
)

type Node struct {
	prev  *Node
	next  *Node
	value int
}

type LinkedList2 struct {
	head *Node
	tail *Node
}

func (l *LinkedList2) AddInTail(item Node) {
	if l.head == nil {
		l.head = &item
		l.head.next = nil
		l.head.prev = nil
	} else {
		l.tail.next = &item
		item.prev = l.tail

	}
	l.tail = &item
	l.tail.next = nil
}

func (l *LinkedList2) Count() int {
	count := 0
	tempNode := l.head

	for tempNode != nil {
		count++
		tempNode = tempNode.next
	}
	return count
}

func (l *LinkedList2) Find(n int) (Node, error) {
	tempNode := l.head
	for tempNode != nil {
		if tempNode.value == n {
			return *tempNode, nil
		}
		tempNode = tempNode.next
	}
	return Node{value: -1, next: nil}, errors.New("node is not finding")
}

func (l *LinkedList2) FindAll(n int) []Node {
	var nodes []Node
	tempNode := l.head
	for tempNode != nil {
		if tempNode.value == n {
			nodes = append(nodes, *tempNode)
		}
		tempNode = tempNode.next
	}
	return nodes
}

func (l *LinkedList2) Delete(n int, all bool) {
	if l.head == nil {
		return
	}
	tempNode := l.head
	for tempNode != nil {
		deleted := false
		if tempNode.value == n && tempNode == l.head {
			l.head = tempNode.next
			if l.head != nil {
				l.head.prev = nil
			}
			deleted = true
		} else if tempNode.value == n && tempNode == l.tail {
			l.tail = tempNode.prev
			if l.tail != nil {
				l.tail.next = nil
			}
			deleted = true
		} else if tempNode.value == n {
			prevNode := tempNode.prev
			nextNode := tempNode.next
			prevNode.next = nextNode
			nextNode.prev = prevNode
			deleted = true
		}
		if !all && deleted {
			return
		}

		tempNode = tempNode.next
	}
}

func (l *LinkedList2) Insert(after *Node, add Node) {

}

func (l *LinkedList2) InsertFirst(first Node) {

}

func (l *LinkedList2) Clean() {
	l.head = nil
	l.tail = nil
}

func GetLinkedList(values []int) *LinkedList2 {
	var resultLL LinkedList2 // resulting linked list
	for _, value := range values {
		resultLL.AddInTail(Node{
			value: value,
		})
	}
	return &resultLL
}

func EqualLists(l1 *LinkedList2, l2 *LinkedList2) bool {
	if l1.head == nil &&
		l2.head == nil {
		return true
	}

	if l1.head.value != l2.head.value {
		return false
	}
	if l1.tail.value != l2.tail.value {
		return false
	}

	countL1, countL2 := l1.Count(), l2.Count()
	if countL1 == countL2 {
		tempL1, tempL2 := l1.head, l2.head
		for tempL1 != nil && tempL2 != nil {
			if tempL1.value != tempL2.value {
				return false
			}
			tempL1 = tempL1.next
			tempL2 = tempL2.next
		}

		return true
	}

	return false
}
