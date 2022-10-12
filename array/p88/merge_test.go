package p88

import "fmt"

func Example_one() {
	nums1, m := []int{1, 2, 3, 0, 0, 0}, 3
	nums2, n := []int{2, 5, 6}, 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
	// Output: [1 2 2 3 5 6]
}

func Example_two() {
	nums1, m := []int{1}, 1
	nums2, n := []int{}, 0
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
	// Output: [1]
}

func Example_three() {
	nums1, m := []int{0}, 0
	nums2, n := []int{1}, 1
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
	// Output: [1]
}
