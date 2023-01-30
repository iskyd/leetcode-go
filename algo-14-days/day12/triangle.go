package main

import (
	"fmt"
)

var cache = map[string]int{}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minimumTotalRecursive(triangle [][]int, i int, j int) int {
	if i == len(triangle)-1 {
		return triangle[i][j]
	}
	key := fmt.Sprintf("%d-%d", i, j)

	if v, ok := cache[key]; ok {
		return v
	}

	cache[key] = Min(minimumTotalRecursive(triangle, i+1, j), minimumTotalRecursive(triangle, i+1, j+1)) + triangle[i][j]

	return cache[key]
}

func minimumTotal(triangle [][]int) int {
	cache = map[string]int{}
	return minimumTotalRecursive(triangle, 0, 0)
}

func main() {
	fmt.Println(minimumTotal([][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}))
}
