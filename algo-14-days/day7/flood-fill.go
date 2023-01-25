package main

import (
	"fmt"
)

type stack [][]int

func (s stack) Push(v []int) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, []int) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) IsEmpty() bool {
	return len(s) == 0
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	s := make(stack, 0)
	v := []int{sr, sc}
	s = s.Push([]int{sr, sc})
	originalColor := image[sr][sc]

	if originalColor == color {
		return image
	}

	for !s.IsEmpty() {
		s, v = s.Pop()
		r, c := v[0], v[1]
		if r < 0 || r >= len(image) || c < 0 || c >= len(image[0]) || image[r][c] != originalColor {
			continue
		}

		image[r][c] = color
		s = s.Push([]int{r - 1, c})
		s = s.Push([]int{r + 1, c})
		s = s.Push([]int{r, c - 1})
		s = s.Push([]int{r, c + 1})
	}

	return image
}

func main() {
	image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	fmt.Println(floodFill(image, 0, 0, 0))
	image = [][]int{{0, 0, 0}, {0, 0, 0}}
	fmt.Println(floodFill(image, 0, 0, 0))
}
