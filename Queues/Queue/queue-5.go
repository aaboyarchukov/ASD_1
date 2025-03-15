package main

import (
	"fmt"
	_ "os"
)

type QNode[T any] struct {
	value T
	next  *QNode[T]
}

type Queue[T any] struct {
	head *QNode[T]
	tail *QNode[T]
	size int
}

func (q *Queue[T]) Size() int {
	return q.size
}

// t = O(1)
func (q *Queue[T]) Dequeue() (T, error) {
	var result T
	if q.size == 0 {
		return result, fmt.Errorf("queue is empty")
	}

	result = q.head.value
	next := q.head.next
	q.head = next

	if next == nil {
		q.tail = next
	}

	q.size--

	return result, nil
}

// t = O(1)
func (q *Queue[T]) Enqueue(itm T) {
	node := &QNode[T]{
		value: itm,
		next:  nil,
	}

	if q.head == nil {
		q.head = node
	} else {
		q.tail.next = node
	}

	q.tail = node
	q.size++
}

func GetQueue[T any](values []T) *Queue[T] {
	var result Queue[T]

	for _, item := range values {
		result.Enqueue(item)
	}

	return &result
}

func EqualQueue[T comparable](q1 *Queue[T], q2 *Queue[T]) bool {
	if q1.size != q2.size {
		return false
	}

	if q1 == nil && q1 == q2 {
		return true
	}

	tempNode1, tempNode2 := q1.head, q2.head

	for tempNode1 != nil {
		if tempNode1.value != tempNode2.value {
			return false
		}
		tempNode1 = tempNode1.next
		tempNode2 = tempNode2.next
	}
	return true
}
