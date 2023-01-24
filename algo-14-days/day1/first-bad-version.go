package main

import "fmt"

func isBadVersion(version int) bool {
	if version >= 2 {
		return true
	}
	return false
}

func firstBadVersion(n int) int {
	var left, right int = 1, n
	for left < right {
		var mid int = left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func main() {
	var res int = firstBadVersion(5)
	fmt.Println(res)
}
