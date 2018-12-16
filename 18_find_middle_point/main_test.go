package _8_find_middle_point

import (
	"fmt"
	"golang-coding-interview-questions-and-solutions/17_linked_list"
	"testing"
)

func Test(t *testing.T) {
	list := linked_list.NewLinkedList()
	list.InsertFirst(1)
	list.InsertFirst(2)
	list.InsertFirst(3)
	list.InsertFirst(4)

	// list : 4 > 3 > 2 > 1

	fmt.Println(middle(list))
}
