package main

import "fmt"

func main() {
	var nums [3]int

	fmt.Println(len(nums))
	nums[0] = 10
	nums[1] = 22
	nums[2] = 999999999999999999

	fmt.Println(nums)

	newNum := [3]int{1, 2, 3}
	fmt.Println(newNum)

	//2nd arr

	matrix := [2][2]int{{1, 2}, {23, 45}}

	fmt.Println(matrix)
}
