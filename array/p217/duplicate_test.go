package p217

import "fmt"

func Example_one() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(nums))
	// Output: true
}

func Example_two() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(nums))
	// Output: false
}

func Example_three() {
	nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(containsDuplicate(nums))
	// Output: true
}
