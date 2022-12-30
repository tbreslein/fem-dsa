// Package search funcs
package search

// LinearSearch Just walk the array from the beginning till we find the thing we are looking for.
//
// Time comp: O(n), since we manually walk every single element
func LinearSearch[T comparable](arr *[]T, x T) int {
	for i, a := range *arr {
		if a == x {
			return i
		}
	}
	return -1
}
