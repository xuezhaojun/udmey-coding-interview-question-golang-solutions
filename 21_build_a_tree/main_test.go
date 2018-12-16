package _1_build_a_tree

import "testing"

func TestTree_BFS(t *testing.T) {
	tree := NewTree()
	tree.Root = NewNode(1)
	tree.Root.Add(2)
	tree.Root.Add(3)
	tree.Root.Children[2].Add(4)
	tree.Root.Children[2].Add(5)
	tree.Root.Children[2].Add(6)
	tree.Root.Children[3].Add(7)
	tree.Root.Children[2].Children[4].Add(8)
	//    1
	//  2   3
	// 466   7
	// 8
	if tree.BFS(9) == true {
		t.Error("no 9")
	}
	if tree.BFS(6) == false {
		t.Error("has 6")
	}
}

func TestTree_DFS(t *testing.T) {
	tree := NewTree()
	tree.Root = NewNode(1)
	tree.Root.Add(2)
	tree.Root.Add(3)
	tree.Root.Children[2].Add(4)
	tree.Root.Children[2].Add(5)
	tree.Root.Children[2].Add(6)
	tree.Root.Children[3].Add(7)
	tree.Root.Children[2].Children[4].Add(8)
	//    1
	//  2   3
	// 466   7
	// 8
	if tree.DFS(9) == true {
		t.Error("no 9")
	}
	if tree.DFS(6) == false {
		t.Error("has 6")
	}
}
