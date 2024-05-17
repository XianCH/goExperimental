package cache_14n

import "container/list"

type Node struct {
	Value any
	Prev  *Node
	Next  *Node
}

func test() {
	ll := list.New()
	ll.Len()
}
