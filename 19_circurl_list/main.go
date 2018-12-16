package _9_circurl_list

import (
	"golang-coding-interview-questions-and-solutions/17_linked_list"
)

// 检测一个链表是否成环 loop
// 对应 两道 leetcode 的题，题号 141 ， 142
// 这个也是典型的要使用fast,slow指针的问题

// fast遇到null ，则表示无环
// fast遇到slow，则表示有环

// leetCode 8ms
func DelectLoop(node *linked_list.Node) bool {
	if node == nil || node.Next == nil {
		return false
	}
	slow, fast := node, node
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
