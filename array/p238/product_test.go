package p238

import "fmt"

func Example_one() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
	// Output: [24 12 8 6]
}

func Example_two() {
	nums := []int{-1, 1, 0, -3, 3}
	fmt.Println(productExceptSelf(nums))
	// Output: [0 0 9 0 0]
}
