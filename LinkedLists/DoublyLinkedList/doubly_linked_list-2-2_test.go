package main

import (
	"fmt"
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

// tests for DLL2

func TestAddInTailDLL2(t *testing.T) {
	tests := []struct {
		name  string
		input DummyLinkedList2
		node  DummyNode
		want  DummyLinkedList2
	}{
		{"Test1: ", GetDLL2([]int{}), DummyNode{value: 1}, GetDLL2([]int{1})},
		{"Test2: ", GetDLL2([]int{22, 1, 3, 12}), DummyNode{value: 10}, GetDLL2([]int{22, 1, 3, 12, 10})},
		{"Test3: ", GetDLL2([]int{22}), DummyNode{value: 11}, GetDLL2([]int{22, 11})},
	}

	for _, tempTest := range tests {
		test := tempTest
		test.input.AddInTail(test.node)
		if !EqualDLL2(test.input, test.want) {
			t.Errorf("failed %s: insert first node", test.name)
		}
	}
}

func TestCountDLL2(t *testing.T) {
	tests := []struct {
		name  string
		input DummyLinkedList2
		want  int
	}{
		{"Test1: ", GetDLL2([]int{}), 0},
		{"Test2: ", GetDLL2([]int{22, 1, 3, 12}), 4},
		{"Test3: ", GetDLL2([]int{22}), 1},
	}

	for _, test := range tests {
		if test.input.Count() != test.want {
			t.Errorf("failed %s: wrong size", test.name)
			fmt.Println(test.input.Count())
		}
	}
}

func TestInsertDLL2(t *testing.T) {
	tests := []struct {
		name  string
		input DummyLinkedList2
		node  DummyNode
		after DummyNode
		want  DummyLinkedList2
	}{
		{"Test1", GetDLL2([]int{1, 2, 3, 4}),
			DummyNode{
				isDummy: false,
				value:   5,
			}, DummyNode{
				isDummy: false,
				value:   2,
			}, GetDLL2([]int{1, 2, 5, 3, 4})},
		{"Test2", GetDLL2([]int{1}),
			DummyNode{
				isDummy: false,
				value:   5,
			}, DummyNode{
				isDummy: false,
				value:   1,
			}, GetDLL2([]int{1, 5})},
		{"Test3", GetDLL2([]int{}),
			DummyNode{
				isDummy: false,
				value:   5,
			}, DummyNode{
				isDummy: false,
				value:   1,
			}, GetDLL2([]int{5})},
		{"Test4", GetDLL2([]int{1, 2, 3, 4}),
			DummyNode{
				isDummy: false,
				value:   5,
			}, DummyNode{
				isDummy: false,
				value:   5,
			}, GetDLL2([]int{1, 2, 3, 4})},
	}

	for _, test := range tests {
		test.input.Insert(&test.after, test.node)

		if !EqualDLL2(test.input, test.want) {
			t.Errorf("%s failed: wrong insert", test.name)
		}
	}
}

func TestInsertFirstDLL2(t *testing.T) {
	tests := []struct {
		name  string
		input DummyLinkedList2
		node  DummyNode
		want  DummyLinkedList2
	}{
		{"Test1", GetDLL2([]int{1, 2, 3, 4}),
			DummyNode{
				isDummy: false,
				value:   5,
			}, GetDLL2([]int{5, 1, 2, 3, 4})},
	}

	for _, test := range tests {
		test.input.InsertFirst(test.node)

		if !EqualDLL2(test.input, test.want) {
			t.Errorf("%s failed: wrong insert first", test.name)
		}
	}
}

func TestDeleteDLL2(t *testing.T) {
	tests := []struct {
		name  string
		input DummyLinkedList2
		value int
		all   bool
		want  DummyLinkedList2
	}{
		{"Test1", GetDLL2([]int{1, 2, 3, 4}),
			2, false, GetDLL2([]int{1, 3, 4})},

		{"Test2", GetDLL2([]int{1, 2, 3, 4}),
			2, true, GetDLL2([]int{1, 3, 4})},

		{"Test3", GetDLL2([]int{1, 2, 3, 2, 2, 4}),
			2, true, GetDLL2([]int{1, 3, 4})},

		{"Test4", GetDLL2([]int{1, 2, 3, 4}),
			5, true, GetDLL2([]int{1, 2, 3, 4})},

		{"Test5", GetDLL2([]int{}),
			5, true, GetDLL2([]int{})},

		{"Test6", GetDLL2([]int{2}),
			2, true, GetDLL2([]int{})},

		{"Test7", GetDLL2([]int{2, 2, 2, 2}),
			2, false, GetDLL2([]int{2, 2, 2})},

		{"Test8", GetDLL2([]int{2, 2, 2, 2}),
			2, true, GetDLL2([]int{})},
	}

	for _, test := range tests {
		test.input.Delete(test.value, test.all)

		if !EqualDLL2(test.input, test.want) {
			t.Errorf("%s failed: wrong delete", test.name)
		}
	}
}
