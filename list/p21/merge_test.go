package p21

import (
	"fmt"

	"github.com/bwilczynski/leetcode/list"
)

func Example_one() {
	list1, list2 := list.New([]int{1, 2, 4}), list.New([]int{1, 3, 4})
	fmt.Println(mergeTwoLists(list1, list2))
	// Output: [1 1 2 3 4 4]
}

func Example_two() {
	list1, list2 := list.New([]int{}), list.New([]int{})
	fmt.Println(mergeTwoLists(list1, list2))
	// Output: []
}

func Example_three() {
	list1, list2 := list.New([]int{}), list.New([]int{0})
	fmt.Println(mergeTwoLists(list1, list2))
	// Output: [0]
}
