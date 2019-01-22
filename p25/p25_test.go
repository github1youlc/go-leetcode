/*
 * Copyright (c) 2018-2118
 * Author: linceyou
 * LastModified: 18-12-14 下午10:01
 */

package p25

import (
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	result := reverseKGroup(head, 3)
	for cur := result; cur != nil; cur = cur.Next {
		t.Log(cur.Val)
	}
}

func Test_reverse(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}

	result := reverse(head)
	for cur := result; cur != nil; cur = cur.Next {
		t.Log(cur.Val)
	}
}

func Test_getKth(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}

	kth := getKth(head, 2)
	t.Log(kth.Val)
}
