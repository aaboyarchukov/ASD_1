package main

import (
	"testing"
)

func TestDeleted(t *testing.T) {
	tests := []struct {
		name  string
		input *LinkedList
		value int
		want  *LinkedList
	}{
		{"Test1: ", GetLinkedList([]int{}), 6, GetLinkedList([]int{})},
		{"Test2: ", GetLinkedList([]int{1, 2, 3, 4, 5}), 5, GetLinkedList([]int{1, 2, 3, 4})},
		{"Test3: ", GetLinkedList([]int{1, 2, 3}), 4, GetLinkedList([]int{1, 2, 3})},
	}

	for _, tempTest := range tests {
		tempTest.input.Delete(tempTest.value, false)
	}
}

func TestDeletedAll(t *testing.T)  {}
func TestCLean(t *testing.T)       {}
func TestFindAll(t *testing.T)     {}
func TestCount(t *testing.T)       {}
func TestInsertAfter(t *testing.T) {}

func TestAdditionLL(t *testing.T) {}

func equalLists(l1 LinkedList, l2 LinkedList) bool {

	return false
}
