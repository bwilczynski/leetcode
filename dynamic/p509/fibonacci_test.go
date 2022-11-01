package p509

import "fmt"

func Example_one() {
	n := 2
	fmt.Println(fib(n))
	// Output: 1
}

func Example_two() {
	n := 3
	fmt.Println(fib(n))
	// Output: 2
}

func Example_three() {
	n := 4
	fmt.Println(fib(n))
	// Output: 3
}

func Example_four() {
	n := 30
	fmt.Println(fib(n))
	// Output: 832040
}
