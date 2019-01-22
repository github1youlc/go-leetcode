/*
 * Copyright (c) 2018-2118
 * Author: linceyou
 * LastModified: 18-12-14 下午9:31
 */

package p24

import (
	"testing"
)

func Test_swapPairs(t *testing.T) {
	head := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}

	result := swapPairs(head)
	for cur := result; cur != nil; cur = cur.Next {
		t.Log(cur.Val)
	}
}
