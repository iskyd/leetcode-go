package main

import (
	"fmt"
)

func backtrack(n int, remain int, comb []int, next int, ans *[][]int) {
	if remain == 0 {
		c := make([]int, len(comb))
		copy(c, comb)
		*ans = append(*ans, c)
		return
	}

	for i := next; i <= n; i++ {
		comb = append(comb, i)
		backtrack(n, remain-1, comb, i+1, ans)
		comb = comb[:len(comb)-1]
	}
}

func combine(n int, k int) [][]int {
	ans := [][]int{}
	backtrack(n, k, []int{}, 1, &ans)

	return ans
}

func main() {
	fmt.Println(combine(4, 2))
}
