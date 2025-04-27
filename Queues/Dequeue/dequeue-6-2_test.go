package main

import (
	"testing"
)

func TestPallindrom(t *testing.T) {
	tests := []struct {
		name string
		row  string
		want bool
	}{
		{"Test1", "abba", true},
		{"Test2", "abcba", true},
		{"Test3", "abcxba", false},
		{"Test4", "a", true},
		{"Test5", "ab", false},
		{"Test6", "aa", true},
		{"Test7", "", true},
	}

	for _, test := range tests {
		if IsPallindrom(test.row) != test.want {
			t.Errorf("%s: failed check pallindrom", test.name)
		}
	}
}

func TestGetMin(t *testing.T) {
	tests := []struct {
		name    string
		deq     MinDeque[int]
		wantMin int
		err     bool
	}{
		{"Test1", GetMinDequeue([]int{3, 1, 2, 4, 6, 1}), 1, false},
		{"Test2", GetMinDequeue([]int{}), 1, true},
	}

	for _, test := range tests {
		min, err_get_min := test.deq.GetMin()

		if !test.err && err_get_min != nil {
			t.Errorf("%s: failed get min, raise error", test.name)
			continue
		}

		if test.err && err_get_min == nil {
			t.Errorf("%s: failed get min, raise error", test.name)
			continue
		}

		if min != test.wantMin && err_get_min == nil {
			t.Errorf("%s: failed get min", test.name)
		}
	}
}
