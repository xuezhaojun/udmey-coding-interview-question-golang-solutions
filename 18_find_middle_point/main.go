package _8_find_middle_point

import "golang-coding-interview-questions-and-solutions/17_linked_list"

// 链表常用方法：
// fast,show point
// fast 到终点， show 刚好跑一半
func middle(ls *linked_list.LinkedList) (data interface{}) {
	head := ls.Head()
	if head == nil {
		return nil
	}

	fast, show := head, head

	// fast 一次移动2步，show 一次移动1步
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		show = show.Next
	}

	return show.Data
}
