package tree

import "math"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

const NilValue = math.MinInt

func New(a []int) *Node {
	return createNode(a, 0)
}

func createNode(a []int, i int) *Node {
	var root *Node
	if i < len(a) && a[i] != NilValue {
		root = &Node{Val: a[i]}
		root.Left = createNode(a, 2*i+1)
		root.Left = createNode(a, 2*i+2)
	}
	return root
}
