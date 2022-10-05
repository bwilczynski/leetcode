package p1

import "fmt"

func Example_one() {
	nums, target := []int{2, 7, 11, 15}, 9
	fmt.Println(twoSum(nums, target))
	// Output: [0 1]
}

func Example_two() {
	nums, target := []int{3, 2, 4}, 6
	fmt.Println(twoSum(nums, target))
	// Output: [1 2]
}

func Example_three() {
	nums, target := []int{3, 3}, 6
	fmt.Println(twoSum(nums, target))
	// Output: [0 1]
}
