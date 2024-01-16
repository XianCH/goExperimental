package main

import (
	"fmt"
	"net/http"
	"strings"
)

type TrieNode struct {
	children map[string]*TrieNode
	isEnd    bool
	handler  func(http.ResponseWriter, *http.Request)
}

type Router struct {
	root *TrieNode
}

func NewRouter() *Router {
	return &Router{root: &TrieNode{children: make(map[string]*TrieNode)}}
}

func (r *Router) AddRoute(path string, handler func(http.ResponseWriter, *http.Request)) {
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

func (r *Router) FindHandler(path string) (func(http.ResponseWriter, *http.Request), bool) {
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

	router.AddRoute("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Handler for /home")
	})

	router.AddRoute("/user/profile", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Handler for /user/profile")
	})

	router.AddRoute("/user/settings", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Handler for /user/settings")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler, found := router.FindHandler(r.URL.Path)
		if found {
			handler(w, r)
		} else {
			http.NotFound(w, r)
		}
	})

	fmt.Println("Server is listening on :8080")
	http.ListenAndServe(":8080", nil)
}
