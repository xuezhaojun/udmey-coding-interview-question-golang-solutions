package _2_bst

import "testing"

func TestBSTTree_Search(t *testing.T) {
	tree := NewBSTTree()
	tree.Add(10)
	tree.Add(0)
	tree.Add(12)
	tree.Add(-1)
	tree.Add(11)
	t.Log(tree.Search(-1) == true)
	t.Log(tree.Search(-2) == false)
	t.Log(tree.Search(12) == true)
}

func TestNewBSTNode(t *testing.T) {
	t1 := TreeNode{
		Val: 1,
	}
	t3 := TreeNode{
		Val: 3,
	}
	t2 := TreeNode{
		Val:   2,
		Left:  &t1,
		Right: &t3,
	}
	t.Log(validBST(&t2, nil, nil))
}
