package main

import (
	"fmt"
)

func hammingWeight(num uint32) int {
	c := 0
	for num != 0 {
		num &= num - 1
		c++
	}

	return c
}

func main() {
	fmt.Println(hammingWeight(00000000000000000000000000001011))
	fmt.Println(hammingWeight(00000000000000000000000010000000))
}
