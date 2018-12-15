package _4_queue

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	q := new(queue)
	q.Enqueue(1)
	q.Enqueue(12)
	q.Enqueue(123)
	fmt.Println(q)
	fmt.Println(q.Dequeue())
	fmt.Println(q)
	fmt.Println(q.Dequeue())
	fmt.Println(q)
	fmt.Println(q.Dequeue())
}
