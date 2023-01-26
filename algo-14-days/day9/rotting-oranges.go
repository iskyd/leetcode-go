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

func orangesRotting(grid [][]int) int {
	queue := make(queue, 0)
	fresh := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				queue = queue.Push([]int{i, j})
			} else if grid[i][j] == 1 {
				fresh++
			}
		}
	}

	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	v := []int{}
	ans := 0
	for !queue.IsEmpty() && fresh > 0 {
		ans++
		l := len(queue)

		for i := 0; i < l; i++ {
			queue, v = queue.Pop()
			m, n := v[0], v[1]
			for _, dir := range dirs {
				if x, y := m+dir[0], n+dir[1]; x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] == 1 {
					grid[x][y] = 2
					fresh--
					queue = queue.Push([]int{x, y})
				}
			}
		}
	}

	if fresh > 0 {
		return -1
	}
	return ans
}

func main() {
	// grid := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	// grid := [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}
	// grid := [][]int{{0, 2}}
	grid := [][]int{{2, 1, 1}, {1, 1, 1}, {0, 1, 2}}
	fmt.Println(orangesRotting(grid))
}
