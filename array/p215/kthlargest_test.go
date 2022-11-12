package p215

import "fmt"

func Example_one() {
	nums, k := []int{3, 2, 1, 5, 6, 4}, 2
	fmt.Println(findKthLargest(nums, k))
	// Output: 5
}

func Example_two() {
	nums, k := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4
	fmt.Println(findKthLargest(nums, k))
	// Output: 4
}
