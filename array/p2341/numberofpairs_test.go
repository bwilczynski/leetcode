package p2341

import "fmt"

func Example_one() {
	nums := []int{1, 3, 2, 1, 3, 2, 2}
	fmt.Println(numberOfPairs(nums))
	// Output: [3 1]
}

func Example_two() {
	nums := []int{1, 1}
	fmt.Println(numberOfPairs(nums))
	// Output: [1 0]
}

func Example_three() {
	nums := []int{0}
	fmt.Println(numberOfPairs(nums))
	// Output: [0 1]
}
