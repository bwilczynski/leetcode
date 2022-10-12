package p1089

import "fmt"

func Example_one() {
	arr := []int{1, 0, 2, 3, 0, 4, 5, 0}
	duplicateZeros(arr)
	fmt.Println(arr)
	// Output: [1 0 0 2 3 0 0 4]
}

func Example_two() {
	arr := []int{1, 2, 3}
	duplicateZeros(arr)
	fmt.Println(arr)
	// Output: [1 2 3]
}
