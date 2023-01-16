// Package queue funcs
package queue

import "github.com/tbreslein/fem-dsa/src/list"

// Queue FIFO list structure
type Queue[T any] struct {
	list *list.List[T]
}

// NewQueue Construct a new queue
func NewQueue[T any]() Queue[T] {
	list := list.NewList[T]()
	return Queue[T]{list: &list}
}

// Push push a value into the queue
func (q *Queue[T]) Push(val T) {
	q.list.PushBack(val)
}

// Pop pop a value of the queue
func (q *Queue[T]) Pop(val T) {
	q.list.PopFront()
}

// Peek returns a copy of the value at the front of the queue, unless the queue is empty
func (q *Queue[T]) Peek() (T, error) {
	return q.list.PeekFront()
}

// ToString convert the queue to a string
func (q *Queue[T]) ToString() string {
	return q.list.ToString()
}
