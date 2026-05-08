package main

import "fmt"

// slices -> dynamic array
// most used construct
func main() {
	// uninitialized slice is nil
	var nums []int

	nums = append(nums, 10)

	if nums != nil {
		fmt.Println("")
	}

	// capacity -> max nums of elements can fit (when cap hit cap become doble from last capacity)
	// init with 0 not nil
	var newNum = make([]int, 0)

	newNum = append(newNum, 10)
	newNum = append(newNum, 23)

	// fmt.Println(newNum)

	chai := []string{}
	chai = append(chai, "elychi")
	chai = append(chai, "ginger")

	// chai[1] = ""
	// chai = append(chai, "elychi3d")

	// fmt.Println(cap(chai))

	var nums2 = make([]string, len(chai))

	// copy function
	copy(nums2, chai)

	fmt.Print(chai, nums2)

	fmt.Println(chai[0:1])

}
