package main

import (
	"constraints"
	"errors"
	"fmt"
	_ "os"
	"strings"
)

func PrintList[T constraints.Ordered](l OrderedList[T]) {
	fmt.Printf("head: %v, tail: %v\n", l.head.value, l.tail.value)
	node := l.head

	for node != nil {
		fmt.Printf("%v ", node.value)
		node = node.next
	}
	fmt.Println()
}

type Node[T constraints.Ordered] struct {
	prev  *Node[T]
	next  *Node[T]
	value T
}

type OrderedList[T constraints.Ordered] struct {
	head *Node[T]
	tail *Node[T]
	// base       []T
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

	size := l.Count()

	if size == 0 {
		l.head = &node
		l.tail = &node
		return
	}

	compareHead, compareTail := l.Compare(node.value, l.head.value), l.Compare(node.value, l.tail.value)

	if (l._ascending && (compareTail == 1 || compareTail == 0)) ||
		(!l._ascending && (compareTail == -1 || compareTail == 0)) {
		l.tail.next = &node
		node.prev = l.tail
		l.tail = &node
		return
	}
	if (l._ascending && (compareHead == -1 || compareHead == 0)) ||
		(!l._ascending && (compareHead == 1 || compareHead == 0)) {
		node.next = l.head
		l.head.prev = &node
		l.head = &node
		return
	}

	left, right := l.head, l.head.next
	for right != nil {
		compareNodeAndLeft := l.Compare(node.value, left.value)
		compareNodeAndRight := l.Compare(node.value, right.value)

		asc := (compareNodeAndLeft == 1 || compareNodeAndLeft == 0) &&
			(compareNodeAndRight == -1 || compareNodeAndRight == 0)

		desc := (compareNodeAndLeft == -1 || compareNodeAndLeft == 0) &&
			(compareNodeAndRight == 1 || compareNodeAndRight == 0)

		if asc || desc {
			node.prev = left
			node.next = right
			left.next = &node
			right.prev = &node
			return
		}

		left = left.next
		right = right.next
	}
}

func (l *OrderedList[T]) AddAtBase(item T) {

}
func (l *OrderedList[T]) Find(n T) (Node[T], error) {
	tempNode := l.head
	for tempNode != nil {
		if l.Compare(tempNode.value, n) == 0 {
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
		compareNode := l.Compare(tempNode.value, n)
		if compareNode == 0 && tempNode == l.head {
			l.head = tempNode.next
			if l.head != nil {
				l.head.prev = nil
			}
			deleted = true
		} else if compareNode == 0 && tempNode == l.tail {
			l.tail = tempNode.prev
			if l.tail != nil {
				l.tail.next = nil
			}
			deleted = true
		} else if compareNode == 0 {
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

	var valueStr1, valueStr2 string
	flagStr := false

	if value, ok := any(v1).(string); ok {
		valueStr1 = strings.Trim(value, " ")
		flagStr = true
	}

	if value, ok := any(v2).(string); ok {
		valueStr2 = strings.Trim(value, " ")
		flagStr = true
	}

	switch flagStr {
	case false:
		if v1 < v2 {
			return -1
		}
		if v1 > v2 {
			return +1
		}

	case true:
		if valueStr1 < valueStr2 {
			return -1
		}
		if valueStr1 > valueStr2 {
			return +1
		}
	}

	return 0
}

func GetOrderedList[T constraints.Ordered](asc bool, values []T) *OrderedList[T] {
	l := &OrderedList[T]{
		head:       nil,
		tail:       nil,
		_ascending: asc,
	}

	for _, item := range values {
		l.Add(item)
	}

	return l
}

func EqualOrderedLists[T constraints.Ordered](l1 *OrderedList[T], values []T) bool {

	if l1 == nil {
		return false
	}

	if l1.Count() != len(values) {
		return false
	}

	node1 := l1.head

	for indx := 0; node1 != nil; indx++ {
		if l1.Compare(node1.value, values[indx]) != 0 {
			return false
		}

		if node1.next != nil && l1._ascending &&
			l1.Compare(node1.next.value, node1.value) == -1 {
			return false
		}

		if node1.next != nil && !l1._ascending &&
			l1.Compare(node1.next.value, node1.value) == 1 {
			return false
		}

		node1 = node1.next
	}

	return true
}
