package p746

import "fmt"

func Example_one() {
	cost := []int{10, 15, 20}
	fmt.Println(minCostClimbingStairs(cost))
	// Output: 15
}

func Example_two() {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println(minCostClimbingStairs(cost))
	// Output: 6
}
