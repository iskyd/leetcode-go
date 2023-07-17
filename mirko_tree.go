package main

import (
	"fmt"
)

func printTree(n int) {
	if n == 0 {
		return
	}
	
	for i := 0; i <= n-1; i++ {
		for j := 0; j < n - i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < (i+1)*2-1; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}

	for j := 0; j < n; j++ {
		fmt.Print(" ")
	}
	fmt.Println("|")
}

func main() {
	printTree(3)
	fmt.Println("")
	printTree(1)
	fmt.Println("")
	printTree(0)
	fmt.Println("")
}
