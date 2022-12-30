// Package search funcs
package search

import "golang.org/x/exp/constraints"

// LinearSearch Just walk the array from the beginning till we find the thing we are looking for, and return its index.
//
// Time comp: O(n), since we manually walk every single element
func LinearSearch[T comparable](haystack *[]T, needle T) int {
	for i, a := range *haystack {
		if a == needle {
			return i
		}
	}
	return -1
}

// BinarySearch Assuming that the haystack is a sorted array, binary search the needle and return its index.
//
// Binary search starts at the middle of the array and checks whether it is equal, smaller or greater than our needle.
// If it is equal, return the index. If the needle is greater, we know we should continue our search in the upper half
// of the array.
// With that knowledge, we jump to the middle of that part of the array, and do that check again, and recurse that until
// we find our value.
//
// Time comp: O(log(n)), since we can half the size of the haystack with each step, and we only have to look at a single
// element in each step.
func BinarySearch[T constraints.Ordered](haystack *[]T, needle T) int {
  upperBound := len(*haystack) - 1
  lowerBound := 0
	for lowerBound < upperBound {
    pos := (upperBound + lowerBound) / 2
		if (*haystack)[pos] == needle {
			return pos
		} else if (*haystack)[pos] < needle {
      // already know that the needle is not at pos, so put the lowerBound past it
      lowerBound = pos + 1
		} else {
      // upperBound is exclusive
      upperBound = pos
		}
	}
  return -1
}
