package p14

import "fmt"

func Example_one() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
	// Output: fl
}

func Example_two() {
	strs := []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(strs))
	// Output:
}

func Example_three() {
	strs := []string{"ab", "a"}
	fmt.Println(longestCommonPrefix(strs))
	// Output: a
}
