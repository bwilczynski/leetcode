package p2

import "github.com/bwilczynski/leetcode/list"

type ListNode = list.Node

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	a, b := l1.Val, l2.Val
	res := &ListNode{}
	head := res
	rem := 0
	for {
		head.Val = ((a + b + rem) % 10)
		rem = (a + b + rem) / 10

		if l1.Next == nil && l2.Next == nil {
			if rem != 0 {
				head.Next = &ListNode{Val: rem}
			}
			break
		}
		head.Next = &ListNode{}
		head = head.Next

		a = 0
		if l1.Next != nil {
			l1 = l1.Next
			a = l1.Val
		}
		b = 0
		if l2.Next != nil {
			l2 = l2.Next
			b = l2.Val
		}
	}
	return res
}
