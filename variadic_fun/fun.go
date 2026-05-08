package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 345, 5, 6}
	multiParams(nums)
	fmt.Println(1, 2, 3, 4)
	println(multiParams(2, 3, 5, "6", 66, true))
}

func multiParams(nums ...interface{}) int {
	return len(nums)
}
