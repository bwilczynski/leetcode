package p2

import (
	"fmt"

	"github.com/bwilczynski/leetcode/list"
)

func Example_one() {
	l1, l2 := list.New([]int{2, 4, 3}), list.New([]int{5, 6, 4})
	fmt.Println(addTwoNumbers(l1, l2))
	// Output: [7 0 8]
}

func Example_two() {
	l1, l2 := list.New([]int{0}), list.New([]int{0})
	fmt.Println(addTwoNumbers(l1, l2))
	// Output: [0]
}

func Example_three() {
	l1, l2 := list.New([]int{9, 9, 9, 9, 9, 9, 9}), list.New([]int{9, 9, 9, 9})
	fmt.Println(addTwoNumbers(l1, l2))
	// Output: [8 9 9 9 0 0 0 1]
}
