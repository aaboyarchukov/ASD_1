package main

import (
	"testing"
)

// func printTesting(l *LinkedList2, t *testing.T) {
// 	temp := l.head
// 	for temp != nil {
// 		t.Errorf("%d ", temp.value)
// 		temp = temp.next
// 	}
// 	t.Error()
// }

func TestReverse(t *testing.T) {
	tests := []struct {
		name  string
		input *LinkedList2
		want  *LinkedList2
	}{
		{"Test1: ", GetLinkedList([]int{1, 2, 3, 4}), GetLinkedList([]int{4, 3, 2, 1})},
		{"Test2: ", GetLinkedList([]int{}), GetLinkedList([]int{})},
		{"Test3: ", GetLinkedList([]int{1}), GetLinkedList([]int{1})},
	}

	for _, tempTest := range tests {
		test := tempTest
		test.input.ReverseList()
		if !EqualLists(test.input, test.want) {
			t.Errorf("failed %s: with reverse list", test.name)
		}
	}
}

func TestCyclic(t *testing.T) {
	tests := []struct {
		name  string
		input *LinkedList2
		item  Node
		cycle bool
		want  bool
	}{
		{"Test1: ", GetLinkedList([]int{1, 2, 3, 4}), Node{value: 2}, true, true},
		{"Test2: ", GetLinkedList([]int{}), Node{value: 2}, false, false},
		{"Test3: ", GetLinkedList([]int{1}), Node{value: 1}, true, true},
	}

	for _, tempTest := range tests {
		test := tempTest

		test.input.GetCyclicList(test.item)
		t.Logf("%s", test.name)
		if test.input.CyclicList() != test.want {
			t.Errorf("failed %s: with checking cycle", test.name)
		}

	}
}
func TestSort(t *testing.T) {
	tests := []struct {
		name  string
		input *LinkedList2
		want  *LinkedList2
	}{
		{"Test1", GetLinkedList([]int{3, 2, 1, 4}), GetLinkedList([]int{1, 2, 3, 4})},
		{"Test2", GetLinkedList([]int{1}), GetLinkedList([]int{1})},
		{"Test3", GetLinkedList([]int{}), GetLinkedList([]int{})},
		{"Test4", GetLinkedList([]int{7, 1, 3, 2, 6, 3, 4, 8, 8}), GetLinkedList([]int{1, 2, 3, 3, 4, 6, 7, 8, 8})},
	}

	for _, tempTest := range tests {
		test := tempTest
		test.input.SortList()

		if !EqualLists(test.input, test.want) {
			t.Errorf("failed %s: sort list", test.name)
		}
	}
}
func TestSortMerge(t *testing.T) {
	tests := []struct {
		name string
		l1   *LinkedList2
		l2   *LinkedList2
		l3   *LinkedList2
	}{
		{"Test1", GetLinkedList([]int{3, 2, 4, 1}), GetLinkedList([]int{7, 9, 1}), GetLinkedList([]int{1, 1, 2, 3, 4, 7, 9})},
		{"Test2", GetLinkedList([]int{3}), GetLinkedList([]int{1}), GetLinkedList([]int{1, 3})},
		{"Test3", GetLinkedList([]int{}), GetLinkedList([]int{}), GetLinkedList([]int{})},
	}

	for _, tempTest := range tests {
		test := tempTest
		result := SortMergeList(test.l1, test.l2)

		if !EqualLists(test.l3, result) {
			t.Errorf("failed %s: sort merge lists", test.name)
		}
	}
}
