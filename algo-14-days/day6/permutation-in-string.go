package main

import (
	"fmt"
)

func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	m1 := [26]int{}
	for i := 0; i < len(s1); i++ {
		m1[s1[i]-'a']++
	}

	var i, j int = 0, 0
	m2 := [26]int{} // current window
	lengthm2 := 0

	for j < len(s2) {
		m2[s2[j]-'a']++
		lengthm2++

		if lengthm2 == len(s1) {
			if m1 == m2 {
				return true
			}

			m2[s2[i]-'a']--
			lengthm2--

			i++
			j++
		} else if lengthm2 < len(s1) {
			j++
		} else {
			// Impossible branch
		}
	}

	return false
}

func main() {
	s1 := "ab"
	s2 := "eidbaooo"

	fmt.Println(checkInclusion(s1, s2))
}
