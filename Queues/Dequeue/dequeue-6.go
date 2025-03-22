package main

import (
	"fmt"
	_ "os"
)

type DequeNode[T any] struct {
	value T
	next  *DequeNode[T]
}

type Deque[T any] struct {
	head *DequeNode[T]
	tail *DequeNode[T]
	size int
}

func (d *Deque[T]) Size() int {
	return d.size
}

func (d *Deque[T]) AddFront(itm T) {
	node := &DequeNode[T]{
		value: itm,
		next:  nil,
	}

	if d.head == nil {
		d.tail = node
	} else {
		node.next = d.head
	}

	d.head = node
	d.size++

}

func (d *Deque[T]) AddTail(itm T) {
	node := &DequeNode[T]{
		value: itm,
		next:  nil,
	}

	if d.tail == nil {
		d.head = node
	} else {
		d.tail.next = node
	}

	d.tail = node

	d.size++
}

func (d *Deque[T]) RemoveFront() (T, error) {
	var result T
	if d.size == 0 {
		return result, fmt.Errorf("dequeue is empty")
	}

	result = d.head.value
	if d.size == 1 {
		d.head, d.tail = nil, nil
	} else {
		currentNode := d.head
		d.head = d.head.next
		currentNode.next = nil
	}

	d.size--

	return result, nil
}

func (d *Deque[T]) RemoveTail() (T, error) {
	var result T
	if d.size == 0 {
		return result, fmt.Errorf("dequeue is empty")
	}

	result = d.tail.value
	if d.size == 1 {
		d.head = nil
		d.tail = nil
	} else {
		tempNode := d.head
		for tempNode.next != d.tail {
			tempNode = tempNode.next
		}
		tempNode.next = nil
		d.tail = tempNode
	}

	d.size--
	return result, nil
}

func GetDequeue[T any](values []T) *Deque[T] {
	var result Deque[T]
	for _, value := range values {
		result.AddTail(value)
	}
	return &result
}

func EqualDequeue[T comparable](d1 *Deque[T], d2 *Deque[T]) bool {
	if d1.size != d2.size {
		return false
	}

	if d1 == nil && d2 == nil {
		return true
	}

	node1, node2 := d1.head, d2.head

	for node1 != nil {
		if node1.value != node2.value {
			return false
		}
		node1 = node1.next
		node2 = node2.next
	}
	return true
}
