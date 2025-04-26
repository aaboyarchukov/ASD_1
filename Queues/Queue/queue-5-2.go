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

type QueueStacks[T any] struct {
	enqueBody *Stack[T]
	dequeBody *Stack[T]
	cap       int
}

// t = O(1), mem = O(1)
func (q *QueueStacks[T]) Enqueue(itm T) {
	q.enqueBody.Push(itm)
	q.cap += 1
}

// t = o(1) mem = O(1)
func (q *QueueStacks[T]) Dequeue() (T, error) {
	if q.dequeBody.size == 0 {
		q.dequeBody.base = make([]T, q.enqueBody.size)
		for i := range q.cap {
			q.dequeBody.base[i] = q.enqueBody.base[q.cap-1-i]
		}
		q.dequeBody.size = q.enqueBody.size
	}

	value, err := q.dequeBody.Pop()
	if err != nil {
		return value, err
	}

	q.cap -= 1

	return value, nil
}

func GetQueueStacks[T any](values []T) QueueStacks[T] {
	var result QueueStacks[T] = QueueStacks[T]{
		enqueBody: &Stack[T]{
			base: make([]T, 0),
		},
		dequeBody: &Stack[T]{
			base: make([]T, 0),
		},
	}

	for _, item := range values {
		result.Enqueue(item)
	}

	return result
}

func EqualStack[T comparable](st1 []T, st2 []T) bool {
	if len(st1) != len(st2) {
		return false
	}

	for i := 0; i < len(st1); i++ {
		if st1[i] != st2[i] {
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
// можно попробовать сделать по-другому -  через правую сторону
type QueueArr[T any] struct {
	head int
	tail int
	cap  int
	body []T
}

func Init[T any](sz int) QueueArr[T] {
	body := make([]T, sz)
	return QueueArr[T]{
		cap:  0,
		head: 0,
		tail: 0,
		body: body,
	}
}

func (qa *QueueArr[T]) IsFull() bool {
	return qa.cap >= len(qa.body)
}

func (qa *QueueArr[T]) Enqueue(value T) {
	if qa.IsFull() {
		return
	}

	qa.body[qa.tail] = value
	if qa.tail == 0 {
		qa.tail = len(qa.body) - 1
	} else {
		qa.tail -= 1
	}

	qa.cap += 1
}

func (qa *QueueArr[T]) Dequeue() {
	if qa.head == 0 {
		qa.head = len(qa.body) - 1
	} else {
		qa.head -= 1
	}

	qa.cap -= 1
}

func EqualCircleQueue[T comparable](body1 []T, body2 []T) bool {
	if len(body1) != len(body2) {
		return false
	}

	for indx := range len(body2) {
		if body1[indx] != body2[indx] {
			return false
		}
	}
	return true
}
