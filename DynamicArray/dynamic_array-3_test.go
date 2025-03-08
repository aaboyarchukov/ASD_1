package main

import (
	"fmt"
	"testing"
)

// upgrade - x2, downgrade - x1.5
// min cap = 16
// if after downgrade len(buffer) < 16 => new len(buffer) = 16
// if len = cap => upgrade, if len / cap < 50% => downgrade

func TestMakeArray(t *testing.T) {
	tests := []struct {
		name  string
		array *DynArray[int]
		input int
		want  int
	}{
		{"Test1: ", &DynArray[int]{
			count:    1,
			capacity: 16,
			array:    []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		}, 3, 16},
		{"Test2: ", &DynArray[int]{
			count:    16,
			capacity: 16,
			array:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		}, 32, 32},
		{"Test3: ", &DynArray[int]{
			count:    1,
			capacity: 16,
			array:    []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		}, 20, 20},
	}

	for _, tempTest := range tests {
		test := tempTest
		test.array.MakeArray(test.input)
		if test.array.capacity != test.want {
			t.Errorf("failed %s: make array", test.name)
		}

	}
}

func TestAppend(t *testing.T) {
	tests := []struct {
		name  string
		input *DynArray[int]
		want  *DynArray[int]
		item  int
	}{
		{"Test1", &DynArray[int]{
			count:    2,
			capacity: 16,
			array:    []int{1, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
			&DynArray[int]{
				count:    3,
				capacity: 16,
				array:    []int{1, 2, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			}, 3},
		{"Test2", &DynArray[int]{
			count:    16,
			capacity: 16,
			array:    []int{1, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
			&DynArray[int]{
				count:    17,
				capacity: 32,
				array: []int{1, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			}, 5},
		{"Test3", &DynArray[int]{
			count:    0,
			capacity: 16,
			array:    []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
			&DynArray[int]{
				count:    1,
				capacity: 16,
				array:    []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			}, 1},
	}

	for _, tempTest := range tests {
		test := tempTest
		test.input.Append(test.item)
		if !EqualArrays(test.input, test.want) {
			t.Errorf("failed %s: append element", test.name)
		}
	}
}

func TestInsert(t *testing.T) {
	tests := []struct {
		name  string
		input *DynArray[int]
		want  *DynArray[int]
		item  int
		indx  int
	}{
		{"Test1", &DynArray[int]{
			count:    2,
			capacity: 16,
			array:    []int{1, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
			&DynArray[int]{
				count:    3,
				capacity: 16,
				array:    []int{1, 2, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			}, 3, 2},
		{"Test2", &DynArray[int]{
			count:    16,
			capacity: 16,
			array:    []int{1, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
			&DynArray[int]{
				count:    17,
				capacity: 32,
				array: []int{1, 2, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			}, 5, 4},
		{"Test3", &DynArray[int]{
			count:    16,
			capacity: 16,
			array:    []int{1, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
			&DynArray[int]{
				count:    17,
				capacity: 32,
				array: []int{1, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			}, 4, 16},
		{"Test4", &DynArray[int]{
			count:    32,
			capacity: 32,
			array: []int{1, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
			&DynArray[int]{
				count:    33,
				capacity: 64,
				array: []int{1, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					5, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			}, 5, 16},
	}

	for _, tempTest := range tests {
		test := tempTest
		test.input.Insert(test.item, test.indx)
		if !EqualArrays(test.input, test.want) {
			t.Errorf("failed %s: insert element", test.name)
		}
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		name  string
		input *DynArray[int]
		want  *DynArray[int]
		indx  int
		err   error
	}{
		{"Test1", &DynArray[int]{
			count:    2,
			capacity: 16,
			array:    []int{1, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
			&DynArray[int]{
				count:    1,
				capacity: 16,
				array:    []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			}, 1, nil},
		{"Test2", &DynArray[int]{
			count:    2,
			capacity: 16,
			array:    []int{1, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
			&DynArray[int]{
				count:    2,
				capacity: 16,
				array:    []int{1, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			}, -1, fmt.Errorf("bad index '%d'", -1)},
		{"Test3", &DynArray[int]{
			count:    18,
			capacity: 64,
			array: []int{1, 2, 0, 0, 4, 0, 0, 5, 0, 0, 0, 3, 0, 0, 1, 0,
				7, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
			&DynArray[int]{
				count:    17,
				capacity: 42,
				array: []int{1, 2, 0, 0, 4, 0, 0, 0, 0, 0, 3, 0, 0, 1, 0, 7,
					8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			}, 7, nil},
	}

	for _, tempTest := range tests {
		test := tempTest

		err := test.input.Remove(test.indx)

		if test.err == nil && err != nil {
			t.Errorf("failed %s: remove element raise error", test.name)
			continue
		}

		if !EqualArrays(test.input, test.want) {
			t.Errorf("failed %s: remove element", test.name)
			continue
		}
	}

	cycleTests := []struct {
		name  string
		input *DynArray[int]
		want  *DynArray[int]
		indx  int
		cycle int
		err   error
	}{
		{"Test1", &DynArray[int]{
			count:    4,
			capacity: 16,
			array:    []int{1, 2, 3, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
			&DynArray[int]{
				count:    0,
				capacity: 16,
				array:    []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			}, 1, 4, nil},
		{"Test2", &DynArray[int]{
			count:    12,
			capacity: 16,
			array:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 0, 0, 0, 0},
		},
			&DynArray[int]{
				count:    8,
				capacity: 16,
				array:    []int{1, 2, 3, 4, 5, 6, 7, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			}, 1, 4, nil},
	}

	for _, tempTest := range cycleTests {
		test := tempTest
		var err error

		for range test.cycle {
			err = test.input.Remove(test.input.count - 1)

		}

		if test.err == nil && err != nil {
			t.Errorf("failed %s: remove element raise error", test.name)
			continue
		}

		if !EqualArrays(test.input, test.want) {
			t.Errorf("failed %s: remove element", test.name)
			continue
		}
	}
}
func TestGetItem(t *testing.T) {
	tests := []struct {
		name    string
		input   *DynArray[int]
		indx    int
		element int
	}{{"Test1", &DynArray[int]{
		count:    1,
		capacity: 16,
		array:    []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}, 0, 1},
		{"Test2", &DynArray[int]{
			count:    1,
			capacity: 16,
			array:    []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		}, -1, 0}}

	for _, tempTest := range tests {
		test := tempTest
		tempElement, _ := test.input.GetItem(test.indx)

		if tempElement != test.element {
			t.Errorf("failed %s: get item from array", tempTest.name)
		}
	}
}

func TestEqualsArrays(t *testing.T) {
	tests := []struct {
		name string
		arr1 *DynArray[int]
		arr2 *DynArray[int]
		want bool
	}{
		{"Test1: ", &DynArray[int]{
			count:    2,
			capacity: 16,
			array:    []int{1, 2},
		}, &DynArray[int]{
			count:    2,
			capacity: 16,
			array:    []int{1, 2},
		}, true},
		{"Test2: ", &DynArray[int]{
			count:    4,
			capacity: 16,
			array:    []int{1, 2, 3, 4},
		}, &DynArray[int]{
			count:    2,
			capacity: 16,
			array:    []int{1, 2},
		}, false},
		{"Test3: ", &DynArray[int]{
			count:    0,
			capacity: 16,
			array:    []int{},
		}, &DynArray[int]{
			count:    0,
			capacity: 16,
			array:    []int{},
		}, true},
	}

	for _, tempTest := range tests {
		test := tempTest
		if EqualArrays(test.arr1, test.arr2) != test.want {
			t.Errorf("failed %s: arrays are not equal", tempTest.name)
		}
	}
}
