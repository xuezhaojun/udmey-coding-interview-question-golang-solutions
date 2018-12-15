package _6_queue_from_stack

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	q := NewQueue()
	q.Add(1)
	q.Add(12)
	q.Add(123)
	fmt.Println(q)
	fmt.Println(q.Remove())
	fmt.Println(q)
	fmt.Println(q.Remove())
	fmt.Println(q)
	fmt.Println(q.Remove())
}
