package p20

import "fmt"

func Example_one() {
	s := "()"
	fmt.Println(isValid(s))
	// Output: true
}

func Example_two() {
	s := "()[]{}"
	fmt.Println(isValid(s))
	// Output: true
}

func Example_three() {
	s := "(]"
	fmt.Println(isValid(s))
	// Output: false
}
