/*
 * Copyright (c) 2018-2118
 * Author: linceyou
 * LastModified: 18-12-14 下午10:01
 */

package p25

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var result *ListNode
	if k == 1 {
		return head
	}

	var tail *ListNode
	begin := head
	kth := getKth(begin, k)
	if kth == nil {
		return head
	}

	for kth != nil {
		left := kth.Next
		kth.Next = nil

		r := reverse(begin)
		if result == nil {
			result = r
			tail = begin
		} else {
			tail.Next = r
			tail = begin
		}
		begin = left
		kth = getKth(begin, k)
	}

	tail.Next = begin
	return result
}

func getKth(head *ListNode, k int) *ListNode {
	result := head
	for i := 1; i < k; i ++ {
		if result == nil {
			return nil
		} else {
			result = result.Next
		}
	}

	return result
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var newCur *ListNode
	cur := head.Next
	head.Next = nil
	for cur != nil {
		newCur = cur.Next
		cur.Next = head
		head = cur
		cur = newCur
	}
	return head
}
