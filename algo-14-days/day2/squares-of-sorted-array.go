package main

import (
	"fmt"
)

func sortedSquares(nums []int) []int {
	var res = make([]int, len(nums))
	var left, right int = 0, len(nums) - 1
	for left <= right {
		var leftSquare, rightSquare int = nums[left] * nums[left], nums[right] * nums[right]
		if leftSquare > rightSquare {
			res[right-left] = leftSquare
			left++
		} else {
			res[right-left] = rightSquare
			right--
		}
	}

	return res
}

func main() {
	arr := []int{-4, -1, 0, 3, 10}
	var res []int = sortedSquares(arr)
	fmt.Println(res)
}
