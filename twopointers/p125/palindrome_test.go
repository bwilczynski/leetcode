package p125

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tc := []struct {
		name   string
		input  string
		output bool
	}{
		{name: "example1", input: "A man, a plan, a canal: Panama", output: true},
		{name: "example2", input: "race a car", output: false},
		{name: "example3", input: " ", output: true},
	}
	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			res := isPalindrome(tt.input)
			if res != tt.output {
				t.Fatalf("isPalindrome(%q) should be %t, got %t.", tt.input, tt.output, res)
			}
		})
	}
}
