package main

import (
	"fmt"
)

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	m := make(map[string]int)
	max := 0
	l := 0
	for i := 0; i < len(s); i++ {
		if _, found := m[string(s[i])]; found {
			if m[string(s[i])] >= l {
				l = m[string(s[i])] + 1
			} else {
				max = Max(max, i-l+1)
			}
		} else {
			max = Max(max, i-l+1)
		}
		m[string(s[i])] = i
	}

	return max
}

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
