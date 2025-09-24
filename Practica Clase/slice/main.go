package main

import (
	"fmt"
	"slices"
)

func main() {
	slc := make([]int, 5)
	slc[0] = 42
	slc[3] = 17
	fmt.Println("slc:", slc)
	slices.SortFunc(slc, func(a, b int) int {
		return b - a
	})
	fmt.Println("sorted:", slc)
}
