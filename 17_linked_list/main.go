package _7_linked_list

import (
	"fmt"
	"strings"
)

// 最常在面试中出现的数据结构

// 如何实现一个linked list
type Node struct {
	Data interface{}
	Next *Node
}

func NewNode(data interface{}, nextNode *Node) *Node {
	return &Node{
		Data: data,
		Next: nextNode,
	}
}

// 整理一下实现一个LinkedList，要考虑的各种特殊情况：
// 1. 新增即换头
// 2. 头很重要： 对head是否等于nil的判断，对于 删除，新增操作等
// 3. 头尾一体的特殊情况： 在删除尾巴的时候，如果头尾一体，即只有一个节点，也要当作特殊情况处理
// 3. size 方法，如果不是用遍历的方式，则对应所有的新增和删除都要添加对size的操作
// 4. 带有index都要考虑到3种情况，头，尾，中间，并且遍历的时候，要找index-1的位置，这个位置是操作位置
type LinkedList struct {
	head *Node
	size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) InsertFirst(data interface{}) {
	newNode := NewNode(data, nil)
	if l.head != nil {
		newNode.Next = l.head
	}
	l.head = newNode
	l.size += 1
}

// InsertLast 和size一样，有两种实现方式：
// O(n) 的方式是，每次遍历到最后一个节点，然后进行新增
// O(1) 的方法是，保存一个tail的标记，但是则需要在 Insert,Remove 等多个方法中维护这个标记,就会很复杂
func (l *LinkedList) InsertLast(data interface{}) {
	newNode := NewNode(data, nil)
	if l.head == nil {
		l.head = newNode
		return
	}
	node := l.head
	for node.Next != nil {
		node = node.Next
	}
	node.Next = newNode

	l.size += 1
}

func (l *LinkedList) GetFirst() interface{} {
	// 记住要判断head是否为nil
	if l.head == nil {
		return nil
	}
	return l.head.Data
}

func (l *LinkedList) GetLast() interface{} {
	node := l.head
	if node == nil {
		return nil
	}
	for node.Next != nil {
		node = node.Next
	}
	return node.Data
}

func (l *LinkedList) Head() *Node {
	return l.head
}

func (l *LinkedList) Size() int {
	// 用遍历的一边的方式获取linked的大小
	//size := 0
	//var node *Node
	//node = l.head
	//for node != nil {
	//	size += 1
	//	node = node.Next
	//}
	return l.size
}

func (l *LinkedList) Clear() {
	node := l.head
	if node == nil {
		return
	}
	for node != nil {
		nextnode := node.Next
		node.Next = nil
		node = nextnode
	}
	// 将 head 设置为空（重要）,同时要将 size 设置为 0
	l.head = nil
	l.size = 0
}

func (l *LinkedList) RemoveFirst() {
	if l.head == nil {
		return
	}

	oldhead := l.head
	l.head = oldhead.Next
	oldhead.Next = nil

	l.size -= 1
}

func (l *LinkedList) RemoveLast() {
	defer func() {
		l.size -= 1
	}()

	if l.head == nil {
		return
	}
	// 只有一个节点的情况
	if l.head.Next == nil {
		l.head = nil
		return
	}

	node := l.head
	// 重要，不能这么写，应为当只有一个node的时候，不存在倒数第二个node
	// 永远只找倒数第一，不找倒数第二
	for node.Next.Next != nil {
		node = node.Next
	}

	node.Next = nil
}

func (l *LinkedList) GetAt(index int) (data interface{}) {
	if l.size == 0 || index > l.size-1 || index < 0 {
		return nil
	}
	node := l.head
	count := 0
	for node.Next != nil {
		if count == index {
			break
		}
		node = node.Next
		count += 1
	}
	return node.Data
}

func (l *LinkedList) RemoveAt(index int) {
	if l.size == 0 || index > l.size-1 || index < 0 {
		return
	}
	// 2种情况
	// 1. 有一个节点
	// 2. 有多个节点

	defer func() {
		l.size -= 1
	}()

	// 只有一个节点
	if index == 0 && l.size == 1 {
		l.RemoveFirst()
		return
	}

	// 有多个节点
	// 1. 删除头
	if index == 0 {
		l.RemoveFirst()
		return
	}
	// 2. 删除尾
	if index == l.size-1 {
		l.RemoveLast()
		return
	}
	// 3. 删除其他,找到index之前的那个节点
	node := l.head
	count := 0
	for node.Next != nil {
		if count == index-1 {
			break
		}
		node = node.Next
		count += 1
	}
	removeedNode := node.Next
	node.Next = node.Next.Next
	removeedNode.Next = nil
	return
}

func (l *LinkedList) String() string {
	var result []string
	var node *Node
	node = l.head
	for node != nil {
		result = append(result, fmt.Sprintf("%#v", node.Data))
		node = node.Next
	}
	return strings.Join(result, " -> ")
}
