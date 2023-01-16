// Package list funcs
package list

import (
	"bytes"
  "errors"
	"fmt"
	"strings"
)

type node[T any] struct {
	value T
	next  *node[T]
	prev  *node[T]
}

// List double linked list
type List[T any] struct {
	head *node[T]
	tail *node[T]
}

// NewList create a new empty list
func NewList[T any]() List[T] {
	return List[T]{head: nil, tail: nil}
}

// PeekFront returns copy of value at head or an error
func (l *List[T]) PeekFront() (T, error) {
	if l.head == nil { return *new(T), errors.New("cannot peek at empty list") }
	return l.head.value, nil
}

// PeekBack returns copy of value at tail or an error
func (l *List[T]) PeekBack() (T, error) {
	if l.tail == nil { return *new(T), errors.New("cannot peek at empty list") }
	return l.tail.value, nil
}

// PushFront push new val to the front of the list
func (l *List[T]) PushFront(val T) {
	newNode := node[T]{value: val, next: l.head, prev: nil}
	if l.head == nil && l.tail == nil {
		l.tail = &newNode
	}
	l.head = &newNode
}

// PushBack push new val to the back of the list
func (l *List[T]) PushBack(val T) {
	newNode := node[T]{value: val, next: nil, prev: l.tail}
	if l.head == nil && l.tail == nil {
		l.head = &newNode
	}
	l.tail.next = &newNode
	l.tail = &newNode
}

// PopFront pop the front of the list and return the value, unless the list is empty
func (l *List[T]) PopFront() (T, error) {
	if l.head == nil { return *new(T), errors.New("cannot pop of empty list") }
	val := l.head.value
	if l.head.next != nil {
		l.head.next.prev = nil
		l.head = l.head.next
	} else {
		// this is the last element
		l.head = nil
		l.tail = nil
	}
	return val, nil
}

// PopBack pop the back of the list and return the value, unless the list is empty
func (l *List[T]) PopBack() (T, error) {
	if l.tail == nil { return *new(T), errors.New("cannot peek at empty list") }
	val := l.tail.value
	if l.tail.prev != nil {
		l.tail = l.tail.prev
		l.tail.next = nil
	} else {
		// this is the last element
		l.head = nil
		l.tail = nil
	}
	return val, nil
}

// ToString Convert the list to a string
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
