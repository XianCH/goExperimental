package btree

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	Root *Node
	//树中的总节点数
	Size int
}

func (bt *BinaryTree) Insert(data int) *BinaryTree {
	if bt.Root == nil {
		newNode := &Node{Data: data}
		bt.Root = newNode
	} else {
		bt.Root.insert(data)
	}
	bt.Size++
	return bt
}

func (node *Node) insert(data int) {
	if data < node.Data {
		if node.Left == nil {
			newNode := &Node{Data: data, Right: nil, Left: nil}
			node.Left = newNode
		}
		node.Left.insert(data)
	} else {
		if node.Right == nil {
			newNode := &Node{Data: data, Right: nil, Left: nil}
			node.Right = newNode
		}
		node.Right.insert(data)
	}
}

// TraverseInOrder 中序遍历二叉树
func (b *BinaryTree) TraverseInOrder(f func(int)) {
	if b.Root != nil {
		b.Root.traverseInOrder(f)
	}
}

// traverseInOrder 中序遍历节点（辅助函数）
func (n *Node) traverseInOrder(f func(int)) {
	if n != nil {
		n.Left.traverseInOrder(f)
		f(n.Data)
		n.Right.traverseInOrder(f)
	}
}

func (bt *BinaryTree) SizeOf() int {
	return bt.Size
}
