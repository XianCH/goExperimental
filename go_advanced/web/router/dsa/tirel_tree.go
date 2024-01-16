package dsa

import (
	"fmt"
)

type TirelNode struct {
	children map[rune]*TirelNode
	isEnd    bool
}

type TireTree struct {
	root *TirelNode
}

func NewTireTree() *TireTree {
	return &TireTree{root: &TirelNode{children: make(map[rune]*TirelNode)}}
}

// Insert node to TireTree
func (t *TireTree) Insert(word string) {
	node := t.root
	for _, char := range word {
		if node.children == nil {
			node.children = make(map[rune]*TirelNode)
		}
		child, exit := node.children[char]
		if !exit {
			child = &TirelNode{children: make(map[rune]*TirelNode)}
			node.children[char] = child
		}
		node = child
	}
	node.isEnd = true
}

// search tirl if exit the word
func (t *TireTree) Search(word string) bool {
	node := t.searchPrefix(word)
	return node != nil && node.isEnd
}

// search tire prefix by word
func (t *TireTree) searchPrefix(word string) *TirelNode {
	node := t.root
	for _, char := range word {
		child, exit := node.children[char]
		if !exit {
			return nil
		}
		node = child
	}
	return node
}

// StartsWith 检查 Trie 中是否有以给定前缀开头的单词
func (t *TireTree) StartsWith(prefix string) bool {
	return t.searchPrefix(prefix) != nil
}

// printTrie 打印 Trie 树结构
func printTrie(node *TirelNode, prefix string) {
	if node.isEnd {
		fmt.Println(prefix)
	}
	for char, child := range node.children {
		printTrie(child, prefix+string(char))
	}
}
func test_tire() {
	trie := NewTireTree()

	// 插入单词
	trie.Insert("apple")
	trie.Insert("app")
	trie.Insert("apricot")

	// 查找单词
	fmt.Println(trie.Search("apple"))   // true
	fmt.Println(trie.Search("app"))     // true
	fmt.Println(trie.Search("apricot")) // true
	fmt.Println(trie.Search("ap"))      // false

	// 检查前缀
	fmt.Println(trie.StartsWith("ap"))     // true
	fmt.Println(trie.StartsWith("appl"))   // true
	fmt.Println(trie.StartsWith("banana")) // false
}
