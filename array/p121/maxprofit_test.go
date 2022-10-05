package p121

import "fmt"

func Example_one() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
	// Output: 5
}

func Example_two() {
	prices := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))
	// Output: 0
}
