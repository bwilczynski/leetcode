package p53

import "fmt"

func Example_one() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
	// Output: 6
}

func Example_two() {
	nums := []int{1}
	fmt.Println(maxSubArray(nums))
	// Output: 1
}

func Example_three() {
	nums := []int{5, 4, -1, 7, 8}
	fmt.Println(maxSubArray(nums))
	// Output: 23
}

func Example_four() {
	nums := []int{-1}
	fmt.Println(maxSubArray(nums))
	// Output: -1
}
