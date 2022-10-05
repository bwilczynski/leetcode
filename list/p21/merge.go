package p21

import "github.com/bwilczynski/leetcode/list"

type ListNode = list.Node

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	p1, p2 := list1, list2
	if list2.Val < list1.Val {
		p1, p2 = list2, list1
	}
	head := p1

	for p1 != nil {
		if p1.Next == nil {
			p1.Next = p2
			break
		}
		if p2.Val < p1.Next.Val {
			tmp := p1.Next
			p1.Next = p2
			p2 = tmp
		}
		p1 = p1.Next
	}
	return head
}
