package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
  foo := NewList[int]()
  assert.Equal(t, foo.ToString(), "[ ]", "construct an empty list")

  foo.PushFront(10)
	assert.Equal(t, foo.head.value, 10, "head is assigned")
	assert.Equal(t, foo.tail.value, 10, "tail is assigned")
	assert.Nil(t, foo.head.next, "head.next is not pointing anywhere")
	assert.Nil(t, foo.head.prev, "head.prev is not pointing anywhere")
	assert.Nil(t, foo.tail.next, "tail.next is not pointing anywhere")
	assert.Nil(t, foo.tail.prev, "tail.prev is not pointing anywhere")
  assert.Equal(t, foo.ToString(), "[ 10 ]", "converts to string correctly")

  foo.PushFront(9)
  assert.Equal(t, foo.ToString(), "[ 9 10 ]", "PushFront works")

  foo.PushBack(11)
  assert.Equal(t, foo.ToString(), "[ 9 10 11 ]", "PushBack works")

  foo.PushFront(8)
  assert.Equal(t, foo.ToString(), "[ 8 9 10 11 ]", "another PushFront works too")

  foo.PushBack(12)
  assert.Equal(t, foo.ToString(), "[ 8 9 10 11 12 ]", "another PushBack works too")

  foo.PopFront()
  assert.Equal(t, foo.ToString(), "[ 9 10 11 12 ]", "PopFront works")

  foo.PopFront()
  assert.Equal(t, foo.ToString(), "[ 10 11 12 ]", "another PopFront works too")

  foo.PopBack()
  assert.Equal(t, foo.ToString(), "[ 10 11 ]", "PopFront works")

  foo.PopBack()
  assert.Equal(t, foo.ToString(), "[ 10 ]", "another PopFront works too")

  foo.PopBack()
  assert.Equal(t, foo.ToString(), "[ ]", "a last PopFront works too")
}
