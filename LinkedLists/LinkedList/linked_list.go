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
	var prev *Node
	deleted := false

	for tempNode != nil {
		if tempNode.value == n && tempNode == l.head {
			l.head = tempNode.next
			deleted = true
		} else if tempNode.value == n && tempNode == l.tail {
			prev.next = nil
			l.tail = prev
			deleted = true
		} else if tempNode.value == n {
			prev.next = tempNode.next
			deleted = true
		}
		if !all && deleted {
			break
		}
		prev = tempNode
		tempNode = tempNode.next
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

func PrintLL(LL *LinkedList) {
	temp := LL.head
	for temp != nil {
		fmt.Printf("%d ", temp.value)
		temp = temp.next
	}
	fmt.Println()
}

func GetLinkedList(values []int) *LinkedList {
	var resultLL LinkedList // resulting linked list
	for _, value := range values {
		resultLL.AddInTail(Node{
			value: value,
		})
	}
	return &resultLL
}
func main() {
	l := GetLinkedList([]int{22, 2, 77, 6, 22, 43, 22, 76, 89})
	fmt.Println("Before deleting: ")
	PrintLL(l)
	fmt.Println("After deleting: ")
	l.Delete(89, false)
	PrintLL(l)
	fmt.Println("After deleting: ")
	l.Delete(22, true)
	PrintLL(l)
	fmt.Println("After deleting: ")
	l.Delete(6, true)
	PrintLL(l)

}
