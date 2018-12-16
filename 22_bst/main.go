package _2_bst

// 二分查找树 binary search tree
// 二分查找树的典型特征《必记》:
// 1. 只能有最多，两个节点(左，右节点)
// 2. 左 < 父 < 右

// 一个二分查找树，主要的方法就是两个:
// Add 用于构建树
// Search 用于查找某一个元素是否存在

// 面试中：最最常见的一个关于二分查找的问题，如何验证一个二分树，是否满足二分查找树

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

// 验证一个二分树是否满足BST
// 牢记一点，并不是 左右子树都是二分树，父树就是二分树的，特例 [10,5,15,null,null,6,20]
// 二分查找树的特性并不是 ： 左节点 < 父 < 右节点，这是bstnode的特性
// 二分查找树的特性是 ： ** 左树的最大 < 父 < 右树的最小 且 左树，右树，也是二分查找树 **
// 所以问题就转化为了，求一颗树的最大和最小
// 对应 leetCode 98 题,数据结构就采用 leetCode上的数据结构了

// todo 优化： 目前我按照 大于左边最大，小于右边最小来计算，会导致会对 min, max 进行多次的重复计算，好处是思路，代码清晰
// 优化思路是，** 左子树要小于的值 = 父节点要大于的值， 右子树要大于的值 = 父节点要小于的值 ** 重要
// 即 isV(node *TreeNode, min int, minExit bool, max int, maxExit bool)
// 然后 对于左右就要在多判断集中情况
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(node *TreeNode) bool {
	if node == nil {
		return true
	}
	// 左
	var leftStatus bool
	switch {
	case node.Left != nil && node.Val <= min(node.Left):
		leftStatus = false
	case node.Left != nil && node.Val > min(node.Left):
		leftStatus = isValidBST(node.Left)
	default:
		leftStatus = true
	}
	// 右
	var rightStatus bool
	switch {
	case node.Right != nil && node.Val >= max(node.Right):
		rightStatus = false
	case node.Right != nil && node.Val < max(node.Right):
		rightStatus = isValidBST(node.Right)
	default:
		rightStatus = true
	}
	// 必须同时为真
	return leftStatus && rightStatus
}

// 最小
func min(node *TreeNode) int {
	for node.Right != nil {
		node = node.Right
	}
	return node.Val
}

// 最大
func max(node *TreeNode) int {
	for node.Left != nil {
		node = node.Left
	}
	return node.Val
}

// 之前的思维，太过于宏观，不够微观，不够一步一步来
// 判断第一个节点的时候，不用去想，是否满足 大于左边整个子树的最大值，也不考虑右边，只是考虑自己
// 把烦人的判断交给后面的节点
// 比如它的左节点（以及左节点所有的子节点），就要满足不能超过上限，就是当前节点的值，而它的右边节点（以及所有的子节点），要满足一个下限，就是不能低于当前节点的值
// 此时来描述一个二分查找树：
// 《重要》
// 我大于（右子节点），则我的所有孩子，不论左右，都会大于
// 我小于（左子节点），则我的所有孩子，不论左右，都会小于
// 《重要》
func validBST(node *TreeNode, min *int, max *int) bool {
	switch {
	case node == nil:
		return true
	case min != nil && node.Val <= *min:
		return false
	case max != nil && node.Val >= *max:
		return false
	case node.Left != nil && !validBST(node.Left, min, &node.Val):
		return false
	case node.Right != nil && !validBST(node.Right, &node.Val, max):
		return false
	default:
		return true
	}
}
