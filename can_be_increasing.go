package main

import (
	"fmt"
)

func canBeIncreasing(nums []int) bool {
	cnt := 0
    for i := 1; i < len(nums) && cnt < 2; i++ {
    	if nums[i - 1] >= nums[i] {
    		cnt++
    		if i > 1 && nums[i - 2] >= nums[i] {
    			nums[i] = nums[i - 1]
    		}
    	}
    }

    return cnt < 2
}

func main() {
	fmt.Println(canBeIncreasing([]int{1,2,10,5,7}))
	fmt.Println(canBeIncreasing([]int{2,3,1,2}))
	fmt.Println(canBeIncreasing([]int{1,1,1}))
	fmt.Println(canBeIncreasing([]int{105,924,32,968}))
}