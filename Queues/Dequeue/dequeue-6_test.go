package main

import (
	"fmt"
	"testing"
)

func PrintDequeue[T any](d Deque[T]) {
	node := d.head

	for node != nil {
		fmt.Printf("%v ", node.value)
		node = node.next
	}
	fmt.Println()
}
func TestSize(t *testing.T) {
	tests := []struct {
		name  string
		input *Deque[any]
		want  int
	}{
		{"Test1", GetDequeue([]any{1, "2", false}), 3},
		{"Test2", GetDequeue([]any{}), 0},
		{"Test3", GetDequeue([]any{1}), 1},
	}

	for _, test := range tests {
		if test.input.size != test.want {
			t.Errorf("failed %s:\noutput: %v\twant: %v", test.name, test.input.size, test.want)
		}
	}
}
func TestAddFront(t *testing.T) {
	tests := []struct {
		name  string
		input *Deque[any]
		value any
		want  *Deque[any]
	}{
		{"Test1", GetDequeue([]any{1, 2, "true", false}), 3, GetDequeue([]any{1, 2, "true", false, 3})},
		{"Test2", GetDequeue([]any{1}), 3, GetDequeue([]any{1, 3})},
		{"Test3", GetDequeue([]any{}), 3, GetDequeue([]any{3})},
	}

	for _, test := range tests {
		test.input.AddTail(test.value)

		if !EqualDequeue(test.input, test.want) {
			t.Errorf("failed %s: add in tail\noutput:", test.name)
			PrintDequeue(*test.input)
			t.Errorf("want: ")
			PrintDequeue(*test.want)
		}
	}
}
func TestRemoveFront(t *testing.T) {
	tests := []struct {
		name  string
		input *Deque[any]
		want  *Deque[any]
		value any
		err   error
	}{
		{"Test1", GetDequeue([]any{1, 2, "true", false}), GetDequeue([]any{2, "true", false}), 1, nil},
		{"Test2", GetDequeue([]any{1}), GetDequeue([]any{}), 1, nil},
		{"Test3", GetDequeue([]any{}), GetDequeue([]any{}), nil, fmt.Errorf("dequeue is empty")},
	}

	for _, test := range tests {
		value, err := test.input.RemoveFront()

		if value != test.value {
			t.Errorf("failed %s: remove front\noutput: %v\nwant: %v", test.name, value, test.value)
		}

		if !EqualDequeue(test.input, test.want) {
			t.Errorf("failed %s: remove front\noutput:", test.name)
			PrintDequeue(*test.input)
			t.Errorf("want: ")
			PrintDequeue(*test.want)
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
func TestAddTail(t *testing.T) {
	tests := []struct {
		name  string
		input *Deque[any]
		value any
		want  *Deque[any]
	}{
		{"Test1", GetDequeue([]any{1, 2, "true", false}), 3, GetDequeue([]any{3, 1, 2, "true", false})},
		{"Test2", GetDequeue([]any{1}), 3, GetDequeue([]any{3, 1})},
		{"Test3", GetDequeue([]any{}), 3, GetDequeue([]any{3})},
	}

	for _, test := range tests {
		test.input.AddFront(test.value)

		if !EqualDequeue(test.input, test.want) {
			t.Errorf("failed %s: add in front\noutput:", test.name)
			PrintDequeue(*test.input)
			t.Errorf("want: ")
			PrintDequeue(*test.want)
		}
	}
}
func TestRemoveTail(t *testing.T) {
	tests := []struct {
		name  string
		input *Deque[any]
		want  *Deque[any]
		value any
		err   error
	}{
		{"Test1", GetDequeue([]any{1, 2, "true", false}), GetDequeue([]any{1, 2, "true"}), false, nil},
		{"Test2", GetDequeue([]any{1}), GetDequeue([]any{}), 1, nil},
		{"Test3", GetDequeue([]any{}), GetDequeue([]any{}), nil, fmt.Errorf("dequeue is empty")},
	}

	for _, test := range tests {
		value, err := test.input.RemoveTail()

		if value != test.value {
			t.Errorf("failed %s: remove tail\noutput: %v\nwant: %v", test.name, value, test.value)
		}

		if !EqualDequeue(test.input, test.want) {
			t.Errorf("failed %s: remove tail\noutput:", test.name)
			PrintDequeue(*test.input)
			t.Errorf("want: ")
			PrintDequeue(*test.want)
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
