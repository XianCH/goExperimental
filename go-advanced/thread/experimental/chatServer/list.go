package chatServer

import "fmt"

type Node struct {
	data any
	next *Node
}

type List struct {
	Header *Node
}

func NewList() *List {
	return &List{
		Header: nil,
	}
}

func (l *List) isEmpty() bool {
	return l.Header == nil
}

func (l *List) Insert(data any) {
	node := &Node{data: data, next: nil}
	if l.isEmpty() {
		l.Header = node
		return
	}
	current := l.Header
	for {
		if current == nil {
			current.next = node
			return
		}
		current = current.next
	}
}

func (l *List) Display() {
	if l.isEmpty() {
		fmt.Println("list is empty")
		return
	}

	current := l.Header
	for {
		fmt.Println(current.data)
		if current.next == nil {
			return
		}
		current = current.next
	}
}
