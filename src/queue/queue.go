// Package queue funcs
package queue

import "github.com/tbreslein/fem-dsa/src/list"

type Queue[T any] struct {
	list *list.List[T]
}

func NewQueue[T any]() Queue[T] {
	list := list.NewList[T]()
	return Queue[T]{list: &list}
}

func (q *Queue[T]) Push(val T) {
	q.list.PushBack(val)
}

func (q *Queue[T]) Pop(val T) {
	q.list.PopFront()
}

func (q *Queue[T]) Peek() T {
	return q.list.PeekFront()
}

func (q *Queue[T]) ToString() {
	q.list.ToString()
}
