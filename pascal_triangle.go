package main

import (
	"fmt"
)

func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	if numRows <= 0 {
		return res
	}

	res[0] = []int{1}

	for i := 1; i < numRows; i++ {
		row := make([]int, i + 1)
		row[0] = 1

		for j := 1; j < i; j++ {
			row[j] = res[i - 1][j - 1] + res[i - 1][j]
		}

		row[i] = 1

		res[i] = row
	}

	return res
}
