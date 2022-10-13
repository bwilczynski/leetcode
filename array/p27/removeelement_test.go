package p27

import "fmt"

func Example_one() {
	nums, val := []int{3, 2, 2, 3}, 3
	k := removeElement(nums, val)
	fmt.Println(nums[:k])
	// Output: [2 2]
}

func Example_two() {
	nums, val := []int{0, 1, 2, 2, 3, 0, 4, 2}, 2
	k := removeElement(nums, val)
	fmt.Println(nums[:k])
	// Output: [0 1 4 0 3]
}

func Example_three() {
	nums, val := []int{3, 3}, 3
	k := removeElement(nums, val)
	fmt.Println(nums[:k])
	// Output: []
}

func Example_four() {
	nums, val := []int{}, 0
	k := removeElement(nums, val)
	fmt.Println(nums[:k])
	// Output: []
}

func Example_five() {
	nums, val := []int{1}, 2
	k := removeElement(nums, val)
	fmt.Println(nums[:k])
	// Output: [1]
}

func Example_six() {
	nums, val := []int{1}, 1
	k := removeElement(nums, val)
	fmt.Println(nums[:k])
	// Output: []
}
