package btree

import (
	"testing"
)

func TestBinaryTree(t *testing.T) {
	binaryTree := &BinaryTree{Root: nil, Size: 0}
	binaryTree.Insert(5)
	binaryTree.Insert(3)
	binaryTree.Insert(7)
	binaryTree.Insert(2)
	binaryTree.Insert(4)
	binaryTree.Insert(6)

	expected := []int{2, 3, 4, 5, 6, 7, 8}
	result := make([]int, 0)
	binaryTree.TraverseInOrder(func(i int) {
		result = append(result, i)
	})
	if !equalSlices(expected, result) {
		t.Errorf("In-order traversal result is incorrect. Got: %v, Expected: %v", result, expected)
	}
	if binaryTree.SizeOf() != len(expected) {
		t.Errorf("Tree size is incorrect. Got: %d, Expected: %d", binaryTree.SizeOf(), len(expected))
	}
}

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
