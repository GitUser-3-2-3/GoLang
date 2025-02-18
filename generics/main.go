package main

import "fmt"

func splitSlice[T any](s []T) ([]T, []T) {
	centre := len(s) / 2
	return s[:centre], s[centre:]
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	firstHalf, secondHalf := splitSlice(s)
	fmt.Print(firstHalf, secondHalf)

	// generate a string example
	s2 := []string{"a", "b", "c", "d", "e"}
	firstHalf2, secondHalf2 := splitSlice(s2)
	fmt.Print(firstHalf2, secondHalf2)
}
