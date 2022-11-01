package p11

import "fmt"

func Example_one() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
	// Output: 49
}

func Example_two() {
	height := []int{1, 1}
	fmt.Println(maxArea(height))
	// Output: 1
}
