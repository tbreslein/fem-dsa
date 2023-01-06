// Package list funcs
package list

import (
	"bytes"
	"fmt"
	"strings"
)

type node[T any] struct {
	value T
	next  *node[T]
	prev  *node[T]
}

type List[T any] struct {
	head *node[T]
	tail *node[T]
}

func NewList[T any]() List[T] {
	return List[T]{head: nil, tail: nil}
}

func (l *List[T]) PeekFront() T {
  // if head == nil { return error... }
  return l.head.value
}

func (l *List[T]) PeekBack() T {
  // if tail == nil { return error... }
  return l.tail.value
}

func (l *List[T]) PushFront(val T) {
	newNode := node[T]{value: val, next: l.head, prev: nil}
  if l.head == nil && l.tail == nil {
    l.tail = &newNode
  }
	l.head = &newNode
}

func (l *List[T]) PushBack(val T) {
	newNode := node[T]{value: val, next: nil, prev: l.tail}
  if l.head == nil && l.tail == nil {
    l.head = &newNode
  }
	l.tail.next = &newNode
	l.tail = &newNode
}

func (l *List[T]) PopFront() T {
	// if l.head == nil {
	// 	return , errors.New("list is empty")
	// }
	val := l.head.value
  if l.head.next != nil {
    l.head.next.prev = nil
    l.head = l.head.next
  } else {
    // this is the last element
    l.head = nil
    l.tail = nil
  }
	return val
}

func (l *List[T]) PopBack() T {
	// if l.head == nil {
	// 	return , errors.New("list is empty")
	// }
	val := l.tail.value
  if l.tail.prev != nil {
    l.tail = l.tail.prev
    l.tail.next = nil
  } else {
    // this is the last element
    l.head = nil
    l.tail = nil
  }
	return val
}

func (l *List[T]) ToString() string {
  var buf bytes.Buffer
  buf.WriteString("[ ")
  n := l.head
  for n != nil {
    buf.WriteString(fmt.Sprintf("%v ", n.value))
    n = n.next
  }
  buf.WriteString("]\n")
  return strings.TrimSpace(buf.String())
}
