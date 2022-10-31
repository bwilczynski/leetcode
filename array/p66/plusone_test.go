package p66

import "fmt"

func Example_one() {
	digits := []int{1, 2, 3}
	fmt.Println(plusOne(digits))
	// Output: [1 2 4]
}

func Example_two() {
	digits := []int{4, 3, 2, 1}
	fmt.Println(plusOne(digits))
	// Output: [4 3 2 2]
}

func Example_three() {
	digits := []int{9}
	fmt.Println(plusOne(digits))
	// Output: [1 0]
}
