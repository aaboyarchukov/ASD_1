package main

import (
	"constraints"
	"errors"
	"testing"
)

func TestSize(t *testing.T) {
	type testType[T constraints.Ordered] struct {
		name     string
		input    *OrderedList[T]
		wantSize int
	}

	testsInt := []testType[int64]{
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
		{
			name:     "Test4",
			input:    GetOrderedList(true, []int64{1, 1, 1, 1}),
			wantSize: 4,
		},
	}
	testsFloat := []testType[float64]{
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
		{
			name:     "Test4",
			input:    GetOrderedList(true, []float64{1.1, 1.1, 1.1}),
			wantSize: 3,
		},
	}
	testsString := []testType[string]{
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
		{
			name:     "Test4",
			input:    GetOrderedList(true, []string{"1", "1", "1", "1"}),
			wantSize: 4,
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

func TestAdd(t *testing.T) {
	type testType[T constraints.Ordered] struct {
		name  string
		input *OrderedList[T]
		value []T
		want  []T
	}

	testsInt := []testType[int64]{
		{
			name:  "Test1",
			input: GetOrderedList(true, []int64{9, 1, 5, 2, 7, 4}),
			value: []int64{3},
			want:  []int64{1, 2, 3, 4, 5, 7, 9},
		},
		{
			name:  "Test2",
			input: GetOrderedList(true, []int64{1}),
			value: []int64{3},
			want:  []int64{1, 3},
		},
		{
			name:  "Test3",
			input: GetOrderedList(true, []int64{}),
			value: []int64{3},
			want:  []int64{3},
		},
		{
			name:  "Test4",
			input: GetOrderedList(false, []int64{9, 1, 5, 2, 7, 4}),
			value: []int64{3},
			want:  []int64{9, 7, 5, 4, 3, 2, 1},
		},
		{
			name:  "Test5",
			input: GetOrderedList(false, []int64{1}),
			value: []int64{3},
			want:  []int64{3, 1},
		},
		{
			name:  "Test6",
			input: GetOrderedList(false, []int64{}),
			value: []int64{3},
			want:  []int64{3},
		},
		{
			name:  "Test7",
			input: GetOrderedList(true, []int64{9, 1, 5, 2, 7, 3, 4}),
			value: []int64{3},
			want:  []int64{1, 2, 3, 3, 4, 5, 7, 9},
		},
		{
			name:  "Test8",
			input: GetOrderedList(false, []int64{9, 1, 5, 2, 7, 3, 4}),
			value: []int64{3},
			want:  []int64{9, 7, 5, 4, 3, 3, 2, 1},
		},
		{
			name:  "Test9",
			input: GetOrderedList(true, []int64{1, 1, 1, 1}),
			value: []int64{1},
			want:  []int64{1, 1, 1, 1, 1},
		},
		{
			name:  "Test10",
			input: GetOrderedList(false, []int64{1, 1, 1, 1}),
			value: []int64{1},
			want:  []int64{1, 1, 1, 1, 1},
		},
		{
			name:  "Test10",
			input: GetOrderedList(false, []int64{1, 1, 1, 1}),
			value: []int64{1},
			want:  []int64{1, 1, 1, 1, 1},
		},
		{
			name:  "Test11",
			input: GetOrderedList(true, []int64{}),
			value: []int64{100, 99, 98, 97, 96, 95, 94, 93, 92, 91, 90, 89, 88, 87, 86, 85, 84, 83, 82, 81, 80, 79, 78, 77, 76, 75, 74, 73, 72, 71, 70, 69, 68, 67, 66, 65, 64, 63, 62, 61, 60, 59, 58, 57, 56, 55, 54, 53, 52, 51, 50, 49, 48, 47, 46, 45, 44, 43, 42, 41, 40, 39, 38, 37, 36, 35, 34, 33, 32, 31, 30, 29, 28, 27, 26, 25, 24, 23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			want:  []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100},
		},
		{
			name:  "Test12",
			input: GetOrderedList(false, []int64{}),
			value: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100},
			want:  []int64{100, 99, 98, 97, 96, 95, 94, 93, 92, 91, 90, 89, 88, 87, 86, 85, 84, 83, 82, 81, 80, 79, 78, 77, 76, 75, 74, 73, 72, 71, 70, 69, 68, 67, 66, 65, 64, 63, 62, 61, 60, 59, 58, 57, 56, 55, 54, 53, 52, 51, 50, 49, 48, 47, 46, 45, 44, 43, 42, 41, 40, 39, 38, 37, 36, 35, 34, 33, 32, 31, 30, 29, 28, 27, 26, 25, 24, 23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
	}
	testsFloat := []testType[float64]{
		{
			name:  "Test1",
			input: GetOrderedList(true, []float64{2.1, 1.2, 5.6, 1.1}),
			value: []float64{3.4},
			want:  []float64{1.1, 1.2, 2.1, 3.4, 5.6},
		},
		{
			name:  "Test2",
			input: GetOrderedList(true, []float64{2.1}),
			value: []float64{3.4},
			want:  []float64{2.1, 3.4},
		},
		{
			name:  "Test3",
			input: GetOrderedList(true, []float64{}),
			value: []float64{3.4},
			want:  []float64{3.4},
		},
		{
			name:  "Test4",
			input: GetOrderedList(false, []float64{2.1, 1.2, 5.6, 1.1}),
			value: []float64{3.4},
			want:  []float64{5.6, 3.4, 2.1, 1.2, 1.1},
		},
		{
			name:  "Test5",
			input: GetOrderedList(false, []float64{2.1}),
			value: []float64{3.4},
			want:  []float64{3.4, 2.1},
		},
		{
			name:  "Test6",
			input: GetOrderedList(false, []float64{}),
			value: []float64{3.4},
			want:  []float64{3.4},
		},
		{
			name:  "Test7",
			input: GetOrderedList(false, []float64{1.1, 1.1, 1.1, 1.1}),
			value: []float64{1.1},
			want:  []float64{1.1, 1.1, 1.1, 1.1, 1.1},
		},
		{
			name:  "Test8",
			input: GetOrderedList(true, []float64{1.1, 1.1, 1.1, 1.1}),
			value: []float64{1.1},
			want:  []float64{1.1, 1.1, 1.1, 1.1, 1.1},
		},
		{
			name:  "Test9",
			input: GetOrderedList(true, []float64{1.1, 1.1, 1.1, 1.1, 5.4, 5.4, 3.2, 1.1, 1.1, 1.1}),
			value: []float64{1.1},
			want:  []float64{1.1, 1.1, 1.1, 1.1, 1.1, 1.1, 1.1, 1.1, 3.2, 5.4, 5.4},
		},
		{
			name:  "Test10",
			input: GetOrderedList(false, []float64{1.1, 1.1, 1.1, 1.1, 5.4, 5.4, 3.2, 1.1, 1.1, 1.1}),
			value: []float64{1.1},
			want:  []float64{5.4, 5.4, 3.2, 1.1, 1.1, 1.1, 1.1, 1.1, 1.1, 1.1, 1.1},
		},
	}
	testsString := []testType[string]{
		{
			name:  "Test1",
			input: GetOrderedList(true, []string{"  1", "3  ", " 2"}),
			value: []string{"12"},
			want:  []string{"1 ", "12  ", " 2", "   3"},
		},
		{
			name:  "Test2",
			input: GetOrderedList(true, []string{"1"}),
			value: []string{"12"},
			want:  []string{"1", "12"},
		},
		{
			name:  "Test3",
			input: GetOrderedList(true, []string{}),
			value: []string{"12"},
			want:  []string{"12"},
		},
		{
			name:  "Test4",
			input: GetOrderedList(false, []string{"  1", "3  ", " 2"}),
			value: []string{"12"},
			want:  []string{"   3", "2", "12  ", " 1"},
		},
		{
			name:  "Test5",
			input: GetOrderedList(false, []string{"1"}),
			value: []string{"12"},
			want:  []string{"12", "1"},
		},
		{
			name:  "Test6",
			input: GetOrderedList(false, []string{}),
			value: []string{"12"},
			want:  []string{"12"},
		},
		{
			name:  "Test7",
			input: GetOrderedList(false, []string{"1  ", "1", " 1"}),
			value: []string{"1"},
			want:  []string{"1  ", "1", " 1", "1"},
		},
		{
			name:  "Test8",
			input: GetOrderedList(true, []string{"1  ", "1", " 1"}),
			value: []string{"1"},
			want:  []string{"1  ", "1", " 1", "1"},
		},
		{
			name:  "Test9",
			input: GetOrderedList(true, []string{"1", "1", "1"}),
			value: []string{"1"},
			want:  []string{"1", "1", "1", "1"},
		},
		{
			name:  "Test10",
			input: GetOrderedList(true, []string{"1", "1", "2", "3", "3", "3", "1"}),
			value: []string{"1"},
			want:  []string{"1", "1", "1", "1", "2", "3", "3", "3"},
		},
		{
			name:  "Test11",
			input: GetOrderedList(true, []string{"a", " b", "v", "  a"}),
			value: []string{"a"},
			want:  []string{"a", "  a", "a", "b", "v"},
		},
	}

	for _, test := range testsInt {
		for _, item := range test.value {
			test.input.Add(item)
		}

		if !EqualOrderedLists(test.input, test.want) {
			t.Errorf("failed %s: not equals lists. Result is: ", test.name)
			PrintList(*test.input)
		}
		PrintList(*test.input)
	}
	for _, test := range testsFloat {
		for _, item := range test.value {
			test.input.Add(item)
		}
		if !EqualOrderedLists(test.input, test.want) {
			t.Errorf("failed %s: not equals lists. Result is: ", test.name)
			PrintList(*test.input)
		}
	}
	for _, test := range testsString {
		for _, item := range test.value {
			test.input.Add(item)
		}
		if !EqualOrderedLists(test.input, test.want) {
			t.Errorf("failed %s: not equals lists. Result is: ", test.name)
			PrintList(*test.input)
		}
	}
}

func TestFind(t *testing.T) {
	type testType[T constraints.Ordered] struct {
		name  string
		input *OrderedList[T]
		want  []T
		value T
		err   error
	}

	testsInt := []testType[int64]{
		{
			name:  "Test1",
			input: GetOrderedList(true, []int64{9, 1, 5, 2, 3, 7, 4}),
			want:  []int64{1, 2, 3, 4, 5, 7, 9},
			value: 3,
			err:   nil,
		},
		{
			name:  "Test2",
			input: GetOrderedList(true, []int64{9, 1, 5, 2, 7, 4}),
			want:  []int64{1, 2, 4, 5, 7, 9},
			value: 3,
			err:   errors.New("node is not finding"),
		},
		{
			name:  "Test3",
			input: GetOrderedList(true, []int64{9, 1, 5, 2, 3, 3, 3, 7, 4}),
			want:  []int64{1, 2, 3, 3, 3, 4, 5, 7, 9},
			value: 3,
			err:   nil,
		},
		{
			name:  "Test4",
			input: GetOrderedList(false, []int64{9, 1, 5, 2, 3, 7, 4}),
			want:  []int64{9, 7, 5, 4, 3, 2, 1},
			value: 3,
			err:   nil,
		},
		{
			name:  "Test5",
			input: GetOrderedList(false, []int64{9, 1, 5, 2, 7, 4}),
			want:  []int64{9, 7, 5, 4, 2, 1},
			value: 3,
			err:   errors.New("node is not finding"),
		},
		{
			name:  "Test6",
			input: GetOrderedList(false, []int64{9, 1, 5, 2, 3, 3, 3, 7, 4}),
			want:  []int64{9, 7, 5, 4, 3, 3, 3, 2, 1},
			value: 3,
			err:   nil,
		},
	}
	testsFloat := []testType[float64]{
		{
			name:  "Test1",
			input: GetOrderedList(true, []float64{2.1, 1.2, 5.6, 1.1, 3.4}),
			want:  []float64{1.1, 1.2, 2.1, 3.4, 5.6},
			value: 3.4,
			err:   nil,
		},
		{
			name:  "Test2",
			input: GetOrderedList(true, []float64{2.1, 1.2, 5.6, 1.1}),
			want:  []float64{1.1, 1.2, 2.1, 5.6},
			value: 3.4,
			err:   errors.New("node is not finding"),
		},
		{
			name:  "Test3",
			input: GetOrderedList(true, []float64{2.1, 1.2, 5.6, 3.4, 3.4, 3.4, 1.1}),
			want:  []float64{1.1, 1.2, 2.1, 3.4, 3.4, 3.4, 5.6},
			value: 3.4,
			err:   nil,
		},
		{
			name:  "Test4",
			input: GetOrderedList(false, []float64{2.1, 1.2, 5.6, 1.1, 3.4}),
			want:  []float64{5.6, 3.4, 2.1, 1.2, 1.1},
			value: 3.4,
			err:   nil,
		},
		{
			name:  "Test5",
			input: GetOrderedList(false, []float64{2.1, 1.2, 5.6, 1.1}),
			want:  []float64{5.6, 2.1, 1.2, 1.1},
			value: 3.4,
			err:   errors.New("node is not finding"),
		},
		{
			name:  "Test6",
			input: GetOrderedList(false, []float64{2.1, 1.2, 5.6, 3.4, 3.4, 3.4, 1.1}),
			want:  []float64{5.6, 3.4, 3.4, 3.4, 2.1, 1.2, 1.1},
			value: 3.4,
			err:   nil,
		},
	}
	testsString := []testType[string]{
		{
			name:  "Test1",
			input: GetOrderedList(true, []string{"12  ", "  1", "3  ", " 2"}),
			want:  []string{"1 ", "12  ", " 2", "   3"},
			value: "12",
			err:   nil,
		},
		{
			name:  "Test2",
			input: GetOrderedList(true, []string{"  1", "3  ", " 2"}),
			want:  []string{"1 ", " 2", "   3"},
			value: "12",
			err:   errors.New("node is not finding"),
		},
		{
			name:  "Test3",
			input: GetOrderedList(true, []string{"12  ", "  1", "3  ", " 2", "  1", "  1"}),
			want:  []string{"1 ", "  1", "  1", "12  ", " 2", "   3"},
			value: "1",
			err:   nil,
		},
		{
			name:  "Test4",
			input: GetOrderedList(false, []string{"12  ", "  1", "3  ", " 2"}),
			want:  []string{"  3", "2  ", "12", " 1 "},
			value: "12",
			err:   nil,
		},
		{
			name:  "Test5",
			input: GetOrderedList(false, []string{"  1", "3  ", " 2"}),
			want:  []string{"  3", "2  ", " 1 "},
			value: "12",
			err:   errors.New("node is not finding"),
		},
		{
			name:  "Test6",
			input: GetOrderedList(false, []string{"12  ", "  1", "3  ", " 2", "  1", "  1"}),
			want:  []string{"  3", "2  ", "12", " 1 ", " 1 ", " 1 "},
			value: "1",
			err:   nil,
		},
	}

	for _, test := range testsInt {
		node, err := test.input.Find(test.value)

		if test.input.Compare(node.value, test.value) != 0 {
			t.Errorf("failed %s: wrong node. Result is: %v, want: %v", test.name, node.value, test.value)
		}

		if !EqualOrderedLists(test.input, test.want) {
			t.Errorf("failed %s: not equals lists. Result is: ", test.name)
			PrintList(*test.input)
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
	for _, test := range testsFloat {
		node, err := test.input.Find(test.value)

		if test.input.Compare(node.value, test.value) != 0 {
			t.Errorf("failed %s: wrong node. Result is: %v, want: %v", test.name, node.value, test.value)
		}

		if !EqualOrderedLists(test.input, test.want) {
			t.Errorf("failed %s: not equals lists. Result is: ", test.name)
			PrintList(*test.input)
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
	for _, test := range testsString {
		node, err := test.input.Find(test.value)

		if test.input.Compare(node.value, test.value) != 0 {
			t.Errorf("failed %s: wrong node. Result is: %v, want: %v", test.name, node.value, test.value)
		}

		if !EqualOrderedLists(test.input, test.want) {
			t.Errorf("failed %s: not equals lists. Result is: ", test.name)
			PrintList(*test.input)
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

func TestDelete(t *testing.T) {
	type testType[T constraints.Ordered] struct {
		name  string
		input *OrderedList[T]
		value T
		want  []T
	}

	testsInt := []testType[int64]{
		{
			name:  "Test1",
			input: GetOrderedList(true, []int64{9, 1, 5, 2, 7, 4, 3}),
			value: 3,
			want:  []int64{1, 2, 4, 5, 7, 9},
		},
		{
			name:  "Test2",
			input: GetOrderedList(true, []int64{1}),
			value: 1,
			want:  []int64{},
		},
		{
			name:  "Test3",
			input: GetOrderedList(false, []int64{9, 1, 5, 2, 7, 3, 4}),
			value: 3,
			want:  []int64{9, 7, 5, 4, 2, 1},
		},
		{
			name:  "Test5",
			input: GetOrderedList(false, []int64{1, 3}),
			value: 1,
			want:  []int64{3},
		},
		{
			name:  "Test6",
			input: GetOrderedList(true, []int64{9, 1, 5, 2, 7, 3, 4, 3}),
			value: 3,
			want:  []int64{1, 2, 3, 4, 5, 7, 9},
		},
		{
			name:  "Test7",
			input: GetOrderedList(false, []int64{9, 1, 5, 2, 7, 3, 4, 3}),
			value: 3,
			want:  []int64{9, 7, 5, 4, 3, 2, 1},
		},
	}
	testsFloat := []testType[float64]{
		{
			name:  "Test1",
			input: GetOrderedList(true, []float64{2.1, 1.2, 5.6, 1.1, 3.4}),
			value: 3.4,
			want:  []float64{1.1, 1.2, 2.1, 5.6},
		},
		{
			name:  "Test2",
			input: GetOrderedList(true, []float64{2.1}),
			value: 2.1,
			want:  []float64{},
		},
		{
			name:  "Test3",
			input: GetOrderedList(true, []float64{2.1, 3.4}),
			value: 3.4,
			want:  []float64{2.1},
		},
	}
	testsString := []testType[string]{
		{
			name:  "Test1",
			input: GetOrderedList(true, []string{"  1", "3  ", " 12"}),
			value: "12",
			want:  []string{"1 ", "   3"},
		},
		{
			name:  "Test2",
			input: GetOrderedList(true, []string{"1"}),
			value: "1",
			want:  []string{},
		},
		{
			name:  "Test3",
			input: GetOrderedList(true, []string{"   3", "4  "}),
			value: "4",
			want:  []string{"3"},
		},
	}

	for _, test := range testsInt {
		test.input.Delete(test.value)
		if !EqualOrderedLists(test.input, test.want) {
			t.Errorf("failed %s: not equals lists. Result is: ", test.name)
			PrintList(*test.input)
		}
	}
	for _, test := range testsFloat {
		test.input.Delete(test.value)
		if !EqualOrderedLists(test.input, test.want) {
			t.Errorf("failed %s: not equals lists. Result is: ", test.name)
			PrintList(*test.input)
		}
	}
	for _, test := range testsString {
		test.input.Delete(test.value)
		if !EqualOrderedLists(test.input, test.want) {
			t.Errorf("failed %s: not equals lists. Result is: ", test.name)
			PrintList(*test.input)
		}
	}
}
