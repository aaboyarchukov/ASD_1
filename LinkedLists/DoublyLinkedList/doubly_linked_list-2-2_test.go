package main

import "testing"

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
		cycle bool
		want  bool
	}{
		{"Test1: ", GetLinkedList([]int{1, 2, 3, 4}), true, true},
		{"Test2: ", GetLinkedList([]int{}), false, false},
		{"Test3: ", GetLinkedList([]int{}), true, false},
		{"Test4: ", GetLinkedList([]int{1}), true, true},
		{"Test5: ", GetLinkedList([]int{1}), false, false},
	}

	for _, tempTest := range tests {
		test := tempTest
		if test.cycle && test.input.Count() > 0 {

		}

		if test.input.CyclicList() != test.want {
			t.Errorf("failed %s: with checking cycle", test.name)
		}

	}
}
func TestSort(t *testing.T) {

}
func TestSortMerge(t *testing.T) {

}
func TestInsertDummy(t *testing.T) {

}
func TestDeleteDummy(t *testing.T) {

}
func TestCrateDummyList(t *testing.T) {

}
