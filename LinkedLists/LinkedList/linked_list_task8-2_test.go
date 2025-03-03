package main

import "testing"

func TestAdditionLL(t *testing.T) {
	tests := []struct {
		name    string
		inputL1 *LinkedList
		inputL2 *LinkedList
		wantL3  *LinkedList
	}{
		{"Test1: ", GetLinkedList([]int{}), GetLinkedList([]int{}), GetLinkedList([]int{})},
		{"Test2: ", GetLinkedList([]int{22, 3, 2, 45, 6}), GetLinkedList([]int{10, 11, 1, 2, 3}), GetLinkedList([]int{32, 14, 3, 47, 9})},
		{"Test3: ", GetLinkedList([]int{22, 3, 2, 45, 6}), GetLinkedList([]int{10, 11}), nil},
		{"Test4: ", GetLinkedList([]int{22}), GetLinkedList([]int{10}), GetLinkedList([]int{32})},
		{"Test5: ", GetLinkedList([]int{22}), GetLinkedList([]int{10, 11}), nil},
	}

	for _, tempTest := range tests {
		test := tempTest
		resultL := GetAdditionalLists(test.inputL1, test.inputL2)
		if !EqualLists(resultL, test.wantL3) {
			t.Errorf("failed %s: additional lists", test.name)
		}
	}
}
