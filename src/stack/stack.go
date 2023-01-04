// Package stack funcs
package stack

import "github.com/tbreslein/fem-dsa/src/list"

type Stack[T any] struct {
	list *list.List[T]
}

func NewStack[T any]() Stack[T] {
	list := list.NewList[T]()
	return Stack[T]{list: &list}
}

func (s *Stack[T]) Push(val T) {
	s.list.PushFront(val)
}

func (s *Stack[T]) Pop(val T) {
	s.list.PopFront()
}

func (s *Stack[T]) Peek() T {
	return s.list.PeekFront()
}

func (s *Stack[T]) ToString() {
	s.list.ToString()
}
