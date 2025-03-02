package main

import (
	"testing"
)

//	func printTesting(l *LinkedList, t *testing.T) {
//		temp := l.head
//		for temp != nil {
//			t.Errorf("%d ", temp.value)
//			temp = temp.next
//		}
//		t.Error()
//	}
func TestDeleted(t *testing.T) {
	tests := []struct {
		name  string
		input *LinkedList
		value int
		all   bool
		head  int
		tail  int
		want  *LinkedList
	}{
		{"Test1: ", GetLinkedList([]int{}), 6, false, -1, -1, GetLinkedList([]int{})},
		{"Test2: ", GetLinkedList([]int{22, 2, 77, 6, 22, 76, 89}), 22, true, 2, 89, GetLinkedList([]int{2, 77, 6, 76, 89})},
		{"Test3: ", GetLinkedList([]int{1}), 1, false, -1, -1, GetLinkedList([]int{})},
		{"Test4: ", GetLinkedList([]int{22, 2, 77, 6, 22, 76, 89}), 6, false, 22, 89, GetLinkedList([]int{22, 2, 77, 22, 76, 89})},
		{"Test5: ", GetLinkedList([]int{22, 2, 77, 6, 22, 76, 89}), 6, true, 22, 89, GetLinkedList([]int{22, 2, 77, 22, 76, 89})},
	}

	for _, tempTest := range tests {
		tempTest.input.Delete(tempTest.value, tempTest.all)
		if !EqualLists(tempTest.input, tempTest.want) {
			t.Errorf("failed deleting node with value %d", tempTest.value)
		}
	}
}

func TestDeletedAll(t *testing.T)  {}
func TestCLean(t *testing.T)       {}
func TestFindAll(t *testing.T)     {}
func TestCount(t *testing.T)       {}
func TestInsertAfter(t *testing.T) {}

func TestAdditionLL(t *testing.T) {}
