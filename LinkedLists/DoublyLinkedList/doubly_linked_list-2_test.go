package main

import (
	"errors"
	"testing"
)

//	func printTesting(l *LinkedList2, t *testing.T) {
//		temp := l.head
//		for temp != nil {
//			t.Errorf("%d ", temp.value)
//			temp = temp.next
//		}
//		t.Error()
//	}

func TestFindAll(t *testing.T) {
	tests := []struct {
		name  string
		input *LinkedList2
		value int
		want  int
	}{
		{"Test1: ", GetLinkedList([]int{}), 1, 0},
		{"Test2: ", GetLinkedList([]int{22, 2, 77, 6, 22, 76, 89}), 22, 2},
		{"Test3: ", GetLinkedList([]int{1}), 1, 1},
	}

	for _, tempTest := range tests {
		tempNodes := tempTest.input.FindAll(tempTest.value)
		if len(tempNodes) != tempTest.want {
			t.Errorf("failed %s: wrong finding nodes", tempTest.name)
		}
	}
}

func TestFind(t *testing.T) {
	tests := []struct {
		name     string
		input    *LinkedList2
		findNode Node
		want     Node
		err      error
	}{
		{"Test1: ", GetLinkedList([]int{}), Node{value: 1}, Node{value: -1}, errors.New("node is not finding")},
		{"Test2: ", GetLinkedList([]int{22, 1, 4, 5, 33}), Node{value: 1}, Node{value: 1}, nil},
		{"Test3: ", GetLinkedList([]int{22, 1, 4, 5, 33}), Node{value: 10}, Node{value: -1}, errors.New("node is not finding")},
		{"Test4: ", GetLinkedList([]int{22}), Node{value: 1}, Node{value: -1}, errors.New("node is not finding")},
		{"Test5: ", GetLinkedList([]int{22}), Node{value: 22}, Node{value: 22}, nil},
	}

	for _, tempTest := range tests {
		tempNode, errFind := tempTest.input.Find(tempTest.findNode.value)
		if tempNode.value != tempTest.want.value && errFind != tempTest.err {
			t.Errorf("failed %s: find value: %v", tempTest.name, tempTest.findNode)

		}
	}
}

func TestCount(t *testing.T) {
	tests := []struct {
		name      string
		input     *LinkedList2
		wantCount int
	}{
		{"Test1: ", GetLinkedList([]int{}), 0},
		{"Test2: ", GetLinkedList([]int{22, 2, 77, 6, 22, 76, 89}), 7},
		{"Test3: ", GetLinkedList([]int{1}), 1},
	}

	for _, tempTest := range tests {
		len := tempTest.input.Count()
		if len != tempTest.wantCount {
			t.Errorf("failed %s: wrong count list", tempTest.name)
		}
	}
}

func TestClean(t *testing.T) {
	tests := []struct {
		name  string
		input *LinkedList2
		want  *LinkedList2
	}{
		{"Test1: ", GetLinkedList([]int{}), GetLinkedList([]int{})},
		{"Test2: ", GetLinkedList([]int{22, 2, 77, 6, 22, 76, 89}), GetLinkedList([]int{})},
		{"Test3: ", GetLinkedList([]int{1}), GetLinkedList([]int{})},
	}

	for _, tempTest := range tests {
		tempTest.input.Clean()
		if !EqualLists(tempTest.input, tempTest.want) {
			t.Errorf("failed %s: clean list", tempTest.name)
		}
	}
}

func TestDeleted(t *testing.T) {
	tests := []struct {
		name   string
		input  *LinkedList2
		values []int
		all    bool
		want   *LinkedList2
	}{
		{"Test1: ", GetLinkedList([]int{}), []int{6}, false, GetLinkedList([]int{})},
		{"Test2: ", GetLinkedList([]int{22, 2, 77, 6, 22, 76, 77, 89}), []int{22, 77}, true, GetLinkedList([]int{2, 6, 76, 89})},
		{"Test3: ", GetLinkedList([]int{1}), []int{1}, false, GetLinkedList([]int{})},
		{"Test4: ", GetLinkedList([]int{22, 2, 77, 6, 22, 76, 89}), []int{6, 2, 89}, false, GetLinkedList([]int{22, 77, 22, 76})},
		{"Test5: ", GetLinkedList([]int{22, 2, 77, 6, 22, 76, 89}), []int{6}, true, GetLinkedList([]int{22, 2, 77, 22, 76, 89})},
		{"Test6: ", GetLinkedList([]int{22, 2, 77, 6, 6, 22, 76, 89}), []int{6}, true, GetLinkedList([]int{22, 2, 77, 22, 76, 89})},
	}

	for _, tempTest := range tests {
		test := tempTest
		for _, value := range test.values {
			test.input.Delete(value, test.all)
		}

		if !EqualLists(test.input, test.want) {
			t.Errorf("failed %s: deleting node with values %v", test.name, test.values)
		}
	}
}

func TestInsertFirst(t *testing.T) {
	tests := []struct {
		name  string
		input *LinkedList2
		node  Node
		want  *LinkedList2
	}{
		{"Test1: ", GetLinkedList([]int{}), Node{value: 1}, GetLinkedList([]int{1})},
		{"Test2: ", GetLinkedList([]int{22, 1, 3, 12}), Node{value: 10}, GetLinkedList([]int{10, 22, 1, 3, 12})},
		{"Test3: ", GetLinkedList([]int{22}), Node{value: 11}, GetLinkedList([]int{11, 22})},
	}

	for _, tempTest := range tests {
		test := tempTest
		test.input.InsertFirst(test.node)
		if !EqualLists(test.input, test.want) {
			t.Errorf("failed %s: insert first node", test.name)
		}
	}
}
func TestInsert(t *testing.T) {
	tests := []struct {
		name      string
		input     *LinkedList2
		afterNode *Node
		node      Node
		want      *LinkedList2
	}{
		{"Test1: ", GetLinkedList([]int{22, 10, 3, 12}), &Node{value: 10}, Node{value: 1}, GetLinkedList([]int{22, 10, 1, 3, 12})},
		{"Test2: ", GetLinkedList([]int{22, 10, 3, 12}), &Node{value: 12}, Node{value: 1}, GetLinkedList([]int{22, 10, 3, 12, 1})},
		{"Test3: ", GetLinkedList([]int{22}), &Node{value: 22}, Node{value: 11}, GetLinkedList([]int{22, 11})},
	}

	for _, tempTest := range tests {
		test := tempTest
		test.input.Insert(test.afterNode, test.node)
		if !EqualLists(test.input, test.want) {
			t.Errorf("failed %s: insert node after node", test.name)
		}
	}
}
