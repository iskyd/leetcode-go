package main

import "fmt"

func search(nums []int, target int) int {
	var left, right int = 0, len(nums) - 1
	for left <= right {
		var mid int = (left + right) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	arr := []int{1, 2, 3, 4}
	var res int = search(arr, 3)
	fmt.Println(res)
}
