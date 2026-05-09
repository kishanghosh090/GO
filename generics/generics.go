package main

import "fmt"

// func printSlice(items []int) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// func printSlice[T int | string](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

func printSlice[T comparable](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

type ApiResponse[T any] struct {
	status  int
	data    T
	message string
	err     error
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 5, 5, 5}
	str := []string{"chai", "elychi"}

	printSlice(nums)
	printSlice(str)
}
