package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 43, 456, 657, 4}

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	for idx, num := range nums {
		println(idx, num)
	}

	m := map[string]string{"name": "kishan", "age": "20"}

	for k, v := range m {
		println(k, v)
	}

	// 255 -> 1 byte
	// 2 byte

	for i, c := range "hello from chai code" {
		fmt.Println(i, c) // c -> unicode
		fmt.Println(string(c))
	}
}
