package main

import (
	"fmt"
)

type Node struct {
	next  *Node
	value int
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (l *LinkedList) AddInTail(item Node) {
	if l.head == nil {
		l.head = &item
	} else {
		l.tail.next = &item
	}
	l.tail = &item
}

func (l *LinkedList) Count() int {
	var count int
	tempNode := l.head
	for tempNode != nil {
		count++
		tempNode = tempNode.next
	}
	return count
}

// error не nil, если узел не найден
func (l *LinkedList) Find(n int) (Node, error) {
	tempNode := l.head
	for tempNode != nil {
		if tempNode.value == n {
			return *tempNode, nil
		}
		tempNode = tempNode.next
	}
	return Node{value: -1, next: nil}, fmt.Errorf("node with value %d not finding", n)
}

func (l *LinkedList) FindAll(n int) []Node {
	var nodes []Node
	tempNode := l.head
	for tempNode != nil {
		if tempNode.value == n {
			nodes = append(nodes, *tempNode)
		}
		tempNode = tempNode.next
	}
	return nodes
}

func (l *LinkedList) Delete(n int, all bool) {
	if l.head == nil {
		return
	}

	tempNode := l.head
	prev := Node{
		next: l.head,
	}
	if all {

	} else {
		for tempNode != nil {
			if tempNode.value == n {
				if prev.next == l.head {
					l.head = tempNode.next
				} else if prev.next == tempNode {
					prev.next = tempNode.next
					if tempNode == l.tail {
						l.tail = &prev
					}
				}
				tempNode = tempNode.next
			}
		}

	}
}

func (l *LinkedList) Insert(after *Node, add Node) {
	tempNode := l.head
	for tempNode != after {
		tempNode = tempNode.next
	}
	if tempNode == l.tail {
		l.AddInTail(add)
	} else {
		nextNode := after.next
		after.next = &add
		add.next = nextNode
	}

}

func (l *LinkedList) InsertFirst(first Node) {
	if l.head == nil {
		l.tail = &first
	} else {
		first.next = l.head
	}
	l.head = &first

}

func (l *LinkedList) Clean() {
	l.head = nil
	l.tail = nil
}

// func PrintLL(LL *LinkedList) {
// 	temp := LL.head
// 	for temp != nil {
// 		fmt.Println(temp.value)
// 		temp = temp.next
// 	}
// }
