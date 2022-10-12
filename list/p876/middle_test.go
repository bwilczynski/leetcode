package p876

import (
	"fmt"

	"github.com/bwilczynski/leetcode/list"
)

func Example_one() {
	head := list.New([]int{1, 2, 3, 4, 5})
	fmt.Println(middleNode(head))
	// Output: [3 4 5]
}

func Example_two() {
	head := list.New([]int{1, 2, 3, 4, 5, 6})
	fmt.Println(middleNode(head))
	// Output: [4 5 6]
}
