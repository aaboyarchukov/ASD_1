package main

import (
	"constraints"
	"testing"
)

func TestSize(t *testing.T) {
	type TestSize[T constraints.Ordered] struct {
		name     string
		input    *OrderedList[T]
		wantSize int
	}

	testsInt := []TestSize[int64]{
		{
			name:     "Test1",
			input:    GetOrderedList(true, []int64{2, 1}),
			wantSize: 2,
		},
		{
			name:     "Test2",
			input:    GetOrderedList(true, []int64{2}),
			wantSize: 1,
		},
		{
			name:     "Test3",
			input:    GetOrderedList(true, []int64{}),
			wantSize: 0,
		},
	}
	testsFloat := []TestSize[float64]{
		{
			name:     "Test1",
			input:    GetOrderedList(true, []float64{2.1, 1.2}),
			wantSize: 2,
		},
		{
			name:     "Test2",
			input:    GetOrderedList(true, []float64{2.1}),
			wantSize: 1,
		},
		{
			name:     "Test3",
			input:    GetOrderedList(true, []float64{}),
			wantSize: 0,
		},
	}
	testsString := []TestSize[string]{
		{
			name:     "Test1",
			input:    GetOrderedList(true, []string{"1", "2"}),
			wantSize: 2,
		},
		{
			name:     "Test2",
			input:    GetOrderedList(true, []string{"1"}),
			wantSize: 1,
		},
		{
			name:     "Test3",
			input:    GetOrderedList(true, []string{}),
			wantSize: 0,
		},
	}

	for _, test := range testsInt {
		count := test.input.Count()
		if test.wantSize != count {
			t.Errorf("failed %s: wrong size. Result is %v, want: %v", test.name, count, test.wantSize)
		}
	}
	for _, test := range testsFloat {
		count := test.input.Count()
		if test.wantSize != count {
			t.Errorf("failed %s: wrong size. Result is %v, want: %v", test.name, count, test.wantSize)
		}
	}
	for _, test := range testsString {
		count := test.input.Count()
		if test.wantSize != count {
			t.Errorf("failed %s: wrong size. Result is %v, want: %v", test.name, count, test.wantSize)
		}
	}
}
