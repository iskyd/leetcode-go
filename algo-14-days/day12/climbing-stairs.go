package main

import (
	"fmt"
)

var cache = map[int]int{}

func climbStairs(n int) int {
	if n == 0 {
		return 1
	}

	if v, ok := cache[n]; ok {
		return v
	}

	if n >= 1 {
		if _, ok := cache[n-1]; !ok {
			cache[n-1] = climbStairs(n - 1)
		}
	}

	if n >= 2 {
		if _, ok := cache[n-2]; !ok {
			cache[n-2] = climbStairs(n - 2)
		}
	}

	return cache[n-1] + cache[n-2]
}

func main() {
	fmt.Println(climbStairs(44))
}
