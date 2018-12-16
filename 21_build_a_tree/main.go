package _1_build_a_tree

import "container/list"

// 能写出一个tree
// 能写出一个深度遍历

type Node struct {
	Data     int           // 为了简化起见，这里就用 int 类型了
	Children map[int]*Node // 使用map，为了方便删除
}

func NewNode(data int) *Node {
	return &Node{
		Data:     data,
		Children: make(map[int]*Node),
	}
}

func (n *Node) Add(data int) {
	node := NewNode(data)
	n.Children[data] = node
}

func (n *Node) Remove(data int) {
	delete(n.Children, data)
}

type Tree struct {
	Root *Node
}

func NewTree() *Tree {
	return &Tree{
		Root: nil,
	}
}

// Breadth First 根据出生的时间
// 如果你要返回每一层的节点数，则需要使用BFS
// BF 要用到 queue FIFO
func (t *Tree) BFS(data int) bool {
	l := list.New()
	l.PushBack(t.Root)
	for l.Len() != 0 {
		front := l.Front()
		if front.Value.(*Node).Data == data {
			return true
		}
		for _, node := range front.Value.(*Node).Children {
			l.PushBack(node)
		}
		l.Remove(front)
	}
	return false
}

// Depth Frist 根据深度
// DF 则要想到 递归
func (t *Tree) DFS(data int) bool {
	return search(t.Root, data)
}

func search(node *Node, data int) bool {
	// 叶子
	if len(node.Children) == 0 {
		if node.Data == data {
			return true
		}
		return false
	}
	// 非叶子
	var result bool
	for _, child := range node.Children {
		result = result || search(child, data)
	}
	return result
}
