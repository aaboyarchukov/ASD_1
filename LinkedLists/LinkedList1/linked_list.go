package main

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
	return 0
}

// error не nil, если узел не найден
func (l *LinkedList) Find(n int) (Node, error) {
	return Node{value: -1, next: nil}, nil
}

func (l *LinkedList) FindAll(n int) []Node {
	var nodes []Node
	return nodes
}

func (l *LinkedList) Delete(n int, all bool) {
	if l.head == nil {
		return
	}

	if all {

	} else {
		temp := l.head
		for temp != nil {
			if temp.value == n {
				print()
				break
			}
		}

	}
}

func (l *LinkedList) Insert(after *Node, add Node) {

}

func (l *LinkedList) InsertFirst(first Node) {

}

func (l *LinkedList) Clean() {

}
