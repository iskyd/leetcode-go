package main

import (
	"fmt"
)

func contains(comb []int, num int) bool {
	for _, n := range comb {
		if n == num {
			return true
		}
	}
	return false
}

func backtrack(nums []int, comb []int, ans *[][]int) {
	if len(comb) == len(nums) {
		c := make([]int, len(comb))
		copy(c, comb)
		*ans = append(*ans, c)
		return
	}

	for i := 0; i < len(nums); i++ {
		if contains(comb, nums[i]) {
			continue
		}
		comb = append(comb, nums[i])
		backtrack(nums, comb, ans)
		comb = comb[:len(comb)-1]
	}
}

func permute(nums []int) [][]int {
	ans := [][]int{}
	backtrack(nums, []int{}, &ans)
	return ans
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}
