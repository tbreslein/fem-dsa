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
type List[T comparable] struct {
	head   *node[T]
	tail   *node[T]
	Length int
}

// NewList create a new empty list
func NewList[T comparable]() List[T] {
	return List[T]{head: nil, tail: nil, Length: 0}
}

// PeekFront returns copy of value at head or an error
func (l *List[T]) PeekFront() (T, error) {
	if l.head == nil {
		return *new(T), errors.New("cannot peek at empty list")
	}
	return l.head.value, nil
}

// PeekBack returns copy of value at tail or an error
func (l *List[T]) PeekBack() (T, error) {
	if l.tail == nil {
		return *new(T), errors.New("cannot peek at empty list")
	}
	return l.tail.value, nil
}

// PushFront push new val to the front of the list
func (l *List[T]) PushFront(val T) {
	newNode := node[T]{value: val, next: l.head, prev: nil}
	l.Length++
	if l.head == nil && l.tail == nil {
		l.tail = &newNode
	}
	if l.head != nil {
		l.head.prev = &newNode
	}
	l.head = &newNode
}

// PushBack push new val to the back of the list
func (l *List[T]) PushBack(val T) {
	newNode := node[T]{value: val, next: nil, prev: l.tail}
	l.Length++
	if l.head == nil && l.tail == nil {
		l.head = &newNode
	}
	if l.tail != nil {
		l.tail.next = &newNode
	}
	l.tail = &newNode
}

// PopFront pop the front of the list and return the value, unless the list is empty
func (l *List[T]) PopFront() (T, error) {
	if l.head == nil {
		return *new(T), errors.New("cannot pop of empty list")
	}
	val := l.head.value
	if l.head.next != nil {
		l.head.next.prev = nil
		l.head = l.head.next
	} else {
		// this is the last element
		l.head = nil
		l.tail = nil
	}
	l.Length--
	return val, nil
}

// PopBack pop the back of the list and return the value, unless the list is empty
func (l *List[T]) PopBack() (T, error) {
	if l.tail == nil {
		return *new(T), errors.New("cannot peek at empty list")
	}
	val := l.tail.value
	if l.tail.prev != nil {
		l.tail = l.tail.prev
		l.tail.next = nil
	} else {
		// this is the last element
		l.head = nil
		l.tail = nil
	}
	l.Length--
	return val, nil
}

// InsertAt inserts value at index; may return an error
func (l *List[T]) InsertAt(val T, idx int) error {
	if idx > l.Length {
		return errors.New("cannot insert out of bounds of list")
	} else if idx == l.Length {
		l.PushBack(val)
		return nil
	} else if idx == 0 {
		l.PushFront(val)
		return nil
	}
	l.Length++

	// find node in front of idx
	currentNode := l.head
	for i := 0; i < idx; i++ {
		currentNode = currentNode.next
	}
	// attach new node first...
	newNode := node[T]{value: val, next: currentNode.next, prev: currentNode}

	// ... then break old links
	currentNode.prev.next = &newNode
	currentNode.prev = &newNode
	return nil
}

// Remove finds the value and removes it from the list
func (l *List[T]) Remove(val T) error {
	if l.Length == 0 {
		return errors.New("cannot remove items from an empty list")
		// the else if l.head.value == val will be handled by the for loop directly
	} else if l.tail.value == val {
		l.PopBack()
	}

	l.Length--
	currentNode := l.head
	for i := 0; i < l.Length; i++ {
		if i == l.Length-1 && currentNode.value != val {
			return errors.New("unable to find element in list")
		} else if currentNode.value == val {
			break
		}
		currentNode = currentNode.next
	}

	prevNode := currentNode.prev
	nextNode := currentNode.next
	currentNode.prev.next = nextNode
	currentNode.next.prev = prevNode

	if currentNode == l.head {
		l.head = currentNode.next
	}
	if currentNode == l.tail {
		l.tail = currentNode.prev
	}

	// now that nothing is pointing at currentNode anymore, we need to make sure that currentNode isn't pointing
	// to anything either
	currentNode.next = nil
	currentNode.prev = nil

	// now currentNode will be swept by the GC

	return nil
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
