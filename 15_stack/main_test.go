package _5_stack

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	s := new(stack)
	s.Push(1)
	s.Push(12)
	s.Push(123)
	fmt.Println(s)
	fmt.Println(s.Pop())
	fmt.Println(s)
	fmt.Println(s.Pop())
	fmt.Println(s)
	fmt.Println(s.Pop())
}
