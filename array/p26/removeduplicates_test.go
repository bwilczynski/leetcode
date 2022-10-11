package p26

import "fmt"

func Example_one() {
	nums := []int{1, 1, 2}
	k := removeDuplicates(nums)
	fmt.Printf("%d, nums = %v", k, nums[:k])
	// Output: 2, nums = [1 2]
}

func Example_two() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k := removeDuplicates(nums)
	fmt.Printf("%d, nums = %v", k, nums[:k])
	// Output: 5, nums = [0 1 2 3 4]
}
