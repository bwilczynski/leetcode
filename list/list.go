package list

import (
	"fmt"
)

// Definition for singly-linked list.
type Node struct {
	Val  int
	Next *Node
}

func New(a []int) *Node {
	if len(a) == 0 {
		return nil
	}
	head := &Node{Val: a[0]}
	res := head
	for _, v := range a[1:] {
		head.Next = &Node{Val: v}
		head = head.Next
	}
	return res
}

func (l *Node) String() string {
	var a []int
	for l != nil {
		a = append(a, l.Val)
		l = l.Next
	}
	return fmt.Sprintf("%v", a)
}
