package main

import (
	"fmt"
)

var cache = map[int]int{}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func robRecursive(nums []int, i int) int {
	if i < 0 {
		return 0
	}

	if v, ok := cache[i]; ok {
		return v
	}

	cache[i] = Max(robRecursive(nums, i-2)+nums[i], robRecursive(nums, i-1))

	return cache[i]
}

func rob(nums []int) int {
	cache = map[int]int{}
	return robRecursive(nums, len(nums)-1)
}

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))

	// fmt.Println(cache)
}
