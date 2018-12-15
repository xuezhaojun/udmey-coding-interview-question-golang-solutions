package _6_queue_from_stack

import (
	"errors"
	"fmt"
)

type Queue interface {
	// 将 i 添加到 队尾
	Add(i int)
	// 返回队首的i, 如果队空，则返回error
	Remove() (i int, err error)
}

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

// queue from stack 是一个通过 stack 模拟 queue 中FIFO表现的队列
// 两个stack , 一个负责存， 一个负责取
type queueFromStack struct {
	fs *stack // 负责存的
	ss *stack // 负责取的
}

func NewQueue() *queueFromStack {
	return &queueFromStack{
		fs: new(stack),
		ss: new(stack),
	}
}

func (q *queueFromStack) Add(i int) {
	q.fs.Push(i)
}

// 这一步是O(n)的时间复杂度
func (q *queueFromStack) Remove() (ret int, err error) {
	for {
		pop, _ := q.fs.Pop()
		// 取走队尾
		if len(q.fs.slice) == 0 {
			ret = pop
			break
		}
		q.ss.Push(pop)
	}
	for len(q.ss.slice) > 0 {
		pop, _ := q.ss.Pop()
		q.fs.Push(pop)
	}
	return
}

func (q *queueFromStack) String() string {
	return fmt.Sprintf("%#v", q.fs)
}
