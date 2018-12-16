package _0_step_back_from_tail

import "golang-coding-interview-questions-and-solutions/17_linked_list"

// 找到距离tail n 个位置的节点
// 最直觉的算法是通过，size，但是，如果不用size呢？
// 很漂亮的算法，fast先行

func StepBack(head *linked_list.Node, num int) *linked_list.Node {
	fast, slow := head, head
	for i := 0; i < num; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
