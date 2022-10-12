package p876

import "github.com/bwilczynski/leetcode/list"

type ListNode = list.Node

func middleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
