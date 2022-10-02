package p680

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tc := []struct {
		name   string
		input  string
		output bool
	}{
		{name: "example1", input: "aba", output: true},
		{name: "example2", input: "abca", output: true},
		{name: "example3", input: "abc", output: false},
		{name: "example4", input: "ebcbbececabbacecbbcbe", output: true},
	}
	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			res := validPalindrome(tt.input)
			if res != tt.output {
				t.Fatalf("validPalindrome(%q) should be %t, got %t.", tt.input, tt.output, res)
			}
		})
	}
}
