package p70

import "fmt"

func Example_one() {
	n := 2
	fmt.Println(climbStairs(n))
	// Output: 2
}

func Example_two() {
	n := 3
	fmt.Println(climbStairs(n))
	// Output: 3
}

func Example_three() {
	n := 45
	fmt.Println(climbStairs(n))
	// Output: 1836311903
}
