package main

import (
	"fmt"
)

type queue [][]int

func (q queue) Push(v []int) queue {
	return append(q, v)
}

func (q queue) Pop() (queue, []int) {
	l := len(q)
	return q[1:l], q[0]
}

func (q queue) IsEmpty() bool {
	return len(q) == 0
}

func updateMatrix(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])
	queue := make(queue, 0)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				queue = queue.Push([]int{i, j})
			} else {
				mat[i][j] = -1
			}
		}
	}

	v := []int{}
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for !queue.IsEmpty() {
		queue, v = queue.Pop()
		i, j := v[0], v[1]

		for _, dir := range dirs {
			x, y := i+dir[0], j+dir[1]
			if x >= 0 && x < m && y >= 0 && y < n && mat[x][y] == -1 {
				mat[x][y] = mat[i][j] + 1
				queue = queue.Push([]int{x, y})
			}
		}
	}

	return mat
}

func main() {
	mat := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	fmt.Println(updateMatrix(mat))
}
