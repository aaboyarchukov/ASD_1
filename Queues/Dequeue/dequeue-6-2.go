package main

import (
	"fmt"
	"slices"
	"strings"

	"constraints"
)

// task 2
// Same time complition
// Use doubly linked list

// task 3
// tests are added

// task 4
// Is pallindrom
// move from tail to head and comparison values

// t = O(n), where n = len(row), mem = O(1)
func RowToDequeue(row string) Deque[string] {
	var result Deque[string]

	for _, item := range row {
		result.AddTail(string(item))
	}

	return result
}

// t = O(n), where n = len(row), mem = O(n), where n = len(row)
func IsPallindrom(row string) bool {
	d := RowToDequeue(row)

	for d.size > 1 {
		head, err_head := d.RemoveFront()
		tail, err_tail := d.RemoveTail()
		if err_head != nil || err_tail != nil {
			return false
		}

		if strings.Compare(head, tail) != 0 {
			return false
		}
	}

	return true
}

// task 5
// Min element for O(1)
// Add min element for every node
type MinDequeNode[T constraints.Ordered] struct {
	value T
	next  *MinDequeNode[T]
	min   T
}

type MinDeque[T constraints.Ordered] struct {
	head *MinDequeNode[T]
	tail *MinDequeNode[T]
	size int
}

func (d *MinDeque[T]) Size() int {
	return d.size
}

func (d *MinDeque[T]) AddFront(itm T) {
	node := &MinDequeNode[T]{
		value: itm,
		next:  nil,
	}

	min, err_get_min := d.GetMin()
	if err_get_min != nil || itm < min {
		min = itm
	}

	node.min = min

	if d.head == nil {
		d.tail = node
	} else {
		node.next = d.head
	}

	d.head = node
	d.size++

}

func (d *MinDeque[T]) AddTail(itm T) {
	node := &MinDequeNode[T]{
		value: itm,
		next:  nil,
	}

	min, err_get_min := d.GetMin()
	if err_get_min != nil || itm < min {
		min = itm
	}

	node.min = min

	if d.tail == nil {
		d.head = node
	} else {
		d.tail.next = node
	}

	d.tail = node
	d.size++

}

func (d *MinDeque[T]) RemoveFront() (T, error) {
	var resultNode *MinDequeNode[T] = &MinDequeNode[T]{}

	if d.size == 0 {
		return resultNode.value, fmt.Errorf("dequeue is empty")
	}

	resultNode = d.head
	if d.size == 1 {
		d.head, d.tail = nil, nil
	} else {
		currentNode := d.head
		d.head = d.head.next
		currentNode.next = nil
	}

	d.size--

	return resultNode.value, nil
}

func (d *MinDeque[T]) RemoveTail() (T, error) {
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

func (d *MinDeque[T]) GetMin() (T, error) {
	var result T
	if d.Size() == 0 {
		return result, fmt.Errorf("dequeue is empty")
	}

	if d.head.min < d.tail.min {
		return d.head.min, nil
	}

	return d.tail.min, nil
}

func GetMinDequeue[T constraints.Ordered](values []T) MinDeque[T] {
	var result MinDeque[T]

	for _, item := range values {
		result.AddTail(item)
	}

	return result
}

// task 6
// Dequeue with dynamic array

type DequeueArray[T constraints.Ordered] struct {
	storageFront []T
	storageTail  []T
	size         int
}

func (da *DequeueArray[T]) Size() int {
	return da.size
}
func (da *DequeueArray[T]) AddTail(itm T) {
	da.storageTail = append(da.storageTail, itm)
	da.size++
}
func (da *DequeueArray[T]) AddFront(itm T) {
	da.storageFront = append(da.storageFront, itm)
	da.size++
}
func (da *DequeueArray[T]) RemoveTail() (T, error) {
	var result T = da.storageTail[da.size-1]
	da.storageTail = slices.Delete(da.storageTail, da.size-1, da.size)
	da.size--
	return result, nil
}

func (da *DequeueArray[T]) RemoveFront() (T, error) {
	var result T = da.storageFront[da.size-1]
	da.storageFront = slices.Delete(da.storageFront, da.size-1, da.size)
	da.size--
	return result, nil
}
