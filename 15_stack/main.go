package _5_stack

import "errors"

type Stack interface {
	Push(i int)
	Pop() (i int, err error)
}

type stack struct {
	slice []int
}

func (s *stack) Push(i int) {
	s.slice = append(s.slice, i)
}

func (s *stack) Pop() (i int, err error) {
	if len(s.slice) == 0 {
		return i, errors.New("no int in stack")
	}
	i = s.slice[len(s.slice)-1]
	s.slice = s.slice[:len(s.slice)-1]
	return
}
