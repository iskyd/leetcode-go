package main

import (
	"fmt"
)

func runningSum(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		res[i] = nums[i] + res[i - 1]
	}

	return res
}

func main() {
	fmt.Println(runningSum([]int{3,1,2,10,1}))
}