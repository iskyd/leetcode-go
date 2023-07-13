package main

import (
	"fmt"
)

func richestCustomer(accounts [][]int) int {
	maxWealth := 0

	for i := 0; i < len(accounts); i++ {
		currentCustomerWealth := 0
		for j := 0; j < len(accounts[i]); j++ {
			currentCustomerWealth += accounts[i][j]
		}

		if currentCustomerWealth >= maxWealth {
			maxWealth = currentCustomerWealth
		}
	}

	return maxWealth
}

func main() {
	fmt.Println(richestCustomer([][]int{
		{7, 1, 3},
		{2, 8, 7},
		{1, 9, 5},
	}))
}