package p867

import "fmt"

func Example_one() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(transpose(matrix))
	// Output: [[1 4 7] [2 5 8] [3 6 9]]
}

func Example_two() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(transpose(matrix))
	// Output: [[1 4] [2 5] [3 6]]
}
