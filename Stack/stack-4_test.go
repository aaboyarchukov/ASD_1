package main

import (
	"fmt"
	"testing"
)

func TestPop(t *testing.T) {
	tests := []struct {
		name  string
		input *Stack[any]
		want  *Stack[any]
		err   error
	}{
		{"Test1", GetStack([]any{1, "row", 2.4, true}), GetStack([]any{1, "row", 2.4}), nil},
		{"Test2", GetStack([]any{}), GetStack([]any{}), fmt.Errorf("this element is not in stack")},
	}

	for _, tempTest := range tests {
		test := tempTest
		peek, err_peek := test.input.Peek()
		pop, err_pop := test.input.Pop()

		if test.err != nil && (err_pop.Error() != test.err.Error() || err_peek.Error() != test.err.Error()) {
			t.Errorf("failed %s", test.name)
			continue
		}

		if pop != peek || !EqualStack(test.input, test.want) {
			t.Errorf("failed %s", test.name)
			continue
		}
	}
}
func TestPeek(t *testing.T) {
	tests := []struct {
		name  string
		input *Stack[any]
		want  any
		err   error
	}{
		{"Test1", GetStack([]any{1, "row", 2.4, true}), true, nil},
		{"Test2", GetStack([]any{}), nil, fmt.Errorf("this element is not in stack")},
	}

	for _, tempTest := range tests {
		test := tempTest
		peek, err_peek := test.input.Peek()

		if test.err != nil && err_peek.Error() != test.err.Error() {
			t.Errorf("failed %s", test.name)
			continue
		}

		if test.want != peek {
			t.Errorf("failed %s", test.name)
			continue
		}
	}
}
func TestPush(t *testing.T) {
	tests := []struct {
		name  string
		input *Stack[any]
		item  any
		want  *Stack[any]
	}{
		{"Test1", GetStack([]any{1, "row", 2.4, true}), false, GetStack([]any{1, "row", 2.4, true, false})},
		{"Test2", GetStack([]any{}), 1, GetStack([]any{1})},
	}

	for _, tempTest := range tests {
		test := tempTest
		test.input.Push(test.item)

		if !EqualStack(test.input, test.want) {
			t.Errorf("failed %s", test.name)
			continue
		}
	}
}
func TestSize(t *testing.T) {
	tests := []struct {
		name  string
		input *Stack[any]
		want  int
	}{
		{"Test1", GetStack([]any{1, "row", 2.4, true}), 4},
		{"Test2", GetStack([]any{}), 0},
	}

	for _, tempTest := range tests {
		test := tempTest

		if test.input.Size() != test.want {
			t.Errorf("failed %s", test.name)
			continue
		}
	}
}
