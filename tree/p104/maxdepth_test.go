package p104

import (
	"testing"

	"github.com/bwilczynski/leetcode/tree"
)

func TestMaxDepth(t *testing.T) {
	tc := []struct {
		name   string
		input  []int
		output int
	}{
		{name: "example1", input: []int{3, 9, 20, tree.NilValue, tree.NilValue, 15, 7}, output: 3},
		{name: "example2", input: []int{1, tree.NilValue, 2}, output: 2},
	}
	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			res := maxDepth(tree.New(tt.input))
			if res != tt.output {
				t.Fatalf("maxDepth(%v) should be %d, got %d.", tt.input, tt.output, res)
			}
		})
	}
}
