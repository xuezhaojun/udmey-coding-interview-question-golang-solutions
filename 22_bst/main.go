package _2_bst

// 二分查找树 binary search tree
// 二分查找树的典型特征《必记》:
// 1. 只能有最多，两个节点(左，右节点)
// 2. 左 < 父 < 右

// 一个二分查找树，主要的方法就是两个:
// Add 用于构建树
// Search 用于查找某一个元素是否存在
type BSTNode struct {
	Data  int
	Left  *BSTNode
	Right *BSTNode
}

func NewBSTNode(data int) *BSTNode {
	return &BSTNode{
		Data: data,
	}
}

func (node *BSTNode) add(data int) {
	switch {
	case data < node.Data && node.Left != nil:
		node.Left.add(data)
	case node.Data < data && node.Right != nil:
		node.Right.add(data)
	case data < node.Data && node.Left == nil:
		node.Left = NewBSTNode(data)
	case node.Data < data && node.Right == nil:
		node.Right = NewBSTNode(data)
	}
}

func (node *BSTNode) search(data int) bool {
	switch {
	case node.Data == data:
		return true
	case data < node.Data && node.Left != nil:
		return node.Left.search(data)
	case node.Data < data && node.Right != nil:
		return node.Right.search(data)
	default:
		return false
	}
}

type BSTTree struct {
	Root *BSTNode
}

func NewBSTTree() *BSTTree {
	return &BSTTree{}
}

func (tree *BSTTree) Add(data int) {
	if tree.Root == nil {
		tree.Root = NewBSTNode(data)
		return
	}
	tree.Root.add(data)
}

func (tree *BSTTree) Search(data int) bool {
	if tree.Root == nil {
		return false
	}
	return tree.Root.search(data)
}
