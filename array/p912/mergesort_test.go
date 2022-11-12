package p912

import "fmt"

func Example_one() {
	nums := []int{5, 2, 3, 1}
	fmt.Println(sortArray(nums))
	// Output: [1 2 3 5]
}

func Example_two() {
	nums := []int{5, 1, 1, 2, 0, 0}
	fmt.Println(sortArray(nums))
	// Output: [0 0 1 1 2 5]
}
