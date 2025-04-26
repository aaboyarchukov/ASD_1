package main

import (
	"fmt"
	"testing"
)

func PrintQueue[T any](q Queue[T]) {
	node := q.head
	for node != nil {
		fmt.Printf("%v ", node.value)
		node = node.next
	}
	fmt.Println()
}
func TestCircleQueue(t *testing.T) {
	tests := []struct {
		name  string
		input *Queue[any]
		n     int
		want  *Queue[any]
		err   error
	}{
		{"Test1", GetQueue([]any{1, "2", false}), 2, GetQueue([]any{false, 1, "2"}), nil},
		{"Test2", GetQueue([]any{1}), 1, GetQueue([]any{1}), nil},
		{"Test3", GetQueue([]any{}), 1, GetQueue([]any{}), fmt.Errorf("queue is empty")},
		{"Test4", GetQueue([]any{1, "2", false, 3, 4}), 2, GetQueue([]any{false, 3, 4, 1, "2"}), nil},
	}

	for _, test := range tests {
		err := Circle(test.input, test.n)
		if !EqualQueue(test.input, test.want) {
			t.Errorf("failed %s: circle queue, output is:\n", test.name)
			fmt.Println("input: ")
			PrintQueue(*test.input)
			fmt.Println("want: ")
			PrintQueue(*test.want)
		}

		switch test.err {
		case nil:
			if err != test.err {
				t.Errorf("failed %s: different errors", test.name)
			}
		default:
			if err.Error() != test.err.Error() {
				t.Errorf("failed %s: different errors", test.name)
			}
		}
	}
}

func TestReverseQueue(t *testing.T) {
	tests := []struct {
		name  string
		input *Queue[any]
		want  *Queue[any]
	}{
		{"Test1", GetQueue([]any{1, "2", false}), GetQueue([]any{false, "2", 1})},
		{"Test2", GetQueue([]any{1}), GetQueue([]any{1})},
		{"Test3", GetQueue([]any{}), GetQueue([]any{})},
		{"Test4", GetQueue([]any{1, "2", false, 3, 4}), GetQueue([]any{4, 3, false, "2", 1})},
	}

	for _, test := range tests {
		test.input = ReverseQueue(test.input)
		if !EqualQueue(test.input, test.want) {
			t.Errorf("failed %s: circle queue, output is:\n", test.name)
			fmt.Println("input: ")
			PrintQueue(*test.input)
			fmt.Println("want: ")
			PrintQueue(*test.want)
		}
	}
}

func TestQueueStacksEnque(t *testing.T) {
	tests := []struct {
		name  string
		input QueueStacks[any]
		value any
		want  QueueStacks[any]
	}{
		{"Test1", GetQueueStacks([]any{1, "2", false}), 3, GetQueueStacks([]any{1, "2", false, 3})},
	}

	for _, test := range tests {
		test.input.Enqueue(test.value)
		if !EqualStack(test.input.enqueBody.base, test.want.enqueBody.base) {
			t.Errorf("failed %s: enqueue", test.name)
		}
	}
}
func TestQueueStacksDeque(t *testing.T) {
	tests := []struct {
		name  string
		input QueueStacks[any]
		want  QueueStacks[any]
		err   error
		value any
	}{
		{"Test1", GetQueueStacks([]any{1, "2", false}), GetQueueStacks([]any{"2", false}), nil, 1},
	}

	for _, test := range tests {
		value, err := test.input.Dequeue()
		if value != test.value {
			t.Errorf("failed %s: dequeue, diffrent values", test.name)
		}

		switch test.err {
		case nil:
			if err != test.err {
				t.Errorf("failed %s: different errors", test.name)
			}
		default:
			if err.Error() != test.err.Error() {
				t.Errorf("failed %s: different errors", test.name)
			}
		}

	}
}
func TestQueueArrEnqueue(t *testing.T) {
	tests := []struct {
		name  string
		input QueueArr[any]
		value any
		want  QueueArr[any]
	}{
		{"Test1", QueueArr[any]{
			cap:  3,
			body: []any{1, 0, 0, 3, 4, 5},
			head: 0,
			tail: 2,
		}, 2, QueueArr[any]{
			cap:  4,
			body: []any{1, 0, 2, 3, 4, 5},
			head: 0,
			tail: 1,
		}},
	}

	for _, test := range tests {
		test.input.Enqueue(test.value)
		if !EqualCircleQueue(test.input.body, test.want.body) {
			t.Errorf("failed %s: enqueue value", test.name)
		}
	}
}
