package main

import (
	"fmt"
)

func splitIntSlice(s []int) ([]int, []int) {
	mid := len(s) / 2

	return s[:mid], s[mid:]
}

func splitStringSlice(s []string) ([]string, []string) {
	mid := len(s) / 2

	return s[:mid], s[mid:]
}

// function using generics
func splitSlice[T any](s []T) ([]T, []T) {
	mid := len(s) / 2

	return s[:mid], s[mid:]
}

func genericsExample() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(splitIntSlice(nums))
	strs := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	fmt.Println(splitStringSlice(strs))

	fmt.Println("-----------------------------------")

	fmt.Println(splitSlice[int](nums))
	fmt.Println(splitSlice[string](strs))
}
