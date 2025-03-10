package main

import (
	"fmt"
	"testing"
)

func TestFrontStackPop(t *testing.T) {
	tests := []struct {
		name    string
		input   *StackFront[any]
		want    *StackFront[any]
		element any
		err     error
	}{
		{"Test1", GetStackFront([]any{2, "+", 1.3, false}), GetStackFront([]any{2, "+", 1.3}), false, nil},
		{"Test2", GetStackFront([]any{}), GetStackFront([]any{}), nil, fmt.Errorf("this element is not in stack")},
	}

	for _, test := range tests {
		pop, err_pop := test.input.Pop()
		if !EqualStackFront(test.input, test.want) {
			t.Errorf("failed %s: stacks are not equal", test.name)
		}

		if pop != test.element {
			t.Errorf("failed %s: wrong pop element", test.name)
		}

		switch err_pop {
		case nil:
			if err_pop != test.err {
				t.Errorf("failed %s: wrong err", test.name)
			}
		default:
			if err_pop.Error() != test.err.Error() {
				t.Errorf("failed %s: wrong err", test.name)
			}
		}

	}
}
func TestFrontStackPush(t *testing.T) {
	tests := []struct {
		name    string
		input   *StackFront[any]
		want    *StackFront[any]
		element any
	}{
		{"Test1", GetStackFront([]any{2, "+", 1.3}), GetStackFront([]any{2, "+", 1.3, false}), false},
		{"Test2", GetStackFront([]any{}), GetStackFront([]any{1}), 1},
	}

	for _, test := range tests {
		test.input.Push(test.element)
		if !EqualStackFront(test.input, test.want) {
			t.Errorf("failed %s: stacks are not equal", test.name)
		}

	}
}
func TestFrontStackPeek(t *testing.T) {
	tests := []struct {
		name    string
		input   *StackFront[any]
		want    *StackFront[any]
		element any
		err     error
	}{
		{"Test1", GetStackFront([]any{2, "+", 1.3, false}), GetStackFront([]any{2, "+", 1.3}), false, nil},
		{"Test2", GetStackFront([]any{}), GetStackFront([]any{}), nil, fmt.Errorf("this element is not in stack")},
	}

	for _, test := range tests {
		peek, err_pop := test.input.Peek()

		if peek != test.element {
			t.Errorf("failed %s: wrong peek element", test.name)
		}

		switch err_pop {
		case nil:
			if err_pop != test.err {
				t.Errorf("failed %s: wrong err", test.name)
			}
		default:
			if err_pop.Error() != test.err.Error() {
				t.Errorf("failed %s: wrong err", test.name)
			}
		}

	}
}

func TestFrontStackSize(t *testing.T) {
	tests := []struct {
		name  string
		input *StackFront[any]
		want  int
	}{
		{"Test1", GetStackFront([]any{2, "+", 1.3, false}), 4},
		{"Test2", GetStackFront([]any{}), 0},
	}

	for _, test := range tests {
		size := test.input.Size()

		if size != test.input.size {
			t.Errorf("failed %s: wrong size answer", test.name)
		}

	}
}

func TestBracketsSequence(t *testing.T) {
	tests := []struct {
		name     string
		sequence string
		want     bool
	}{
		{"Test1", "((()))", true},
		{"Test2", "((())(()))))", false},
		{"Test3", "((()[)][[]]])", false},
		{"Test4", "([{(([]))}])({[][()]})", true},
	}

	for _, tempTest := range tests {
		test := tempTest
		result := BracketsSequence(test.sequence)

		if result != test.want {
			t.Errorf("failed %s", test.name)
		}
	}
}

func TestMinMean(t *testing.T) {
	tests := []struct {
		name  string
		input *MinStack
		min   int
		mean  float32
	}{
		{"Test1", GetMinStack([]int{2, 1, 3, 0}), 0, 6.0 / 4.0},
	}

	for _, tempTest := range tests {
		test := tempTest
		min := test.input.GetMin()
		mean := test.input.GetMean()

		if min != test.min {
			t.Errorf("failed %s: get min", test.name)
		}

		if mean != test.mean {
			t.Errorf("failed %s: get mean", test.name)
		}
	}
}

func TestResultExpression(t *testing.T) {
	tests := []struct {
		name   string
		input  []any
		result int
	}{
		{"Test1", []any{"=", "*", 3, "+", 2, 1}, 9},
		{"Test2", []any{"=", "+", 9, "*", 5, "+", 2, 8}, 59},
	}

	for _, test := range tests {
		value, err := ResultPostfixRecord[any](test.input)

		if err != nil {
			t.Errorf("failed %s: raise error", test.name)
			continue
		}

		if value != test.result {
			t.Errorf("failed %s", test.name)
		}
	}
}
