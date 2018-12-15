package _4_queue

import (
	"errors"
	"fmt"
)

// 队列
// 面试中有可能会问到，如何用golang实现一个队列
// 要实现队列，首先要指导队列是什么？
// 表现特征：FIFO
// 含有的方法： Enqueue , Dequeue
// 下面实现一个 int 的 queue

type Queue interface {
	// 将 i 添加到 队尾
	Enqueue(i int)
	// 返回队首的i, 如果队空，则返回error
	Dequeue() (i int, err error)
}

type queue struct {
	slice []int
}

func (q *queue) Enqueue(i int) {
	q.slice = append(q.slice, i)
}

func (q *queue) Dequeue() (i int, err error) {
	if len(q.slice) == 0 {
		return 0, errors.New("no int in queue")
	}
	// 获取队首
	i = q.slice[0]
	// 删除队首
	q.slice = q.slice[1:]
	return
}

func (q *queue) String() string {
	return fmt.Sprintf("%#v", q.slice)
}
