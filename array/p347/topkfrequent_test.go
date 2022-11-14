package p347

import "fmt"

func Example_one() {
	nums, k := []int{1, 1, 1, 2, 2, 3}, 2
	fmt.Println(topKFrequent(nums, k))
	// Output: [1 2]
}

func Example_two() {
	nums, k := []int{1}, 1
	fmt.Println(topKFrequent(nums, k))
	// Output: [1]
}
