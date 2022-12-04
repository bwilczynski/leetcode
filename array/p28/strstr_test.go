package p28

import "fmt"

func Example_one() {
	haystack, needle := "sadbutsad", "sad"
	fmt.Println(strStr(haystack, needle))
	// Output: 0
}

func Example_two() {
	haystack, needle := "leetcode", "leeto"
	fmt.Println(strStr(haystack, needle))
	// Output: -1
}
