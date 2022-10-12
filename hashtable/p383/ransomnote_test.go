package p383

import "fmt"

func Example_one() {
	ransomNote, magazine := "a", "b"
	fmt.Println(canConstruct(ransomNote, magazine))
	// Output: false
}

func Example_two() {
	ransomNote, magazine := "aa", "ab"
	fmt.Println(canConstruct(ransomNote, magazine))
	// Output: false
}

func Example_three() {
	ransomNote, magazine := "aa", "aab"
	fmt.Println(canConstruct(ransomNote, magazine))
	// Output: true
}
