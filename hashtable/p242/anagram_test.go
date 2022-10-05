package p242

import "fmt"

func Example_one() {
	s, t := "anagram", "nagaram"
	fmt.Println(isAnagram(s, t))
	// Output: true
}

func Example_two() {
	s, t := "rat", "car"
	fmt.Println(isAnagram(s, t))
	// Output: false
}
