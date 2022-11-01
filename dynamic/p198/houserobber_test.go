package p198

import "fmt"

func Example_one() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(rob(nums))
	// Output: 4
}

func Example_two() {
	nums := []int{2, 7, 9, 3, 1}
	fmt.Println(rob(nums))
	// Output: 12
}
