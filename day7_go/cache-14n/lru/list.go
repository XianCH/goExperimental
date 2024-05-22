package lru

import (
	"errors"
	"sync"
)

type Node struct {
	Value      any
	Prev, Next *Node
}

type DoublyLinkedList struct {
	root *Node
	len  int
	mu   sync.RWMutex
}

func (dl *DoublyLinkedList) Init() *DoublyLinkedList {
	dl.root = &Node{}
	dl.root.Next = dl.root
	dl.root.Prev = dl.root
	dl.len = 0
	return dl
}

func NewList() *DoublyLinkedList {
	return new(DoublyLinkedList).Init()
}

func (dl *DoublyLinkedList) IsEmpty() bool {
	return dl.len == 0
}

func (dl *DoublyLinkedList) Len() int {
	dl.mu.RLock() // Use read lock
	defer dl.mu.RUnlock()
	return dl.len
}

// AddFront inserts a node at the front of the list
func (dl *DoublyLinkedList) AddFront(value any) {
	dl.mu.Lock()
	defer dl.mu.Unlock()
	newNode := &Node{Value: value}
	if dl.root.Next == nil {
		newNode.Next = dl.root
		newNode.Prev = dl.root
		dl.root.Next = newNode
		dl.root.Prev = newNode
	} else {
		newNode.Next = dl.root.Next
		newNode.Prev = dl.root
		dl.root.Next.Prev = newNode
		dl.root.Next = newNode
		// Update the previous first node's Prev pointer
		dl.root.Next.Next.Prev = newNode
	}
	dl.len++
}

// AddBack inserts a node at the back of the list
func (dl *DoublyLinkedList) AddBack(value any) {
	dl.mu.Lock()
	defer dl.mu.Unlock()
	newNode := &Node{Value: value}
	newNode.Prev = dl.root.Prev
	newNode.Next = dl.root
	dl.root.Prev.Next = newNode
	dl.root.Prev = newNode
	dl.len++
}

// GetFront returns the front node of the list
func (dl *DoublyLinkedList) GetFront() (*Node, error) { // Change return type to *Node
	dl.mu.RLock()
	defer dl.mu.RUnlock()
	if dl.IsEmpty() {
		return nil, errors.New("List is empty")
	}
	return dl.root.Next, nil
}

// GetBack returns the back node of the list
func (dl *DoublyLinkedList) GetBack() (*Node, error) { // Change return type to *Node
	dl.mu.RLock()
	defer dl.mu.RUnlock()
	if dl.IsEmpty() {
		return nil, errors.New("List is empty")
	}
	return dl.root.Prev, nil
}

// RemoveFront removes the front node from the list
func (dl *DoublyLinkedList) RemoveFront() error {
	dl.mu.Lock()
	defer dl.mu.Unlock()
	if dl.IsEmpty() {
		return errors.New("List is empty")
	}
	if dl.root.Next == dl.root { // if only root node exists
		dl.len = 0
		return nil
	}
	firstNode := dl.root.Next
	dl.root.Next = firstNode.Next
	firstNode.Next.Prev = dl.root
	firstNode.Next = nil // Set Next to nil to allow GC
	firstNode.Prev = nil // Set Prev to nil to allow GC
	dl.len--
	return nil
}

// RemoveBack removes the back node from the list

func (dl *DoublyLinkedList) RemoveBack() error {
	dl.mu.Lock()
	defer dl.mu.Unlock()
	if dl.IsEmpty() {
		return errors.New("List is empty")
	}
	if dl.root.Prev == dl.root { // if only root node exists
		dl.len = 0
		return nil
	}
	lastNode := dl.root.Prev
	dl.root.Prev = lastNode.Prev
	lastNode.Prev.Next = dl.root
	lastNode.Next = nil // Set Next to nil to allow GC
	lastNode.Prev = nil // Set Prev to nil to allow GC
	dl.len--
	return nil
}

func (dl *DoublyLinkedList) MoveToFront(node *Node) {
	dl.mu.Lock()
	defer dl.mu.Unlock()

	// if the node is the first
	if dl.root.Next == node {
		return
	}

	//Remove the node
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev

	//insert the node to front
	dl.root.Next.Prev = node
	node.Next = dl.root.Next
	node.Prev = dl.root
	dl.root.Next = node
}

func (dl *DoublyLinkedList) MoveToBack(node *Node) {
	dl.mu.Lock()
	defer dl.mu.Unlock()

	if dl.root.Prev == node {
		return
	}
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev

	dl.root.Prev.Next = node
	node.Next = dl.root
	node.Prev = dl.root.Prev
	dl.root.Prev = node
}

func (dl *DoublyLinkedList) PushNodeToFront(v any) *Node {
	dl.mu.Lock()
	defer dl.mu.Unlock()

	newNode := &Node{Value: v}
	if dl.root.Next == nil {
		// If the list is empty
		newNode.Next = dl.root
		newNode.Prev = dl.root
		dl.root.Next = newNode
		dl.root.Prev = newNode
	} else {
		newNode.Next = dl.root.Next
		newNode.Prev = dl.root
		dl.root.Next.Prev = newNode
		dl.root.Next = newNode
	}
	dl.len++
	return newNode
}

func (dl *DoublyLinkedList) RemoveNode(node *Node) {
	dl.mu.Lock()
	defer dl.mu.Unlock()

	if node == dl.root {
		return
	}

	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	node.Next = nil // Set Next to nil to allow GC
	node.Prev = nil // Set Prev to nil to allow GC
	dl.len--
}
