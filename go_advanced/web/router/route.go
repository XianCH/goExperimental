package main

import (
	"fmt"
	"strings"
)

type TrieNode struct {
	children map[string]*TrieNode
	isEnd    bool
	handler  func()
}

type Router struct {
	root *TrieNode
}

func NewRouter() *Router {
	return &Router{root: &TrieNode{children: make(map[string]*TrieNode)}}
}

func (r *Router) AddRoute(path string, handler func()) {
	node := r.root
	segments := strings.Split(path, "/")

	for _, segment := range segments {
		if segment == "" {
			continue
		}

		child, exists := node.children[segment]
		if !exists {
			child = &TrieNode{children: make(map[string]*TrieNode)}
			node.children[segment] = child
		}
		node = child
	}
	node.isEnd = true
	node.handler = handler
}

func (r *Router) FindHandler(path string) (func(), bool) {
	node := r.root
	segments := strings.Split(path, "/")

	for _, segment := range segments {
		if segment == "" {
			continue
		}

		child, exists := node.children[segment]
		if !exists {
			return nil, false
		}
		node = child
	}
	if node.isEnd {
		return node.handler, true
	}
	return nil, false
}

func main() {
	router := NewRouter()

	router.AddRoute("/home", func() {
		fmt.Println("Handler for /home")
	})

	router.AddRoute("/user/profile", func() {
		fmt.Println("Handler for /user/profile")
	})

	router.AddRoute("/user/settings", func() {
		fmt.Println("Handler for /user/settings")
	})

	paths := []string{"/home", "/user/profile", "/user/settings", "/unknown"}

	for _, path := range paths {
		handler, found := router.FindHandler(path)
		if found {
			handler()
		} else {
			fmt.Printf("No handler found for %s\n", path)
		}
	}
}
