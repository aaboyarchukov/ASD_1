package main

import (
	"constraints"
	"errors"
	_ "os"
)

type Node[T constraints.Ordered] struct {
	prev  *Node[T]
	next  *Node[T]
	value T
}

type OrderedList[T constraints.Ordered] struct {
	head       *Node[T]
	tail       *Node[T]
	_ascending bool
}

func (l *OrderedList[T]) Count() int {
	count := 0
	tempNode := l.head

	for tempNode != nil {
		count++
		tempNode = tempNode.next
	}

	return count
}

func (l *OrderedList[T]) Add(item T) {
	node := Node[T]{
		value: item,
	}

	left := l.head
	right := l.tail

	if node.value < left.value {
		left.prev = &node
		l.head = left.prev
		l.head.next = left
		return
	}

	if node.value > right.value {
		right.next = &node
		l.tail = &node
		node.prev = right
		return
	}

	for left != right {
		if l.Compare(node.value, left.value) == -1 {
			nextToLeft := left.next
			left.next = &node
			node.next = nextToLeft
			nextToLeft.prev = &node
			node.prev = left
			return
		}

		if l.Compare(node.value, left.value) == 1 {
			prevToRight := right.prev
			left.next = &node
			node.next = prevToRight
			prevToRight.prev = &node
			node.prev = left
			return
		}

		left = left.next
		right = right.prev
	}
}

func (l *OrderedList[T]) Find(n T) (Node[T], error) {
	tempNode := l.head
	for tempNode != nil {
		if tempNode.value == n {
			return *tempNode, nil
		}
		tempNode = tempNode.next
	}

	return Node[T]{value: n, next: nil, prev: nil}, errors.New("node is not finding")
}

func (l *OrderedList[T]) Delete(n T) {
	if l.head == nil {
		return
	}

	if l.Count() == 1 {
		l.Clear(false)
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

		if deleted {
			break
		}
		tempNode = tempNode.next
	}
}

func (l *OrderedList[T]) Clear(asc bool) {
	l.head = nil
	l.tail = nil
	l._ascending = asc
}

func (l *OrderedList[T]) GetArray() []T {
	result := make([]T, 0, l.Count())

	tempNode := l.head

	for tempNode != nil {
		result = append(result, tempNode.value)
		tempNode = tempNode.next
	}

	return result
}

func (l *OrderedList[T]) Compare(v1 T, v2 T) int {

	if v1 < v2 {
		return -1
	}
	if v1 > v2 {
		return +1
	}
	return 0
}

func GetOrderedList[T constraints.Ordered](asc bool, values []T) *OrderedList[T] {
	var l OrderedList[T]
	l._ascending = asc

	for _, item := range values {
		l.Add(item)
	}

	return &l
}
