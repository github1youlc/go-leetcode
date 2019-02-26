/*
 * Copyright (c) 2018-2118
 * Author: linceyou
 * LastModified: 18-12-14 下午9:31
 */

package p24

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var result *ListNode
	var t *ListNode

	cur := head

	var left *ListNode
	for cur != nil && cur.Next != nil {
		left = cur.Next.Next
		cur.Next.Next = cur
		if t == nil {
			result = cur.Next
			t = cur
		} else {
			t.Next = cur.Next
			t = cur
		}
		cur = left
	}

	t.Next = cur

	return result
}
