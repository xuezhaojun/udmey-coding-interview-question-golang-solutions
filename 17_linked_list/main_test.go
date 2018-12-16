package linked_list

import (
	"fmt"
	"testing"
)

// 插入测试
func TestLinkedList_InsertFirst(t *testing.T) {
	list := NewLinkedList()
	list.InsertFirst(1)
	fmt.Println(list, list.Size())
	list.InsertFirst(2)
	fmt.Println(list, list.Size())
	list.InsertFirst(3)
	fmt.Println(list, list.Size())
	fmt.Println(list.GetFirst())
	fmt.Println(list.GetLast())
	list.Clear()
	fmt.Println(list)
}

// 插入尾测试
func TestLinkedList_InsertLast(t *testing.T) {
	list := NewLinkedList()
	list.InsertLast(1)
	fmt.Println(list, list.Size())
	list.InsertLast(2)
	fmt.Println(list, list.Size())
	list.InsertLast(3)
	fmt.Println(list, list.Size())
	fmt.Println(list.GetFirst())
	fmt.Println(list.GetLast())
	list.Clear()
	fmt.Println(list)
}

// 删除头测试
func TestLinkedList_RemoveFirst(t *testing.T) {
	list := NewLinkedList()
	list.InsertFirst(1)
	fmt.Println(list, list.Size())
	list.InsertFirst(2)
	fmt.Println(list, list.Size())
	list.InsertFirst(3)
	fmt.Println(list, list.Size())
	list.RemoveFirst()
	fmt.Println(list, list.Size())
	list.RemoveFirst()
	fmt.Println(list, list.Size())
	list.RemoveFirst()
}

// 删除尾测试
func TestLinkedList_RemoveLast(t *testing.T) {
	list := NewLinkedList()
	fmt.Println(list, list.Size())
	list.InsertFirst(1)
	fmt.Println(list, list.Size())
	list.InsertFirst(2)
	fmt.Println(list, list.Size())
	list.InsertFirst(3)
	fmt.Println(list, list.Size())
	list.RemoveLast()
	fmt.Println(list, list.Size())
	list.RemoveLast()
	fmt.Println(list, list.Size())
	list.RemoveLast()
	fmt.Println(list, list.Size())
}

// 获取第X个的测试
func TestLinkedList_GetAt(t *testing.T) {
	list := NewLinkedList()
	list.InsertFirst(1)
	list.InsertFirst(2)
	list.InsertFirst(3)
	fmt.Println(list.GetAt(0))
	fmt.Println(list.GetAt(1))
	fmt.Println(list.GetAt(2))
	fmt.Println(list, list.Size())
}

// 删除第X个的测试
func TestLinkedList_RemoveAt(t *testing.T) {
	list := NewLinkedList()
	list.InsertFirst(1)
	list.InsertFirst(2)
	list.InsertFirst(3)
	fmt.Println(list, list.Size())
	list.RemoveAt(1)
	fmt.Println(list, list.Size())
}
