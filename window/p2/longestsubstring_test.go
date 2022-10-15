package p2

import "fmt"

func Example_one() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
	// Output: 3
}

func Example_two() {
	s := "bbbbb"
	fmt.Println(lengthOfLongestSubstring(s))
	// Output: 1
}

func Example_three() {
	s := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
	// Output: 3
}

func Example_four() {
	s := ""
	fmt.Println(lengthOfLongestSubstring(s))
	// Output: 0
}

func Example_five() {
	s := " "
	fmt.Println(lengthOfLongestSubstring(s))
	// Output: 1
}
