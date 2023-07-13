package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	i := 0
    for _, v := range nums {
    	if v != val {
    		nums[i] = v
    		i++
    	}
    }

    return i
}

func main() {
	nums := []int{3,2,2,3}
	fmt.Println(removeElement(nums, 3))
	fmt.Println(nums)
}