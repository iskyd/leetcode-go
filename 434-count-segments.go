package main

import (
	"fmt"
)

// Example 1:

// Input: s = "Hello, my name is John"
// Output: 5
// Explanation: The five segments are ["Hello,", "my", "name", "is", "John"]
// Example 2:

// Input: s = "Hello"
// Output: 1

func countSegments(s string) int {
	res := 0
	for _, character := range s {
		fmt.Println(character)
	}

	return res
}

func main() {
	fmt.Println(countSegments("Hello, my name is John"))
	//fmt.Println(countSegments(""))
	//fmt.Println(countSegments("                "))
	//fmt.Println(countSegments("a"))
	//fmt.Println(countSegments("Of all the gin joints in all the towns in all the world,   "))
}