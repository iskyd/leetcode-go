package main

import (
	"fmt"
	"strings"
)

func reverseString(s string) string {
	var res string
	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}
	return res
}

func reverseWords(s string) string {
	splitted := strings.Split(s, " ")
	for i := 0; i < len(splitted); i++ {
		splitted[i] = reverseString(splitted[i])
	}

	return strings.Join(splitted, " ")
}

func main() {
	s := "Let's take LeetCode contest"
	fmt.Println(reverseWords(s))
}
