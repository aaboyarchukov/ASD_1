package main

import "fmt"

// task 3
// Circle queue
// t = O(n), where n is data input number, mem = O(1)
func Circle[T any](q *Queue[T], n int) error {
	if q.head == nil {
		return fmt.Errorf("queue is empty")
	}

	if q.size == 1 {
		return nil
	}

	for range n {
		node, err := q.Dequeue()
		if err != nil {
			return fmt.Errorf("queue is empty")
		}
		q.Enqueue(node)
	}

	return nil
}

// task 4
// Queue with two stacks
//

type Stack[T any] struct {
	base []T
	size int
}

func (st *Stack[T]) Size() int {
	return st.size
}

func (st *Stack[T]) Peek() (T, error) {
	var result T
	if st.size == 0 {
		return result, fmt.Errorf("stack is empty")
	}

	result = st.base[st.size-1]

	return result, nil
}

// t = O(1)
func (st *Stack[T]) Pop() (T, error) {
	var result T
	result, err_peek := st.Peek()
	if err_peek != nil {
		return result, err_peek
	}

	st.base = st.base[:st.size-1]
	st.size--

	return result, nil
}

// t = O(1)*, *with amortization time, when needed allocation
func (st *Stack[T]) Push(itm T) {
	st.base = append(st.base, itm)
	st.size++
}

func RemoveStack[T any](dest Stack[T], src Stack[T]) {

}

type QueueStacks[T any] struct {
	enqueBody Stack[T]
	dequeBody Stack[T]
}

// t = O(1), mem = O(1)
func (q *QueueStacks[T]) Enqueue(itm T) {
	q.enqueBody.Push(itm)
}

// t = o(1) mem = O(1)
func (q *QueueStacks[T]) Dequeue() (T, error) {
	if q.dequeBody.size == 0 {
		q.dequeBody.base = make([]T, q.enqueBody.size)
		copy(q.dequeBody.base, q.enqueBody.base) // it does not work!!!
		fmt.Println(q.dequeBody.base, q.enqueBody.base)
	}
	value, err := q.dequeBody.Pop()
	if err != nil {
		return value, err
	}

	return value, nil
}

func GetQueueStacks[T any](values []T) *QueueStacks[T] {
	var result QueueStacks[T]

	for _, item := range values {
		result.Enqueue(item)
	}

	return &result
}

func EqualStack[T comparable](st1 *Stack[T], st2 *Stack[T]) bool {
	if st1.size != st2.size {
		return false
	}

	for i := 0; i < st1.size; i++ {
		if st1.base[i] != st2.base[i] {
			return false
		}
	}

	return true
}

// task 5
// Reverse queue
// t = O(n), where n = queue.size, meme = O(m), where m = queue.size
func ReverseQueue[T any](q *Queue[T]) *Queue[T] {
	buffer := make([]T, 0, q.size)
	var result Queue[T]

	node := q.head
	for node != nil {
		buffer = append(buffer, node.value)
		node = node.next
	}

	for indx := range buffer {
		result.Enqueue(buffer[q.size-indx-1])
	}

	return &result
}

// task 6
type QueueArr[T any] struct {
	head int
	tail int
	body []T
}

func GetCircleQueue[T any](values []T) *QueueArr[T] {
	var result QueueArr[T]
	sizeArray := len(values)
	result.head = 1
	result.tail = sizeArray - 1

	body := make([]T, sizeArray+1)
	for indx, value := range values {
		if indx == sizeArray {
			break
		}

		body[indx+1] = value
	}

	result.body = body

	return &result
}

func EqualCircleQueue[T comparable](body1 *[]T, body2 *[]T) bool {
	if len(*body1) != len(*body2)+1 {
		return false
	}

	for indx := range len(*body2) {
		if (*body1)[indx+1] != (*body2)[indx] {
			return false
		}
	}
	return true
}
