// Package sort funcs
package sort

import (
	"golang.org/x/exp/constraints"
)

func BubbleSort[T constraints.Ordered](arr *[]T) {
  for i := 0; i < len(*arr); i++ {
    for j := 0; j < len(*arr) - i - 1; j++ {
      if (*arr)[j] > (*arr)[j+1] {
        temp := (*arr)[j+1]
        (*arr)[j+1] = (*arr)[j]
        (*arr)[j] = temp
      }
    }
  }
}
