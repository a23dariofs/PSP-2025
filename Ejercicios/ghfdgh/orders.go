package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Error: no has introducido números.")
		os.Exit(1)
	}

	numbers := make([]int, len(args))

	// Convertir y validar
	for i, a := range args {
		n, err := strconv.Atoi(a)
		if err != nil {
			fmt.Println("Error:", a, "no es un número válido.")
			os.Exit(1)
		}
		numbers[i] = n
	}

	fmt.Println("Ascending:")
	fmt.Println(sortAscending(numbers))

	fmt.Println("Descending:")
	fmt.Println(sortDescending(numbers))

	fmt.Println("By even and odd:")
	fmt.Println(sortEvenOdd(numbers))
}

func sortAscending(nums []int) []int {
	result := append([]int{}, nums...)
	sort.Ints(result)
	return result
}

func sortDescending(nums []int) []int {
	result := append([]int{}, nums...)
	sort.Sort(sort.Reverse(sort.IntSlice(result)))
	return result
}

func sortEvenOdd(nums []int) []int {
	result := append([]int{}, nums...)
	sort.Slice(result, func(i, j int) bool {
		if result[i]%2 == 0 && result[j]%2 != 0 {
			return true
		}
		if result[i]%2 != 0 && result[j]%2 == 0 {
			return false
		}
		return result[i] < result[j]
	})
	return result
}
