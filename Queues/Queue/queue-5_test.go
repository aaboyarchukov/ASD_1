package main

import (
	"fmt"
	"testing"
)

func TestDequeue(t *testing.T) {
	tests := []struct {
		name  string
		input *Queue[any]
		want  *Queue[any]
		value any
		err   error
	}{
		{"Test1", GetQueue([]any{1, "row", true}), GetQueue([]any{"row", true}), 1, nil},
		{"Test2", GetQueue([]any{}), GetQueue([]any{}), nil, fmt.Errorf("queue is empty")},
		{"Test3", GetQueue([]any{"row"}), GetQueue([]any{}), "row", nil},
	}

	for _, test := range tests {
		value, err_deq := test.input.Dequeue()

		if value != test.value {
			t.Errorf("failed %s: vlaue are not equal", test.name)
		}
		if !EqualQueue(test.input, test.want) {
			t.Errorf("failed %s: delete last item, queues are not equal", test.name)
		}

		switch test.err {
		case nil:
			if err_deq != test.err {
				t.Errorf("failed %s: different errors", test.name)
			}
		default:
			if err_deq.Error() != test.err.Error() {
				t.Errorf("failed %s: different errors", test.name)
			}
		}
	}

}
func TestSize(t *testing.T) {
	tests := []struct {
		name  string
		input *Queue[any]
		want  int
	}{
		{"Test1", GetQueue([]any{1, "row", true}), 3},
		{"Test2", GetQueue([]any{}), 0},
		{"Test3", GetQueue([]any{1}), 1},
	}

	for _, test := range tests {
		testSize := test.input.size

		if testSize != test.want {
			t.Errorf("failed %s: get size", test.name)
		}
	}
}
func TestEnqueue(t *testing.T) {
	tests := []struct {
		name  string
		input *Queue[any]
		value any
		want  *Queue[any]
	}{
		{"Test1", GetQueue([]any{1, "row", true}), false, GetQueue([]any{1, "row", true, false})},
		{"Test2", GetQueue([]any{1}), false, GetQueue([]any{1, false})},
		{"Test2", GetQueue([]any{}), false, GetQueue([]any{false})},
	}

	for _, test := range tests {
		test.input.Enqueue(test.value)

		if !EqualQueue(test.input, test.want) {
			t.Errorf("failed %s: insert item at the end", test.name)
		}
	}
}
