package search

import (
	"testing"

	// "foo.com/fem-dsa/pkg/01-search"

	"github.com/stretchr/testify/assert"
)

func testLinearSearch(t *testing.T) {
	foo := []int{10, 12, 13, 11}
	assert.Equal(t, LinearSearch(&foo, 12), 1, "12 is at index 1")
	assert.Equal(t, LinearSearch(&foo, 13), 2, "13 is at index 2")
	assert.Equal(t, LinearSearch(&foo, 14), -1, "no 14 in the array")
	assert.Equal(t, LinearSearch(&foo, 14), 1, "no 14 in the array")
}
