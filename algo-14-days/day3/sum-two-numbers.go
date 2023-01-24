package main

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	var left, right int = 0, len(numbers) - 1
	for left < right {
		var sum int = numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return []int{-1, -1}
}

func main() {
	arr := []int{2, 7, 11, 15}
	target := 9
	var res []int = twoSum(arr, target)
	fmt.Println(res)
}
