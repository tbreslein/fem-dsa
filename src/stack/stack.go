// Package stack funcs
package stack

import "github.com/tbreslein/fem-dsa/src/list"

type Stack[T any] struct {
	List *list.List[T]
}

func NewStack[T any]() Stack[T] {
	list := list.NewList[T]()
	return Stack[T]{List: &list}
}

func (s *Stack[T]) Push(val T) {
	s.List.PushFront(val)
}

func (s *Stack[T]) Pop() T {
	return s.List.PopFront()
}

func (s *Stack[T]) Peek() T {
	return s.List.PeekFront()
}

func (s *Stack[T]) ToString() {
	s.List.ToString()
}
