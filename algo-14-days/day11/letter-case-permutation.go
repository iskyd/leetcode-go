package main

import (
	"fmt"
)

func backtrack(s string, comb []byte, ans *[]string, j int) {
	if len(comb) == len(s) {
		*ans = append(*ans, string(comb))
		return
	}

	if j < len(s) {
		if s[j] >= '0' && s[j] <= '9' {
			comb = append(comb, s[j])
			backtrack(s, comb, ans, j+1)
			comb = comb[:len(comb)-1]
		} else {
			comb = append(comb, s[j])
			backtrack(s, comb, ans, j+1)
			comb = comb[:len(comb)-1]

			comb = append(comb, byte(s[j]^32))
			backtrack(s, comb, ans, j+1)
			comb = comb[:len(comb)-1]
		}
	}
}

func letterCasePermutation(s string) []string {
	ans := []string{}
	backtrack(s, []byte{}, &ans, 0)
	return ans
}

func main() {
	s := "a1b2"
	fmt.Println(letterCasePermutation(s))
}
