package main

import (
	"testing"
)

func TestDeleted(t *testing.T) {
	tests := []struct {
		name  string
		input LinkedList
		value int
		want  LinkedList
	}{
		{"Test1: ", getLinkedList([]int{}), 6, getLinkedList([]int{})},
		{"Test2: ", getLinkedList([]int{1, 2, 3, 4, 5}), 5, getLinkedList([]int{1, 2, 3, 4})},
		{"Test3: ", getLinkedList([]int{1, 2, 3}), 4, getLinkedList([]int{1, 2, 3})},
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

func getLinkedList(values []int) LinkedList {
	var resultLL LinkedList // resulting linked list
	for _, value := range values {
		resultLL.AddInTail(Node{
			value: value,
		})
	}
	return resultLL
}

func equalLists(l1 LinkedList, l2 LinkedList) bool {

	return false
}
